package main

import (
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

// loadOpenAPISpec loads and parses an OpenAPI specification from a file.
// It validates the file exists, reads its contents, and parses it using the kin-openapi library.
// Returns the parsed OpenAPI spec or an error if the file is not found or invalid.
func loadOpenAPISpec(filename string) (*openapi3.T, error) {
	// Check if the input file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	// Read the swagger file
	swaggerData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Parse the swagger file
	swagger, err := openapi3.NewLoader().LoadFromData(swaggerData)
	if err != nil {
		return nil, err
	}

	return swagger, nil
}
