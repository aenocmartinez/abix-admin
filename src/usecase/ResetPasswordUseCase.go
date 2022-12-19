package usecase

import (
	"abix360/shared"
	"abix360/src/view/dto"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResetPasswordUseCase struct{}

func (useCase *ResetPasswordUseCase) Execute(c *gin.Context, userDto dto.UserDto) interface{} {
	findUserUseCase := FindUserUseCase{}
	user := findUserUseCase.Execute(c, userDto.Id)
	if !user.Exists() {
		return `{"error":"el usuario no existe"}`
	}

	token := shared.GetTokenRequest(c)
	url := shared.Config().Appauth.Server + "/" + shared.Config().Appauth.ResetPassword
	payload := strings.NewReader("{\"id\":" + strconv.Itoa(int(userDto.Id)) + ", \"password\":\"" + userDto.Password + "\" }")

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return shared.ResponseJSON(body)
}
