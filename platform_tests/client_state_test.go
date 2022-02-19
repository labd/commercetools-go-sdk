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

func TestStateCreate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.StateDraft
		requestBody string
	}{
		{
			desc: "default",
			input: platform.StateDraft{
				Key:  "state",
				Type: platform.StateTypeEnum("ProductState"),
				Name: &platform.LocalizedString{
					"en": "State machine",
				},
				Description: &platform.LocalizedString{
					"en": "This is a state machine. Beep beep boop.",
				},
				Initial: ctutils.BoolRef(true),
				Transitions: []platform.StateResourceIdentifier{
					{
						ID: ctutils.StringRef("1"),
					},
				},
			},
			requestBody: `{
				"key": "state",
				"type": "ProductState",
				"name": {
					"en": "State machine"
				},
				"description": {
					"en": "This is a state machine. Beep beep boop."
				},
				"initial": true,
				"transitions": [
					{
						"typeId": "state",
						"id": "1"
					}
				]
			}`,
		},
		{
			desc: "no transitions",
			input: platform.StateDraft{
				Key:  "state",
				Type: platform.StateTypeEnum("ProductState"),
				Name: &platform.LocalizedString{
					"en": "State machine",
				},
				Description: &platform.LocalizedString{
					"en": "This is a state machine. Beep beep boop.",
				},
				Initial: ctutils.BoolRef(true),
			},
			requestBody: `{
				"key": "state",
				"type": "ProductState",
				"name": {
					"en": "State machine"
				},
				"description": {
					"en": "This is a state machine. Beep beep boop."
				},
				"initial": true
			}`,
		},
		{
			desc: "empty transitions",
			input: platform.StateDraft{
				Key:  "state",
				Type: platform.StateTypeEnum("ProductState"),
				Name: &platform.LocalizedString{
					"en": "State machine",
				},
				Description: &platform.LocalizedString{
					"en": "This is a state machine. Beep beep boop.",
				},
				Initial:     ctutils.BoolRef(true),
				Transitions: []platform.StateResourceIdentifier{},
			},
			requestBody: `{
				"key": "state",
				"type": "ProductState",
				"name": {
					"en": "State machine"
				},
				"description": {
					"en": "This is a state machine. Beep beep boop."
				},
				"initial": true,
				"transitions": []
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.ResponseData{StatusCode: 201, Body: "{}"}

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").States().Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)

			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStateUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.StateUpdate
		requestBody string
	}{
		{
			desc: "Change key",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateChangeKeyAction{
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
			desc: "Set name",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateSetNameAction{
						Name: platform.LocalizedString{
							"en": "New name",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setName",
						"name": {
							"en": "New name"
						}
					}
				]
			}`,
		},
		{
			desc: "Set description",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateSetDescriptionAction{
						Description: platform.LocalizedString{
							"en": "New description",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setDescription",
						"description": {
							"en": "New description"
						}
					}
				]
			}`,
		},
		{
			desc: "Change type",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateChangeTypeAction{
						Type: "OrderState",
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeType",
						"type": "OrderState"
					}
				]
			}`,
		},
		{
			desc: "Change initial",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateChangeInitialAction{
						Initial: true,
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "changeInitial",
						"initial": true
					}
				]
			}`,
		},
		{
			desc: "Set transitions",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateSetTransitionsAction{
						Transitions: []platform.StateResourceIdentifier{
							{
								ID: ctutils.StringRef("1"),
							},
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setTransitions",
						"transitions": [
							{
								"typeId": "state",
								"id": "1"
							}
						]
					}
				]
			}`,
		},
		{
			desc: "Set transitions (empty)",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateSetTransitionsAction{
						Transitions: []platform.StateResourceIdentifier{},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setTransitions",
						"transitions": []
					}
				]
			}`,
		},
		{
			desc: "Unset transitions",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateSetTransitionsAction{},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setTransitions"
					}
				]
			}`,
		},
		{
			desc: "Set roles",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateSetRolesAction{
						Roles: []platform.StateRoleEnum{
							"ReviewIncludedInStatistics",
							"Return",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "setRoles",
						"roles": [
							"ReviewIncludedInStatistics",
							"Return"
						]
					}
				]
			}`,
		},
		{
			desc: "Add roles",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateAddRolesAction{
						Roles: []platform.StateRoleEnum{
							"Return",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "addRoles",
						"roles": [
							"Return"
						]
					}
				]
			}`,
		},
		{
			desc: "Remove roles",
			input: platform.StateUpdate{
				Version: 2,
				Actions: []platform.StateUpdateAction{
					&platform.StateRemoveRolesAction{
						Roles: []platform.StateRoleEnum{
							"ReviewIncludedInStatistics",
						},
					},
				},
			},
			requestBody: `{
				"version": 2,
				"actions": [
					{
						"action": "removeRoles",
						"roles": [
							"ReviewIncludedInStatistics"
						]
					}
				]
			}`,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.ResponseData{StatusCode: 200, Body: `
				{
					"version": 1,
					"actions": [
						{
							"action": "setName",
							"name": {
								"en": "New Name"
							}
						}
					]
				}
			`}

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.
				WithProjectKey("unittest").
				States().
				WithId("1234").
				Post(tC.input).
				Execute(context.TODO())
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/states/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStateDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		States().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyStatesByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/states/1234", output.URL.Path)
}

func TestStateGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2015-01-21T09:22:03.906Z")

	client, server := testutil.MockClient(t, testutil.Fixture("state.json", 200), nil, nil)
	defer server.Close()

	input := platform.State{
		ID:      "7c2e2694-aefe-43d7-888e-6a99514caaca",
		Version: 1,
		Key:     "Initial",
		Type:    platform.StateTypeEnum("LineItemState"),
		Roles:   []platform.StateRoleEnum{},
		Name: &platform.LocalizedString{
			"en": "Initial",
		},
		Description: &platform.LocalizedString{
			"en": "Initial is the first that (custom) line item gets after it's creation",
		},
		BuiltIn:        true,
		Initial:        true,
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.WithProjectKey("unittest").States().WithId("1234").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, *result)
}

func TestStateQuery(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

	client, server := testutil.MockClient(t, response, &output, nil)
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
