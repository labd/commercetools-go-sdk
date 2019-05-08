package commercetools_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestCartCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("cart.json"), &output, nil)
	defer server.Close()

	input := &commercetools.CartDraft{
		Currency: "EUR",
		ShippingAddress: &commercetools.Address{
			Country: "DE",
		},
	}

	cart, err := client.CartCreate(input)
	fmt.Printf("%#v", cart)
	assert.Nil(t, err)

	assert.Equal(t, commercetools.CountryCode("DE"), cart.ShippingAddress.Country)
}

func TestCartUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.CartUpdateInput
		requestBody string
	}{
		{
			desc: "Set locale",
			input: &commercetools.CartUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.CartUpdateAction{
					&commercetools.CartSetLocaleAction{
						Locale: "DE",
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

			client, server := testutil.MockClient(t, "{}", &output, nil)
			defer server.Close()

			_, err := client.CartUpdate(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/carts/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestCartDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.CartDelete("debd77c9-ef14-4595-8614-5d123429f9e1", 2, true)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	params.Add("dataErasure", "true")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/carts/debd77c9-ef14-4595-8614-5d123429f9e1", output.URL.Path)
}

func TestCartGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("cart.json"), nil, nil)
	defer server.Close()

	cart, err := client.CartGetByID("debd77c9-ef14-4595-8614-5d123429f9e1")

	assert.Nil(t, err)
	assert.Equal(t, "debd77c9-ef14-4595-8614-5d123429f9e1", cart.ID)
}

func TestCartQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.CartQuery(&queryInput)

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}

func TestCartReplicate(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.ReplicaCartDraft{
		Reference: commercetools.CartReference{
			ID: "debd77c9-ef14-4595-8614-5d123429f9e1",
		},
	}
	_, err := client.CartReplicate(&queryInput)

	assert.Nil(t, err)

	expectedJSON := `{
        "reference": {
        	"typeId": "cart",
        	"id": "debd77c9-ef14-4595-8614-5d123429f9e1"
        }
    }`
	assert.Equal(t, "POST", output.Method)
	assert.Equal(t, "/unittest/carts/replicate", output.URL.Path)
	assert.JSONEq(t, expectedJSON, output.JSON)
}
