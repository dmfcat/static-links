package main

import (
	"os"

	"go.yaml.in/yaml/v4"
)

// Demarshal yaml into struct
func config_parse(raw []byte) Config {
	var config Config
	err := yaml.Unmarshal(raw, &config)

	if err != nil {
		os.Exit(1)
	}

	return config
}
