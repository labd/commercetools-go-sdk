package main

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Update` function
func generateServiceUpdate(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {
	updateStructID := fmt.Sprintf("%sInput", method.Name)
	updateObjectType := fmt.Sprintf("%sUpdateAction", method.ReturnType)

	// Generate the `type <resource>With<ID>Input struct {}`
	identifier := CreateCodeName(method.Parameters[len(method.Parameters)-1])
	c := jen.Commentf("%s is input for function %s", updateStructID, method.Name).Line()
	c.Type().Id(updateStructID).Struct(
		jen.Id(identifier).String(),
		jen.Id("Version").Int(),
		jen.Id("Actions").Index().Id(updateObjectType),
	).Line()

	c.Func().Params(jen.Id("input").Op("*").Id(updateStructID)).Id("Validate").Params().Parens(jen.Id("error")).Block(

		// Validate if input identifier is empty
		jen.If(jen.Id("input").Op(".").Id(identifier).Op("==").Lit("")).Block(
			jen.Return(
				jen.Qual("fmt", "Errorf").Call(
					jen.Lit(fmt.Sprintf("no valid value for %s given", identifier)),
				),
			),
		),

		// Validate if update actions are empty
		jen.If(jen.Id("len").Call(jen.Id("input").Op(".").Id("Actions")).Op("==").Lit(0)).Block(
			jen.Return(
				jen.Qual("fmt", "Errorf").Call(
					jen.Lit("no update actions specified"),
				),
			),
		),

		jen.Return(jen.Nil()),
	).Line()

	extraParams := generateServiceArgumentCode(method)
	methodParamsList := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}
	methodParamsList = append(methodParamsList, extraParams[:len(extraParams)-1]...)
	methodParamsList = append(methodParamsList,
		jen.Id("input").Op("*").Id(updateStructID),
		jen.Id("opts").Op("...").Id("RequestOption"),
	)
	clientMethod := "update"

	structReceiver := jen.Id("client").Op("*").Id("Client")
	description := fmt.Sprintf("for type %s", method.ReturnType)
	// if httpMethod.Description != "" {
	// 	description = httpMethod.Description
	// }
	returnObject := objects[method.ReturnType]
	returnParams := createServiceReturnList(method, returnObject)

	c.Commentf("%s %s", method.Name, description).Line()
	c.Func().Params(structReceiver).Id(method.Name).Params(jen.List(methodParamsList...)).Parens(returnParams).Block(
		// if err := input.Validate(); err != nil {
		// 	return nil, err
		// }
		jen.If(
			jen.Err().Op(":=").Id("input").Op(".").Id("Validate").Call(),
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(jen.Nil(), jen.Err()),
		).Line(),

		// params := url.Values{}
		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		).Line(),

		jen.Id("endpoint").Op(":=").Add(generateServicePathCode(method)),
		jen.Id("err").Op("=").Id("client").Op(".").Id(clientMethod).Call(
			jen.Id("ctx"),
			jen.Id("endpoint"),
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
