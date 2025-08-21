package main

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// generateMarkdown converts an OpenAPI specification to markdown
// Sample output:
// # API Title
// API description
//
// ## Paths
// | Path | Operations |
// ...
//
// ## Definitions
// ### User
// ...
func generateMarkdown(swagger *openapi3.T) string {
	var sb strings.Builder

	sb.WriteString(generateInfoSection(swagger.Info))
	sb.WriteString(generatePathsTable(swagger.Paths))
	sb.WriteString(generatePathsDocumentation(swagger.Paths))
	sb.WriteString(generateDefinitionsSection(swagger.Extensions))

	return sb.String()
}

// generateInfoSection generates the info section with title, description, and contact
// Sample output:
// # My API
// This is a sample API
//
// Support Team
// https://example.com
// support@example.com
func generateInfoSection(info *openapi3.Info) string {
	var sb strings.Builder

	if info != nil {
		sb.WriteString("# " + info.Title)
		sb.WriteString(newlineDouble)
		sb.WriteString(info.Description)
		sb.WriteString(newlineDouble)

		if info.Contact != nil {
			sb.WriteString(info.Contact.Name)
			sb.WriteString(newlineDouble)
			sb.WriteString(info.Contact.URL)
			sb.WriteString(newlineDouble)
			sb.WriteString(info.Contact.Email)
			sb.WriteString(newlineDouble)
		}
	}

	return sb.String()
}

// generateDefinitionsSection generates the definitions section from extensions
// Sample output:
// ## Definitions
//
// ### <span id="/definitions/User">User</span>
// ...
func generateDefinitionsSection(extensions map[string]any) string {
	var sb strings.Builder

	sb.WriteString(definitionsSection)
	sb.WriteString(newlineDouble)
	definitions, ok := extensions["definitions"]

	if ok {
		definitionsMap, ok := definitions.(map[string]any)
		if ok {
			sb.WriteString(markDownDefinitions(definitionsMap))
		}
	}

	return sb.String()
}

// oneleline removes newlines and trims whitespace from a string
// Sample: "Line 1\nLine 2" -> "Line 1 Line 2"
func oneleline(input string) string {
	return strings.TrimSpace(strings.ReplaceAll(input, "\n", " "))
}