package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database struct {
		Source string
	}
	Auth struct {
		SecretSource string
	}
}

func GetConfig(configSource string) (*Config, error) {
	configFile, err := os.Open(configSource)
	defer configFile.Close()

	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)

	config := new(Config)

	err = jsonParser.Decode(config)

	return config, err
}
