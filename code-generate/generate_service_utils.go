package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

func createServiceReturnList(method ServiceMethod, ro RamlType) jen.Code {
	if ro.isInterface() {
		return jen.List(
			jen.Id("result").Id(method.ReturnType),
			jen.Err().Error(),
		)
	}
	return jen.List(
		jen.Id("result").Op("*").Id(method.ReturnType),
		jen.Err().Error(),
	)
}

func generateServicePathCode(method ServiceMethod) jen.Code {
	// Replace placeholders with %s.
	// E.g.  "/in-store/key={storeKey}/carts/{ID}" => "/in-store/key=%s/carts/%s"
	var paramCodes []jen.Code

	var parameters []string
	if method.Type == "update" {
		parameters = append(parameters, method.Parameters[:len(method.Parameters)-1]...)
		lastParam := CreateCodeName(method.Parameters[len(method.Parameters)-1])
		parameters = append(parameters, "input."+lastParam)
	} else {
		parameters = method.Parameters
	}

	path := method.Path
	for _, param := range method.Parameters {
		path = strings.Replace(path, fmt.Sprintf("{%s}", param), "%s", 1)
	}

	for _, param := range parameters {
		if param == "ID" {
			param = "id"
		}
		paramCodes = append(paramCodes, jen.Id(param))
	}

	if len(paramCodes) < 1 {
		return jen.Lit(path[1:])
	}
	return jen.Qual("fmt", "Sprintf").Call(jen.Lit(path[1:]), jen.List(paramCodes...))

}

func generateServiceArgumentCode(method ServiceMethod) []jen.Code {
	var paramCodes []jen.Code
	for _, param := range method.Parameters {
		if param == "ID" {
			param = "id"
		}
		paramCodes = append(paramCodes, jen.Id(param).String())
	}
	return paramCodes
}
