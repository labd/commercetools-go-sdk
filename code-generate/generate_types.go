package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func addCodeObjects(f *jen.File, objects map[string]jen.Code) {
	keys := []string{}
	for key := range objects {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		f.Add(objects[key])
	}
}

func generateString(object RamlType) *jen.Statement {
	c := jen.Comment(fmt.Sprintf("%s is of type string", object.CodeName)).Line()
	return c.Type().Id(object.CodeName).String().Line()
}

func generateMap(object RamlType) *jen.Statement {
	c := jen.Comment(fmt.Sprintf("%s is a map", object.CodeName)).Line()
	mapType := c.Type().Id(object.CodeName).Map(jen.String())
	if len(object.Attributes) != 1 {
		mapType = mapType.String()
	} else {
		attribute := object.Attributes[0]
		if attribute.TypeName == "string" {
			mapType = mapType.String()
		} else {
			mapType = mapType.Interface()
		}
	}

	return mapType.Line()
}

func generateEnum(object RamlType) *jen.Statement {
	c := jen.Comment(fmt.Sprintf("%s is an enum type", object.CodeName)).Line()
	c = c.Type().Id(object.CodeName).String().Line()

	// TODO
	c = c.Comment(fmt.Sprintf("Enum values for %s", object.CodeName)).Line()
	c.Const().DefsFunc(func(g *jen.Group) {
		for _, value := range object.EnumValues {
			name := strcase.ToCamel(value)
			g.Id(object.CodeName + name).Id(object.CodeName).Op("=").Lit(value)
		}
	})

	return c.Line()
}

func generateStruct(object RamlType) *jen.Statement {
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

	c := jen.Null()

	if object.Discriminator != "" || object.isInterface() {
		if discriminatorAttr != nil {
			c = c.Comment(fmt.Sprintf("%s uses %s as discriminator attribute", object.InterfaceName, discriminatorAttr.Name)).Line()
		}
		c = c.Type().Id(object.InterfaceName).Interface().Line()

		if len(fields) > 0 && !object.isInterface() {
			c = c.Type().Id(object.CodeName).Struct(fields...)
		}
	} else {
		if object.TypeObject != nil {
			if object.TypeObject.isInterface() {
				c = c.Comment(fmt.Sprintf("%s implements the interface %s", object.CodeName, object.TypeObject.InterfaceName)).Line()
			} else {
				c = c.Comment(fmt.Sprintf("%s is of type %s", object.CodeName, object.TypeObject.CodeName)).Line()
			}
		} else {
			c = c.Comment(fmt.Sprintf("%s is a standalone struct", object.CodeName)).Line()
		}

		c = c.Type().Id(object.CodeName).Struct(fields...)
	}

	// Create custom MarshalJson method
	if object.DiscriminatorValue != "" && discriminatorAttr != nil {
		marshalJSON := generateStructMarshalJSON(&object, discriminatorAttr)
		c.Add(marshalJSON)
	}
	if object.Discriminator != "" {
		mappingFunc := generateStructDiscriminatorMapping(&object)
		c.Add(mappingFunc)
	}

	// Generate UnmarshalJSON method
	if len(fields) > 0 && !object.isInterface() {
		discriminatedAttributes := object.getDiscriminatedAttributes()
		if len(discriminatedAttributes) > 0 {
			unmarshalJSON := generateStructUnmarshalJSON(&object, discriminatedAttributes)
			c.Add(unmarshalJSON)
		}
	}

	// If the CodeName ends with Error then add an error method
	if strings.HasSuffix(object.CodeName, "Error") {
		c = c.Line().Line()
		c.Func().Parens(jen.Id("obj").Id(object.CodeName)).Id("Error").Params().String().Block(
			jen.Return(jen.Id("obj").Dot("Message")),
		)
	}

	if object.CodeName == "ErrorResponse" {
		c = c.Line().Line()
		c.Func().Parens(jen.Id("obj").Id(object.CodeName)).Id("Error").Params().String().Block(
			jen.Return(jen.Id("obj").Dot("Message")),
		)
	}

	return c.Line()
}

// Generate the *DiscriminatorMapping function which translates a typed json
// object to a specific struct based on the discriminator field
func generateStructDiscriminatorMapping(object *RamlType) *jen.Statement {

	interfaceName := object.InterfaceName
	c := jen.Line()
	c = c.Func().Id(object.discriminatorFunctionName()).Params(jen.Id("input").Interface()).Parens(jen.List(jen.Id(interfaceName), jen.Error())).Block(

		jen.Var().Id("discriminator").String(),

		jen.If(
			jen.List(jen.Id("data"), jen.Id("ok")).
				Op(":=").
				Id("input").Assert(jen.Map(jen.String()).Interface()),
			jen.Id("ok"),
		).Block(
			jen.List(jen.Id("discriminator"), jen.Id("ok")).
				Op("=").
				Id("data").Index(jen.Lit(object.Discriminator)).Assert(jen.String()),
			jen.If(jen.Op("!").Id("ok")).Block(
				jen.Return(jen.Nil(), jen.Qual("errors", "New").Call(
					jen.Lit(fmt.Sprintf("Error processing discriminator field '%s'", object.Discriminator)),
				)),
			),
		).Else().Block(
			jen.Return(jen.Nil(), jen.Qual("errors", "New").Call(
				jen.Lit("Invalid data"),
			)),
		),

		jen.Switch(jen.Id("discriminator")).BlockFunc(func(g *jen.Group) {
			for _, child := range object.Children {
				g.Case(jen.Lit(child.DiscriminatorValue)).Block(
					jen.Id("new").Op(":=").Id(child.CodeName).Values(),
					jen.Err().Op(":=").Id("decodeStruct").Call(jen.Id("input"), jen.Op("&").Id("new")),
					jen.If(jen.Err().Op("!=").Nil().Block(jen.Return(jen.List(jen.Nil(), jen.Err())))),
					jen.CustomFunc(jen.Options{}, func(g *jen.Group) {
						discriminatedAttributes := child.getDiscriminatedAttributes()
						if len(discriminatedAttributes) == 0 {

						} else {
							for _, attr := range discriminatedAttributes {
								if attr.Many {
									stmt := jen.For(jen.Id("i").Op(":=").Range().Id("new").Dot(attr.CodeName)).Block(
										jen.List(jen.Id("new").Dot(attr.CodeName).Index(jen.Id("i")), jen.Err()).
											Op("=").
											Id(attr.TypeObject.discriminatorFunctionName()).
											Call(
												jen.Id("new").Dot(attr.CodeName).Index(jen.Id("i")),
											),
									)
									g.Add(stmt.Line())
								} else {
									stmt := jen.If(jen.Id("new").Dot(attr.CodeName).Op("!=").Nil()).Block(
										jen.List(jen.Id("new").Dot(attr.CodeName), jen.Err()).
											Op("=").
											Id(attr.TypeObject.discriminatorFunctionName()).
											Call(
												jen.Id("new").Dot(attr.CodeName),
											),
										jen.If(jen.Err().Op("!=").Nil().Block(
											jen.Return().List(jen.Nil(), jen.Err()),
										)),
									)
									g.Add(stmt.Line())
								}
							}
						}
						g.Return(jen.List(jen.Id("new"), jen.Nil()))
					}),
				)
			}
		}),

		jen.Return(jen.List(jen.Nil(), jen.Nil())),
	)

	return c
}

func generateStructMarshalJSON(object *RamlType, discriminatorAttr *RamlTypeAttribute) *jen.Statement {
	c := jen.Line()
	c = c.Comment(fmt.Sprintf("MarshalJSON override to set the discriminator value")).Line()
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

	return c
}

// Generate the UnmarshalJSON method for the Structs
func generateStructUnmarshalJSON(object *RamlType, discriminatedAttributes []*RamlTypeAttribute) *jen.Statement {
	c := jen.Line()
	c = c.Comment("UnmarshalJSON override to deserialize correct attribute types based ").Line()
	c = c.Comment("on the discriminator value").Line()
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
					// Process regular many
					stmt := jen.For(jen.Id("i").Op(":=").Range().Id("obj").Dot(attr.CodeName)).Block(
						jen.Var().Err().Error(),
						jen.List(jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")), jen.Err()).
							Op("=").
							Id(attr.TypeObject.discriminatorFunctionName()).
							Call(
								jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")),
							),
						jen.If(jen.Err().Op("!=").Nil().Block(
							jen.Return(jen.Err()),
						)),
					)
					g.Add(stmt.Line())

				} else if attr.TypeObject != nil {
					// Process regular
					stmt := jen.If(jen.Id("obj").Dot(attr.CodeName).Op("!=").Nil()).Block(
						jen.Var().Err().Error(),
						jen.List(jen.Id("obj").Dot(attr.CodeName), jen.Err()).
							Op("=").
							Id(attr.TypeObject.discriminatorFunctionName()).
							Call(
								jen.Id("obj").Dot(attr.CodeName),
							),
						jen.If(jen.Err().Op("!=").Nil().Block(
							jen.Return(jen.Err()),
						)),
					)
					g.Add(stmt.Line())

				} else if attr.ItemsObject != nil {
					// Process Array type
					stmt := jen.For(jen.Id("i").Op(":=").Range().Id("obj").Dot(attr.CodeName)).Block(
						jen.Var().Err().Error(),
						jen.List(jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")), jen.Err()).
							Op("=").
							Id(attr.ItemsObject.discriminatorFunctionName()).
							Call(
								jen.Id("obj").Dot(attr.CodeName).Index(jen.Id("i")),
							),
						jen.If(jen.Err().Op("!=").Nil().Block(
							jen.Return(jen.Err()),
						)),
					)
					g.Add(stmt.Line())
				}
			}
		}),
		jen.Return(jen.Nil()),
	)

	return c
}

func generateStructField(object RamlTypeAttribute) jen.Code {
	if strings.HasPrefix(object.Name, "/") {
		return jen.Line()
	}
	code := jen.Id(object.CodeName)

	if object.Name == "transitions" && object.TypeName == "StateResourceIdentifier" {
		code = code.Op("*")
	}

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
		typeCode := getNativeTypeCode(object)
		code.Add(typeCode)

	}

	jsonTags := []string{object.Name}
	if object.Optional {
		if object.TypeName != "boolean" {
			jsonTags = append(jsonTags, "omitempty")
		}
	}

	code = code.Tag(map[string]string{"json": strings.Join(jsonTags, ",")})
	return code
}

func getNativeTypeCode(object RamlTypeAttribute) jen.Code {
	switch object.TypeName {
	case "string":
		return jen.String()
	case "boolean":
		return jen.Bool()
	case "number":
		if object.Format == "int64" {
			return jen.Int()
		} else {
			code := jen.Empty()
			if object.Minimum != nil && *object.Minimum == 0 {
				code = code.Op("*")
			}
			return code.Float64()
		}
	case "integer":
		return jen.Int()
	case "int64":
		return jen.Int()
	case "array":
		if object.ItemsObject != nil {
			if object.ItemsObject.InterfaceName != object.ItemsObject.CodeName {
				return jen.Index().Id(object.ItemsObject.InterfaceName)
			} else {
				return jen.Index().Id(object.ItemsObject.CodeName)
			}
		} else {
			code := jen.Index()
			switch object.Items {
			case "string":
				return code.String()
			case "number":
				return code.Float64()
			default:
				return code.Interface()
			}
		}
	case "datetime":
		code := jen.Empty()
		if object.Optional {
			code = code.Op("*")
		}
		return code.Qual("time", "Time")
	case "date-only":
		return jen.Op("*").Id("Date")
	default:
		return jen.Interface()
	}
}
