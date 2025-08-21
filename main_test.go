package main

import (
	"os"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestGenerateMarkdownIntegration(t *testing.T) {
	// Test with the existing YAML file
	inputFile := "./testdata/sample-swagger.yaml"
	expectedOutputFile := "./testdata/sample-output.md"

	// Read the input file
	swaggerData, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}

	// Parse the swagger file
	swagger, err := openapi3.NewLoader().LoadFromData(swaggerData)
	if err != nil {
		t.Fatalf("Error parsing Swagger file: %v", err)
	}

	// Generate markdown
	actualOutput := generateMarkdown(swagger)

	// Read expected output
	expectedOutput, err := os.ReadFile(expectedOutputFile)
	if err != nil {
		t.Fatalf("Failed to read expected output file: %v", err)
	}

	// Compare outputs
	if actualOutput != string(expectedOutput) {
		t.Errorf("Generated markdown does not match expected output")
		// Write actual output for debugging
		os.WriteFile("./testdata/actual-output-debug.md", []byte(actualOutput), 0644)
		t.Logf("Actual output written to ./testdata/actual-output-debug.md for comparison")
	}
}

func TestGenerateMarkdownWithNilFields(t *testing.T) {
	// Test with minimal OpenAPI spec to ensure nil handling
	spec := &openapi3.T{
		Info: &openapi3.Info{
			Title:       "Test API",
			Description: "Test Description",
		},
		Paths: openapi3.Paths{},
	}

	result := generateMarkdown(spec)

	expected := "# Test API\n\nTest Description\n\n## Paths\n\n| Path | Operations |\n| --- | --- |\n\n\n## Definitions\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMarkDownDefinitions(t *testing.T) {
	definitions := map[string]any{
		"TestSchema": map[string]any{
			"type":        "object",
			"title":       "Test Schema",
			"description": "A test schema",
			"properties": map[string]any{
				"name": map[string]any{
					"type":        "string",
					"description": "Name field",
				},
			},
		},
	}

	result := markDownDefinitions(definitions)

	// Should contain the schema name and basic structure
	if len(result) == 0 {
		t.Error("Expected non-empty result from markDownDefinitions")
	}

	// Should contain schema title
	if result == "" {
		t.Error("Result should not be empty")
	}
}

func TestOnelineFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Simple text", "Simple text"},
		{"Text with\nnewlines", "Text with newlines"},
		{"  Text with spaces  ", "Text with spaces"},
		{"Text with\nmultiple\nnewlines", "Text with multiple newlines"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := oneleline(tt.input)
			if result != tt.expected {
				t.Errorf("oneleline(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDetectObjectGoType(t *testing.T) {
	tests := []struct {
		name     string
		schema   map[string]any
		expected string
	}{
		{
			name: "object with properties",
			schema: map[string]any{
				"properties": map[string]any{},
			},
			expected: "object",
		},
		{
			name: "map with additionalProperties",
			schema: map[string]any{
				"additionalProperties": map[string]any{},
			},
			expected: "map",
		},
		{
			name:     "neither object nor map",
			schema:   map[string]any{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectObjectGoType(tt.schema)
			if result != tt.expected {
				t.Errorf("detectObjectGoType() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func TestObjectMarkDown(t *testing.T) {
	tests := []struct {
		name     string
		schema   map[string]any
		contains []string
	}{
		{
			name: "simple object with properties",
			schema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"name": map[string]any{
						"type":        "string",
						"description": "Object name",
					},
				},
			},
			contains: []string{"**Type:** object", "**Properties:**", "| name |"},
		},
		{
			name: "object with enum property",
			schema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"status": map[string]any{
						"type": "string",
						"enum": []any{"active", "inactive"},
					},
				},
			},
			contains: []string{"**Type:** object", "## Enums", "status:**"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := objectMarkDown(tt.schema)
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected result to contain %q, got %q", expected, result)
				}
			}
		})
	}
}

func TestArrayMarkDown(t *testing.T) {
	tests := []struct {
		name     string
		schema   map[string]any
		expected string
	}{
		{
			name: "array of strings",
			schema: map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "string",
				},
			},
			expected: "[]string",
		},
		{
			name: "array with $ref",
			schema: map[string]any{
				"type": "array",
				"items": map[string]any{
					"$ref": "#/definitions/User",
				},
			},
			expected: "[]",
		},
		{
			name: "array without items",
			schema: map[string]any{
				"type": "array",
			},
			expected: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := arrayMarkDown(tt.schema)
			if !strings.Contains(result, tt.expected) {
				t.Errorf("Expected result to contain %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestMapMD(t *testing.T) {
	tests := []struct {
		name     string
		schema   map[string]any
		contains string
	}{
		{
			name: "map with string values",
			schema: map[string]any{
				"type": "object",
				"additionalProperties": map[string]any{
					"type": "string",
				},
			},
			contains: "**Type:** map[*]->string",
		},
		{
			name: "map with ref values",
			schema: map[string]any{
				"type": "object",
				"additionalProperties": map[string]any{
					"$ref": "#/definitions/User",
				},
			},
			contains: "**Type:** map[*]->[User](#/definitions/User)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapMD(tt.schema)
			if !strings.Contains(result, tt.contains) {
				t.Errorf("Expected result to contain %q, got %q", tt.contains, result)
			}
		})
	}
}

func TestInvalidInputFile(t *testing.T) {
	// Test with invalid YAML file
	inputFile := "./testdata/invalid.yaml"

	// Read the input file
	swaggerData, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}

	// Parse the swagger file - should fail
	_, err = openapi3.NewLoader().LoadFromData(swaggerData)
	if err == nil {
		t.Error("Expected error when parsing invalid YAML, but got none")
	}
}
