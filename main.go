// Package main provides a CLI tool to convert Swagger/OpenAPI specifications to Markdown documentation.
// The tool reads YAML or JSON OpenAPI files and generates structured Markdown with
// API overview, path documentation, and schema definitions.
package main

import (
	"flag"
	"fmt"
	"os"
)

// main is the CLI entry point
func main() {
	input := flag.String("i", "", "input swagger file")
	output := flag.String("o", "", "output markdown file")
	flag.Parse()
	
	if *input == "" || *output == "" {
		flag.Usage()
		return
	}

	// Load and parse the OpenAPI specification
	swagger, err := loadOpenAPISpec(*input)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Input file not found: %s\n", *input)
		} else {
			fmt.Printf("Error parsing Swagger file: %v\n", err)
		}
		return
	}

	// Generate markdown
	markdown := generateMarkdown(swagger)

	// Save the Markdown to a file
	err = os.WriteFile(*output, []byte(markdown), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Printf("Markdown file generated: %s\n", *output)
}