package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Get` function
func generateServiceGet(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {
	structReceiver := jen.Id("client").Op("*").Id("Client")

	// Generate function parameters
	extraParams := generateServiceArgumentCode(method)
	methodParams := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}
	methodParams = append(methodParams, extraParams...)
	methodParams = append(
		methodParams,
		jen.Id("opts").Op("...").Id("RequestOption"),
	)
	returnObject := objects[method.ReturnType]
	returnParams := createServiceReturnList(method, returnObject)

	description := fmt.Sprintf("for type %s", method.ReturnType)
	if method.Description != "" {
		description = method.Description
	}
	c := jen.Commentf("%s %s", method.Name, description).Line()
	c.Func().Params(structReceiver).Id(method.Name).Params(jen.List(methodParams...)).Parens(returnParams).BlockFunc(func(d *jen.Group) {

		d.Id("params").Op(":=").Qual("net/url", "Values").Block()

		// for _, opt := range opts {
		// 	opt(&params)
		// }
		d.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		)

		d.Id("endpoint").Op(":=").Add(generateServicePathCode(method))
		d.Id("err").Op("=").Id("client").Op(".").Id("get").Call(
			jen.Id("ctx"),
			jen.Id("endpoint"),
			jen.Id("params"),
			jen.Op("&").Id("result"),
		)
		d.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		)

		if returnObject.isInterface() {
			d.Return(jen.Id(returnObject.discriminatorFunctionName()).Call(jen.Id("result")))
		} else {
			d.Return(jen.Id("result"), jen.Nil())
		}
	}).Line()
	return c
}
