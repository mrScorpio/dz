package validationApi

import (
	"encoding/json"
	"log"
	"os"
)

type VldConfig struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func LoadConfigFromFile() *VldConfig {
	jFile, err := os.ReadFile("cfg")
	if err != nil {
		log.Println(err.Error())
	}

	var confItem VldConfig
	err = json.Unmarshal(jFile, &confItem)
	if err != nil {
		log.Println(err.Error())
	}
	return &confItem
}
