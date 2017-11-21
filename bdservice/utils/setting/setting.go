package setting

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//PRO
const settingsFile = "appsettings.json"

//PRE
//const settingsFile = "pre-appsettings.json"

//DEV
//const settingsFile = "dev-appsettings.json"

type Setting struct {
	Mongo Mongo
	Log   Log
}

type Mongo struct {
	Hosts string `json:"hosts"`
	Db    string `json:"db"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
}

type Log struct {
	Folder string `json:"folder"`
	File   string `json:"file"`
	Ext    string `json:"ext"`
}

func GetSettings() Setting {
	return parseSettings(settingsFile)
}

func parseSettings(settingsFile string) Setting {
	f, err := ioutil.ReadFile(settingsFile)
	if err != nil {
		log.Fatal("Unable to access settings file: ", err)
	}

	var setting Setting
	err = json.Unmarshal(f, &setting)

	if err != nil {
		log.Fatal(err)
	}

	return setting
}
