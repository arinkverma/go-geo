package db

import (
	"os"
	"log"
	"strings"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const (
	CITY_KEY = "idx:cities"
)

type RedisContext struct{
	pool *redis.Pool
}

func (self RedisContext) Get () redis.Conn{
	return self.pool.Get()
}

func (self RedisContext) GeoAdd (latitude string, longitude string, value string) (int, error){
	conn := self.Get()
	defer conn.Close()
	return redis.Int(conn.Do("GEOADD", CITY_KEY, longitude, latitude, value))
}

func (self RedisContext) Ping () (string, error){
	conn := self.Get()
	defer conn.Close()
	return redis.String(conn.Do("PING"))
}

func (self RedisContext) GeoRadius (countryCode string, latitude string, longitude string) (string, error){
	conn := self.Get()
	defer conn.Close()
	log.Printf("GeoRadius: [countryCode:%s, latitude %s, longitude %s]", countryCode, latitude, longitude)
	values, err := redis.Strings(conn.Do("GEORADIUS", CITY_KEY, longitude, latitude, 1000, "km", "ASC"))
	value := ""
	for _, record := range values {
		if strings.HasSuffix(record, fmt.Sprintf(":%s", strings.ToUpper(countryCode))) {
			value = record
			break
		}
	}
	return value, err
}

func(self RedisContext) waitForRedis() {
	conn := self.Get()
	defer conn.Close()
	for true {
		log.Println("Pinging to Redis")
		pong, err := self.Ping()
		if err == nil {
			log.Printf("Redis says %s", pong)
			break
		}
	}
}

func MakeRedisContext() *RedisContext{
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
	ctx := &RedisContext{pool}
	ctx.waitForRedis()
	return ctx
}
