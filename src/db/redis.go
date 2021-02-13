package db

import (
	"log"
	"fmt"
	"strings"
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
	values, err := redis.Strings(conn.Do("GEORADIUS", CITY_KEY, longitude, latitude, 1000, "km", "ASC"))
	if err != nil{
		return "", err
	}
	for _, record := range values {
		if strings.HasSuffix(record, fmt.Sprintf(":%s", countryCode)) {
			return record, err
		}
	}
	return "", fmt.Errorf("Can't find any city within 1000km and country %s", countryCode)
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
				log.Fatal("ERROR: fail init redis: %s", err.Error())
			}
			return conn, err
		},
	}
	ctx := &RedisContext{pool}
	ctx.waitForRedis()
	return ctx
}
