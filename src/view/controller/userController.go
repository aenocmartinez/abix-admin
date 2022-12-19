package controller

import (
	"abix360/src/usecase"
	"abix360/src/view/dto"
	formrequest "abix360/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllUsers(c *gin.Context) {
	useCase := usecase.ListUsersUseCase{}
	users := useCase.Execute(c)

	c.JSON(http.StatusOK, gin.H{"result": users})
}

func CreateUser(c *gin.Context) {
	var req formrequest.CreateUserFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	useCase := usecase.CreateUserUseCase{}
	result := useCase.Execute(dto.CreateUserDto{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func ActiveUser(c *gin.Context) {
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

	useCase := usecase.ActiveUserUseCase{}
	result := useCase.Execute(c, int64(id))

	c.JSON(http.StatusOK, gin.H{"result": result})
}