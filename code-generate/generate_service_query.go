package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// Create the `<method>Query` function
func queryResourceCode(resource *RamlType, resourceService ResourceService, urlPathName string) (code *jen.Statement) {
	queryName := fmt.Sprintf("%sQuery", resource.CodeName)
	returnParams := jen.Id("client").Op("*").Id("Client")
	queryParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("input").Op("*").Id("QueryInput"),
	)

	c := jen.Commentf("%s allows querying for type %s", queryName, resourceService.ResourceType).Line()
	c.Func().Params(returnParams).Id(queryName).Params(queryParams).Parens(jen.List(jen.Id("result").Op("*").Id(resourceService.ResourceQueryType), jen.Err().Error())).Block(
		jen.Id("err").Op("=").Id("client").Op(".").Id("Query").Call(
			jen.Id("ctx"),
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
