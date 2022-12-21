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
	r.GET("/user/active/:id", controller.ActiveUser)
	r.GET("/user/:id", controller.ViewUser)
	r.PUT("/reset-password", controller.ResetPassword)

	// Collections
	r.GET("/collections", controller.AllCollections)
	r.POST("/collection", controller.CreateCollection)
	r.GET("/collection/:id", controller.ViewCollection)
	r.PUT("/collection", controller.UpdateCollection)
	r.DELETE("/collection/:id", controller.DeleteCollection)

	// Fields-Collection
	r.POST("/collection/add-field", controller.AddFieldToCollection)
	r.DELETE("/collection/remove-field/:idCollection/:idField", controller.RemoveFieldToCollection)
	r.GET("/collection/fields/:idCollection", controller.AllFieldsOfCollections)

	// Fields
	r.POST("/field", controller.CreateField)
	r.GET("/fields", controller.AllFields)
	r.GET("/field/:id", controller.ViewField)
	r.DELETE("/field/:id", controller.DeleteField)
	r.PUT("/field", controller.UpdateField)
	r.GET("/search-fields", controller.SearchFields)

	r.Run(":8082")
}
