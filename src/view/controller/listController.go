package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var req formrequest.CreateListFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.CreateListUseCase{}
	code, err := useCase.Execute(dto.ListDto{
		Name:   req.Name,
		Values: req.Values,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "success"})
}

func UpdateList(c *gin.Context) {
	var req formrequest.UpdateListFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecase.UpdateListUseCase{}
	code, err := useCase.Execute(dto.ListDto{
		Id:     req.Id,
		Name:   req.Name,
		Values: req.Values,
	})

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "success"})
}

func DeleteList(c *gin.Context) {
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

	useCase := usecase.DeleteListUseCase{}
	code, err := useCase.Execute(int64(id))

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(code, gin.H{"message": "success"})
}

func ViewList(c *gin.Context) {
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

	useCase := usecase.ViewListUseCase{}
	list, err := useCase.Execute(int64(id))

	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": list})
}

func SearchList(c *gin.Context) {
	// var name string = c.Param("name")
	var name string = c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro no válido"})
		return
	}

	useCase := usecase.SearchListUseCase{}
	lists, err := useCase.Execute(name)

	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": lists})
}

func AllList(c *gin.Context) {
	useCase := usecase.AllListUseCase{}
	lists, err := useCase.Execute()
	if err != nil {
		c.JSON(202, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"result": lists})
}
