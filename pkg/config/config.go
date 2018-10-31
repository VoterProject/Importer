package config

import (
	"encoding/json"
	"io/ioutil"
)

type DBConfig struct {
	DSN string `json:"db"`
	PA  string `json:"pa"`
	WA  string `json:"wa"`
}

func ConfigFromFile(filename string) DBConfig {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	config := DBConfig{}
	jsonErr := json.Unmarshal(data, &config)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return config
}
