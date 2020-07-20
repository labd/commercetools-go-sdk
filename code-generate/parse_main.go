package main

import (
	yaml "gopkg.in/yaml.v2"
)

func parseYaml(data yaml.MapSlice) (objects []RamlType, resources []ResourceService) {
	types := getPropertyValue(data, "types")
	for _, mapItem := range types.(yaml.MapSlice) {
		obj := createRamlType(mapItem.Key.(string), mapItem.Value.(yaml.MapSlice))
		objects = append(objects, obj)
	}

	resources = make([]ResourceService, 0)
	apiResources := getPropertyValue(data, "/{projectKey}")
	for _, apiResource := range apiResources.(yaml.MapSlice) {
		resourceService := createResourceService(apiResource)
		if resourceService != nil {
			resources = append(resources, *resourceService)
		}
	}

	return
}
