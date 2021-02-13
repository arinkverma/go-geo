package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"go-geo/db"
	"go-geo/service"
)


const DataFile = "cities15000.txt"

func main() {
	log.Println("Starting geo micro-service")
	redisCtx := db.MakeRedisContext()
	dataCtx := db.MakeDataContext(redisCtx, DataFile)
	serviceCtx := service.MakeServiceContext(redisCtx)

	dataCtx.InitData()
	r := gin.Default()
	r.GET("/ping", serviceCtx.Ping)
	r.GET("/resolve/:countryCode/:latlon", serviceCtx.ResolveLatLon)
	r.Run() // listen and serve on 0.0.0.0:8080
}
