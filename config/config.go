package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	DBHost         string `json:"database.host"`
	DBUser         string `json:"database.username"`
	DBPassword     string `json:"database.password"`
	DatabaseUrl    string `json:"database.url"`
	DBPort         string `json:"database.port"`
	DBName         string `json:"database.name"`
	DBSchema       string `json:"database.schema"`
	ServerPort     string `json:"server.port"`
	DBMaxIdleConns int    `json:"database.maxIdleConnection"`
	DBMaxOpenConns int    `json:"database.maxOpenConnection"`
	TLSCache       string `json:"tls.certificate.cache"`
}

func LoadConfig(filename string) (*Configuration, error) {

	//read json file using os function from go
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//decode json file
	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, err
}
