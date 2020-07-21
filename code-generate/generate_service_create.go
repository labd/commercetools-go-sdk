package main

import (
	"github.com/dave/jennifer/jen"
)

// Generate the `<service>Create` function
func generateServiceCreate(funcName string, resourceService ResourceService, urlPathName string) (code *jen.Statement) {
	structReceiver := jen.Id("client").Op("*").Id("Client")
	createParams := jen.List(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("draft").Op("*").Id(resourceService.ResourceDraftType),
		jen.Id("opts").Op("...").Id("RequestOption"),
	)
	returnParams := jen.List(
		jen.Id("result").Op("*").Id(resourceService.ResourceType),
		jen.Err().Error())

	c := jen.Commentf("%s creates a new instance of type %s", funcName, resourceService.ResourceType).Line()

	// func (client *Client) ResourceCreate(ctx context.Context, draft *ResourceDraft, opts ...RequestOption) (result *APIClient, err error) {
	c.Func().Params(structReceiver).Id(funcName).Params(createParams).Parens(returnParams).Block(

		// params := url.Values{}
		// for _, opt := range opts {
		// 	opt(&params)
		// }
		jen.Id("params").Op(":=").Qual("net/url", "Values").Block(),
		jen.For(jen.List(jen.Id("_"), jen.Id("opt")).Op(":=").Range().Id("opts")).Block(
			jen.Id("opt").Call(jen.Op("&").Id("params")),
		),
		jen.Line(),

		// err = client.Create(ctx, ProductURLPath, params, draft, &result)
		jen.Id("err").Op("=").Id("client").Op(".").Id("Create").Call(
			jen.Id("ctx"),
			jen.Id(urlPathName),
			jen.Id("params"),
			jen.Id("draft"),
			jen.Op("&").Id("result"),
		),

		// if err != nil {
		// 	return nil, err
		// }
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),

		// return result, nil
		jen.Return(jen.Id("result"), jen.Nil()),
	).Line()
	return c
}
