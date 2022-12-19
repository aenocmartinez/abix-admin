package database

import (
	"abix360/shared"
	"database/sql"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type IConnectionDB interface {
	Close()
	Conn() *sql.DB
}

type ConnectDB struct {
	source IConnectionDB
}

var lock = &sync.Mutex{}
var instance *ConnectDB

func getConfig() *configYAML {
	content, err := os.ReadFile(shared.GetRootPath() + "/database/cfg-database.yml")
	if err != nil {
		log.Fatal("dao / datasource / Config / os.ReadFile: ", err)
	}

	var configFile configYAML
	err = yaml.Unmarshal(content, &configFile)
	if err != nil {
		log.Fatal("dao / datasource / Config / yaml.Unmarshal: ", err)
	}
	return &configFile
}

var instanceSetting *settings

func DataSource() *settings {
	if instanceSetting == nil {
		if instanceSetting == nil {
			instanceSetting = &settings{
				config: *getConfig(),
			}
		}
	}
	return instanceSetting
}

func Instance() *ConnectDB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &ConnectDB{
				source: getSource(),
			}
		}
	}
	return instance
}

func (c *ConnectDB) Source() IConnectionDB {
	return c.source
}
