package commercetools_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTaxCategoryCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("tax-category.create.json"), &output, nil)
	defer server.Close()

	input := &commercetools.TaxCategoryDraft{
		Name:        "Tax category",
		Key:         "tax-category",
		Description: "Tax category description.",
		Rates: []commercetools.TaxRateDraft{
			{
				Name:            "Tax rate low",
				Amount:          6.00,
				IncludedInPrice: false,
				Country:         "NL",
			},
			{
				Name:            "Tax rate high",
				Amount:          21.00,
				IncludedInPrice: false,
				Country:         "NL",
			},
		},
	}

	fmt.Println(output)

	_, err := client.TaxCategoryCreate(input)
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
	testCases := []struct {
		desc        string
		input       *commercetools.TaxCategoryUpdateInput
		requestBody string
	}{
		{
			desc: "Change name",
			input: &commercetools.TaxCategoryUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TaxCategoryUpdateAction{
					&commercetools.TaxCategoryChangeNameAction{
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
			input: &commercetools.TaxCategoryUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TaxCategoryUpdateAction{
					&commercetools.TaxCategorySetKeyAction{
						Key: "tax-category",
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
			input: &commercetools.TaxCategoryUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TaxCategoryUpdateAction{
					&commercetools.TaxCategorySetDescriptionAction{
						Description: "A tax category description",
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
			input: &commercetools.TaxCategoryUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TaxCategoryUpdateAction{
					&commercetools.TaxCategoryAddTaxRateAction{
						TaxRate: &commercetools.TaxRateDraft{
							Name:            "High tax rate",
							Amount:          21.14,
							IncludedInPrice: true,
							Country:         commercetools.CountryCode("US"),
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
			input: &commercetools.TaxCategoryUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TaxCategoryUpdateAction{
					&commercetools.TaxCategoryReplaceTaxRateAction{
						TaxRateID: "2",
						TaxRate: &commercetools.TaxRateDraft{
							Name:            "High tax rate",
							Amount:          21.14,
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
			input: &commercetools.TaxCategoryUpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.TaxCategoryUpdateAction{
					&commercetools.TaxCategoryRemoveTaxRateAction{
						TaxRateID: "2",
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

			client, server := testutil.MockClient(t, testutil.Fixture("tax-category.update.json"), &output, nil)
			defer server.Close()

			_, err := client.TaxCategoryUpdate(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/tax-categories/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestTaxCategoryDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.TaxCategoryDeleteByID("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/tax-categories/1234", output.URL.Path)
}

func TestTaxCategoryDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.TaxCategoryDeleteByKey("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/tax-categories/key=1234", output.URL.Path)
}

func TestTaxCategoryGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-02-24T15:33:40.811Z")

	client, server := testutil.MockClient(t, testutil.Fixture("tax-category.json"), nil, nil)
	defer server.Close()

	input := &commercetools.TaxCategory{
		ID:      "c60f7377-2643-4e99-adb5-b2909657444d",
		Version: 1,
		Name:    "test-tax-category",
		Rates: []commercetools.TaxRate{
			{
				Name:            "test-tax-category",
				Amount:          0.2,
				IncludedInPrice: true,
				Country:         "DE",
				ID:              "vWTk7VjT",
				SubRates:        []commercetools.SubRate{},
			},
		},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.TaxCategoryGetByID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}
