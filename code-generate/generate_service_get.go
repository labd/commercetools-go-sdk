package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Create the `<method>Get` function
func getResourceHTTPMethod(resource *RamlType, resourceService ResourceService, resourceMethod ResourceMethod, httpMethod ResourceHTTPMethod) (code *jen.Statement) {
	methodName := fmt.Sprintf("%sGet%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
	returnParams := jen.Id("client").Op("*").Id("Client")
	methodParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id(resourceMethod.PathParameterName).String(),
		jen.Id("opts").Op("...").Id("RequestOption"),
	)

	urlPath := fmt.Sprintf("%s%s", resourceService.BasePath, resourceMethod.Path)

	description := fmt.Sprintf("for type %s", resourceService.ResourceType)
	if httpMethod.Description != "" {
		description = httpMethod.Description
	}
	c := jen.Commentf("%s %s", methodName, description).Line()
	c.Func().Params(returnParams).Id(methodName).Params(methodParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),

		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		),

		jen.Id("err").Op("=").Id("client").Op(".").Id(strings.Title(httpMethod.HTTPMethod)).Call(
			jen.Id("ctx"),
			jen.Qual("strings", "Replace").Call(jen.Lit(urlPath), jen.Lit(fmt.Sprintf("{%s}", resourceMethod.PathParameterName)), jen.Id(resourceMethod.PathParameterName), jen.Lit(1)),
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
