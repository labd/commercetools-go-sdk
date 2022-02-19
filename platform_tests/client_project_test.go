package platform_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestProjectGet(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2013-11-01T15:14:34.325Z")

	client, server := testutil.MockClient(t, testutil.Fixture("project.json", 200), nil, nil)
	defer server.Close()

	input := &platform.Project{
		Version:    1,
		Key:        "test-project",
		Name:       "Some project name",
		Countries:  []string{"DE", "US"},
		Currencies: []string{"EUR"},
		Languages:  []string{"en"},
		CreatedAt:  timestamp,
		TrialUntil: ctutils.StringRef("2013-11-01T15:14:34.325Z"),
		Messages: platform.MessagesConfiguration{
			Enabled: false,
		},
		ShippingRateInputType: platform.CartScoreType{},
	}

	result, err := client.WithProjectKey("test-project").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestProjectUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.ProjectUpdate
		requestBody string
	}{
		{
			desc: "Change name",
			input: platform.ProjectUpdate{
				Version: 2,
				Actions: []platform.ProjectUpdateAction{
					&platform.ProductChangeNameAction{
						Name: platform.LocalizedString{
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
			input: platform.ProjectUpdate{
				Version: 2,
				Actions: []platform.ProjectUpdateAction{
					&platform.ProjectChangeCurrenciesAction{
						Currencies: []string{"EUR"},
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
			input: platform.ProjectUpdate{
				Version: 2,
				Actions: []platform.ProjectUpdateAction{
					&platform.ProjectChangeCountriesAction{
						Countries: []string{"EN", "NL"},
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
			input: platform.ProjectUpdate{
				Version: 2,
				Actions: []platform.ProjectUpdateAction{
					&platform.ProjectChangeLanguagesAction{
						Languages: []string{"en", "nl"},
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
			input: platform.ProjectUpdate{
				Version: 2,
				Actions: []platform.ProjectUpdateAction{
					&platform.ProjectChangeMessagesEnabledAction{
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
			input: platform.ProjectUpdate{
				Version: 2,
				Actions: []platform.ProjectUpdateAction{
					&platform.ProjectSetShippingRateInputTypeAction{
						ShippingRateInputType: platform.CartScoreType{},
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

			client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest", output.URL.Path)
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
			object: platform.CartValueType{},
			json: `{
				"type": "CartValue"
			}`,
		},
		{
			desc: "Cart classification type",
			object: platform.CartClassificationType{
				Values: []platform.CustomFieldLocalizedEnumValue{
					{
						Key: "test",
						Label: platform.LocalizedString{
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
			object: platform.CartScoreType{},
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

			p := platform.Project{}

			err := p.UnmarshalJSON([]byte(buf))

			assert.Nil(t, err)
			assert.Equal(t, tC.object, p.ShippingRateInputType)
		})
	}
}
