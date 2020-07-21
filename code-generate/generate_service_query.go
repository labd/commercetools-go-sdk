package main

import (
	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Query` function
func generateServiceQuery(funcName string, resourceService ResourceService, urlPathName string) (code *jen.Statement) {
	structReceiver := jen.Id("client").Op("*").Id("Client")
	queryParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("input").Op("*").Id("QueryInput"),
	)
	returnParams := jen.List(
		jen.Id("result").Op("*").Id(resourceService.ResourceQueryType),
		jen.Err().Error())

	c := jen.Commentf("%s allows querying for type %s", funcName, resourceService.ResourceType).Line()
	c.Func().Params(structReceiver).Id(funcName).Params(queryParams).Parens(returnParams).Block(
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
