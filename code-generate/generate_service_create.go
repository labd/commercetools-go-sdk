package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// Create the `<method>Create` function
func createResourceCode(resource *RamlType, resourceService ResourceService, urlPathName string) (code *jen.Statement) {
	createName := fmt.Sprintf("%sCreate", resource.CodeName)
	returnParams := jen.Id("client").Op("*").Id("Client")
	createParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("draft").Op("*").Id(resourceService.ResourceDraftType),
		jen.Id("opts").Op("...").Id("RequestOption"),
	)

	c := jen.Commentf("%s creates a new instance of type %s", createName, resourceService.ResourceType).Line()
	c.Func().Params(returnParams).Id(createName).Params(createParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceType), jen.Err().Error())).Block(
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),

		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		),

		// err = client.Create(ctx, ProductURLPath, params, draft, &result)
		jen.Id("err").Op("=").Id("client").Op(".").Id("Create").Call(
			jen.Id("ctx"),
			jen.Id(urlPathName),
			jen.Id("params"),
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
