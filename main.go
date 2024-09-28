package main

import (
	"fmt"
	"tutorial/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LocalConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	


	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT)) // listen and serve on 0.0.0.0:8080
}