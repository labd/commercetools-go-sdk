package main

import (
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func createRamlType(name string, properties yaml.MapSlice) RamlType {

	typeName := getPropertyString(properties, "type")

	object := RamlType{
		Name:               name,
		TypeName:           typeName,
		CodeName:           CreateCodeName(name),
		InterfaceName:      CreateCodeName(name),
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

		// If the raml type has one attribute with a wildcard name (//) then
		// turn it into a map keeping the type
		if len(object.Attributes) == 1 && object.Attributes[0].Name == "//" {
			object.asMap = true
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

	var itemsObject *RamlType
	if strings.HasPrefix(name, "/") {
		placeholder := getPropertyValue(properties.(yaml.MapSlice), "(placeholderParam)")
		if placeholder != nil {
			name = getPropertyString(properties.(yaml.MapSlice), "(placeholderParam)", "paramName")
			itemsObject = &RamlType{
				Name:     getPropertyString(properties.(yaml.MapSlice), "(placeholderParam)", "placeholder"),
				TypeName: typeName,
			}
		}
	}

	codeName := CreateCodeName(name)
	if codeName == "Error" {
		codeName = "ErrorMessage"
	}

	object := RamlTypeAttribute{
		Name:        name,
		TypeName:    typeName,
		CodeName:    codeName,
		Optional:    optional,
		Many:        many,
		ItemsObject: itemsObject,
	}

	if props, isOk := properties.(yaml.MapSlice); isOk {
		object.Items = getPropertyString(props, "items")
		object.Format = getPropertyString(props, "format")
		object.Minimum = getPropertyInt(props, "minimum")
		object.Maximum = getPropertyInt(props, "maximum")
	}
	return object
}
