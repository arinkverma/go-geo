package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Starting geo micro-service")
	redisCtx := makeRedisPool()
	dataCtx := DataContext{
		redisCtx: redisCtx,
	}
	dataCtx.InitData()

	serviceCtx := &ServiceContext{
		redisCtx: redisCtx,
	}

	r := gin.Default()
	r.GET("/ping", serviceCtx.Ping)
	r.GET("/resolve/:countryCode/:latlon", serviceCtx.ResolveLatLon)
	r.Run() // listen and serve on 0.0.0.0:8080
}

