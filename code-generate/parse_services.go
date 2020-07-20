package main

import (
	"strings"

	yaml "gopkg.in/yaml.v2"
)

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

func createResourceService(apiResource yaml.MapItem) *ResourceService {
	if !strings.HasPrefix(apiResource.Key.(string), "/") {
		return nil
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
	return &resourceService
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
