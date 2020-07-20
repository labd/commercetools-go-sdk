package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
)

type ResourceIdentifier struct {
	ArgName         string
	PublicName      string
	endpointPattern string
}

func (ri *ResourceIdentifier) createEndpointCode(fromInput bool) jen.Code {
	var source jen.Code
	if fromInput {
		source = jen.Id("input").Op(".").Id(ri.PublicName)
	} else {
		source = jen.Id(ri.ArgName)
	}
	return jen.Id("endpoint").Op(":=").Qual("fmt", "Sprintf").Call(jen.Lit(ri.endpointPattern), source)
}

func createResourceIdentifier(resourceService ResourceService, resourceMethod ResourceMethod) ResourceIdentifier {
	ri := ResourceIdentifier{}

	ri.PublicName = "ID"
	if strings.Contains(resourceMethod.MethodName, "Key") {
		ri.PublicName = "Key"
	}
	ri.ArgName = strings.ToLower(ri.PublicName)

	if ri.PublicName == "ID" {
		ri.endpointPattern = fmt.Sprintf("%s/%%s", resourceService.BasePath)
	} else if ri.PublicName == "Key" {
		ri.endpointPattern = fmt.Sprintf("%s/key=%%s", resourceService.BasePath)
	} else {
		log.Fatalf("Unrecognized PublicName %v", ri.PublicName)
	}

	return ri
}
