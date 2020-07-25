package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Action` function
func generateServiceAction(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {
	methodParamList := []jen.Code{
		jen.Id("ctx").Qual("context", "Context"),
	}
	extraParams := generateServiceArgumentCode(method)
	methodParamList = append(methodParamList, extraParams...)
	if method.InputType != "" {
		methodParamList = append(methodParamList, jen.Id("value").Op("*").Id(method.InputType))
	}
	methodParamList = append(methodParamList, jen.Id("opts").Op("...").Id("RequestOption"))

	var clientMethod string
	switch method.HTTPMethod {
	case "post":
		clientMethod = "create"
	case "get":
		clientMethod = "get"
	}

	structReceiver := jen.Id("client").Op("*").Id("Client")
	description := fmt.Sprintf("for type %s", method.InputType)
	if method.Description != "" {
		description = method.Description
	}

	if method.ReturnType == "" {
		log.Fatalf("Empty return type for %s", method.Name)
	}

	returnObject := objects[method.ReturnType]
	returnParams := createServiceReturnList(method, returnObject)

	// Generate the `type <resource>Input struct {}`
	c := jen.Empty()
	if len(method.QueryParameters) > 0 {
		c.Commentf("%s is input for function %s", method.InputType, method.Name).Line()
		c.Type().Id(method.InputType).StructFunc(func(d *jen.Group) {
			for _, param := range method.QueryParameters {

				line := jen.Empty()

				line.Id(param.CodeName)
				typeCode := getNativeTypeCode(param)

				if param.ItemsObject != nil {
					line.Map(jen.String()).Add(typeCode)
				} else {
					line.Add(typeCode)
				}

				// Generate the url tag
				urlTags := []string{param.Name}
				if param.Optional || param.TypeName == "boolean" {
					urlTags = append(urlTags, "omitempty")
				}
				tagVal := map[string]string{"url": strings.Join(urlTags, ",")}
				line.Tag(tagVal)

				d.Add(line)
			}
		}).Line()

	}

	c.Commentf("%s %s", method.Name, description).Line()
	c.Func().Params(structReceiver).Id(method.Name).Params(jen.List(methodParamList...)).Parens(returnParams).BlockFunc(func(b *jen.Group) {

		// params := url.Values{} | SerializeQueryParams(value)
		// for _, opt := range opts {
		// 	opt(&params)
		// }
		if len(method.QueryParameters) > 0 {
			b.Id("params").Op(":=").Id("serializeQueryParams").Call(jen.Id("value"))
		} else {
			b.Id("params").Op(":=").Qual("net/url", "Values").Block()
		}

		b.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		).Line()

		b.Id("endpoint").Op(":=").Add(generateServicePathCode(method))

		b.Id("err").Op("=").Id("client").Op(".").Id(clientMethod).CallFunc(func(g *jen.Group) {
			g.Id("ctx")
			g.Id("endpoint")
			g.Id("params")

			if method.HTTPMethod == "post" {
				if method.InputType != "" {
					g.Id("value")
				} else {
					g.Nil()
				}
			}

			g.Op("&").Id("result")
		})
		b.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		)
		b.Return(jen.Id("result"), jen.Nil())
	}).Line()

	return c
}
