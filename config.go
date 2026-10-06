package main

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

// Generate a new default `config.yaml`
func config_gen() {
	wd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Generating `config.yaml` in %s\n", wd)
	fmt.Println("See documentation for more info")
	f, _ := os.Create("config.yaml")

	f.Write(get_data())

	defer f.Close()
}

// Add default data
func get_data() []uint8 {
	config := Config{
		Title:    "My Title",
		Subtitle: "My Subtitle",
		Colours: Colours{
			Background:     "#1d1d1b",
			Title:          "#ffffff",
			Subtitle:       "#d6d6d6",
			LinkBackground: "#404040",
			LinkTitle:      "#ffffff",
			Icon:           "#d6d6d6",
			Accent:         "red",
			Border:         "white",
		},
		Links: []Links{
			{Title: "Link 1", URL: "https://example.com"},
			{Title: "Link 2", URL: "https://example.com"},
			{Title: "Link 3", URL: "https://example.com"},
			{Title: "link 4", URL: "https://example.com"},
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
	Colours  Colours `yaml:"colours"`
	Links    []Links `yaml:"links"`
}

type Colours struct {
	Background     string `yaml:"background"`
	Title          string `yaml:"title-colour"`
	Subtitle       string `yaml:"subtitle-colour"`
	LinkBackground string `yaml:"link-background"`
	LinkTitle      string `yaml:"link-title-colour"`
	Icon           string `yaml:"icon-colour"`
	Accent         string `yaml:"accent-colour"`
	Border         string `yaml:"border-colour"`
}

type Links struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
}
