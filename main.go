package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// convert a swagger file to markdown file
// usage: swaggermd -i swagger.json -o swagger.md
func main() {
	input := flag.String("i", "", "input swagger file")
	output := flag.String("o", "", "output markdown file")
	flag.Parse()
	if *input == "" || *output == "" {
		flag.Usage()
		return
	}

	// check if the input file exists
	if _, err := os.Stat(*input); os.IsNotExist(err) {
		fmt.Printf("Input file not found: %s\n", *input)
		return
	}

	// read the swagger file
	swaggerData, err := os.ReadFile(*input)
	if err != nil {
		panic(err)
	}

	// parse the swagger file
	swagger, err := openapi3.NewLoader().LoadFromData(swaggerData)
	if err != nil {
		fmt.Printf("Error parsing Swagger file: %v\n", err)
		return
	}

	// generate markdown
	markdown := generateMarkdown(swagger)

	// Save the Markdown to a file
	err = os.WriteFile(*output, []byte(markdown), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Markdown file generated: %s\n", *output)
}

func generateMarkdown(swagger *openapi3.T) string {
	var sb strings.Builder

	if swagger.Info != nil {
		sb.WriteString("# " + swagger.Info.Title + "\n\n")
		sb.WriteString(swagger.Info.Description + "\n\n")

		if swagger.Info.Contact != nil {
			sb.WriteString(swagger.Info.Contact.Name + "\n\n")
			sb.WriteString(swagger.Info.Contact.URL + "\n\n")
			sb.WriteString(swagger.Info.Contact.Email + "\n\n")
		}
	}

	// Generate a table of all paths with the link to the path
	sb.WriteString("## Paths\n\n")
	sb.WriteString("| Path | Operations |\n")
	sb.WriteString("| --- | --- |\n")

	// Sort paths by path name
	paths := make([]string, 0, len(swagger.Paths))
	for path := range swagger.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		pathItem := swagger.Paths[path]
		sb.WriteString("| [" + path + "](#path" + strings.ToLower(path) + ") | ")
		operations := pathItem.Operations()
		methods := make([]string, 0, len(operations))
		for method := range operations {
			methods = append(methods, method)
		}
		sort.Strings(methods)
		sb.WriteString(strings.Join(methods, ", ") + " |\n")

	}

	sb.WriteString("\n\n")

	// Generate Markdown for each path
	for _, path := range paths {
		pathItem := swagger.Paths[path]
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
			sb.WriteString("### " + method + "\n\n")
			sb.WriteString(operation.Description + "\n\n")
			sb.WriteString("**Parameters:**\n\n")
			sb.WriteString("| Name | Required | Type | Description | Example |\n")
			sb.WriteString("| --- | --- | --- | --- | --- |\n")

			for _, parameter := range operation.Parameters {
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

			if operation.RequestBody != nil {
				sb.WriteString("\n\n")
				sb.WriteString("**Request Body:**\n\n")
				sb.WriteString("| Name | Required | Type | Description | Example |\n")
				sb.WriteString("| --- | --- | --- | --- | --- |\n")

				for _, mediaType := range operation.RequestBody.Value.Content {
					sb.WriteString("| " + mediaType.Schema.Value.Title + " | ")
					sb.WriteString(fmt.Sprintf("%v", mediaType.Schema.Value.Required) + " | ")
					sb.WriteString(mediaType.Schema.Value.Type + " | ")
					sb.WriteString(oneleline(mediaType.Schema.Value.Description) + " | ")
					sb.WriteString(fmt.Sprintf("%v", mediaType.Schema.Value.Example) + " |\n")
				}
			}

			sb.WriteString("\n**Responses:**\n\n")
			sb.WriteString("| Status Code | Description |\n")
			sb.WriteString("| --- | --- |\n")

			// Sort responses by status code
			statusCodes := make([]string, 0, len(operation.Responses))
			for statusCode := range operation.Responses {
				statusCodes = append(statusCodes, statusCode)
			}
			sort.Strings(statusCodes)

			for _, statusCode := range statusCodes {
				response := operation.Responses[statusCode]
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

			sb.WriteString("\n\n---\n\n") // Add a separator between each method
		}
	}

	sb.WriteString("## Definitions\n\n")
	definitions, ok := swagger.Extensions["definitions"]

	if ok {
		definitionsMap, ok := definitions.(map[string]any)
		if ok {
			sb.WriteString(markDownDefinitions(definitionsMap))
		}
	}

	return sb.String()
}

func markDownDefinitions(definitionsMap map[string]any) string {
	var sb strings.Builder

	// Sort definitions by name
	definitionNames := make([]string, 0, len(definitionsMap))
	for name := range definitionsMap {
		definitionNames = append(definitionNames, name)
	}
	sort.Strings(definitionNames)

	for _, name := range definitionNames {
		schema := definitionsMap[name]
		sb.WriteString(fmt.Sprintf("### <span id=\"/definitions/%s\">%s</span>\n\n", name, name))
		sb.WriteString(fmt.Sprintf("<a id=\"/definitions/%s\"></a>\n\n", name))

		schemaMap, ok := schema.(map[string]any)
		if !ok {
			continue
		}

		if title, ok := schemaMap["title"].(string); ok {
			sb.WriteString(title + "\n\n")
		}

		if description, ok := schemaMap["description"].(string); ok {
			sb.WriteString(description + "\n\n")
		}

		switch schemaMap["type"] {
		case "object":
			sb.WriteString(objectMarkDown(schemaMap) + "\n\n")
		case "array":
			sb.WriteString(arrayMarkDown(schemaMap) + "\n\n")
		default:
			sb.WriteString(fmt.Sprintf("**Type:** %s\n\n", schemaMap["type"]))
		}

		sb.WriteString("\n\n---\n\n")
	}

	return sb.String()
}

func oneleline(input string) string {
	return strings.TrimSpace(strings.ReplaceAll(input, "\n", " "))
}

func detectObjectGoType(schemaMap map[string]any) string {
	if _, ok := schemaMap["properties"]; ok {
		return "object"
	}

	if _, ok := schemaMap["additionalProperties"]; ok {
		return "map"
	}

	return ""
}

func objectMarkDown(schemaMap map[string]any) string {
	goType := detectObjectGoType(schemaMap)

	switch goType {
	case "object":
		return objectMD(schemaMap)
	case "map":
		return mapMD(schemaMap)
	default:
		return ""
	}
}

func objectMD(schemaMap map[string]any) string {
	properties, ok := schemaMap["properties"]
	if !ok {
		return ""
	}

	propertiesMap, ok := properties.(map[string]any)
	if !ok {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("**Type:** object\n\n")
	sb.WriteString("**Properties:**\n\n")
	sb.WriteString("| Name | Type | Description | Example |\n")
	sb.WriteString("| --- | --- | --- | --- |\n")

	enumInformations := make(map[string][]string)

	// Sort properties by name
	propertyNames := make([]string, 0, len(propertiesMap))
	for name := range propertiesMap {
		propertyNames = append(propertyNames, name)
	}
	sort.Strings(propertyNames)

	for _, propertyName := range propertyNames {
		property := propertiesMap[propertyName]
		propertyMap, ok := property.(map[string]any)
		if !ok {
			continue
		}

		typ := ""
		if propertyType, ok := propertyMap["type"].(string); ok {
			typ = propertyType
		}

		if refLink, ok := propertyMap["$ref"].(string); ok {
			refText := strings.TrimPrefix(refLink, "#/definitions/")
			typ = fmt.Sprintf("[%s](%s)", refText, refLink)
		} else if typ == "array" {
			typ = arrayMarkDown(propertyMap)
		}

		description := ""
		if propertyDescription, ok := propertyMap["description"].(string); ok {
			description = oneleline(propertyDescription)
		}

		example := ""
		if propertyExample, ok := propertyMap["example"].(string); ok {
			example = propertyExample
		}

		enums := []string{}
		if propertyEnums, ok := propertyMap["enum"].([]any); ok {
			for _, enum := range propertyEnums {
				if enumValue, ok := enum.(string); ok {
					enums = append(enums, enumValue)
				}
			}
			enumInformations[propertyName] = enums
			typ = fmt.Sprintf("%s ([enums](#/enums/%s))", typ, propertyName)
		}

		sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", propertyName, typ, description, example))
	}

	if len(enumInformations) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString("## Enums\n\n")

		// Sort enums by name
		enumNames := make([]string, 0, len(enumInformations))
		for name := range enumInformations {
			enumNames = append(enumNames, name)
		}
		sort.Strings(enumNames)

		for _, enumName := range enumNames {
			enumValues := enumInformations[enumName]
			sb.WriteString(fmt.Sprintf("**<span id=\"/enums/%s\"></span>%s:**\n\n", enumName, enumName))
			sb.WriteString("| " + enumName + " |\n")
			sb.WriteString("| --- |\n")
			sb.WriteString("|" + strings.Join(enumValues, ", ") + "|\n\n")
		}
	}

	return sb.String()
}

func mapMD(schemaMap map[string]any) string {
	sb := strings.Builder{}

	additionalProperties, ok := schemaMap["additionalProperties"]
	if !ok {
		return ""
	}

	additionalPropertiesMap, ok := additionalProperties.(map[string]any)
	if !ok {
		return ""
	}

	sb.WriteString("**Type:** map[*]->")

	if ref, ok := additionalPropertiesMap["$ref"].(string); ok {
		objectName := strings.ReplaceAll(ref, "#/definitions/", "")
		sb.WriteString("[" + objectName + "](" + ref + ")")
	} else if additionalPropertiesMap["type"] == "array" {
		sb.WriteString(arrayMarkDown(additionalPropertiesMap))
	} else if additionalPropertiesMap["type"] == "object" {
		sb.WriteString(objectMarkDown(additionalPropertiesMap))
	} else if mapValueText, ok := additionalPropertiesMap["type"].(string); ok {
		sb.WriteString(mapValueText)
	}

	return sb.String()
}

func arrayMarkDown(schemaMap map[string]any) string {
	sb := strings.Builder{}
	sb.WriteString("[]")

	items, ok := schemaMap["items"].(map[string]any)
	if !ok {
		return sb.String()
	}

	if typ, ok := items["type"].(string); ok {
		switch typ {
		case "object":
			sb.WriteString(objectMarkDown(items))
		case "array":
			sb.WriteString(arrayMarkDown(items))
		default:
			sb.WriteString(typ)
		}
	} else if ref, ok := items["$ref"].(string); ok {
		objectName := strings.TrimPrefix(ref, "#/definitions/")
		sb.WriteString(fmt.Sprintf("[%s](%s)", objectName, ref))
	} else if itemsItems, ok := items["items"].(map[string]any); ok {
		if typ, ok := itemsItems["type"].(string); ok {
			switch typ {
			case "object":
				sb.WriteString(objectMarkDown(itemsItems))
			case "array":
				sb.WriteString(arrayMarkDown(itemsItems))
			default:
				sb.WriteString(typ)
			}
		}
	} else {
		sb.WriteString("unknown")
	}

	return sb.String()
}
