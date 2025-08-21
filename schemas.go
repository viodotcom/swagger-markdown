package main

import (
	"fmt"
	"sort"
	"strings"
)

// markDownDefinitions generates markdown for schema definitions
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

		sb.WriteString(sectionSeparator)
	}

	return sb.String()
}

// detectObjectGoType determines if a schema represents an object or a map
func detectObjectGoType(schemaMap map[string]any) string {
	if _, ok := schemaMap["properties"]; ok {
		return "object"
	}

	if _, ok := schemaMap["additionalProperties"]; ok {
		return "map"
	}

	return ""
}

// objectMarkDown generates markdown for object schemas
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

// objectMD generates markdown for object types with properties
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
	sb.WriteString(objectTypeLabel + newlineDouble)
	sb.WriteString(propertiesSection + newlineDouble)
	sb.WriteString(propsTableHeader)

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
			refText := strings.TrimPrefix(refLink, definitionsPrefix)
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
			typ = fmt.Sprintf("%s ([enums](%s%s))", typ, enumsPrefix, propertyName)
		}

		sb.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", propertyName, typ, description, example))
	}

	if len(enumInformations) > 0 {
		sb.WriteString("\n\n")
		sb.WriteString(enumsSection + newlineDouble)

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

// mapMD generates markdown for map types
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

	sb.WriteString(mapTypePrefix)

	if ref, ok := additionalPropertiesMap["$ref"].(string); ok {
		objectName := strings.ReplaceAll(ref, definitionsPrefix, "")
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

// arrayMarkDown generates markdown for array types
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
		objectName := strings.TrimPrefix(ref, definitionsPrefix)
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