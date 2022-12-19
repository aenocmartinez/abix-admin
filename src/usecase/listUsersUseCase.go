package usecase

import (
	"abix360/shared"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListUsersUseCase struct{}

func (useCase *ListUsersUseCase) Execute(c *gin.Context) interface{} {
	token := shared.GetTokenRequest(c)
	url := shared.Config().Appauth.Server + "/" + shared.Config().Appauth.AllUsers

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

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
