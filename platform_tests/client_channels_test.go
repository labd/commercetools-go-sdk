package platform_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestChannelCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 201, Body: "{}"}, &output, nil)
	defer server.Close()

	input := platform.ChannelDraft{
		Key:   "channel-key",
		Roles: []platform.ChannelRoleEnum{"Primary"},
		Name: &platform.LocalizedString{
			"en": "Primary channel",
		},
		Description: &platform.LocalizedString{
			"en": "This is the channel description.",
		},
	}

	fmt.Println(output)

	_, err := client.WithProjectKey("unittest").Channels().Post(input).Execute(context.TODO())
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
		input       platform.ChannelUpdate
		requestBody string
	}{
		{
			desc: "Change key",
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					platform.ChannelChangeKeyAction{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelChangeNameAction{
						Name: platform.LocalizedString{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelChangeDescriptionAction{
						Description: platform.LocalizedString{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelSetRolesAction{
						Roles: []platform.ChannelRoleEnum{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelAddRolesAction{
						Roles: []platform.ChannelRoleEnum{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelRemoveRolesAction{
						Roles: []platform.ChannelRoleEnum{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelSetAddressAction{
						Address: &platform.BaseAddress{
							Country: "NL",
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelSetCustomTypeAction{
						Type: &platform.TypeResourceIdentifier{
							Key: ctutils.StringRef("resource-key"),
							ID:  ctutils.StringRef("2"),
						},
						Fields: &platform.FieldContainer{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelSetCustomFieldAction{
						Name: "custom-field",
						Value: platform.LocalizedString{
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
			input: platform.ChannelUpdate{
				Version: 2,
				Actions: []platform.ChannelUpdateAction{
					&platform.ChannelSetGeoLocationAction{
						GeoLocation: platform.GeoJsonPoint{
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

			client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").Channels().WithId("1234").Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/channels/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestChannelDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Channels().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyChannelsByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/channels/1234", output.URL.Path)
}

func TestChannelGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.ResponseData{
		StatusCode: 200,
		Body: `{
			"id": "1234",
			"key": "commercetools",
			"version": 1,
			"roles": [
			"InventorySupply"
			],
			"createdAt": "2015-05-28T09:48:35.023Z",
			"lastModifiedAt": "2015-05-28T09:48:35.023Z"
			}
		`}, nil, nil)
	defer server.Close()

	timestamp, _ := time.Parse(time.RFC3339, "2015-05-28T09:48:35.023Z")

	input := platform.Channel{
		ID:             "1234",
		Key:            "commercetools",
		Version:        1,
		Roles:          []platform.ChannelRoleEnum{"InventorySupply"},
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.WithProjectKey("unittest").Channels().WithId("1234").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, *result)
}

func TestChannelQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Channels().
		Get().
		WithQueryParams(platform.ByProjectKeyChannelsRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
