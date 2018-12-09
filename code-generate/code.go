package main

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func generateFilename(name string) string {
	if name == "" {
		panic("No package name found")
	}
	return strcase.ToSnake(name) + ".go"
}

func generateCode(objects []RamlType) {
	items := map[string][]*RamlType{}

	for i := range objects {
		item := objects[i]
		items[item.Package] = append(items[item.Package], &item)
	}

	for pkg, packageObjects := range items {
		f := jen.NewFile("commercetools")
		f.HeaderComment("Automatically generated, do not edit")

		for _, object := range packageObjects {
			if object.asMap {
				generateMap(f, *object)
			} else if len(object.EnumValues) > 0 {
				generateEnum(f, *object)
			} else if object.TypeName == "string" {
				generateString(f, *object)
			} else {
				generateStruct(f, *object)
			}
		}

		filename := generateFilename(pkg)
		err := f.Save("commercetools/types_" + filename)
		if err != nil {
			panic(err)
		}
	}
}

func generateString(f *jen.File, object RamlType) {
	f.Type().Id(object.CodeName).String()
}

func generateMap(f *jen.File, object RamlType) {
	f.Type().Id(object.CodeName).Map(jen.String()).String()
}

func generateEnum(f *jen.File, object RamlType) {
	f.Type().Id(object.CodeName).String()

	// TODO
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, value := range object.EnumValues {
			name := strcase.ToCamel(value)
			g.Id(object.CodeName + name).Id(object.CodeName).Op("=").Lit(value)
		}
	})
}

func generateStruct(f *jen.File, object RamlType) {
	c := f.Line()

	discriminatorAttr := object.getDiscriminatorAttribute()

	fields := []jen.Code{}

	// // Embed the parent type in the struct
	// if object.TypeObject != nil {
	// 	field := jen.Op("*").Id(object.TypeObject.CodeName)
	// 	fields = append(fields, field)
	// }

	for _, attr := range object.getAllAttributes() {
		if discriminatorAttr != nil && attr.CodeName == discriminatorAttr.CodeName {
			continue
		}
		field := generateStructField(attr)
		fields = append(fields, field)
	}

	isInterface := false
	// TODO; check if 1 attribute and that is the discriminator
	if len(object.Attributes) == 1 {
		for _, attr := range object.Children {
			if attr.Discriminator != "" {
				isInterface = true
			}
		}
	}

	if object.Discriminator != "" || isInterface {
		c = c.Type().Id(object.InterfaceName).Interface().Line()
		c = c.Type().Id(object.CodeName).Struct(fields...)
	} else {
		c = c.Type().Id(object.CodeName).Struct(fields...)
	}

	// Create custom MarshalJson method
	if object.DiscriminatorValue != "" && discriminatorAttr != nil {
		// fmt.Println(discriminatorAttr.CodeName)

		c = c.Line()
		c = c.Func()

		// (obj *T)
		c = c.Parens(jen.Id("obj").Id(object.CodeName))

		// Function name
		c = c.Id("MarshalJSON").Params()

		// Return params
		c = c.Parens(jen.List(jen.Index().Id("byte"), jen.Id("error")))

		c = c.Block(
			jen.Type().Id("Alias").Id(object.CodeName),
			jen.Return(
				jen.Qual("encoding/json", "Marshal").Call(
					jen.Struct(
						jen.Id(discriminatorAttr.CodeName).Id("string").Tag(map[string]string{"json": discriminatorAttr.Name}),
						jen.Op("*").Id("Alias"),
					).Values(
						jen.Id(discriminatorAttr.CodeName).Op(":").Lit(object.DiscriminatorValue),
						jen.Id("Alias").Op(":").Parens(jen.Op("*").Id("Alias")).Parens(jen.Op("&").Id("obj")),
					),
				),
			),
		)
	}

	// Generate mapping function
	interfaceName := object.InterfaceName
	if object.Discriminator != "" {
		c = c.Line()
		c = c.Func().Id(object.CodeName+"DiscriminatorMapping").Params(jen.Id("input").Id(interfaceName)).Parens(jen.Id(interfaceName)).Block(
			jen.Id("discriminator").Op(":=").Id("input").Assert(jen.Map(jen.String()).Interface()).Index(jen.Lit(object.Discriminator)),
			jen.Switch(jen.Id("discriminator")).BlockFunc(func(g *jen.Group) {
				for _, child := range object.Children {
					g.Case(jen.Lit(child.DiscriminatorValue)).Block(
						jen.Id("new").Op(":=").Id(child.CodeName).Values(),
						jen.Qual("github.com/mitchellh/mapstructure", "Decode").Call(jen.Id("input"), jen.Op("&").Id("new")),
						jen.CustomFunc(jen.Options{}, func(g *jen.Group) {
							// HIER
							discriminatedAttributes := child.getDiscriminatedAttributes()
							for _, attr := range discriminatedAttributes {
								if attr.Many {
									stmt := jen.For(jen.Id("i").Op(":=").Range().Id("new").Dot(attr.CodeName)).Block(
										jen.Id("new").Dot(attr.CodeName).Index(jen.Id("i")).Op("=").Id(attr.TypeObject.CodeName + "DiscriminatorMapping").Call(jen.Id("new").Dot(attr.CodeName).Index(jen.Id("i"))),
									)
									g.Add(stmt.Line())
								} else {
									stmt := jen.If(jen.Id("new").Dot(attr.CodeName).Op("!=").Nil()).Block(
										jen.Id("new").Dot(attr.CodeName).Op("=").Id(attr.TypeObject.CodeName + "DiscriminatorMapping").Call(jen.Id("new").Dot(attr.CodeName)),
									)
									g.Add(stmt.Line())
								}
							}
						}),
						jen.Return(jen.Id("new")),
					)
				}
			}),
			jen.Return(jen.Nil()),
		)
	}

	discriminatedAttributes := object.getDiscriminatedAttributes()

	// Generate UnmarshalJSON method
	if len(discriminatedAttributes) > 0 {
		c = c.Line()
		c = c.Func()

		// (obj *T)
		c = c.Parens(jen.Id("obj").Op("*").Id(object.CodeName))

		// Function name
		c = c.Id("UnmarshalJSON").Params(jen.Id("data").Index().Id("byte"))

		// Return params
		c = c.Parens(jen.Id("error"))

		c = c.Block(
			jen.Type().Id("Alias").Id(object.CodeName),
			jen.If(
				jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(
					jen.Id("data"),
					jen.Parens(jen.Op("*").Id("Alias")).Parens(jen.Id("obj")),
				),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			),
			jen.CustomFunc(jen.Options{}, func(g *jen.Group) {
				for _, attr := range discriminatedAttributes {
					if strings.Contains(attr.CodeName, "/") {
						// Apply on each value in struct

					} else if attr.Many {
						stmt := jen.For(jen.Id("i").Op(":=").Range().Id("obj").Dot(attr.CodeName)).Block(
							jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")).Op("=").Id(attr.TypeObject.CodeName + "DiscriminatorMapping").Call(jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i"))),
						)
						g.Add(stmt.Line())
					} else if attr.TypeObject != nil {
						stmt := jen.If(jen.Id("obj").Dot(attr.CodeName).Op("!=").Nil()).Block(
							jen.Id("obj").Dot(attr.CodeName).Op("=").Id(attr.TypeObject.CodeName + "DiscriminatorMapping").Call(jen.Id("obj").Dot(attr.CodeName)),
						)
						g.Add(stmt.Line())
					} else if attr.ItemsObject != nil {
						stmt := jen.For(jen.Id("i").Op(":=").Range().Id("obj").Dot(attr.CodeName)).Block(
							jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")).Op("=").Id(attr.ItemsObject.CodeName + "DiscriminatorMapping").Call(jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i"))),
						)
						g.Add(stmt.Line())
					}
				}
			}),
			jen.Return(jen.Nil()),
		)
	}

	// If the CodeName ends with Error then add an error method
	if strings.HasSuffix(object.CodeName, "Error") {
		c = c.Line().Line()
		c.Func().Parens(jen.Id("e").Id(object.CodeName)).Id("Error").Params().String().Block(
			jen.Return(jen.Id("e").Dot("Message")),
		)
	}

	if object.CodeName == "ErrorResponse" {
		c = c.Line().Line()
		c.Func().Parens(jen.Id("e").Id(object.CodeName)).Id("Error").Params().String().Block(
			jen.Return(jen.Id("e").Dot("Message")),
		)
	}

	// fmt.Printf("%#v\n", f)
}

func generateStructField(object RamlTypeAttribute) jen.Code {
	if strings.HasPrefix(object.Name, "/") {
		return jen.Line()
	}
	code := jen.Id(object.CodeName)

	if object.Many {
		code = code.Index()
	}

	if object.TypeObject != nil {
		if object.TypeObject.Discriminator == "" {

			if !object.Many && len(object.TypeObject.EnumValues) == 0 && object.TypeObject.TypeName != "string" {
				code = code.Op("*")
			}
			code = code.Id(object.TypeObject.CodeName)
		} else {
			code = code.Id(object.TypeObject.InterfaceName)
		}
	} else {
		switch object.TypeName {
		case "string":
			code = code.String()
		case "boolean":
			code = code.Bool()
		case "number":
			if object.Format == "int64" {
				code = code.Int()
			} else {
				code = code.Float64()
			}
		case "integer":
			code = code.Int()
		case "int64":
			code = code.Int()
		case "array":
			if object.ItemsObject != nil {
				if object.ItemsObject.Discriminator == "" {
					code = code.Index().Id(object.ItemsObject.CodeName)
				} else {
					code = code.Index().Id(object.ItemsObject.InterfaceName)
				}
			} else {
				code = code.Index()
				switch object.Items {
				case "string":
					code = code.String()
				case "number":
					code = code.Float64()
				default:
					code = code.Interface()
				}
			}
		case "datetime":
			code = code.Qual("time", "Time")
		default:
			code = code.Interface()
		}
	}

	jsonTags := []string{object.Name}
	if object.Optional {
		jsonTags = append(jsonTags, "omitempty")
	}

	code = code.Tag(map[string]string{"json": strings.Join(jsonTags, ",")})
	return code
}
