package commercetools_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestChannelCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	input := &commercetools.ChannelDraft{
		Key:   "channel-key",
		Roles: []commercetools.ChannelRoleEnum{"Primary"},
		Name: &commercetools.LocalizedString{
			"en": "Primary channel",
		},
		Description: &commercetools.LocalizedString{
			"en": "This is the channel description.",
		},
	}

	fmt.Println(output)

	_, err := client.ChannelCreate(input)
	assert.Nil(t, err)

	expectedBody := `{
		"key": "channel-key",
		"roles": ["Primary"],
		"name": {
			"en": "Primary channel"
		},
		"description": {
			"en": "This is the channel description."
		}
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestChannelUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.ChannelUpdateWithIDInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelChangeKeyAction{
						Key: "123456",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeKey",
						"key": "123456"
					}
				]
			}`,
		},
		{
			desc: "Change name",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelChangeNameAction{
						Name: &commercetools.LocalizedString{
							"en": "Primary channel",
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
							"en": "Primary channel"
						}
					}
				]
			}`,
		},
		{
			desc: "Change description",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelChangeDescriptionAction{
						Description: &commercetools.LocalizedString{
							"en": "This is the channel description.",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeDescription",
						"description": {
							"en": "This is the channel description."
						}
					}
				]
			}`,
		},
		{
			desc: "Set roles",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelSetRolesAction{
						Roles: []commercetools.ChannelRoleEnum{
							"Primary",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setRoles",
						"roles": ["Primary"]
					}
				]
			}`,
		},
		{
			desc: "Add roles",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelAddRolesAction{
						Roles: []commercetools.ChannelRoleEnum{
							"Primary",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "addRoles",
						"roles": ["Primary"]
					}
				]
			}`,
		},
		{
			desc: "Remove roles",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelRemoveRolesAction{
						Roles: []commercetools.ChannelRoleEnum{
							"Primary",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "removeRoles",
						"roles": ["Primary"]
					}
				]
			}`,
		},
		{
			desc: "Set address",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelSetAddressAction{
						Address: &commercetools.Address{
							Country: commercetools.CountryCode("NL"),
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setAddress",
						"address": {
							"country": "NL"
						}
					}
				]
			}`,
		},
		{
			desc: "Set custom type",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelSetCustomTypeAction{
						Type: &commercetools.TypeResourceIdentifier{
							Key: "resource-key",
							ID:  "2",
						},
						Fields: &commercetools.FieldContainer{
							"description": "my-description",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setCustomType",
						"type": {
							"key": "resource-key",
							"id": "2",
							"typeId": "type"
						},
						"fields": {
							"description": "my-description"
						}
					}
				]
			}`,
		},
		{
			desc: "Set custom field",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelSetCustomFieldAction{
						Name: "custom-field",
						Value: commercetools.LocalizedString{
							"en": "My custom field",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setCustomField",
						"name": "custom-field",
						"value": {
							"en": "My custom field"
						}
					}
				]
			}`,
		},
		{
			desc: "Set geo location",
			input: &commercetools.ChannelUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.ChannelUpdateAction{
					&commercetools.ChannelSetGeoLocationAction{
						GeoLocation: &commercetools.GeoJSONPoint{
							Coordinates: []float64{52.068439, 5.1075891},
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setGeoLocation",
						"geoLocation": {
							"type": "Point",
							"coordinates": [52.068439, 5.1075891]
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

			_, err := client.ChannelUpdateWithID(tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/channels/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestChannelDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.ChannelDeleteWithID("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/channels/1234", output.URL.Path)
}

func TestChannelGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("channels.json"), nil, nil)
	defer server.Close()

	timestamp, _ := time.Parse(time.RFC3339, "2015-05-28T09:48:35.023Z")

	input := &commercetools.Channel{
		ID:             "1234",
		Key:            "commercetools",
		Version:        1,
		Roles:          []commercetools.ChannelRoleEnum{"InventorySupply"},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.ChannelGetWithID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestChannelQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.ChannelQuery(&queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
