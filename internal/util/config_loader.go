package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

type DataBaseConfig struct {
	Location string
}

func LoadConfig() (*DataBaseConfig, error) {
	filename := "./config/db.yaml"
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := DataBaseConfig{Location: filename}
	err = yaml.Unmarshal(file, &config)
	return &config, err
}
