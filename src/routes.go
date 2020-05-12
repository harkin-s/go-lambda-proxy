package main

import (
	"github.com/gin-gonic/gin"
)

// GetRoutes defines all the routes and there function
func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", helloHandler)
	return router
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
