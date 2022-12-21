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

func AddFieldToCollection(c *gin.Context) {
	var req formrequest.AddFieldToCollectionFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.AddFieldToCollectionUseCase{}
	code, err := useCase.Execute(dto.FieldCollectionDto{
		IdCollection: req.IdCollection,
		IdField:      req.IdField,
		Unique:       req.Unique,
		Editable:     req.Editable,
		Required:     req.Required,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(code, gin.H{"message": "se ha agregado el campo a la coleccion"})
}

func RemoveFieldToCollection(c *gin.Context) {
	var strIdCollection string = c.Param("idCollection")
	var strIdField string = c.Param("idField")
	if len(strIdCollection) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro idCollection no válido"})
		return
	}

	idCollection, err := strconv.Atoi(strIdCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro idCollection no válido"})
		return
	}

	if len(strIdField) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro iddField no válido"})
		return
	}

	iddField, err := strconv.Atoi(strIdField)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro iddField no válido"})
		return
	}

	useCase := usecase.RemoveFieldToCollectionUseCase{}
	code, err := useCase.Execute(int64(idCollection), int64(iddField))
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(code, gin.H{"message": "se ha quitado el campo de la coleccion"})
}

func AllFieldsOfCollections(c *gin.Context) {
	var strIdCollection string = c.Param("idCollection")
	if len(strIdCollection) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	idCollection, err := strconv.Atoi(strIdCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}
	useCase := usecase.ListFieldsOfCollectionUseCase{}
	fields, err := useCase.Execute(int64(idCollection))
	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": fields})
}
