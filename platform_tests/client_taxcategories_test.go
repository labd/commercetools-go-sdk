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

func TestTaxCategoryCreate(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{
		StatusCode: 201,
		Body: `
		{
			"name": "test-tax-category",
			"rates": [
			  {
				"name": "test-tax-category",
				"amount": 0.2,
				"includedInPrice": true,
				"country": "DE"
			  }
			]
		  }
		`,
	}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	low_tax_rate := 6.00
	high_tax_rate := 21.00

	input := platform.TaxCategoryDraft{
		Name:        "Tax category",
		Key:         ctutils.StringRef("tax-category"),
		Description: ctutils.StringRef("Tax category description."),
		Rates: []platform.TaxRateDraft{
			{
				Name:            "Tax rate low",
				Amount:          &low_tax_rate,
				IncludedInPrice: false,
				Country:         "NL",
			},
			{
				Name:            "Tax rate high",
				Amount:          &high_tax_rate,
				IncludedInPrice: false,
				Country:         "NL",
			},
		},
	}

	_, err := client.WithProjectKey("unittest").TaxCategories().Post(input).Execute(context.TODO())
	assert.Nil(t, err)

	expectedBody := `{
		"name": "Tax category",
		"key": "tax-category",
		"description": "Tax category description.",
		"rates": [
			{
				"name": "Tax rate low",
				"amount": 6,
				"includedInPrice": false,
				"country": "NL"
			},
			{
				"name": "Tax rate high",
				"amount": 21,
				"includedInPrice": false,
				"country": "NL"
			}
		]
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestTaxCategoryUpdate(t *testing.T) {
	high_tax_rate := 21.14
	testCases := []struct {
		desc        string
		input       platform.TaxCategoryUpdate
		requestBody string
	}{
		{
			desc: "Change name",
			input: platform.TaxCategoryUpdate{
				Version: 2,
				Actions: []platform.TaxCategoryUpdateAction{
					&platform.TaxCategoryChangeNameAction{
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
			input: platform.TaxCategoryUpdate{
				Version: 2,
				Actions: []platform.TaxCategoryUpdateAction{
					&platform.TaxCategorySetKeyAction{
						Key: ctutils.StringRef("tax-category"),
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setKey",
						"key": "tax-category"
					}
				]
			}`,
		},
		{
			desc: "Set description",
			input: platform.TaxCategoryUpdate{
				Version: 2,
				Actions: []platform.TaxCategoryUpdateAction{
					&platform.TaxCategorySetDescriptionAction{
						Description: ctutils.StringRef("A tax category description"),
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setDescription",
						"description": "A tax category description"
					}
				]
			}`,
		},
		{
			desc: "Add tax rate",
			input: platform.TaxCategoryUpdate{
				Version: 2,
				Actions: []platform.TaxCategoryUpdateAction{
					&platform.TaxCategoryAddTaxRateAction{
						TaxRate: platform.TaxRateDraft{
							Name:            "High tax rate",
							Amount:          &high_tax_rate,
							IncludedInPrice: true,
							Country:         "US",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "addTaxRate",
						"taxRate": {
							"name": "High tax rate",
							"amount": 21.14,
							"includedInPrice": true,
							"country": "US"
						}
					}
				]
			}`,
		},
		{
			desc: "Replace tax rate",
			input: platform.TaxCategoryUpdate{
				Version: 2,
				Actions: []platform.TaxCategoryUpdateAction{
					&platform.TaxCategoryReplaceTaxRateAction{
						TaxRateId: "2",
						TaxRate: platform.TaxRateDraft{
							Name:            "High tax rate",
							Amount:          &high_tax_rate,
							IncludedInPrice: true,
							Country:         "US",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "replaceTaxRate",
						"taxRateId": "2",
						"taxRate": {
							"name": "High tax rate",
							"amount": 21.14,
							"includedInPrice": true,
							"country": "US"
						}
					}
				]
			}`,
		},
		{
			desc: "Remove tax rate",
			input: platform.TaxCategoryUpdate{
				Version: 2,
				Actions: []platform.TaxCategoryUpdateAction{
					&platform.TaxCategoryRemoveTaxRateAction{
						TaxRateId: "2",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "removeTaxRate",
						"taxRateId": "2"
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.ResponseData{
				StatusCode: 200,
				Body: `
				{
					"name": "test-tax-category",
					"rates": [
					  {
						"name": "test-tax-category",
						"amount": 0.2,
						"includedInPrice": true,
						"country": "DE"
					  }
					]
				  }
				`,
			}

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").TaxCategories().WithId("1234").Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/tax-categories/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestTaxCategoryDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		TaxCategories().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/tax-categories/1234", output.URL.Path)
}

func TestTaxCategoryDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		TaxCategories().
		WithKey("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyTaxCategoriesKeyByKeyRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/tax-categories/key=1234", output.URL.Path)
}

func TestTaxCategoryGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")
	response := testutil.Fixture("tax-category.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()
	tax_rate := 0.2

	input := &platform.TaxCategory{
		ID:      "c60f7377-2643-4e99-adb5-b2909657444d",
		Version: 1,
		Name:    "test-tax-category",
		Rates: []platform.TaxRate{
			{
				Name:            "test-tax-category",
				Amount:          tax_rate,
				IncludedInPrice: true,
				Country:         "DE",
				ID:              ctutils.StringRef("vWTk7VjT"),
				SubRates:        []platform.SubRate{},
			},
		},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.WithProjectKey("unittest").
		TaxCategories().
		WithId("1234").
		Get().
		Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestTaxCategoryQuery(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		TaxCategories().
		Get().
		Limit(500).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
