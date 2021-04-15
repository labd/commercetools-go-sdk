package main

import (
	yaml "gopkg.in/yaml.v2"
)

// XXX: using a global variable has a smell... however, its just a code
// generator, not production code
var traits yaml.MapSlice

func parseYaml(data yaml.MapSlice) (objects []RamlType, domains []ServiceDomain) {
	types := getPropertyValue(data, "types")
	for _, mapItem := range types.(yaml.MapSlice) {
		obj := createRamlType(mapItem.Key.(string), mapItem.Value.(yaml.MapSlice))
		objects = append(objects, obj)
	}
	traits = getPropertyValue(data, "traits").(yaml.MapSlice)

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

func getTrait(name string) yaml.MapSlice {
	for _, t := range traits {
		if t.Key == name {
			return t.Value.(yaml.MapSlice)
		}
	}
	return nil
}
