package main

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

// Generate a new default `config.yaml`
// TODO: Could add path to where file is created
func config_gen() {
	fmt.Println("Generating `config.yaml` file")
	fmt.Println("See documentation for more info")
	f, _ := os.Create("config.yaml")

	data := get_data()

	f.Write(data)
	fmt.Println(string(data))

	defer f.Close()
}

// Add default data
func get_data() []uint8 {
	config := Config{
		Title:    "My Title",
		Subtitle: "My Subtitle",
		Colours: Colours{
			Title:          "#ffffff",
			Subtitle:       "#d6d6d6",
			Background:     "#1d1d1b",
			LinkBackground: "#404040",
			Icon:           "#d6d6d6",
		},
		Links: []Links{
			{Title: "My Link Title 1", Subtitle: "My Link Subtitle 1", URL: "https://example.com"},
			{Title: "My Link Title 2", Subtitle: "My Link Subtitle 2", URL: "https://example.com"},
		},
	}

	data, _ := yaml.Marshal(&config)
	return data
}

// Demarshal yaml into struct
func config_parse(raw []byte) Config {
	var config Config
	err := yaml.Unmarshal(raw, &config)

	if err != nil {
		os.Exit(1)
	}

	return config
}

// TODO: Move this somewhere else? lol
// Structs for `config.yaml`
type Config struct {
	Title    string  `yaml:"title"`
	Subtitle string  `yaml:"subtitle"`
	Links    []Links `yaml:"links"`
	Colours  Colours `yaml:"colours"`
}

type Colours struct {
	Title          string `yaml:"title-colour"`
	Subtitle       string `yaml:"subtitle-colour"`
	Background     string `yaml:"background"`
	LinkBackground string `yaml:"link-background"`
	Icon           string `yaml:"icon-colour"`
}

type Links struct {
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
	URL      string `yaml:"url"`
}
