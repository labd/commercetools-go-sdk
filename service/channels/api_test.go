package channels_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/service/channels"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestChannelCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()
	svc := channels.New(client)

	input := &channels.ChannelDraft{
		Key:   "channel-key",
		Roles: []channels.ChannelRole{channels.Primary},
		Name: commercetools.LocalizedString{
			"en": "Primary channel",
		},
		Description: commercetools.LocalizedString{
			"en": "This is the channel description.",
		},
	}

	fmt.Println(output)

	_, err := svc.Create(input)
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
		input       *channels.UpdateInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.ChangeKey{
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.ChangeName{
						Name: commercetools.LocalizedString{
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.ChangeDescription{
						Description: commercetools.LocalizedString{
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.SetRoles{
						Roles: []channels.ChannelRole{
							channels.Primary,
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.AddRoles{
						Roles: []channels.ChannelRole{
							channels.Primary,
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.RemoveRoles{
						Roles: []channels.ChannelRole{
							channels.Primary,
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.SetAddress{
						Address: commercetools.Address{},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setAddress",
						"address": {}
					}
				]
			}`,
		},
		{
			desc: "Set custom type",
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.SetCustomType{
						Type: commercetools.ResourceIdentifier{
							Key:    "resource-key",
							TypeID: "2",
						},
						Fields: map[string]string{
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
							"typeId": "2"
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.SetCustomField{
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
			input: &channels.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&channels.SetGeoLocation{
						GeoLocation: commercetools.Point{
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
			svc := channels.New(client)

			_, err := svc.Update(tC.input)
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
	svc := channels.New(client)

	_, err := svc.Delete("1234", 2)
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

	input := &channels.Channel{
		ID:             "1234",
		Key:            "commercetools",
		Version:        1,
		Roles:          []channels.ChannelRole{channels.InventorySupply},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	svc := channels.New(client)
	result, err := svc.GetByID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}
