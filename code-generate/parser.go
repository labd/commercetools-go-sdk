package main

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	yaml "gopkg.in/mikefarah/yaml.v2"
)

// Generate a go identifier for the given value. We also do some extra
// modifications to make golint happy.
func createCodeName(value string) string {
	if strings.HasPrefix(value, "/") {
		return value
	}

	result := strcase.ToCamel(value)
	translateMap := map[string]string{
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

func createRamlType(name string, properties yaml.MapSlice) RamlType {

	typeName := getPropertyString(properties, "type")

	object := RamlType{
		Name:               name,
		TypeName:           typeName,
		CodeName:           createCodeName(name),
		InterfaceName:      createCodeName(name),
		Discriminator:      getPropertyString(properties, "discriminator"),
		DiscriminatorValue: getPropertyString(properties, "discriminatorValue"),
		Package:            getPropertyString(properties, "(package)"),
	}
	if object.Package == "" {
		object.Package = "Base"
	}

	if value := getPropertyValue(properties, "enum"); value != nil {
		for _, item := range value.([]interface{}) {
			object.EnumValues = append(object.EnumValues, item.(string))
		}
	}

	if getPropertyValue(properties, "(asMap)") != nil {
		object.asMap = true
	}

	typeProperties := getPropertyValue(properties, "properties")

	if typeProperties != nil {
		for _, mapItem := range typeProperties.(yaml.MapSlice) {
			attribute := createRamlTypeAttribute(mapItem.Key.(string), mapItem.Value)
			object.Attributes = append(object.Attributes, attribute)
		}
	}
	return object
}

func createRamlTypeAttribute(name string, properties interface{}) RamlTypeAttribute {

	optional := false
	if strings.HasSuffix(name, "?") {
		name = name[:len(name)-1]
		optional = true
	}

	typeName := ""
	if value, isOk := properties.(yaml.MapSlice); isOk {
		typeName = getPropertyString(value, "type")
	} else {
		typeName = properties.(string)
	}

	many := false
	if strings.HasSuffix(typeName, "[]") {
		typeName = typeName[:len(typeName)-2]
		many = true
	}

	codeName := createCodeName(name)
	if codeName == "Error" {
		codeName = "_Error"
	}

	object := RamlTypeAttribute{
		Name:     name,
		TypeName: typeName,
		CodeName: codeName,
		Optional: optional,
		Many:     many,
	}

	if props, isOk := properties.(yaml.MapSlice); isOk {
		object.Items = getPropertyString(props, "items")
		object.Format = getPropertyString(props, "format")
		object.Minimum = getPropertyInt(props, "minimum")
		object.Maximum = getPropertyInt(props, "maximum")
	}
	return object
}

func parseYaml(data yaml.MapSlice) (objects []RamlType) {
	types := getPropertyValue(data, "types")
	for _, mapItem := range types.(yaml.MapSlice) {
		obj := createRamlType(mapItem.Key.(string), mapItem.Value.(yaml.MapSlice))
		objects = append(objects, obj)
	}
	return
}
