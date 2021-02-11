package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

func main() {
	r := gin.Default()
	r.GET("/resolve/:lat/:lon", ResolveLatLon)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ResolveLatLon(c *gin.Context) {
	lat := c.Param("lat")
	lon := c.Param("lon")
	playUrl := fmt.Sprintf("%s%s", lat, lon)
	fmt.Println(playUrl)
	return "hi"
}
