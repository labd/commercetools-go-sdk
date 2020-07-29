package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

const pkgCommercetools = "github.com/labd/commercetools-go-sdk/commercetools"
const pkgTestUtil = "github.com/labd/commercetools-go-sdk/testutil"
const pkgAssert = "github.com/stretchr/testify/assert"

// Entry point to generate services
func generateTests(objects []RamlType, resources []ServiceDomain) {
	typeMapping := map[string]string{}

	for _, item := range objects {
		typeMapping[item.Name] = item.CodeName

		if item.Name == "Message" {
			fmt.Println(item)
		}
	}
	objectsMap := map[string]RamlType{}
	for _, item := range objects {
		objectsMap[item.Name] = item
	}

	for _, resourceService := range resources {

		if len(resourceService.methods) < 1 {
			log.Printf("! No methods found for domain %s\n", resourceService.ContextName)
			continue
		}

		if resourceService.ContextName == "CustomObject" {
			continue // Waiting on fix on api reference
		}

		filename := generateFilename(resourceService.Package, "_generated_test")
		log.Printf("Writing commercetools/%s\n", filename)

		f := jen.NewFile("commercetools_test")
		f.HeaderComment("Automatically generated, do not edit")
		f.ImportName("github.com/stretchr/testify/assert", "assert")
		f.ImportName("github.com/labd/commercetools-go-sdk/commercetools", "commercetools")
		f.ImportName("github.com/labd/commercetools-go-sdk/testutil", "testutil")

		contextNames, resourceMethods := orderServiceMethods(&resourceService, objectsMap)
		for _, contextName := range contextNames {
			value := resourceMethods[contextName]

			for _, method := range value.Get {
				code := generateTestsGet(method, objectsMap)
				f.Add(code)
			}
			for _, method := range value.Delete {
				code := generateTestsDelete(method, objectsMap)
				f.Add(code)
			}
			for _, method := range value.Query {
				code := generateTestsQuery(method, objectsMap)
				f.Add(code)
			}
		}

		err := f.Save("commercetools/service_" + filename)
		if err != nil {
			panic(err)
		}
	}
}

func jsonBlock(value string) jen.Code {
	lines := strings.SplitAfter(value, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	data := strings.Join(append([]string{""}, lines...), "\t")
	data = strings.Trim(data, "\n\t ")
	return jen.Op("`").Op(data).Op("`")
}

// Generate the `<service>Get` function
func generateTestsGet(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {

	methodParams := jen.List(
		jen.Id("t").Op("*").Qual("testing", "T"),
	)
	returnObject := objects[method.ReturnType]

	objName := strcase.ToSnake(method.ReturnType)
	switch objName {
	case "type":
		objName = "rType"
	}

	c := jen.Func().Id("TestGenerated" + method.Name).Params(methodParams).BlockFunc(func(d *jen.Group) {

		d.Id("responseData").Op(":=").Add(jsonBlock(method.ExampleData))
		d.List(jen.Id("client"), jen.Id("server")).Op(":=").Qual(pkgTestUtil, "MockClient").Call(
			jen.Id("t"),
			jen.Id("responseData"),
			jen.Nil(),
			jen.Nil(),
		)

		d.Defer().Id("server").Op(".").Id("Close").Call()

		d.List(jen.Id(objName), jen.Err()).Op(":=").Id("client").Op(".").Id(method.Name).CallFunc(func(d *jen.Group) {
			d.Qual("context", "TODO").Call()
			for range method.Parameters {
				d.Lit("dummy-id")
			}

		})

		d.If(jen.Err().Op("!=").Nil()).Block(
			jen.Id("t").Op(".").Id("Fatal").Call(jen.Err()),
		)

		d.Add(generateAssertStatements(objName, method, returnObject))

	}).Line()
	return c
}

func generateTestsDelete(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {

	methodParams := jen.List(
		jen.Id("t").Op("*").Qual("testing", "T"),
	)
	returnObject := objects[method.ReturnType]

	objName := strcase.ToSnake(method.ReturnType)
	switch objName {
	case "type":
		objName = "rType"
	}

	c := jen.Func().Id("TestGenerated" + method.Name).Params(methodParams).BlockFunc(func(d *jen.Group) {

		d.Id("responseData").Op(":=").Add(jsonBlock(method.ExampleData))
		d.List(jen.Id("client"), jen.Id("server")).Op(":=").Qual(pkgTestUtil, "MockClient").Call(
			jen.Id("t"),
			jen.Id("responseData"),
			jen.Nil(),
			jen.Nil(),
		)

		d.Defer().Id("server").Op(".").Id("Close").Call()

		d.List(jen.Id(objName), jen.Err()).Op(":=").Id("client").Op(".").Id(method.Name).CallFunc(func(d *jen.Group) {
			d.Qual("context", "TODO").Call()
			for range method.Parameters {
				d.Lit("dummy-id")
			}

			// TODO: nasty hack / incomplete API def
			if method.Context != "ApiClient" {
				d.Lit(1)
			}

			if method.HasTrait("dataErasure") {
				d.True()
			}

		})

		d.If(jen.Err().Op("!=").Nil()).Block(
			jen.Id("t").Op(".").Id("Fatal").Call(jen.Err()),
		)
		d.Add(generateAssertStatements(objName, method, returnObject))
	}).Line()
	return c
}

func generateTestsQuery(method ServiceMethod, objects map[string]RamlType) (code *jen.Statement) {
	const pkgTestUtil = "github.com/labd/commercetools-go-sdk/testutil"
	const pkgAssert = "github.com/stretchr/testify/assert"

	methodParams := jen.List(
		jen.Id("t").Op("*").Qual("testing", "T"),
	)
	returnObject := objects[method.ReturnType]

	objName := "queryResult"

	c := jen.Func().Id("TestGenerated" + method.Name).Params(methodParams).BlockFunc(func(d *jen.Group) {

		d.Id("responseData").Op(":=").Add(jsonBlock(method.ExampleData))
		d.List(jen.Id("client"), jen.Id("server")).Op(":=").Qual(pkgTestUtil, "MockClient").Call(
			jen.Id("t"),
			jen.Id("responseData"),
			jen.Nil(),
			jen.Nil(),
		)

		d.Defer().Id("server").Op(".").Id("Close").Call()

		d.Id("input").Op(":=").Qual(pkgCommercetools, "QueryInput").Block()

		d.List(jen.Id(objName), jen.Err()).Op(":=").Id("client").Op(".").Id(method.Name).CallFunc(func(d *jen.Group) {
			d.Qual("context", "TODO").Call()

			for range method.Parameters {
				d.Lit("dummy-id")
			}
			d.Op("&").Id("input")
		})

		d.If(jen.Err().Op("!=").Nil()).Block(
			jen.Id("t").Op(".").Id("Fatal").Call(jen.Err()),
		)
		d.Add(generateAssertStatements(objName, method, returnObject))
	}).Line()
	return c
}

func generateAssertStatements(objName string, method ServiceMethod, returnObject RamlType) *jen.Group {

	c := &jen.Group{}

	c.Qual(pkgAssert, "NotNil").Call(jen.Id("t"), jen.Id(objName)).Line()

	if returnObject.isInterface() {
		return c
	}

	for _, attr := range returnObject.getAllAttributes() {
		if ev := method.getExampleValue(attr.Name); ev != nil {
			if _, ok := ev.(map[string]interface{}); ok {
				continue
			}
			if _, ok := ev.([]interface{}); ok {
				continue
			}

			if _, ok := ev.(float64); ok {
				c.Qual(pkgAssert, "NotNil").Call(
					jen.Id("t"),
					jen.Id(objName).Op(".").Id(attr.CodeName),
				).Line()
			} else if rv, ok := ev.(bool); ok {
				if rv == true {
					c.Qual(pkgAssert, "True").Call(
						jen.Id("t"),
						jen.Id(objName).Op(".").Id(attr.CodeName),
					).Line()
				} else {
					c.Qual(pkgAssert, "False").Call(
						jen.Id("t"),
						jen.Id(objName).Op(".").Id(attr.CodeName),
					).Line()
				}
				continue
			} else {
				c.Qual(pkgAssert, "NotEmpty").Call(
					jen.Id("t"),
					jen.Id(objName).Op(".").Id(attr.CodeName),
				).Line()
			}
		}
	}
	return c
}
