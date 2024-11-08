package Configs

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

var Env = koanf.New(".")

func LoadConfigs() {
	loadDatabases()
}

func loadDatabases() {
	// Load Yaml config.
	if err := Env.Load(file.Provider("Common/Configs/db.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// Check Database variables
	if !Env.Exists("databases") {
		log.Fatalln("databases key not exists in db.yaml config file")
	}
	if !Env.Exists("default_database") {
		log.Fatalln("default_database key not exists in db.yaml config file")
	}

}
