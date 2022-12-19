package shared

import (
	"log"
	"os"
	"regexp"
)

var projectDirName string = "abix-admin"

func GetRootPath() string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, err := os.Getwd()
	if err != nil {
		log.Println("Shared / GetRootPath() / os.Getwd(): ", err)
	}
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath)
}
