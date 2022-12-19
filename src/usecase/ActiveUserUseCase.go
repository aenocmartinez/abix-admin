package usecase

import (
	"abix360/shared"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActiveUserUseCase struct{}

func (useCase *ActiveUserUseCase) Execute(c *gin.Context, id int64) interface{} {
	findUserUseCase := FindUserUseCase{}
	user := findUserUseCase.Execute(c, id)
	if !user.Exists() {
		return shared.ResponseJSON([]byte(`{"error":"el usuario no existe"}`))
	}

	token := shared.GetTokenRequest(c)
	endPoint := shared.Config().Appauth.Unsuscribe
	if !user.State() {
		endPoint = shared.Config().Appauth.ActivateUser
	}

	url := shared.Config().Appauth.Server + "/" + endPoint + "/" + strconv.Itoa(int(id))
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

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
