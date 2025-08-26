package controller

import "github.com/gin-gonic/gin"

func InitProbeRoute(router *gin.RouterGroup) {
	router.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
