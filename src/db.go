package main

import (
	"os"
	"log"
	"strings"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisContext struct{
	pool *redis.Pool
}

func(ctx RedisContext) WaitForRedis() {
	conn := ctx.Get()
	defer conn.Close()
	for true {
		_, err := conn.Do("PING")
		if err == nil {
			break
		}
		log.Println("Can't connect to the Redis database, Retrying...")
	}
}

func (ctx RedisContext) Get () redis.Conn{
	return ctx.pool.Get()
}

func (ctx RedisContext) GeoAdd (latitude string, longitude string, value string) {
	conn := ctx.Get()
	defer conn.Close()
	_, err := conn.Do("GEOADD", "idx:cities", longitude, latitude, value)
	if err != nil{
		log.Fatal(err)
	}
}

func (ctx RedisContext) GeoRadius (latitude string, longitude string, countryCode string) string{
	conn := ctx.Get()
	defer conn.Close()
	log.Printf("latitude %s, longitude %s", latitude, longitude)
	values, err := redis.Strings(conn.Do("GEORADIUS", "idx:cities", longitude, latitude, 1000, "km", "ASC"))
	if err != nil{
		log.Fatal(err)
	}
	for _, record := range values {
		if strings.HasSuffix(record, fmt.Sprintf(":%s", strings.ToUpper(countryCode))) {
			return record
		}
	}
	return ""
}

func makeRedisPool() *RedisContext{
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
	return &RedisContext{pool}
}