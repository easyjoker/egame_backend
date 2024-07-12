package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml:"port"`
}

func LoadConfig() (Config, error) {
	var config Config
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}
