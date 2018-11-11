package catcher

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port     int
	Host     string
	RootHost string `json:"root_host"`
}

func LoadConfiguration(filename string) (*Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)

	return &config, err
}
