package main

import(
	"log"
	"strings"
	"net/http"
		"github.com/gin-gonic/gin"
)

type ServiceContext struct{
	redisCtx *RedisContext
}

func (ctx ServiceContext) Ping(c *gin.Context) {
	// Get a connection
	conn := ctx.redisCtx.Get()
	defer conn.Close()
	// Test the connection
	val, err := conn.Do("PING")
	if err != nil {
	  log.Fatal("Can't connect to the Redis database")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": val,
	})
}

func (ctx ServiceContext) ResolveLatLon(c *gin.Context) {
	latlon := c.Param("latlon")
	cordinates := strings.Split(latlon, ",")
	value := ctx.redisCtx.GeoRadius(cordinates[0], cordinates[1])
	resp := gin.H{}
	if (value != ""){
		data := strings.Split(value, ":")
		resp = gin.H{
			"id": data[0],
			"city": data[1],
			"country": data[2],
		}	
	}
	c.JSON(http.StatusOK, resp)
}
