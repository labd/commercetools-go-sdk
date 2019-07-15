package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func generateFilename(name string) string {
	if name == "" {
		panic("No package name found")
	}
	return strcase.ToSnake(name) + ".go"
}

// Entry point to generate services
func generateResources(objects []RamlType, resources []ResourceService) {
	typeMapping := map[string]*RamlType{}

	for i := range objects {
		item := objects[i]
		typeMapping[item.Name] = &item
	}

	for _, resourceService := range resources {
		f := jen.NewFile("commercetools")
		f.HeaderComment("Automatically generated, do not edit")

		resource := typeMapping[resourceService.ResourceType]

		if resource == nil {
			continue
		}
		// Skip this because it has a weird RAML definition
		if resource.CodeName == "CustomObject" {
			continue
		}
		// TODO: ugly hack
		resourceService.ResourceType = CreateCodeName(resourceService.ResourceType)
		resourceService.ResourceDraftType = CreateCodeName(resourceService.ResourceDraftType)
		resourceService.ResourceQueryType = CreateCodeName(resourceService.ResourceQueryType)

		urlPathName := fmt.Sprintf("%sURLPath", resource.CodeName)
		c := jen.Comment(fmt.Sprintf("%s is the commercetools API path.", urlPathName)).Line()
		c.Const().Id(urlPathName).Op("=").Lit(resourceService.BasePath).Line()
		f.Add(c)

		if resourceService.HasCreate {
			createCode := createResourceCode(resource, resourceService, urlPathName)
			f.Add(createCode)
		}

		if resourceService.HasList {
			queryCode := queryResourceCode(resource, resourceService, urlPathName)
			f.Add(queryCode)
		}

		httpMethodsCode := generateResourceHTTPMethodsCode(resource, resourceService, urlPathName)
		for _, httpMethodCode := range httpMethodsCode {
			f.Add(httpMethodCode)
		}

		filename := generateFilename(resource.Package)
		err := f.Save("commercetools/service_" + filename)
		if err != nil {
			panic(err)
		}
	}
}

func deleteResourceHTTPMethod(resource *RamlType, resourceService ResourceService, resourceMethod ResourceMethod, httpMethod ResourceHTTPMethod) (code *jen.Statement) {
	deleteIdentifier := "ID"
	if strings.Contains(resourceMethod.MethodName, "Key") {
		deleteIdentifier = "key"
	}

	methodName := fmt.Sprintf("%sDelete%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
	urlPath := fmt.Sprintf("%s%s", resourceService.BasePath, resourceMethod.Path)

	deleteWithVersion := true
	// TODO: nasty hack / incomplete API def
	if resourceService.ResourceType == "APIClient" {
		deleteWithVersion = false
	}

	methodIdentifierParam := jen.Id(resourceMethod.PathParameterName).String()
	methodParams := methodIdentifierParam
	setVersionParam := jen.Empty()
	if deleteWithVersion {
		methodParams = jen.List(methodIdentifierParam, jen.Id("version").Int())
		setVersionParam = jen.Id("params").Op(".").Id("Set").Call(jen.Lit("version"), jen.Qual("strconv", "Itoa").Call(jen.Id("version")))
	}
	setDataErasure := jen.Empty()
	if httpMethod.HasTrait("dataErasure") {
		// TODO: nasty hack, assume version is also input
		methodParams = jen.List(methodIdentifierParam, jen.Id("version").Int(), jen.Id("dataErasure").Bool())

		setDataErasure = jen.Id("params").Op(".").Id("Set").Call(jen.Lit("dataErasure"), jen.Qual("strconv", "FormatBool").Call(jen.Id("dataErasure")))
	}
	clientMethod := "Delete"

	returnParams := jen.Id("client").Op("*").Id("Client")

	description := fmt.Sprintf("for type %s", resourceService.ResourceType)
	if httpMethod.Description != "" {
		description = httpMethod.Description
	}
	c := jen.Commentf("%s %s", methodName, description).Line()
	c.Func().Params(returnParams).Id(methodName).Params(methodParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),
		setVersionParam,
		setDataErasure,
		jen.Id("err").Op("=").Id("client").Op(".").Id(clientMethod).Call(
			jen.Qual("strings", "Replace").Call(jen.Lit(urlPath), jen.Lit(fmt.Sprintf("{%s}", resourceMethod.PathParameterName)), jen.Id(deleteIdentifier), jen.Lit(1)),
			jen.Id("params"),
			jen.Op("&").Id("result"),
		),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Id("result"), jen.Nil()),
	).Line()

	return c
}

func postResourceHTTPMethod(resource *RamlType, resourceService ResourceService, resourceMethod ResourceMethod, httpMethod ResourceHTTPMethod) (code *jen.Statement) {
	updateIdentifier := "ID"
	if strings.Contains(resourceMethod.MethodName, "Key") {
		updateIdentifier = "Key"
	}

	methodName := fmt.Sprintf("%sUpdate%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
	updateStructID := fmt.Sprintf("%sUpdate%sInput", resource.CodeName, strings.Title(resourceMethod.MethodName))
	updateObjectType := fmt.Sprintf("%sUpdateAction", resource.CodeName)
	// TODO: Inconsistent api def... :(
	if updateObjectType == "InventoryEntryUpdateAction" {
		updateObjectType = "InventoryUpdateAction"
	}

	c := jen.Commentf("%s is input for function %s", updateStructID, methodName).Line()
	c.Type().Id(updateStructID).Struct(
		jen.Id(updateIdentifier).String(),
		jen.Id("Version").Int(),
		jen.Id("Actions").Index().Id(updateObjectType),
	).Line()
	urlPath := fmt.Sprintf("%s%s", resourceService.BasePath, resourceMethod.Path)

	methodParams := jen.Id("input").Op("*").Id(updateStructID)
	clientMethod := "Update"

	returnParams := jen.Id("client").Op("*").Id("Client")
	/*
		TODO: add this sanity check
		if input.ID == "" {
			return nil, fmt.Errorf("no valid type id passed")
		}
	*/
	description := fmt.Sprintf("for type %s", resourceService.ResourceType)
	if httpMethod.Description != "" {
		description = httpMethod.Description
	}
	c.Commentf("%s %s", methodName, description).Line()
	c.Func().Params(returnParams).Id(methodName).Params(methodParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("err").Op("=").Id("client").Op(".").Id(clientMethod).Call(
			jen.Qual("strings", "Replace").Call(jen.Lit(urlPath), jen.Lit(fmt.Sprintf("{%s}", resourceMethod.PathParameterName)), jen.Id("input").Op(".").Id(updateIdentifier), jen.Lit(1)),
			jen.Nil(),
			jen.Id("input").Op(".").Id("Version"),
			jen.Id("input").Op(".").Id("Actions"),
			jen.Op("&").Id("result"),
		),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Id("result"), jen.Nil()),
	).Line()

	return c
}

func getResourceHTTPMethod(resource *RamlType, resourceService ResourceService, resourceMethod ResourceMethod, httpMethod ResourceHTTPMethod) (code *jen.Statement) {
	methodName := fmt.Sprintf("%sGet%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
	returnParams := jen.Id("client").Op("*").Id("Client")
	methodParams := jen.Id(resourceMethod.PathParameterName).String()

	urlPath := fmt.Sprintf("%s%s", resourceService.BasePath, resourceMethod.Path)

	description := fmt.Sprintf("for type %s", resourceService.ResourceType)
	if httpMethod.Description != "" {
		description = httpMethod.Description
	}
	c := jen.Commentf("%s %s", methodName, description).Line()
	c.Func().Params(returnParams).Id(methodName).Params(methodParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("err").Op("=").Id("client").Op(".").Id(strings.Title(httpMethod.HTTPMethod)).Call(
			jen.Qual("strings", "Replace").Call(jen.Lit(urlPath), jen.Lit(fmt.Sprintf("{%s}", resourceMethod.PathParameterName)), jen.Id(resourceMethod.PathParameterName), jen.Lit(1)),
			jen.Nil(),
			jen.Op("&").Id("result"),
		),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Id("result"), jen.Nil()),
	).Line()
	return c
}

func generateResourceHTTPMethodsCode(resource *RamlType, resourceService ResourceService, urlPathName string) (httpMethodsCode []*jen.Statement) {
	httpMethodsCode = make([]*jen.Statement, 0)
	for _, resourceMethod := range resourceService.ResourceMethods {
		for _, httpMethod := range resourceMethod.HTTPMethods {
			switch httpMethod.HTTPMethod {
			case "get":
				c := getResourceHTTPMethod(resource, resourceService, resourceMethod, httpMethod)
				httpMethodsCode = append(httpMethodsCode, c)
			case "post":
				c := postResourceHTTPMethod(resource, resourceService, resourceMethod, httpMethod)
				httpMethodsCode = append(httpMethodsCode, c)
			case "delete":
				c := deleteResourceHTTPMethod(resource, resourceService, resourceMethod, httpMethod)
				httpMethodsCode = append(httpMethodsCode, c)
			default:
				continue

			}

		}
	}

	return httpMethodsCode
}

func queryResourceCode(resource *RamlType, resourceService ResourceService, urlPathName string) (code *jen.Statement) {
	queryName := fmt.Sprintf("%sQuery", resource.CodeName)
	returnParams := jen.Id("client").Op("*").Id("Client")
	queryParams := jen.Id("input").Op("*").Id("QueryInput")

	c := jen.Commentf("%s allows querying for type %s", queryName, resourceService.ResourceType).Line()
	c.Func().Params(returnParams).Id(queryName).Params(queryParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceQueryType), jen.Err().Error())).Block(
		jen.Id("err").Op("=").Id("client").Op(".").Id("Query").Call(
			jen.Id(urlPathName),
			jen.Id("input").Op(".").Id("toParams").Call(),
			jen.Op("&").Id("result"),
		),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Id("result"), jen.Nil()),
	).Line()
	return c
}

func createResourceCode(resource *RamlType, resourceService ResourceService, urlPathName string) (code *jen.Statement) {
	createName := fmt.Sprintf("%sCreate", resource.CodeName)
	returnParams := jen.Id("client").Op("*").Id("Client")
	createParams := jen.Id("draft").Op("*").Id(resourceService.ResourceDraftType)

	c := jen.Commentf("%s creates a new instance of type %s", createName, resourceService.ResourceType).Line()
	c.Func().Params(returnParams).Id(createName).Params(createParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("err").Op("=").Id("client").Op(".").Id("Create").Call(
			jen.Id(urlPathName),
			jen.Nil(),
			jen.Id("draft"),
			jen.Op("&").Id("result"),
		),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Id("result"), jen.Nil()),
	).Line()
	return c
}

// Entry point to generate code object for a specific RamlType
func generateCode(objects []RamlType) {
	items := map[string][]*RamlType{}

	for i := range objects {
		item := objects[i]
		items[item.Package] = append(items[item.Package], &item)
	}

	for pkg, packageObjects := range items {
		f := jen.NewFile("commercetools")
		f.HeaderComment("Automatically generated, do not edit")

		enumObjects := map[string]jen.Code{}
		mapObjects := map[string]jen.Code{}
		stringObjects := map[string]jen.Code{}
		structObjects := map[string]jen.Code{}

		for _, object := range packageObjects {
			var stmt *jen.Statement
			if object.asMap {
				stmt = generateMap(*object)
				mapObjects[object.CodeName] = stmt
			} else if len(object.EnumValues) > 0 {
				stmt := generateEnum(*object)
				enumObjects[object.CodeName] = stmt
			} else if object.TypeName == "string" {
				stmt := generateString(*object)
				stringObjects[object.CodeName] = stmt
			} else {
				stmt = generateStruct(*object)
				structObjects[object.CodeName] = stmt
			}
		}

		addCodeObjects(f, enumObjects)
		addCodeObjects(f, stringObjects)
		addCodeObjects(f, mapObjects)
		addCodeObjects(f, structObjects)

		filename := generateFilename(pkg)
		err := f.Save("commercetools/types_" + filename)
		if err != nil {
			panic(err)
		}
	}
}

func addCodeObjects(f *jen.File, objects map[string]jen.Code) {
	keys := []string{}
	for key := range objects {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		f.Add(objects[key])
	}
}

func generateString(object RamlType) *jen.Statement {
	c := jen.Comment(fmt.Sprintf("%s is of type string", object.CodeName)).Line()
	return c.Type().Id(object.CodeName).String().Line()
}

func generateMap(object RamlType) *jen.Statement {
	c := jen.Comment(fmt.Sprintf("%s is a map", object.CodeName)).Line()
	mapType := c.Type().Id(object.CodeName).Map(jen.String())
	if len(object.Attributes) != 1 {
		mapType = mapType.String()
	} else {
		attribute := object.Attributes[0]
		if attribute.TypeName == "string" {
			mapType = mapType.String()
		} else {
			mapType = mapType.Interface()
		}
	}

	return mapType.Line()
}

func generateEnum(object RamlType) *jen.Statement {
	c := jen.Comment(fmt.Sprintf("%s is an enum type", object.CodeName)).Line()
	c = c.Type().Id(object.CodeName).String().Line()

	// TODO
	c = c.Comment(fmt.Sprintf("Enum values for %s", object.CodeName)).Line()
	c.Const().DefsFunc(func(g *jen.Group) {
		for _, value := range object.EnumValues {
			name := strcase.ToCamel(value)
			g.Id(object.CodeName + name).Id(object.CodeName).Op("=").Lit(value)
		}
	})

	return c.Line()
}

func generateStruct(object RamlType) *jen.Statement {
	discriminatorAttr := object.getDiscriminatorAttribute()

	fields := []jen.Code{}

	// // Embed the parent type in the struct
	// if object.TypeObject != nil {
	// 	field := jen.Op("*").Id(object.TypeObject.CodeName)
	// 	fields = append(fields, field)
	// }

	for _, attr := range object.getAllAttributes() {
		if discriminatorAttr != nil && attr.CodeName == discriminatorAttr.CodeName {
			continue
		}
		field := generateStructField(attr)
		fields = append(fields, field)
	}

	c := jen.Null()

	if object.Discriminator != "" || object.isInterface() {
		if discriminatorAttr != nil {
			c = c.Comment(fmt.Sprintf("%s uses %s as discriminator attribute", object.InterfaceName, discriminatorAttr.Name)).Line()
		}
		c = c.Type().Id(object.InterfaceName).Interface().Line()

		if len(fields) > 0 && !strings.HasPrefix(object.CodeName, "Abstract") {
			c = c.Type().Id(object.CodeName).Struct(fields...)
		}
	} else {
		if object.TypeObject != nil {
			if object.TypeObject.InterfaceName != object.TypeObject.CodeName {
				c = c.Comment(fmt.Sprintf("%s implements the interface %s", object.CodeName, object.TypeObject.InterfaceName)).Line()
			} else {
				c = c.Comment(fmt.Sprintf("%s is of type %s", object.CodeName, object.TypeObject.CodeName)).Line()
			}
		} else {
			c = c.Comment(fmt.Sprintf("%s is a standalone struct", object.CodeName)).Line()
		}

		c = c.Type().Id(object.CodeName).Struct(fields...)
	}

	// Create custom MarshalJson method
	if object.DiscriminatorValue != "" && discriminatorAttr != nil {
		marshalJSON := generateStructMarshalJSON(&object, discriminatorAttr)
		c.Add(marshalJSON)
	}
	if object.Discriminator != "" {
		mappingFunc := generateStructDiscriminatorMapping(&object)
		c.Add(mappingFunc)
	}

	// Generate UnmarshalJSON method
	if len(fields) > 0 && !strings.HasPrefix(object.CodeName, "Abstract") {
		discriminatedAttributes := object.getDiscriminatedAttributes()
		if len(discriminatedAttributes) > 0 {
			unmarshalJSON := generateStructUnmarshalJSON(&object, discriminatedAttributes)
			c.Add(unmarshalJSON)
		}
	}

	// If the CodeName ends with Error then add an error method
	if strings.HasSuffix(object.CodeName, "Error") {
		c = c.Line().Line()
		c.Func().Parens(jen.Id("obj").Id(object.CodeName)).Id("Error").Params().String().Block(
			jen.Return(jen.Id("obj").Dot("Message")),
		)
	}

	if object.CodeName == "ErrorResponse" {
		c = c.Line().Line()
		c.Func().Parens(jen.Id("obj").Id(object.CodeName)).Id("Error").Params().String().Block(
			jen.Return(jen.Id("obj").Dot("Message")),
		)
	}

	return c.Line()
}

// Generate the *DiscriminatorMapping function which translates a typed json
// object to a specific struct based on the discriminator field
func generateStructDiscriminatorMapping(object *RamlType) *jen.Statement {

	interfaceName := object.InterfaceName
	c := jen.Line()
	c = c.Func().Id(object.discriminatorFunctionName()).Params(jen.Id("input").Interface()).Parens(jen.List(jen.Id(interfaceName), jen.Error())).Block(

		jen.Var().Id("discriminator").String(),

		jen.If(
			jen.List(jen.Id("data"), jen.Id("ok")).
				Op(":=").
				Id("input").Assert(jen.Map(jen.String()).Interface()),
			jen.Id("ok"),
		).Block(
			jen.List(jen.Id("discriminator"), jen.Id("ok")).
				Op("=").
				Id("data").Index(jen.Lit(object.Discriminator)).Assert(jen.String()),
			jen.If(jen.Op("!").Id("ok")).Block(
				jen.Return(jen.Nil(), jen.Qual("errors", "New").Call(
					jen.Lit(fmt.Sprintf("Error processing discriminator field '%s'", object.Discriminator)),
				)),
			),
		).Else().Block(
			jen.Return(jen.Nil(), jen.Qual("errors", "New").Call(
				jen.Lit("Invalid data"),
			)),
		),

		jen.Switch(jen.Id("discriminator")).BlockFunc(func(g *jen.Group) {
			for _, child := range object.Children {
				g.Case(jen.Lit(child.DiscriminatorValue)).Block(
					jen.Id("new").Op(":=").Id(child.CodeName).Values(),
					jen.Err().Op(":=").Qual("github.com/mitchellh/mapstructure", "Decode").Call(jen.Id("input"), jen.Op("&").Id("new")),
					jen.If(jen.Err().Op("!=").Nil().Block(jen.Return(jen.List(jen.Nil(), jen.Err())))),
					jen.CustomFunc(jen.Options{}, func(g *jen.Group) {
						discriminatedAttributes := child.getDiscriminatedAttributes()
						if len(discriminatedAttributes) == 0 {

						} else {
							for _, attr := range discriminatedAttributes {
								if attr.Many {
									stmt := jen.For(jen.Id("i").Op(":=").Range().Id("new").Dot(attr.CodeName)).Block(
										jen.List(jen.Id("new").Dot(attr.CodeName).Index(jen.Id("i")), jen.Err()).
											Op("=").
											Id(attr.TypeObject.discriminatorFunctionName()).
											Call(
												jen.Id("new").Dot(attr.CodeName).Index(jen.Id("i")),
											),
									)
									g.Add(stmt.Line())
								} else {
									stmt := jen.If(jen.Id("new").Dot(attr.CodeName).Op("!=").Nil()).Block(
										jen.List(jen.Id("new").Dot(attr.CodeName), jen.Err()).
											Op("=").
											Id(attr.TypeObject.discriminatorFunctionName()).
											Call(
												jen.Id("new").Dot(attr.CodeName),
											),
										jen.If(jen.Err().Op("!=").Nil().Block(
											jen.Return().List(jen.Nil(), jen.Err()),
										)),
									)
									g.Add(stmt.Line())
								}
							}
						}
						g.Return(jen.List(jen.Id("new"), jen.Nil()))
					}),
				)
			}
		}),

		jen.Return(jen.List(jen.Nil(), jen.Nil())),
	)

	return c
}

func generateStructMarshalJSON(object *RamlType, discriminatorAttr *RamlTypeAttribute) *jen.Statement {
	c := jen.Line()
	c = c.Comment(fmt.Sprintf("MarshalJSON override to set the discriminator value")).Line()
	c = c.Func()

	// (obj *T)
	c = c.Parens(jen.Id("obj").Id(object.CodeName))

	// Function name
	c = c.Id("MarshalJSON").Params()

	// Return params
	c = c.Parens(jen.List(jen.Index().Id("byte"), jen.Id("error")))

	c = c.Block(
		jen.Type().Id("Alias").Id(object.CodeName),
		jen.Return(
			jen.Qual("encoding/json", "Marshal").Call(
				jen.Struct(
					jen.Id(discriminatorAttr.CodeName).Id("string").Tag(map[string]string{"json": discriminatorAttr.Name}),
					jen.Op("*").Id("Alias"),
				).Values(
					jen.Id(discriminatorAttr.CodeName).Op(":").Lit(object.DiscriminatorValue),
					jen.Id("Alias").Op(":").Parens(jen.Op("*").Id("Alias")).Parens(jen.Op("&").Id("obj")),
				),
			),
		),
	)

	return c
}

// Generate the UnmarshalJSON method for the Structs
func generateStructUnmarshalJSON(object *RamlType, discriminatedAttributes []*RamlTypeAttribute) *jen.Statement {
	c := jen.Line()
	c = c.Comment("UnmarshalJSON override to deserialize correct attribute types based ").Line()
	c = c.Comment("on the discriminator value").Line()
	c = c.Func()

	// (obj *T)
	c = c.Parens(jen.Id("obj").Op("*").Id(object.CodeName))

	// Function name
	c = c.Id("UnmarshalJSON").Params(jen.Id("data").Index().Id("byte"))

	// Return params
	c = c.Parens(jen.Id("error"))

	c = c.Block(
		jen.Type().Id("Alias").Id(object.CodeName),
		jen.If(
			jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(
				jen.Id("data"),
				jen.Parens(jen.Op("*").Id("Alias")).Parens(jen.Id("obj")),
			),
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(jen.Err()),
		),
		jen.CustomFunc(jen.Options{}, func(g *jen.Group) {
			for _, attr := range discriminatedAttributes {
				if strings.Contains(attr.CodeName, "/") {
					// Apply on each value in struct
				} else if attr.Many {
					// Process regular many
					stmt := jen.For(jen.Id("i").Op(":=").Range().Id("obj").Dot(attr.CodeName)).Block(
						jen.Var().Err().Error(),
						jen.List(jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")), jen.Err()).
							Op("=").
							Id(attr.TypeObject.discriminatorFunctionName()).
							Call(
								jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")),
							),
						jen.If(jen.Err().Op("!=").Nil().Block(
							jen.Return(jen.Err()),
						)),
					)
					g.Add(stmt.Line())

				} else if attr.TypeObject != nil {
					// Process regular
					stmt := jen.If(jen.Id("obj").Dot(attr.CodeName).Op("!=").Nil()).Block(
						jen.Var().Err().Error(),
						jen.List(jen.Id("obj").Dot(attr.CodeName), jen.Err()).
							Op("=").
							Id(attr.TypeObject.discriminatorFunctionName()).
							Call(
								jen.Id("obj").Dot(attr.CodeName),
							),
						jen.If(jen.Err().Op("!=").Nil().Block(
							jen.Return(jen.Err()),
						)),
					)
					g.Add(stmt.Line())

				} else if attr.ItemsObject != nil {
					// Process Array type
					stmt := jen.For(jen.Id("i").Op(":=").Range().Id("obj").Dot(attr.CodeName)).Block(
						jen.Var().Err().Error(),
						jen.List(jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")), jen.Err()).
							Op("=").
							Id(attr.ItemsObject.discriminatorFunctionName()).
							Call(
								jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")),
							),
						jen.If(jen.Err().Op("!=").Nil().Block(
							jen.Return(jen.Err()),
						)),
					)
					g.Add(stmt.Line())
				}
			}
		}),
		jen.Return(jen.Nil()),
	)

	return c
}

func generateStructField(object RamlTypeAttribute) jen.Code {
	if strings.HasPrefix(object.Name, "/") {
		return jen.Line()
	}
	code := jen.Id(object.CodeName)

	if object.Many {
		code = code.Index()
	}

	if object.TypeObject != nil {
		if object.TypeObject.Discriminator == "" {

			if !object.Many && len(object.TypeObject.EnumValues) == 0 && object.TypeObject.TypeName != "string" {
				code = code.Op("*")
			}
			code = code.Id(object.TypeObject.CodeName)
		} else {
			code = code.Id(object.TypeObject.InterfaceName)
		}
	} else {
		switch object.TypeName {
		case "string":
			code = code.String()
		case "boolean":
			code = code.Bool()
		case "number":
			if object.Format == "int64" {
				code = code.Int()
			} else {
				if object.Minimum != nil && *object.Minimum == 0 {
					code = code.Op("*")
				}
				code = code.Float64()
			}
		case "integer":
			code = code.Int()
		case "int64":
			code = code.Int()
		case "array":
			if object.ItemsObject != nil {
				if object.ItemsObject.InterfaceName != object.ItemsObject.CodeName {
					code = code.Index().Id(object.ItemsObject.InterfaceName)
				} else {
					code = code.Index().Id(object.ItemsObject.CodeName)
				}
			} else {
				code = code.Index()
				switch object.Items {
				case "string":
					code = code.String()
				case "number":
					code = code.Float64()
				default:
					code = code.Interface()
				}
			}
		case "datetime":
			code = code.Qual("time", "Time")
		case "date-only":
			code = code.Id("Date")
		default:
			code = code.Interface()
		}
	}

	jsonTags := []string{object.Name}
	if object.Optional {
		if object.TypeName != "boolean" {
			jsonTags = append(jsonTags, "omitempty")
		}
	}

	code = code.Tag(map[string]string{"json": strings.Join(jsonTags, ",")})
	return code
}
