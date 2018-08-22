package project_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"

	"github.com/labd/commercetools-go-sdk/service/project"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestProjectGet(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2013-11-01T15:14:34.325Z")

	client, server := testutil.MockClient(t, testutil.Fixture("project.json"), nil, nil)
	defer server.Close()

	input := &project.Project{
		Version:    1,
		Key:        "test-project",
		Name:       "Some project name",
		Countries:  []string{"DE", "US"},
		Currencies: []string{"EUR"},
		Languages:  []string{"en"},
		CreatedAt:  timestamp,
		TrialUntil: "2013-11-01T15:14:34.325Z",
		Messages: project.MessagesConfiguration{
			Enabled: false,
		},
		ShippingRateInputType: project.CartScore{},
	}

	svc := project.New(client)
	result, err := svc.Get()
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestProjectUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *project.UpdateInput
		requestBody string
	}{
		{
			desc: "Change name",
			input: &project.UpdateInput{
				Version: 2,
				Actions: commercetools.UpdateActions{
					&project.ChangeName{
						Name: "Changed name",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeName",
						"name": "Changed name"
					}
				]
			}`,
		},
		{
			desc: "Change currencies",
			input: &project.UpdateInput{
				Version: 2,
				Actions: commercetools.UpdateActions{
					&project.ChangeCurrencies{
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
			input: &project.UpdateInput{
				Version: 2,
				Actions: commercetools.UpdateActions{
					&project.ChangeCountries{
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
			input: &project.UpdateInput{
				Version: 2,
				Actions: commercetools.UpdateActions{
					&project.ChangeLanguages{
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
			input: &project.UpdateInput{
				Version: 2,
				Actions: commercetools.UpdateActions{
					&project.ChangeMessagesEnabled{
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
			input: &project.UpdateInput{
				Version: 2,
				Actions: commercetools.UpdateActions{
					&project.SetShippingRateInputType{
						ShippingRateInputType: project.CartScore{},
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
			svc := project.New(client)

			_, err := svc.Update(tC.input)
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
			object: project.CartValue{},
			json: `{
				"type": "CartValue"
			}`,
		},
		{
			desc: "Cart classification type",
			object: project.CartClassification{
				Values: []commercetools.LocalizedEnumValue{
					{
						Key: "test",
						Label: commercetools.LocalizedString{
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
			object: project.CartScore{},
			json: `{
				"type": "CartScore"
			}`,
		},
	}

	for _, tC := range testCases {

		// Validate Marshalling (object -> json)
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

			p := project.Project{}

			err := p.UnmarshalJSON([]byte(buf))

			assert.Nil(t, err)
			assert.Equal(t, tC.object, p.ShippingRateInputType)
		})
	}
}
