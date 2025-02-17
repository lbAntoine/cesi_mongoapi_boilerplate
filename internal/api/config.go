package api

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	MongoDB struct {
		URI      string `yaml:"uri"`
		Database string `yaml:"database"`
		Timeout  int    `yaml:"timeout"`
	} `yaml:"mongodb"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
