package main

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// generateMarkdown converts an OpenAPI specification to markdown
func generateMarkdown(swagger *openapi3.T) string {
	var sb strings.Builder

	sb.WriteString(generateInfoSection(swagger.Info))
	sb.WriteString(generatePathsTable(swagger.Paths))
	sb.WriteString(generatePathsDocumentation(swagger.Paths))
	sb.WriteString(generateDefinitionsSection(swagger.Extensions))

	return sb.String()
}

// generateInfoSection generates the info section with title, description, and contact
func generateInfoSection(info *openapi3.Info) string {
	var sb strings.Builder

	if info != nil {
		sb.WriteString("# " + info.Title + "\n\n")
		sb.WriteString(info.Description + "\n\n")

		if info.Contact != nil {
			sb.WriteString(info.Contact.Name + "\n\n")
			sb.WriteString(info.Contact.URL + "\n\n")
			sb.WriteString(info.Contact.Email + "\n\n")
		}
	}

	return sb.String()
}

// generateDefinitionsSection generates the definitions section from extensions
func generateDefinitionsSection(extensions map[string]any) string {
	var sb strings.Builder

	sb.WriteString(definitionsSection + newlineDouble)
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
func oneleline(input string) string {
	return strings.TrimSpace(strings.ReplaceAll(input, "\n", " "))
}