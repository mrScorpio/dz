package validationApi

import (
	"encoding/json"
	"os"
)

type HashData struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func (d *HashData) SaveJson() error {
	file, err := os.OpenFile("hashdata.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(d)
	return nil
}

func (d *HashData) ReadJson() error {
	file, err := os.OpenFile("hashdata.json", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	json.NewDecoder(file).Decode(d)
	file.Close()
	if os.Remove("hashdata.json") != nil {
		return err
	}
	return nil
}
