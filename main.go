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

	// generate a table of all paths with the link to the path
	sb.WriteString("## Paths\n\n")
	sb.WriteString("| Path | Operations |\n")
	sb.WriteString("| --- | --- |\n")
	// sort paths by path name
	paths := make([]string, 0, len(swagger.Paths))
	for path := range swagger.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		pathItem := swagger.Paths[path]
		sb.WriteString("| [" + path + "](#path" + strings.ToLower(path) + ") | ")
		for method := range pathItem.Operations() {
			sb.WriteString(method + " ")
		}
		sb.WriteString("|\n")
	}

	sb.WriteString("\n\n")

	// Generate Markdown for each path
	for _, path := range paths {
		pathItem := swagger.Paths[path]
		sb.WriteString("## <span id=\"path" + path + "\">" + path + "</span>\n\n")

		// Generate Markdown for each HTTP method in the path
		for method, operation := range pathItem.Operations() {
			sb.WriteString("### " + method + "\n\n")
			sb.WriteString(operation.Description + "\n\n")
			sb.WriteString("**Parameters:**\n\n")
			sb.WriteString("| Name | Required | Type | Description | Example |\n")
			sb.WriteString("| --- | --- | --- | --- | --- |\n")
			for _, parameter := range operation.Parameters {
				if parameter.Value != nil {
					sb.WriteString("| " +
						parameter.Value.Name + " | " +
						fmt.Sprintf("%v", parameter.Value.Required) + " | ")
					typ, ok := parameter.Value.Extensions["type"]
					if ok {
						sb.WriteString(typ.(string) + " | ")
					} else {
						sb.WriteString(" | ")
					}
					sb.WriteString(
						oneleline(parameter.Value.Description) + " | " +
							fmt.Sprintf("%v", parameter.Value.Example) + " |\n")
				}
			}

			sb.WriteString("**Responses:**\n\n")
			sb.WriteString("| Status Code | Description |\n")
			sb.WriteString("| --- | --- |\n")
			for statusCode, response := range operation.Responses {
				description := ""
				if response.Value.Description != nil {
					description = oneleline(*response.Value.Description)
				}
				linkToSchema := ""
				ext, ok := response.Value.Extensions["schema"]
				if ok {
					ref, ok := ext.(map[string]interface{})["$ref"]
					if ok {
						linkToSchema = ref.(string)
					}
				}
				// replace "#/definitions/" with "#definitions/" to make it work in GitHub
				//linkToSchema = strings.ReplaceAll(linkToSchema, "#/", "#")
				sb.WriteString("| " + statusCode + " | [" + description + "](" + linkToSchema + ") |\n")
			}
			sb.WriteString("\n\n---\n\n") // Add a separator between each method
		}
	}

	sb.WriteString("## Definitions\n\n")
	definitions, ok := swagger.Extensions["definitions"]
	if ok {
		for name, schema := range definitions.(map[string]interface{}) {
			sb.WriteString("### <span id=\"/definitions/" + name + "\">" + name + "</span>\n\n")
			sb.WriteString("<a id=\"/definitions/" + name + "\"></a>\n\n")
			schemaMap := schema.(map[string]interface{})
			title, ok := schemaMap["title"].(string)
			if ok {
				sb.WriteString(title + "\n\n")
			}
			description, ok := schemaMap["description"].(string)
			if ok {
				sb.WriteString(description + "\n\n")
			}

			if schemaMap["type"] == "object" {
				sb.WriteString(objectMarkDown(schemaMap) + "\n\n")
			} else if schemaMap["type"] == "array" {
				sb.WriteString(arrayMarkDown(schemaMap) + "\n\n")
			} else {
				sb.WriteString("**Type:** " + schemaMap["type"].(string) + "\n\n")
			}

			sb.WriteString("\n\n---\n\n")
		}
	}

	return sb.String()
}

func oneleline(input string) string {
	return strings.ReplaceAll(input, "\n", " ")
}

func detectObjectGoType(schemaMap map[string]interface{}) string {
	_, ok := schemaMap["properties"]
	goType := "object"
	if !ok {
		_, ok = schemaMap["additionalProperties"]
		goType = "map"
		if !ok {
			return ""
		}
	}
	return goType
}

func objectMarkDown(schemaMap map[string]interface{}) string {
	goType := detectObjectGoType(schemaMap)
	if goType == "object" {
		return objectMD(schemaMap)
	} else if goType == "map" {
		return mapMD(schemaMap)
	}
	return ""
}

func objectMD(schemaMap map[string]interface{}) string {
	sb := strings.Builder{}

	properties, ok := schemaMap["properties"]
	if !ok {
		return ""
	}

	propertiesMap, ok := properties.(map[string]interface{})
	if !ok {
		return ""
	}
	sb.WriteString("**Type:** object\n\n")
	sb.WriteString("**Properties:**\n\n")
	sb.WriteString("| Name | Type | Description | Example |\n")
	sb.WriteString("| --- | --- | --- | --- |\n")
	enumInformations := make(map[string][]string, 0)
	for propertyName, property := range propertiesMap {
		propertyMap := property.(map[string]interface{})
		typ, ok := propertyMap["type"]
		var typeText string
		if ok {
			typeText, _ = typ.(string)
		}

		ref, ok := propertyMap["$ref"]
		var refLink string
		if ok {
			refLink, _ = ref.(string)
		}

		description, ok := propertyMap["description"]
		var descriptionText string
		if ok {
			descriptionText, _ = description.(string)
		}

		example, ok := propertyMap["example"]
		var exampleText string
		if ok {
			exampleText, _ = example.(string)
		}

		enums, ok := propertyMap["enum"]
		var enumsTexts []string
		if ok {
			enumsArray, ok := enums.([]interface{})
			if ok {
				for _, enum := range enumsArray {
					enumsTexts = append(enumsTexts, enum.(string))
				}
				enumInformations[propertyName] = enumsTexts
			}
		}

		if refLink != "" {
			refText := strings.ReplaceAll(refLink, "#/definitions/", "")
			typeText = "[" + refText + "]" + "(" + refLink + ")"
		} else if typeText == "array" {
			items, ok := propertyMap["items"]
			if ok {
				itemsMap := items.(map[string]interface{})
				ref, ok := itemsMap["$ref"]
				if ok {
					refLink, _ = ref.(string)
					refText := strings.ReplaceAll(refLink, "#/definitions/", "")
					typeText = "[][" + refText + "]" + "(" + refLink + ")"
				} else {
					typ, ok := itemsMap["type"]
					if ok {
						arrayTypeText, _ := typ.(string)
						typeText = "[]" + arrayTypeText
					}
				}
			}
		}
		if len(enumsTexts) > 0 {
			typeText = typeText + " ([enums](#/enums/" + propertyName + "))"
		}

		sb.WriteString("| " + propertyName + " | " + typeText + " | " + oneleline(descriptionText) + " | " + exampleText + " |\n")
	}

	if len(enumInformations) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString("## Enums\n\n")
		for enumName, enumValues := range enumInformations {
			sb.WriteString("**<span id=\"/enums/" + enumName + "\"></span>" + enumName + ":**\n\n")
			sb.WriteString("| " + enumName + " |\n")
			sb.WriteString("| --- |\n")
			sb.WriteString("|" + strings.Join(enumValues, ", ") + "|\n\n")
		}
	}

	return sb.String()
}

func mapMD(schemaMap map[string]interface{}) string {
	sb := strings.Builder{}
	additionalProperties, ok := schemaMap["additionalProperties"]
	if !ok {
		return ""
	}
	additionalPropertiesMap, ok := additionalProperties.(map[string]interface{})
	if !ok {
		return ""
	}
	sb.WriteString("**Type:** map[*]->")
	ref, ok := additionalPropertiesMap["$ref"]
	if ok {
		refLink, ok := ref.(string)
		if ok {
			objectName := strings.ReplaceAll(refLink, "#/definitions/", "#")
			sb.WriteString("[" + objectName + "](" + refLink + ")\n\n")
		}
	} else if additionalPropertiesMap["type"] == "array" {
		sb.WriteString(arrayMarkDown(additionalPropertiesMap) + "\n\n")
	} else if additionalPropertiesMap["type"] == "object" {
		sb.WriteString(objectMarkDown(additionalPropertiesMap) + "\n\n")
	} else {
		mapValueText, ok := additionalPropertiesMap["type"].(string)
		if ok {
			sb.WriteString(mapValueText + "\n\n")
		}
	}
	return sb.String()
}

func arrayMarkDown(schemaMap map[string]interface{}) string {
	sb := strings.Builder{}
	sb.WriteString("[]")
	items, ok := schemaMap["items"]
	if ok {
		itemsMap, ok := items.(map[string]interface{})
		if ok {
			ref, ok := itemsMap["$ref"]
			if ok {
				refLink, ok := ref.(string)
				if ok {
					objectName := strings.ReplaceAll(refLink, "#/definitions/", "")
					sb.WriteString("[" + objectName + "](" + refLink + ")\n\n")
				}
			} else {
				itemsMapItems, ok := itemsMap["items"]
				if ok {
					itemsMapItemsMap, ok := itemsMapItems.(map[string]interface{})
					if ok {
						typ, ok := itemsMapItemsMap["type"]
						if ok {
							typText, ok := typ.(string)
							if ok {
								if typText == "object" {
									sb.WriteString(objectMarkDown(itemsMapItemsMap) + "\n\n")
								} else if typText == "array" {
									sb.WriteString(arrayMarkDown(itemsMapItemsMap) + "\n\n")
								} else {
									sb.WriteString(typText + "\n\n")
								}
							}
						}
					}
				} else if itemsMap["type"] == "array" {
					sb.WriteString(arrayMarkDown(itemsMap) + "\n\n")
				} else if itemsMap["type"] == "object" {
					sb.WriteString(objectMarkDown(itemsMap) + "\n\n")
				} else {
					typ, ok := itemsMap["type"]
					if ok {
						typText, ok := typ.(string)
						if ok {
							sb.WriteString(typText + "\n\n")
						}
					}
				}
			}
		}
	}
	return sb.String()
}
