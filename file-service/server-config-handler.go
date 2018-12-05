package main

import (
	"encoding/json"
	"os"
)

//Config holds temporary data in a struct from "server-config.json"
type Config struct {
	ServerDirectory  string `json:"file-directory"`
	ConnectionString string `json:"connection-string"`
}

//LoadConfig loads a config.json file and returns Config / Error
func LoadConfig(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()

	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}
