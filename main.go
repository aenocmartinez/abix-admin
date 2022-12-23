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
	r.GET("/user/active", controller.ActiveUser)
	r.GET("/user", controller.ViewUser)
	r.PUT("/reset-password", controller.ResetPassword)

	// Collections
	r.GET("/collections", controller.AllCollections)
	r.POST("/collection", controller.CreateCollection)
	r.GET("/collection", controller.ViewCollection)
	r.PUT("/collection", controller.UpdateCollection)
	r.DELETE("/collection/:id", controller.DeleteCollection)

	// Fields-Collection
	r.POST("/collection/add-field", controller.AddFieldToCollection)
	r.DELETE("/collection/remove-field", controller.RemoveFieldToCollection)
	r.GET("/collection/fields", controller.AllFieldsOfCollections)
	r.GET("/collection/field", controller.ViewFieldCollection)

	// Fields
	r.POST("/field", controller.CreateField)
	r.GET("/fields", controller.AllFields)
	r.GET("/field", controller.ViewField)
	r.DELETE("/field/:id", controller.DeleteField)
	r.PUT("/field", controller.UpdateField)
	r.GET("/search-fields", controller.SearchFields)

	// Sequences
	r.GET("/sequences", controller.AllSequences)
	r.POST("/sequence", controller.CreateSequence)
	r.GET("/sequence", controller.ViewSequence)
	r.PUT("/sequence", controller.UpdateSequence)
	r.DELETE("/sequence/:id", controller.DeleteSequence)
	r.GET("/search-sequence", controller.SearchSequences)

	// List
	r.GET("/lists", controller.AllList)
	r.POST("/list", controller.CreateList)
	r.GET("/list", controller.ViewList)
	r.PUT("/list", controller.UpdateList)
	r.DELETE("/list/:id", controller.DeleteList)
	r.GET("/search-list", controller.SearchList)

	r.Run(":8082")
}
