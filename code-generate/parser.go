package main

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"

	yaml "gopkg.in/mikefarah/yaml.v2"
)

// CreateCodeName generates a go identifier for the given value. We also do some extra
// modifications to make golint happy.
func CreateCodeName(value string) string {
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

	codeName := CreateCodeName(name)
	if codeName == "Error" {
		codeName = "ErrorMessage"
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

// ResourceHTTPMethod contains data for a specific HTTP method
type ResourceHTTPMethod struct {
	HTTPMethod  string
	Description string
	Traits      []string
}

// HasTrait checks if a ResourceHTTPMethod has a certain trait
func (resourceHTTPMethod *ResourceHTTPMethod) HasTrait(searchTrait string) bool {
	for _, trait := range resourceHTTPMethod.Traits {
		if trait == searchTrait {
			return true
		}
	}
	return false
}

// ResourceMethod contains HTTP CRUD actions for a specific method.
type ResourceMethod struct {
	Path              string
	PathParameterName string
	MethodName        string
	HTTPMethods       []ResourceHTTPMethod
}

// ResourceService contains all the data to CRUD an API resource.
type ResourceService struct {
	BasePath          string
	HasCreate         bool
	HasList           bool
	ResourceMethods   []ResourceMethod
	ResourceType      string
	ResourceQueryType string
	ResourceDraftType string
}

func parseYaml(data yaml.MapSlice) (objects []RamlType, resources []ResourceService) {
	types := getPropertyValue(data, "types")
	for _, mapItem := range types.(yaml.MapSlice) {
		obj := createRamlType(mapItem.Key.(string), mapItem.Value.(yaml.MapSlice))
		objects = append(objects, obj)
	}

	resources = make([]ResourceService, 0)
	apiResources := getPropertyValue(data, "/{projectKey}")
	for _, apiResource := range apiResources.(yaml.MapSlice) {
		if !strings.HasPrefix(apiResource.Key.(string), "/") {
			continue
		}
		resourceService := ResourceService{
			BasePath: apiResource.Key.(string)[1:],
		}
		for _, item := range apiResource.Value.(yaml.MapSlice) {
			key := item.Key.(string)
			if key == "get" {
				resourceService.HasList = true
			}
			if key == "post" {
				resourceService.HasCreate = true
			}
			if key == "type" {
				typeInfo := getPropertyValue(apiResource.Value.(yaml.MapSlice), "type")
				if _, ok := typeInfo.(yaml.MapSlice); !ok {
					continue
				}
				baseDomain := getPropertyValue(typeInfo.(yaml.MapSlice), "baseDomain").(yaml.MapSlice)
				resourceService.ResourceType = getPropertyString(baseDomain, "resourceType")
				resourceService.ResourceQueryType = strings.Split(getPropertyString(baseDomain, "resourceQueryType"), " ")[0]
				resourceService.ResourceDraftType = getPropertyString(baseDomain, "resourceDraft")
			}
			if strings.HasPrefix(key, "/") {
				resourceMethod := createResourceMethod(item.Value.(yaml.MapSlice), key)
				// Only support withId/withKey for now.
				if resourceMethod.MethodName == "withID" || resourceMethod.MethodName == "withKey" {
					resourceService.ResourceMethods = append(resourceService.ResourceMethods, resourceMethod)
				}
			}
		}
		resources = append(resources, resourceService)
	}

	return
}

func createResourceMethod(resourceMethodData yaml.MapSlice, key string) (resourceMethod ResourceMethod) {
	resourceMethod = ResourceMethod{
		Path: key,
	}
	switch key {
	case "/{ID}":
		resourceMethod.PathParameterName = "ID"
	case "/key={key}":
		resourceMethod.PathParameterName = "key"
	}
	for _, methodEntry := range resourceMethodData {
		methodKey := methodEntry.Key.(string)
		if methodKey == "(methodName)" {
			resourceMethod.MethodName = methodEntry.Value.(string)
			if resourceMethod.MethodName == "withId" {
				resourceMethod.MethodName = "withID"
			}
		}

		if methodKey == "get" || methodKey == "post" || methodKey == "delete" {
			resourceHTTPMethod := createHTTPMethod(methodEntry, methodKey)
			resourceMethod.HTTPMethods = append(resourceMethod.HTTPMethods, resourceHTTPMethod)
		}
	}
	return resourceMethod
}

func createHTTPMethod(methodEntry yaml.MapItem, methodKey string) ResourceHTTPMethod {
	resourceHTTPMethod := ResourceHTTPMethod{
		HTTPMethod: methodKey,
	}
	if methodEntry.Value != nil {
		methodInfo := methodEntry.Value.(yaml.MapSlice)
		if methodInfo != nil {
			resourceHTTPMethod.Description = getPropertyString(methodInfo, "description")
			isSection := getPropertyValue(methodInfo, "is")
			if isSection != nil {
				for _, traitInterface := range isSection.([]interface{}) {
					if traitName, ok := traitInterface.(string); ok {
						resourceHTTPMethod.Traits = append(resourceHTTPMethod.Traits, traitName)
					}
				}
			}
		}
	}
	return resourceHTTPMethod
}
