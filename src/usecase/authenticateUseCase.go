package usecase

import "github.com/gin-gonic/gin"

type AuthenticateUseCase struct{}

func (useCase *AuthenticateUseCase) Execute(c *gin.Context) bool {

	return true
}
