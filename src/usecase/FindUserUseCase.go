package usecase

import (
	"abix360/shared"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FindUserUseCase struct{}

func (useCase *FindUserUseCase) Execute(c *gin.Context, id int64) domain.User {
	token := shared.GetTokenRequest(c)
	url := shared.Config().Appauth.Server + "/" + shared.Config().Appauth.FindUser + "/" + strconv.Itoa(int(id))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return domain.User{}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return domain.User{}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return domain.User{}
	}

	strBody := string(body)
	strBody = strings.Replace(strBody, "{\"data\":", "", -1)
	strBody = strings.Replace(strBody, "}}", "}", -1)

	var userDto dto.UserDto
	err = json.Unmarshal([]byte(strBody), &userDto)
	if err != nil {
		return domain.User{}
	}

	user := domain.NewUser()
	user.WithId(userDto.Id)
	user.WithName(userDto.Name)
	user.WithEmail(userDto.Email)
	user.WithState(userDto.State)
	user.WithCreatedAt(userDto.CreatedAt)

	return *user
}
