package main

import (
	"fmt"
	"log"
	"os"
)

// TODO: Create HTML template with Go snippets https://pkg.go.dev/html/template
// TODO: Insert extracted YAML into HTML template
// TODO: Export generated .html file
// TODO: Add Stylesheets
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
		log.Fatal(err)
	}

	// Render template
	render_page(config_parse(raw))

}
