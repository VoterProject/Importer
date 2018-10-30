package sql

import (
	"encoding/json"
	"io/ioutil"
)

type DBConfig struct {
	DSN string `json:"db"`
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
