package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCollection(c *gin.Context) {
	var req formrequest.CreateCollectionFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.CreateCollectionUseCase{}
	code, err := useCase.Execute(req.Name)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "la colección se ha creado con éxito"})
}

func UpdateCollection(c *gin.Context) {
	var req formrequest.UpdateCollectionFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.UpdateCollectionUseCase{}
	code, err := useCase.Execute(dto.CollectionDto{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "la colección se ha creado con éxito"})
}

func ViewCollection(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.ViewCollectionUseCase{}
	dtoCollection, err := useCase.Execute(int64(id))
	if err != nil {
		c.AbortWithStatusJSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": dtoCollection})
}

func DeleteCollection(c *gin.Context) {
	var strId string = c.Param("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.DeleteCollectionUseCase{}
	code, err := useCase.Execute(int64(id))
	if err != nil {
		c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": "la colección ha sido eliminada con éxito"})
}

func AllCollections(c *gin.Context) {
	useCase := usecase.ListCollectionsUseCase{}
	collections := useCase.Execute()
	c.JSON(200, gin.H{"result": collections})
}
