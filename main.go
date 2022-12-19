package main

import (
	"abix360/shared"
	"abix360/src/view/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func validateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "header no valid"})
		}
		c.Next()
	}
}

func validateAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !shared.ValidateToken(c) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Forbidden access"})
		}
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), validateHeader(), validateAuthenticate())

	r.GET("/users", controller.AllUsers)
	r.POST("/user", controller.CreateUser)

	r.Run(":8082")
}
