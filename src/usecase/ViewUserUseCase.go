package usecase

import (
	"abix360/src/view/dto"
	"errors"

	"github.com/gin-gonic/gin"
)

type ViewUserUseCase struct{}

func (useCase *ViewUserUseCase) Execute(c *gin.Context, id int64) (dto.UserDto, error) {
	findUserUseCase := FindUserUseCase{}
	user := findUserUseCase.Execute(c, id)
	if !user.Exists() {
		return dto.UserDto{}, errors.New("el usuario no existe")
	}

	userDto := dto.UserDto{
		Id:        user.Id(),
		Name:      user.Name(),
		Email:     user.Email(),
		State:     user.State(),
		CreatedAt: user.CreatedAt(),
	}

	return userDto, nil
}
