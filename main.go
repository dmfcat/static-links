package main

import (
	"fmt"
	"os"
)

// TODO: Create HTML template with Go snippets https://pkg.go.dev/html/template
// TODO: Insert extracted YAML into HTML template
// TODO: Export generated .html file
//
// STRETCH
// TODO: Add support for goicons https://pkg.go.dev/github.com/dimmerz92/go-icons#section-readme
// TODO: Change config so it can generate a config with whitespace and comments

func main() {
	fileName := "config.yaml"
	raw, err := os.ReadFile(fileName)

	// Generate a new template if it doesn't exist
	if err != nil {
		fmt.Println("No .yaml file detected")
		config_gen()
		os.Exit(1)
	}

	// Parse YAML
	config := config_parse(raw)

	// Parse existing template
	fmt.Printf("%+v\n", config) // Test print

}
