package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// Generate the `<service>DeleteWithID` and `<service>DeleteWithKey` functions
func generateServiceDelete(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {
	deleteWithVersion := true
	// TODO: nasty hack / incomplete API def
	if method.Context == "ApiClient" {
		deleteWithVersion = false
	}

	extraParams := generateServiceArgumentCode(method)
	methodParamList := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}
	methodParamList = append(methodParamList, extraParams...)

	returnObject := objects[method.ReturnType]
	returnParams := createServiceReturnList(method, returnObject)

	setVersionParam := jen.Empty()
	if deleteWithVersion {
		methodParamList = append(methodParamList, jen.Id("version").Int())
		setVersionParam = jen.Id("params").Op(".").Id("Set").Call(
			jen.Lit("version"),
			jen.Qual("strconv", "Itoa").Call(jen.Id("version")))
	}

	setDataErasure := jen.Empty()
	if method.HasTrait("dataErasure") {
		methodParamList = append(methodParamList, jen.Id("dataErasure").Bool())
		setDataErasure = jen.Id("params").Op(".").Id("Set").Call(
			jen.Lit("dataErasure"),
			jen.Qual("strconv", "FormatBool").Call(jen.Id("dataErasure")))
	}
	methodParamList = append(methodParamList, jen.Id("opts").Op("...").Id("RequestOption"))

	methodParams := jen.List(methodParamList...)
	clientMethod := "delete"

	structReceiver := jen.Id("client").Op("*").Id("Client")

	description := fmt.Sprintf("for type %s", method.ReturnType)
	if method.Description != "" {
		description = method.Description
	}
	c := jen.Commentf("%s %s", method.Name, description).Line()
	c.Func().Params(structReceiver).Id(method.Name).Params(methodParams).Parens(returnParams).Block(
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),
		setVersionParam,
		setDataErasure,

		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		),

		jen.Id("endpoint").Op(":=").Add(generateServicePathCode(method)),
		jen.Id("err").Op("=").Id("client").Op(".").Id(clientMethod).Call(
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
