package main

import (
	yaml "gopkg.in/yaml.v2"
)

func parseYaml(data yaml.MapSlice) (objects []RamlType, domains []ServiceDomain) {
	types := getPropertyValue(data, "types")
	for _, mapItem := range types.(yaml.MapSlice) {
		obj := createRamlType(mapItem.Key.(string), mapItem.Value.(yaml.MapSlice))
		objects = append(objects, obj)
	}

	domains = make([]ServiceDomain, 0)
	apiResources := getPropertyValue(data, "/{projectKey}")

	for _, apiResource := range apiResources.(yaml.MapSlice) {
		serviceDomain := createService(apiResource, nil)
		if serviceDomain != nil {
			domains = append(domains, *serviceDomain)
		}

		// resourceService := createResourceService(apiResource)
		// if resourceService != nil {
		// 	resources = append(resources, *resourceService)
		// }
	}

	return
}
