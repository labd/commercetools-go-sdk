package main

import (
	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Query` function
func generateServiceQuery(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {
	structReceiver := jen.Id("client").Op("*").Id("Client")
	methodParamList := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}

	extraParams := generateServiceArgumentCode(method)
	methodParamList = append(methodParamList, extraParams...)
	methodParamList = append(methodParamList,
		jen.Id("input").Op("*").Id("QueryInput"),
	)

	returnObject := objects[method.ReturnType]
	returnParams := createServiceReturnList(method, returnObject)

	c := jen.Commentf("%s allows querying for type %s", method.Name, method.Domain.ResourceType).Line()
	if method.Description != "" {
		c.Comment(method.Description).Line()
	}

	c.Func().Params(structReceiver).Id(method.Name).Params(jen.List(methodParamList...)).Parens(returnParams).Block(
		jen.Id("endpoint").Op(":=").Add(generateServicePathCode(method)),
		jen.Id("err").Op("=").Id("client").Op(".").Id("query").Call(
			jen.Id("ctx"),
			jen.Id("endpoint"),
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
