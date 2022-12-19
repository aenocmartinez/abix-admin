package shared

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type configApp struct {
	Jwt struct {
		Secret string `yaml:"secret"`
	}
	Appauth struct {
		EndPointValidateToken string `yaml:"endPointValidateToken"`
	}
}

func Config() configApp {
	content, err := os.ReadFile(GetRootPath() + "/app/cfg-app.yml")
	if err != nil {
		log.Fatal("abix-admin / shared / GetKeySecret() / os.ReadFile: ", err)
	}

	var config configApp
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("abix-admin / shared / GetKeySecret() / yaml.Unmarshal: ", err)
	}

	return config
}
