package commercetools_test

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestStateCreate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.StateDraft
		requestBody string
	}{
		{
			desc: "default",
			input: &commercetools.StateDraft{
				Key:  "state",
				Type: commercetools.StateTypeEnum("ProductState"),
				Name: &commercetools.LocalizedString{
					"en": "State machine",
				},
				Description: &commercetools.LocalizedString{
					"en": "This is a state machine. Beep beep boop.",
				},
				Initial: true,
				Transitions: &[]commercetools.StateResourceIdentifier{
					{
						ID: "1",
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
			input: &commercetools.StateDraft{
				Key:  "state",
				Type: commercetools.StateTypeEnum("ProductState"),
				Name: &commercetools.LocalizedString{
					"en": "State machine",
				},
				Description: &commercetools.LocalizedString{
					"en": "This is a state machine. Beep beep boop.",
				},
				Initial: true,
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
			input: &commercetools.StateDraft{
				Key:  "state",
				Type: commercetools.StateTypeEnum("ProductState"),
				Name: &commercetools.LocalizedString{
					"en": "State machine",
				},
				Description: &commercetools.LocalizedString{
					"en": "This is a state machine. Beep beep boop.",
				},
				Initial:     true,
				Transitions: &[]commercetools.StateResourceIdentifier{},
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

			client, server := testutil.MockClient(t, "{}", &output, nil)
			defer server.Close()

			_, err := client.StateCreate(context.TODO(), tC.input)
			assert.Nil(t, err)

			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStateUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.StateUpdateWithIDInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateChangeKeyAction{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateSetNameAction{
						Name: &commercetools.LocalizedString{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateSetDescriptionAction{
						Description: &commercetools.LocalizedString{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateChangeTypeAction{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateChangeInitialAction{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateSetTransitionsAction{
						Transitions: &[]commercetools.StateResourceIdentifier{
							{
								ID: "1",
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateSetTransitionsAction{
						Transitions: &[]commercetools.StateResourceIdentifier{},
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateSetTransitionsAction{},
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateSetRolesAction{
						Roles: []commercetools.StateRoleEnum{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateAddRolesAction{
						Roles: []commercetools.StateRoleEnum{
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
			input: &commercetools.StateUpdateWithIDInput{
				ID:      "1234",
				Version: 2,
				Actions: []commercetools.StateUpdateAction{
					&commercetools.StateRemoveRolesAction{
						Roles: []commercetools.StateRoleEnum{
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

			client, server := testutil.MockClient(t, testutil.Fixture("state.update.json"), &output, nil)
			defer server.Close()

			_, err := client.StateUpdateWithID(context.TODO(), tC.input)
			assert.Nil(t, err)
			assert.Equal(t, "/unittest/states/1234", output.URL.Path)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}
}

func TestStateDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.StateDeleteWithID(context.TODO(), "1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/states/1234", output.URL.Path)
}

func TestStateGetByID(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2015-01-21T09:22:03.906Z")

	client, server := testutil.MockClient(t, testutil.Fixture("state.json"), nil, nil)
	defer server.Close()

	input := &commercetools.State{
		ID:      "7c2e2694-aefe-43d7-888e-6a99514caaca",
		Version: 1,
		Key:     "Initial",
		Type:    commercetools.StateTypeEnum("LineItemState"),
		Roles:   []commercetools.StateRoleEnum{},
		Name: &commercetools.LocalizedString{
			"en": "Initial",
		},
		Description: &commercetools.LocalizedString{
			"en": "Initial is the first that (custom) line item gets after it's creation",
		},
		BuiltIn:        true,
		Initial:        true,
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	result, err := client.StateGetWithID(context.TODO(), "1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestStateQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.StateQuery(context.TODO(), &queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
