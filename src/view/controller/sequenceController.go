package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSequence(c *gin.Context) {
	var req formrequest.CreateSequenceFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	useCase := usecase.CreateSequenceUseCase{}
	code, err := useCase.Execute(dto.SequenceDto{
		Name:   req.Name,
		Prefix: req.Prefix,
		Value:  req.Value,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(code, gin.H{"message": "La secuencia se ha creado de forma exitosa"})
}

func UpdateSequence(c *gin.Context) {
	var req formrequest.UpdateSequenceFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	useCase := usecase.UpdateSequenceUseCase{}
	code, err := useCase.Execute(dto.SequenceDto{
		Id:     req.Id,
		Name:   req.Name,
		Prefix: req.Prefix,
		Value:  req.Value,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(code, gin.H{"message": "La secuencia se ha actualizado de forma exitosa"})
}

func DeleteSequence(c *gin.Context) {
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

	useCase := usecase.DeleteSequenceUseCase{}
	code, err := useCase.Execute(int64(id))

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	c.JSON(code, gin.H{"message": "La secuencia se ha eliminado de forma exitosa"})
}

func ViewSequence(c *gin.Context) {
	// var strId string = c.Param("id")
	var strId string = c.Query("id")
	if len(strId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.ViewSequenceUseCase{}
	sequence, err := useCase.Execute(int64(id))

	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": sequence})
}

func AllSequences(c *gin.Context) {
	useCase := usecase.ListSequencesUseCase{}
	sequences, err := useCase.Execute()

	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": sequences})
}

func SearchSequences(c *gin.Context) {
	// var name string = c.Param("name")
	var name string = c.Query("name")
	useCase := usecase.SearchSequenceUseCase{}
	sequences, err := useCase.Execute(name)

	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": sequences})
}
