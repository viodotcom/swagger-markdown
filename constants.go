package main

// Markdown template constants for consistent formatting across the generated documentation.
const (
	// Basic formatting
	newline       = "\n"
	newlineDouble = "\n\n"

	// Table headers and separators (clean content + formatting)
	pathsTableHeader     = "| Path | Operations |" + newline + "| --- | --- |" + newline
	paramsTableHeader    = "| Name | Required | Type | Description | Example |" + newline + "| --- | --- | --- | --- | --- |" + newline
	responsesTableHeader = "| Status Code | Description |" + newline + "| --- | --- |" + newline
	propsTableHeader     = "| Name | Type | Description | Example |" + newline + "| --- | --- | --- | --- |" + newline

	// Section headers (clean content without formatting)
	pathsSection       = "## Paths"
	definitionsSection = "## Definitions"
	parametersSection  = "**Parameters:**"
	requestBodySection = "**Request Body:**"
	responsesSection   = "**Responses:**"
	propertiesSection  = "**Properties:**"
	enumsSection       = "## Enums"

	// Type indicators
	objectTypeLabel = "**Type:** object"
	mapTypePrefix   = "**Type:** map[*]->"

	// Complex formatting
	sectionSeparator = newlineDouble + "---" + newlineDouble

	// URL fragments (no formatting needed)
	definitionsPrefix = "#/definitions/"
	enumsPrefix       = "#/enums/"
	pathPrefix        = "#path"
)
