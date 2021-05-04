package main

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/davecgh/go-spew/spew"
	yaml "gopkg.in/yaml.v2"
)

type ServiceMethod struct {
	Name            string
	Context         string
	HTTPMethod      string
	Path            string
	Parameters      []string
	ReturnType      string
	InputType       string
	Type            string
	Domain          ServiceDomain
	Description     string
	QueryParameters []RamlTypeAttribute
	Traits          []string
	ExampleData     string
	exampleDataMap  map[string]interface{}
}

func (sm *ServiceMethod) getExampleValue(key string) interface{} {
	if sm.exampleDataMap == nil {
		err := json.Unmarshal([]byte(sm.ExampleData), &sm.exampleDataMap)
		if err != nil {
			panic(err)
		}
	}
	for k, v := range sm.exampleDataMap {
		if key == k {
			return v
		}
	}
	return nil
}

// HasTrait checks if a ServiceMethod has a certain trait
func (sm *ServiceMethod) HasTrait(searchTrait string) bool {
	for _, trait := range sm.Traits {
		if trait == searchTrait {
			return true
		}
	}
	return false
}

type ServiceDomain struct {
	Package           string
	ContextName       string // E.g. InStore or Me
	Path              string
	PathParameters    []string
	ResourceDraft     string
	ResourceType      string
	ResourceQueryType string

	methods []ServiceMethod
}

func (sm *ServiceMethod) UpdateTypeNames(types map[string]RamlType) {
	mapped, found := types[sm.InputType]
	if found {
		sm.InputType = mapped.CodeName
	}

	mapped, found = types[sm.ReturnType]
	if found {
		sm.ReturnType = mapped.CodeName
	}

}

func createService(apiResource yaml.MapItem, parent *ServiceDomain) *ServiceDomain {
	currentKey := apiResource.Key.(string)
	if !strings.HasPrefix(currentKey, "/") {
		return nil
	}

	if strings.HasPrefix(currentKey, "/graphql") {
		return nil
	}
	currentData := apiResource.Value.(yaml.MapSlice)

	var path string
	if parent != nil {
		path = parent.Path + currentKey
	} else {
		path = currentKey
	}

	serviceDomain := ServiceDomain{
		Path:              path,
		ResourceDraft:     ExtractCodeName(getPropertyString(currentData, "type", "baseDomain", "resourceDraft")),
		ResourceType:      ExtractCodeName(getPropertyString(currentData, "type", "baseDomain", "resourceType")),
		ResourceQueryType: ExtractCodeName(getPropertyString(currentData, "type", "baseDomain", "resourceQueryType")),
	}

	if parent != nil {
		serviceDomain.PathParameters = append(serviceDomain.PathParameters, parent.PathParameters...)
		serviceDomain.Package = parent.Package
		serviceDomain.ContextName = parent.ContextName
	} else {
		if serviceDomain.ResourceType != "" {
			serviceDomain.Package = serviceDomain.ResourceType
		} else {
			serviceDomain.Package = generatePackageName(currentKey)
		}
	}

	if serviceDomain.ResourceType == "" {
		suffix := generateContextName(currentKey, parent)
		if suffix != serviceDomain.ContextName {
			serviceDomain.ContextName += suffix

		}
	} else {
		if strings.HasPrefix(serviceDomain.ResourceType, serviceDomain.ContextName) {
			serviceDomain.ContextName = serviceDomain.ResourceType
		} else {
			serviceDomain.ContextName += serviceDomain.ResourceType
		}
	}

	params := getParameters(currentData)
	for _, param := range params {
		serviceDomain.PathParameters = append(serviceDomain.PathParameters, param)
	}
	if len(serviceDomain.PathParameters) > 1 {
		panic(serviceDomain.PathParameters)
	}

	itemType := getTypeInformation(apiResource)
	log.Printf("Parsing service %s (%s)\n", serviceDomain.Path, itemType)

	if itemType == "base" {
		for _, value := range currentData {
			if !strings.HasPrefix(value.Key.(string), "/") {
				continue
			}
			child := createService(value, &serviceDomain)
			if child != nil {
				serviceDomain.methods = append(serviceDomain.methods, child.methods...)
			}
		}
		if len(serviceDomain.methods) == 0 {
			if parent == nil {
				method := createActionServiceMethods("", currentData, &serviceDomain)
				if method != nil {
					serviceDomain.methods = append(serviceDomain.methods, *method)
				}
			} else {
				method := createActionServiceMethods(currentKey, currentData, parent)
				if method != nil {
					serviceDomain.methods = append(serviceDomain.methods, *method)
				}
			}
		}
		return &serviceDomain
	}

	if itemType != "baseDomain" {
		log.Fatalf("expected type baseDomain for %s, got %s\n", currentKey, itemType)
	}

	for _, value := range apiResource.Value.(yaml.MapSlice) {
		if _, ok := value.Value.(yaml.MapSlice); !ok {
			continue
		}
		subItemValue := value.Value.(yaml.MapSlice)
		subItemType := getTypeInformation(value)
		subItemParams := getParameters(subItemValue)

		if subItemType == "baseDomain" {
			child := createService(value, &serviceDomain)
			if child != nil {
				serviceDomain.methods = append(serviceDomain.methods, child.methods...)
			}
			continue
		} else if len(subItemParams) > 0 {
			methods := createResourceMethods(value.Key.(string), subItemValue, &serviceDomain)
			serviceDomain.methods = append(serviceDomain.methods, methods...)
		} else {
			method := createDomainMethod(value, &serviceDomain)
			if method != nil {
				serviceDomain.methods = append(serviceDomain.methods, *method)
			}
		}
	}

	return &serviceDomain
}

func createDomainMethod(item yaml.MapItem, sd *ServiceDomain) *ServiceMethod {
	// get -> Query
	// post -> Create
	// actions
	key := item.Key.(string)
	value, ok := item.Value.(yaml.MapSlice)
	if !ok {
		return nil
	}

	var method *ServiceMethod
	switch key {
	case "get":
		method = &ServiceMethod{
			Name:       CreateCodeName(sd.ContextName + "Query"),
			Context:    sd.ContextName,
			Parameters: sd.PathParameters,
			HTTPMethod: "get",
			Path:       sd.Path,
			ReturnType: sd.ResourceQueryType,
			Type:       "query",
		}
	case "post":
		method = &ServiceMethod{
			Name:       CreateCodeName(sd.ContextName + "Create"),
			InputType:  sd.ResourceDraft,
			Context:    sd.ContextName,
			Parameters: sd.PathParameters,
			HTTPMethod: "post",
			Path:       sd.Path,
			ReturnType: sd.ResourceType,
			Type:       "create",
		}
		if sd.ResourceDraft == "CustomerDraft" {
			method.ReturnType = "CustomerSignInResult"
		}
	default:
		if strings.HasPrefix(key, "/") {
			method = createActionServiceMethods(key, value, sd)
		}
	}
	if method != nil {
		parseResourceData(value, method)
	}
	return method
}

func createActionServiceMethods(key string, data yaml.MapSlice, sd *ServiceDomain) *ServiceMethod {

	for _, value := range data {
		if _, ok := value.Value.(yaml.MapSlice); !ok {
			continue
		}
	}

	// Guess the http method to use for this action
	// TODO: is this even ok?
	var methodKey string
	if val := getPropertyValue(data, "post", "body"); val != nil {
		methodKey = "post"
	}
	if val := getPropertyValue(data, "get", "body"); val != nil {
		methodKey = "post"
	}
	if val := getPropertyValue(data, "post", "responses"); val != nil {
		methodKey = "post"
	}
	if val := getPropertyValue(data, "get", "responses"); val != nil {
		methodKey = "get"
	}

	methodSlice := getPropertyValue(data, methodKey)
	if methodSlice == nil {
		spew.Dump(data)
		log.Println("No method found for action")
		return nil
	}
	methodData := methodSlice.(yaml.MapSlice)

	name := CreateCodeName(sd.ContextName)
	if key != "" {

		// FIXME: Add (methodName) in raml specs?
		if sd.ContextName == "Store" && strings.HasSuffix(sd.Path, "/shipping-methods") {
			name += "ShippingMethodsFor"
		}
		name += CreateCodeName(key[1:])
	}
	method := &ServiceMethod{
		Name:       name,
		Context:    sd.ContextName,
		Parameters: sd.PathParameters,
		Path:       sd.Path + key,
		Type:       "action",
		HTTPMethod: methodKey,
		InputType:  getInputType(methodData, sd, name),
		ReturnType: getReturnType(methodData, sd),
	}
	parseResourceData(methodData, method)
	return method
}

func createResourceMethods(path string, source yaml.MapSlice, sd *ServiceDomain) (result []ServiceMethod) {
	methodName := CreateCodeName(getPropertyString(source, "(methodName)"))
	// resourceType := getPropertyString(source, "type", "baseResource", "resourceType")
	paramNames := getParameters(source)

	for _, subvalue := range source {
		key := subvalue.Key.(string)

		value, ok := subvalue.Value.(yaml.MapSlice)
		if !ok {
			continue
		}

		var method *ServiceMethod

		switch key {
		case "get":
			method = &ServiceMethod{
				Context:    sd.ContextName,
				HTTPMethod: "get",
				Name:       CreateCodeName(sd.ContextName + "Get" + methodName),
				Parameters: sd.PathParameters,
				Path:       sd.Path + path,
				ReturnType: sd.ResourceType,
				Type:       "get",
			}
			method.Parameters = append(method.Parameters, paramNames...)
			parseResourceData(value, method)
		case "post":
			method = &ServiceMethod{
				Context:    sd.ContextName,
				HTTPMethod: "post",
				Name:       CreateCodeName(sd.ContextName + "Update" + methodName),
				Parameters: sd.PathParameters,
				Path:       sd.Path + path,
				ReturnType: sd.ResourceType,
				Type:       "update",
			}
			method.Parameters = append(method.Parameters, paramNames...)
			parseResourceData(value, method)
		case "delete":
			method = &ServiceMethod{
				Context:    sd.ContextName,
				HTTPMethod: "delete",
				Name:       CreateCodeName(sd.ContextName + "Delete" + methodName),
				Parameters: sd.PathParameters,
				Path:       sd.Path + path,
				ReturnType: sd.ResourceType,
				Type:       "delete",
			}
			method.Parameters = append(method.Parameters, paramNames...)
			parseResourceData(value, method)
		default:
			if strings.HasPrefix(key, "/") {
				method = createActionServiceMethods(key, value, sd)
			}
		}

		if method != nil {
			result = append(result, *method)
		}
	}
	return result
}

func parseResourceData(val yaml.MapSlice, method *ServiceMethod) {
	isSection := getPropertyValue(val, "is")
	if isSection != nil {
		for _, traitInterface := range isSection.([]interface{}) {
			if traitName, ok := traitInterface.(string); ok {
				method.Traits = append(method.Traits, traitName)
			}
		}
	}
	if val != nil {
		method.Description = getPropertyString(val, "description")
	}

	method.ExampleData = getPropertyString(val, "responses", "200", "body", "application/json", "example")

	if val != nil {
		method.Description = getPropertyString(val, "description")
	}

	qparams := getPropertyValue(val, "queryParameters")
	if qparams != nil {
		for _, item := range qparams.(yaml.MapSlice) {
			obj := createRamlTypeAttribute(item.Key.(string), item.Value)
			method.QueryParameters = append(method.QueryParameters, obj)
		}
	}

	for _, name := range method.Traits {
		trait := getTrait(name)
		if trait == nil {
			continue
		}
		qparams := getPropertyValue(trait, "queryParameters")
		if qparams != nil {
			for _, item := range qparams.(yaml.MapSlice) {
				obj := createRamlTypeAttribute(item.Key.(string), item.Value)
				// traits are supposed to be optional
				obj.Optional = true
				exists := false
				for _, existing := range method.QueryParameters {
					if existing.Name == obj.Name {
						exists = true
						break
					}
				}
				if !exists {
					method.QueryParameters = append(method.QueryParameters, obj)
				}
			}
		}
	}

}

func getTypeInformation(item yaml.MapItem) string {

	val, ok := item.Value.(yaml.MapSlice)
	if !ok {
		return ""
	}

	typeResource := getPropertyValue(val, "type")
	if val, ok := typeResource.(string); ok {
		return val
	} else if val, ok := typeResource.(yaml.MapItem); ok {
		return val.Value.(string)
	} else if val, ok := typeResource.(yaml.MapSlice); ok {
		return val[0].Key.(string)
	}
	return "base"
}

func getInputType(source yaml.MapSlice, sd *ServiceDomain, name string) string {
	if value := getPropertyString(source, "body", "application/json", "type"); value != "" {
		return value
	}
	if value := getPropertyValue(source, "queryParameters"); value != nil {
		return name + "Input"
	}
	return sd.ResourceDraft
}

func getReturnType(source yaml.MapSlice, sd *ServiceDomain) string {
	if value := getPropertyString(source, "responses", "201", "body", "application/json", "type"); value != "" {
		return value
	}
	if value := getPropertyString(source, "responses", "200", "body", "application/json", "type"); value != "" {
		return value
	}
	return sd.ResourceType
}

func getParameters(source yaml.MapSlice) []string {
	paramSlice := getPropertyValue(source, "uriParameters")
	if value, ok := paramSlice.(yaml.MapSlice); ok {
		paramNames := []string{}
		for _, paramItem := range value {
			paramNames = append(paramNames, paramItem.Key.(string))
		}
		return paramNames
	}

	paramName := getPropertyString(source, "type", "baseResource", "uriParameterName")
	if paramName != "" {
		return []string{paramName}
	}

	paramSlice = getPropertyValue(source, "type", "baseResource", "uriParameters")
	if value, ok := paramSlice.(yaml.MapSlice); ok {
		paramNames := []string{}
		for _, paramItem := range value {
			paramNames = append(paramNames, paramItem.Key.(string))
		}
		return paramNames
	}
	return []string{}
}

func generateContextName(value string, parent *ServiceDomain) string {
	re := regexp.MustCompile(`/([^/]+)`)
	matches := re.FindStringSubmatch(value)

	if len(matches) < 2 {
		log.Fatalf("no valid path found")
	}

	switch matches[1] {
	case "in-store":
		return "Store"
	case "graphql":
		return "GraphQL"
	case "login":
		return "Login"
	case "me":
		return "My"
	}
	if parent != nil && parent.ContextName != "" {
		return parent.ContextName
	}
	return "dummy"
}

func generatePackageName(value string) string {
	re := regexp.MustCompile(`/([^/]+)`)
	matches := re.FindStringSubmatch(value)

	if len(matches) < 2 {
		log.Fatalf("no valid path found")
	}

	switch matches[1] {
	case "in-store":
		return "InStore"
	case "graphql":
		return "GraphQL"
	case "login":
		return "Login"
	case "me":
		return "me"
	}
	return "dummy"
}
