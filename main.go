package main

import (
	"embed"
	"log"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var config embed.FS

type myConfig struct {
	DB_URL string `yaml:"DB_URL"`
	DB_ID  string `yaml:"DB_ID"`
	DB_PW  string `yaml:"DB_PW"`
}

func main() {

	var conf myConfig
	yamlFile, err := config.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(conf)
}
