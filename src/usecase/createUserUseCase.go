package usecase

import (
	"abix360/shared"
	"abix360/src/view/dto"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CreateUserUseCase struct{}

func (useCase *CreateUserUseCase) Execute(createUser dto.CreateUserDto) interface{} {

	url := shared.Config().Appauth.Server + "/" + shared.Config().Appauth.CreateUser
	payload := strings.NewReader("{\"name\":\"" + createUser.Name + "\", \"email\":\"" + createUser.Email + "\", \"password\":\"" + createUser.Password + "\"}")

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return shared.ResponseJSON(body)
}
