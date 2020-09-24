package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Server *ServerConfig `json:"server"`
	Zeebe  *ZeebeConfig  `json:"zeebe"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type ZeebeConfig struct {
	Host string `json:"host"`
}

func GetConfig() (*Config, error) {
	jsonFile, err := os.Open("configs/config.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	err = json.Unmarshal(byteValue, &config)
	return &config, err
}
