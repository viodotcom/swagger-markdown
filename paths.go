package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// generatePathsTable generates a markdown table of all paths with links
func generatePathsTable(paths openapi3.Paths) string {
	var sb strings.Builder

	sb.WriteString(pathsSection + newlineDouble)
	sb.WriteString(pathsTableHeader)

	// Sort paths by path name
	pathNames := make([]string, 0, len(paths))
	for path := range paths {
		pathNames = append(pathNames, path)
	}
	sort.Strings(pathNames)

	for _, path := range pathNames {
		pathItem := paths[path]
		sb.WriteString("| [" + path + "](" + pathPrefix + strings.ToLower(path) + ") | ")
		operations := pathItem.Operations()
		methods := make([]string, 0, len(operations))
		for method := range operations {
			methods = append(methods, method)
		}
		sort.Strings(methods)
		sb.WriteString(strings.Join(methods, ", ") + " |\n")
	}

	sb.WriteString(newlineDouble)
	return sb.String()
}

// generatePathsDocumentation generates detailed markdown for each path and operation
func generatePathsDocumentation(paths openapi3.Paths) string {
	var sb strings.Builder

	// Sort paths by path name
	pathNames := make([]string, 0, len(paths))
	for path := range paths {
		pathNames = append(pathNames, path)
	}
	sort.Strings(pathNames)

	// Generate Markdown for each path
	for _, path := range pathNames {
		pathItem := paths[path]
		sb.WriteString("## <span id=\"path" + path + "\">" + path + "</span>\n\n")

		// Generate Markdown for each HTTP method in the path
		operations := pathItem.Operations()
		methods := make([]string, 0, len(operations))
		for method := range operations {
			methods = append(methods, method)
		}
		sort.Strings(methods)
		
		for _, method := range methods {
			operation := operations[method]
			sb.WriteString(generateOperationDocumentation(method, operation))
		}
	}

	return sb.String()
}

// generateOperationDocumentation generates markdown for a single operation
func generateOperationDocumentation(method string, operation *openapi3.Operation) string {
	var sb strings.Builder

	sb.WriteString("### " + method + "\n\n")
	sb.WriteString(operation.Description + "\n\n")
	
	sb.WriteString(generateParametersTable(operation.Parameters))
	sb.WriteString(generateRequestBodyTable(operation.RequestBody))
	sb.WriteString(generateResponsesTable(operation.Responses))
	sb.WriteString(sectionSeparator) // Add a separator between each method

	return sb.String()
}

// generateParametersTable generates a markdown table for operation parameters
func generateParametersTable(parameters openapi3.Parameters) string {
	var sb strings.Builder

	sb.WriteString(parametersSection + newlineDouble)
	sb.WriteString(paramsTableHeader)

	for _, parameter := range parameters {
		if parameter.Value != nil {
			sb.WriteString("| " + parameter.Value.Name + " | ")
			sb.WriteString(fmt.Sprintf("%v", parameter.Value.Required) + " | ")

			typ, ok := parameter.Value.Extensions["type"]
			if ok {
				sb.WriteString(typ.(string) + " | ")
			} else {
				sb.WriteString(" | ")
			}

			sb.WriteString(oneleline(parameter.Value.Description) + " | ")
			sb.WriteString(fmt.Sprintf("%v", parameter.Value.Example) + " |\n")
		}
	}

	return sb.String()
}

// generateRequestBodyTable generates a markdown table for request body
func generateRequestBodyTable(requestBody *openapi3.RequestBodyRef) string {
	if requestBody == nil {
		return ""
	}

	var sb strings.Builder

	sb.WriteString(newlineDouble)
	sb.WriteString(requestBodySection + newlineDouble)
	sb.WriteString(paramsTableHeader)

	for _, mediaType := range requestBody.Value.Content {
		sb.WriteString("| " + mediaType.Schema.Value.Title + " | ")
		sb.WriteString(fmt.Sprintf("%v", mediaType.Schema.Value.Required) + " | ")
		sb.WriteString(mediaType.Schema.Value.Type + " | ")
		sb.WriteString(oneleline(mediaType.Schema.Value.Description) + " | ")
		sb.WriteString(fmt.Sprintf("%v", mediaType.Schema.Value.Example) + " |\n")
	}

	return sb.String()
}

// generateResponsesTable generates a markdown table for responses
func generateResponsesTable(responses openapi3.Responses) string {
	var sb strings.Builder

	sb.WriteString(newline + responsesSection + newlineDouble)
	sb.WriteString(responsesTableHeader)

	// Sort responses by status code
	statusCodes := make([]string, 0, len(responses))
	for statusCode := range responses {
		statusCodes = append(statusCodes, statusCode)
	}
	sort.Strings(statusCodes)

	for _, statusCode := range statusCodes {
		response := responses[statusCode]
		description := ""

		if response.Value.Description != nil {
			description = oneleline(*response.Value.Description)
		}

		linkToSchema := ""
		ext, ok := response.Value.Extensions["schema"]
		if ok {
			ref, ok := ext.(map[string]any)["$ref"]
			if ok {
				linkToSchema = ref.(string)
			}
		}

		sb.WriteString("| " + statusCode + " | [" + description + "](" + linkToSchema + ") |\n")
	}

	return sb.String()
}