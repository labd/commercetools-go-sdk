package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/dave/jennifer/jen"
)

type ResourceMethods struct {
	Name          string
	ServiceDomain ServiceDomain
	Get           []ServiceMethod
	Query         []ServiceMethod
	Create        []ServiceMethod
	Update        []ServiceMethod
	Delete        []ServiceMethod
	Actions       []ServiceMethod
}

func groupByResourceMethods(methods []ServiceMethod) ResourceMethods {
	result := ResourceMethods{}

	sort.Slice(methods, func(i, j int) bool {
		return methods[i].Name < methods[j].Name
	})

	for _, method := range methods {
		switch method.Type {
		case "create":
			result.Create = append(result.Create, method)
		case "update":
			result.Update = append(result.Update, method)
		case "query":
			result.Query = append(result.Query, method)
		case "delete":
			result.Delete = append(result.Delete, method)
		case "get":
			result.Get = append(result.Get, method)
		case "action":
			result.Actions = append(result.Actions, method)
		default:
			log.Fatalf("Unknown method type for %s::%s: '%s'\n", method.Context, method.Name, method.Type)
		}
	}
	return result
}

// Entry point to generate services
func generateServices(objects []RamlType, resources []ServiceDomain) {
	typeMapping := map[string]string{}

	for _, item := range objects {
		typeMapping[item.Name] = item.CodeName

		if item.Name == "Message" {
			fmt.Println(item)
		}
	}
	objectsMap := map[string]RamlType{}
	for _, item := range objects {
		objectsMap[item.Name] = item
	}

	for _, resourceService := range resources {

		if len(resourceService.methods) < 1 {
			log.Printf("! No methods found for domain %s\n", resourceService.ContextName)
			continue
		}

		filename := generateFilename(resourceService.Package, "")
		log.Printf("Writing commercetools/%s\n", filename)

		f := jen.NewFile("commercetools")
		f.HeaderComment("Automatically generated, do not edit")

		contextNames, resourceMethods := orderServiceMethods(&resourceService, objectsMap)
		for _, contextName := range contextNames {
			value := resourceMethods[contextName]

			for _, method := range value.Create {
				code := generateServiceCreate(method, objectsMap)
				f.Add(code)
			}

			for _, method := range value.Query {
				queryCode := generateServiceQuery(method, objectsMap)
				f.Add(queryCode)
			}

			for _, method := range value.Delete {
				code := generateServiceDelete(method, objectsMap)
				f.Add(code)
			}

			for _, method := range value.Get {
				code := generateServiceGet(method, objectsMap)
				f.Add(code)
			}

			for _, method := range value.Update {
				code := generateServiceUpdate(method, objectsMap)
				f.Add(code)
			}

			for _, method := range value.Actions {
				code := generateServiceAction(method, objectsMap)
				f.Add(code)
			}
		}

		err := f.Save("commercetools/service_" + filename)
		if err != nil {
			panic(err)
		}
	}
}

func orderServiceMethods(sd *ServiceDomain, objects map[string]RamlType) ([]string, map[string]ResourceMethods) {
	groupedMethods := map[string][]ServiceMethod{}
	for _, item := range sd.methods {
		item.Domain = *sd
		item.UpdateTypeNames(objects)
		groupedMethods[item.Context] = append(groupedMethods[item.Context], item)
	}

	groupedResourceMethods := map[string]ResourceMethods{}
	contextNames := []string{}

	for key, value := range groupedMethods {
		groupedResourceMethods[key] = groupByResourceMethods(value)
		contextNames = append(contextNames, key)
	}
	sort.Strings(contextNames)

	return contextNames, groupedResourceMethods
}
