package shared

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) bool {
	token := GetTokenRequest(c)
	url := Config().Appauth.Server + "/" + Config().Appauth.ValidateToken
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return string(body) == "{\"isValid\":true}"
}

func GetTokenRequest(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		return ""
	}
	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
	return tokenString
}
