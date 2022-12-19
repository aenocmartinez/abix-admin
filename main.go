package main

import (
	abixauth "abix360/src/infraestructure/abix_auth"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": abixauth.ValidateToken(c)})
	})

	r.Run(":8082")
}
