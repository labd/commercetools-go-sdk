package commercetools_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShippingMethodCreate(t *testing.T) {

	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.create.json"), &output, nil)
	defer server.Close()

	input := &commercetools.ShippingMethodDraft{
		Name:        "Shipping method",
		Key:         "shipping-method",
		Description: "Shipping method description.",
	}

	_, err := client.ShippingMethodCreate(input)
	assert.Nil(t, err)
	expectBody := `{
		"name": "Shipping method",
		"key": "shipping-method",
		"description": "Shipping method description.",
		"isDefault": false,
		"zoneRates": null,
		"taxCategory": null
	}`

	assert.JSONEq(t, expectBody, output.JSON)
}

func TestShippingMethodUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.ShippingMethodUpdateInput
		requestBody string
	}{
		{
			desc: "Change name",
			input: &commercetools.ShippingMethodUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ShippingMethodUpdateAction{
					&commercetools.ShippingMethodChangeNameAction{
						Name: "New name",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeName",
						"name": "New name"
					}
				]
			}`,
		},
		{
			desc: "Set key",
			input: &commercetools.ShippingMethodUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ShippingMethodUpdateAction{
					&commercetools.ShippingMethodSetKeyAction{
						Key: "new-key",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setKey",
						"key": "new-key"
					}
				]
			}`,
		},
		{
			desc: "Set description",
			input: &commercetools.ShippingMethodUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ShippingMethodUpdateAction{
					&commercetools.ShippingMethodSetDescriptionAction{
						Description: "A new shipping method description",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setDescription",
						"description": "A new shipping method description"
					}
				]
			}`,
		},
		{
			desc: "Set default",
			input: &commercetools.ShippingMethodUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ShippingMethodUpdateAction{
					&commercetools.ShippingMethodChangeIsDefaultAction{
						IsDefault: true,
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeIsDefault",
						"isDefault": true
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.update.json"), &output, nil)
			defer server.Close()

			_, err := client.ShippingMethodUpdate(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/shipping-methods/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestShippingMethodDeleteById(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.ShippingMethodDeleteByID("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/shipping-methods/1234", output.URL.Path)
}

func TestShippingMethodGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.json"), nil, nil)
	defer server.Close()

	input := &commercetools.ShippingMethod{
		ID:             "93153b7e-577e-11e9-a0db-4ff9c32a63cc",
		Version:        1,
		Name:           "test shipping method",
		Key:            "test-shipping-method",
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.ShippingMethodGetByID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}
func TestShippingMethodGetByKey(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.json"), nil, nil)
	defer server.Close()

	input := &commercetools.ShippingMethod{
		ID:             "93153b7e-577e-11e9-a0db-4ff9c32a63cc",
		Version:        1,
		Name:           "test shipping method",
		Key:            "test-shipping-method",
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.ShippingMethodGetByKey("test-shipping-method")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestShippingMethodQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.ShippingMethodQuery(&queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
