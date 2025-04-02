package config

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Bot struct {
		Token             string `yaml:"token"`
		Channel_ID        string `yaml:"monitor_channel"`
		Button_Channel_ID string `yaml: "button_channel"`
	} `yaml:"bot"`
	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if path == "" {
		return nil, err
	}

	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
