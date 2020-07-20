package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Update` function
func postResourceHTTPMethod(resource *RamlType, resourceService ResourceService, resourceMethod ResourceMethod, httpMethod ResourceHTTPMethod) (code *jen.Statement) {
	updateIdentifier := "ID"
	if strings.Contains(resourceMethod.MethodName, "Key") {
		updateIdentifier = "Key"
	}

	funcName := fmt.Sprintf("%sUpdate%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
	updateStructID := fmt.Sprintf("%sUpdate%sInput", resource.CodeName, strings.Title(resourceMethod.MethodName))
	updateObjectType := fmt.Sprintf("%sUpdateAction", resource.CodeName)

	c := jen.Commentf("%s is input for function %s", updateStructID, funcName).Line()
	c.Type().Id(updateStructID).Struct(
		jen.Id(updateIdentifier).String(),
		jen.Id("Version").Int(),
		jen.Id("Actions").Index().Id(updateObjectType),
	).Line()
	urlPath := fmt.Sprintf("%s%s", resourceService.BasePath, resourceMethod.Path)

	methodParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("input").Op("*").Id(updateStructID),
		jen.Id("opts").Op("...").Id("RequestOption"),
	)
	clientMethod := "Update"

	structReceiver := jen.Id("client").Op("*").Id("Client")
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
	returnParams := jen.List(
		jen.Id("result").Op("*").Id(resourceService.ResourceType),
		jen.Err().Error(),
	)

	c.Commentf("%s %s", funcName, description).Line()
	c.Func().Params(structReceiver).Id(funcName).Params(methodParams).Parens(returnParams).Block(
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),

		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		),
		jen.Id("err").Op("=").Id("client").Op(".").Id(clientMethod).Call(
			jen.Id("ctx"),
			jen.Qual("strings", "Replace").Call(jen.Lit(urlPath), jen.Lit(fmt.Sprintf("{%s}", resourceMethod.PathParameterName)), jen.Id("input").Op(".").Id(updateIdentifier), jen.Lit(1)),
			jen.Id("params"),
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
