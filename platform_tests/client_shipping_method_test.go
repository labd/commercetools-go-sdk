package platform_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShippingMethodCreate(t *testing.T) {

	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{
		StatusCode: 201,
		Body: `
		{
			"name": "test-shipping-method"
		}`,
	}, &output, nil)
	defer server.Close()

	input := platform.ShippingMethodDraft{
		Name:        "Shipping method",
		Key:         ctutils.StringRef("shipping-method"),
		Description: ctutils.StringRef("Shipping method description."),
		IsDefault:   false,
	}

	_, err := client.WithProjectKey("unittest").ShippingMethods().Post(input).Execute(context.TODO())
	assert.Nil(t, err)
	expectBody := `{
		"name": "Shipping method",
		"key": "shipping-method",
		"description": "Shipping method description.",
		"isDefault": false,
		"zoneRates": null,
		"taxCategory": {
			"typeId": "tax-category"
		}
	}`

	assert.JSONEq(t, expectBody, output.JSON)
}

func TestShippingMethodUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.ShippingMethodUpdate
		requestBody string
	}{
		{
			desc: "Change name",
			input: platform.ShippingMethodUpdate{
				Version: 2,
				Actions: []platform.ShippingMethodUpdateAction{
					&platform.ShippingMethodChangeNameAction{
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
			input: platform.ShippingMethodUpdate{
				Version: 2,
				Actions: []platform.ShippingMethodUpdateAction{
					&platform.ShippingMethodSetKeyAction{
						Key: ctutils.StringRef("new-key"),
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
			input: platform.ShippingMethodUpdate{
				Version: 2,
				Actions: []platform.ShippingMethodUpdateAction{
					&platform.ShippingMethodSetDescriptionAction{
						Description: ctutils.StringRef("A new shipping method description"),
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
			input: platform.ShippingMethodUpdate{
				Version: 2,
				Actions: []platform.ShippingMethodUpdateAction{
					&platform.ShippingMethodChangeIsDefaultAction{
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

	// test ShippingMethodupdateByID
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.ResponseData{
				StatusCode: 200,
				Body: `
				{
					"version": 1,
					"actions": [
					  {
						"action": "changeName",
						"name": "New Name"
					  }
					]
				  }
				 `,
			}

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.
				WithProjectKey("unittest").
				ShippingMethods().
				WithId("1234").
				Post(tC.input).
				Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/shipping-methods/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestShippingMethodDeleteById(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		ShippingMethods().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyShippingMethodsByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/shipping-methods/1234", output.URL.Path)
}

func TestShippingMethodGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.json", 200), nil, nil)
	defer server.Close()

	input := platform.ShippingMethod{
		ID:             "93153b7e-577e-11e9-a0db-4ff9c32a63cc",
		Version:        1,
		Name:           "test shipping method",
		Key:            ctutils.StringRef("test-shipping-method"),
		CreatedAt:      timestamp,
		Description:    ctutils.StringRef(""),
		LastModifiedAt: timestamp,
	}

	result, err := client.WithProjectKey("unittest").
		ShippingMethods().
		WithId("1234").
		Get().
		Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, *result)
}
func TestShippingMethodGetByKey(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("shipping-method.json", 200), nil, nil)
	defer server.Close()

	input := platform.ShippingMethod{
		ID:             "93153b7e-577e-11e9-a0db-4ff9c32a63cc",
		Version:        1,
		Name:           "test shipping method",
		Key:            ctutils.StringRef("test-shipping-method"),
		Description:    ctutils.StringRef(""),
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.WithProjectKey("unittest").
		ShippingMethods().
		WithKey("test-shipping-method").
		Get().
		Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, *result)
}

func TestShippingMethodQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		ShippingMethods().
		Get().
		WithQueryParams(platform.ByProjectKeyShippingMethodsRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
