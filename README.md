# swagger-markdown
Converts a swagger or OpenAPI file to MarkDown

## Installation

### Using go install
```bash
go install github.com/omid/swagger-markdown@latest
```

### From source
```bash
git clone https://github.com/omid/swagger-markdown.git
cd swagger-markdown
go build
```

## Usage

```bash
swagger-markdown -i <inputFile> -o <outputFile>
```

### Example

```bash
swagger-markdown -i ./testdata/sample-swagger.yaml -o ./testdata/sample-output.md
```

### Development Usage

If you're developing or testing locally:
```bash
go run main.go -i ./testdata/sample-swagger.yaml -o ./testdata/sample-output.md
```
