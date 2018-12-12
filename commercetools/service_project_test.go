package commercetools_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestProjectGet(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2013-11-01T15:14:34.325Z")

	client, server := testutil.MockClient(t, testutil.Fixture("project.json"), nil, nil)
	defer server.Close()

	input := &commercetools.Project{
		Version:    1,
		Key:        "test-project",
		Name:       "Some project name",
		Countries:  []commercetools.CountryCode{"DE", "US"},
		Currencies: []commercetools.CurrencyCode{"EUR"},
		Languages:  []commercetools.Locale{"en"},
		CreatedAt:  timestamp,
		TrialUntil: "2013-11-01T15:14:34.325Z",
		Messages: &commercetools.MessageConfiguration{
			Enabled: false,
		},
		ShippingRateInputType: commercetools.CartScoreType{},
	}

	result, err := client.ProjectGet()
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestProjectUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.ProjectUpdateInput
		requestBody string
	}{
		{
			desc: "Change name",
			input: &commercetools.ProjectUpdateInput{
				Version: 2,
				Actions: []commercetools.ProjectUpdateAction{
					&commercetools.ProductChangeNameAction{
						Name: &commercetools.LocalizedString{
							"en": "Changed name",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeName",
						"name": {
							"en": "Changed name"
						}
					}
				]
			}`,
		},
		{
			desc: "Change currencies",
			input: &commercetools.ProjectUpdateInput{
				Version: 2,
				Actions: []commercetools.ProjectUpdateAction{
					&commercetools.ProjectChangeCurrenciesAction{
						Currencies: []commercetools.CurrencyCode{"EUR"},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeCurrencies",
						"currencies": ["EUR"]
					}
				]
			}`,
		},
		{
			desc: "Change countries",
			input: &commercetools.ProjectUpdateInput{
				Version: 2,
				Actions: []commercetools.ProjectUpdateAction{
					&commercetools.ProjectChangeCountriesAction{
						Countries: []commercetools.CountryCode{"EN", "NL"},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeCountries",
						"countries": ["EN", "NL"]
					}
				]
			}`,
		},
		{
			desc: "Change languages",
			input: &commercetools.ProjectUpdateInput{
				Version: 2,
				Actions: []commercetools.ProjectUpdateAction{
					&commercetools.ProjectChangeLanguagesAction{
						Languages: []commercetools.Locale{"en", "nl"},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeLanguages",
						"languages": ["en", "nl"]
					}
				]
			}`,
		},
		{
			desc: "Change messages enabled",
			input: &commercetools.ProjectUpdateInput{
				Version: 2,
				Actions: []commercetools.ProjectUpdateAction{
					&commercetools.ProjectChangeMessagesEnabledAction{
						MessagesEnabled: true,
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeMessagesEnabled",
						"messagesEnabled": true
					}
				]
			}`,
		},
		{
			desc: "Set shipping rate input type",
			input: &commercetools.ProjectUpdateInput{
				Version: 2,
				Actions: []commercetools.ProjectUpdateAction{
					&commercetools.ProjectSetShippingRateInputTypeAction{
						ShippingRateInputType: commercetools.CartScoreType{},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setShippingRateInputType",
						"shippingRateInputType": {
							"type": "CartScore"
						}
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

			_, err := client.ProjectUpdate(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestShippingRateInputType(t *testing.T) {
	testCases := []struct {
		desc   string
		object interface {
			MarshalJSON() ([]byte, error)
		}
		json string
	}{
		{
			desc:   "Cart value type",
			object: commercetools.CartValueType{},
			json: `{
				"type": "CartValue"
			}`,
		},
		{
			desc: "Cart classification type",
			object: commercetools.CartClassificationType{
				Values: []commercetools.CustomFieldLocalizedEnumValue{
					{
						Key: "test",
						Label: &commercetools.LocalizedString{
							"en": "test",
						},
					},
				},
			},
			json: `{
				"type": "CartClassification",
				"values": [
					{
						"key": "test",
						"label": {
							"en": "test"
						}
					}
				]
			}`,
		},
		{
			desc:   "Cart score type",
			object: commercetools.CartScoreType{},
			json: `{
				"type": "CartScore"
			}`,
		},
	}

	for _, tC := range testCases {

		// Validate Marshaling (object -> json)
		t.Run(tC.desc, func(t *testing.T) {
			output, err := tC.object.MarshalJSON()
			assert.Nil(t, err)
			assert.JSONEq(t, tC.json, string(output))
		})

		// Validate Umarshalling (json -> object)
		t.Run(tC.desc, func(t *testing.T) {
			buf := fmt.Sprintf(`
			{
		 		"version": 1,
				"key": "test-project",
				"name": "Some project name",
				"countries": ["DE", "US"],
				"currencies": ["EUR"],
				"languages": ["en"],
				"createdAt": "2013-11-01T15:14:34.325Z",
				"trialUntil": "2013-11-01T15:14:34.325Z",
				"messages": {
					"enabled": false
				},
				"shippingRateInputType": %s
			}
			`, tC.json)

			p := commercetools.Project{}

			err := p.UnmarshalJSON([]byte(buf))

			assert.Nil(t, err)
			assert.Equal(t, tC.object, p.ShippingRateInputType)
		})
	}
}
