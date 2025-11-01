package validationApi

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func LoadConfig() *Config {
	jFile, err := os.ReadFile("cfg")
	if err != nil {
		log.Println(err.Error())
	}

	var confItem Config
	err = json.Unmarshal(jFile, &confItem)
	if err != nil {
		log.Println(err.Error())
	}
	return &confItem
}
