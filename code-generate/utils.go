package main

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

// CreateCodeName generates a go identifier for the given value. We also do some extra
// modifications to make golint happy.
func CreateCodeName(value string) string {
	if strings.HasPrefix(value, "/") {
		return value
	}

	result := strcase.ToCamel(value)
	translateMap := map[string]string{
		"/":        "",
		"Id$":      "ID",
		"Sku$":     "SKU",
		"Uri$":     "URI",
		"Url$":     "URL",
		"Json":     "JSON",
		"IdAction": "IDAction",
		"Ttl":      "TTL",
		"Http":     "HTTP",
		"^Api":     "API",
	}

	for key, value := range translateMap {
		r := regexp.MustCompile(key)
		result = r.ReplaceAllString(result, value)
	}

	return result
}

func ExtractCodeName(value string) string {
	values := strings.Split(value, "|")
	return strings.TrimSpace(values[0])
}

func generateFilename(name string, suffix string) string {
	if name == "" {
		panic("No package name found")
	}
	filename := strcase.ToSnake(name)
	if suffix != "" {
		filename += suffix

	}
	return filename + ".go"
}
