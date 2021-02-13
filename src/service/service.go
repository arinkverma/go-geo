package service

import(
	"fmt"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"go-geo/db"
)

type ServiceContext struct{
	redisCtx *db.RedisContext
}

func (self ServiceContext) Ping(c *gin.Context) {
	value, err := self.redisCtx.Ping()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": value,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprint(err),
		})
	}
}

func (self ServiceContext) ResolveLatLon(c *gin.Context) {
	latlon := c.Param("latlon")
	countryCode := strings.ToUpper(c.Param("countryCode"))
	cordinates := strings.Split(latlon, ",")
	value, err := self.redisCtx.GeoRadius(countryCode, cordinates[0], cordinates[1])
	if err == nil {
		data := strings.Split(value, ":")
		c.JSON(http.StatusOK, gin.H{
			"id": data[0],
			"city": data[1],
			"country": data[2],
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprint(err),
		})
	}
}

func MakeServiceContext(redisCtx *db.RedisContext) *ServiceContext{
	return &ServiceContext{
		redisCtx: redisCtx,
	}
}