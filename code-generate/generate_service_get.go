package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Get` function
func getResourceHTTPMethod(resource *RamlType, resourceService ResourceService, resourceMethod ResourceMethod, httpMethod ResourceHTTPMethod) (code *jen.Statement) {
	funcName := fmt.Sprintf("%sGet%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
	structReceiver := jen.Id("client").Op("*").Id("Client")
	resourceIdentifier := createResourceIdentifier(resourceService, resourceMethod)
	methodParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id(resourceIdentifier.ArgName).String(),
		jen.Id("opts").Op("...").Id("RequestOption"),
	)

	description := fmt.Sprintf("for type %s", resourceService.ResourceType)
	if httpMethod.Description != "" {
		description = httpMethod.Description
	}
	c := jen.Commentf("%s %s", funcName, description).Line()
	c.Func().Params(structReceiver).Id(funcName).Params(methodParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),

		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		),

		resourceIdentifier.createEndpointCode(false),
		jen.Id("err").Op("=").Id("client").Op(".").Id(strings.Title(httpMethod.HTTPMethod)).Call(
			jen.Id("ctx"),
			jen.Id("endpoint"),
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
