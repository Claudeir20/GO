package router

import "github.com/gin-gonic/gin"

func Initializie() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // inicia o servidor na porta 8080
}
