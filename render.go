package main

import (
	"html/template"
	"log"
	"os"
)

func render_page(config Config) {
	// Parse template
	tmpl, err := template.ParseFiles("templates/template.html")

	if err != nil {
		log.Fatal(err)
	}

	// Create output file
	file, err := os.Create("public/index.html")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Render output file
	if err := tmpl.Execute(file, config); err != nil {
		log.Fatal(err)
	}
}
