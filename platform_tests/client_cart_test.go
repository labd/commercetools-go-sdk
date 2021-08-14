package platform_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestCartCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("cart.json", 201), &output, nil)
	defer server.Close()

	input := platform.CartDraft{
		Currency: "EUR",
		ShippingAddress: &platform.BaseAddress{
			Country: "DE",
		},
	}

	cart, err := client.WithProjectKey("tests").Carts().Post(input).Execute(context.TODO())
	fmt.Printf("%#v", cart)
	assert.Nil(t, err)

	assert.Equal(t, "DE", cart.ShippingAddress.Country)
}

func TestCartUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.CartUpdate
		requestBody string
	}{
		{
			desc: "Set locale",
			input: platform.CartUpdate{
				Version: 2,
				Actions: []platform.CartUpdateAction{
					&platform.CartSetLocaleAction{
						Locale: ctutils.StringRef("DE"),
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setLocale",
						"locale": "DE"
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").Carts().WithId("1234").Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/carts/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestCartDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").Carts().
		WithId("debd77c9-ef14-4595-8614-5d123429f9e1").
		Delete().
		WithQueryParams(platform.ByProjectKeyCartsByIDRequestMethodDeleteInput{
			Version:     2,
			DataErasure: ctutils.BoolRef(true),
		}).
		Execute(context.TODO())

	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	params.Add("dataErasure", "true")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/carts/debd77c9-ef14-4595-8614-5d123429f9e1", output.URL.Path)
}

func TestCartGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("cart.json", 200), nil, nil)
	defer server.Close()

	cart, err := client.WithProjectKey("unittest").Carts().WithId("debd77c9-ef14-4595-8614-5d123429f9e1").Get().Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, "debd77c9-ef14-4595-8614-5d123429f9e1", cart.ID)
}

func TestCartQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").Carts().Get().WithQueryParams(platform.ByProjectKeyCartsRequestMethodGetInput{
		Limit: ctutils.IntRef(500),
	}).Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
