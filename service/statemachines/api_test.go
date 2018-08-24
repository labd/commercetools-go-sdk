package statemachines_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/service/statemachines"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestStateCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()
	svc := statemachines.New(client)

	input := &statemachines.StateDraft{
		Key:  "state",
		Type: statemachines.ProductStateType,
		Name: commercetools.LocalizedString{
			"en": "State machine",
		},
		Description: commercetools.LocalizedString{
			"en": "This is a state machine. Beep beep boop.",
		},
		Initial: true,
		Transitions: []commercetools.Reference{
			{
				TypeID: "1",
				ID:     "1",
			},
		},
	}

	fmt.Println(output)

	_, err := svc.Create(input)
	assert.Nil(t, err)

	expectedBody := `{
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
				"typeId": "1",
				"id": "1"
			}
		]
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestStateUpdate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *statemachines.UpdateInput
		requestBody string
	}{
		{
			desc: "Change key",
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.ChangeKey{
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.SetName{
						Name: commercetools.LocalizedString{
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.SetDescription{
						Description: commercetools.LocalizedString{
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.ChangeType{
						Type: statemachines.OrderStateType,
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.ChangeInitial{
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.SetTransitions{
						Transitions: []commercetools.Reference{
							{
								TypeID: "1",
								ID:     "1",
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
								"typeId": "1",
								"id": "1"
							}
						]
					}
				]
			}`,
		},
		{
			desc: "Set roles",
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.SetRoles{
						Roles: []statemachines.StateRole{
							statemachines.ReviewIncludedInStatisticsRole,
							statemachines.ReturnRole,
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.AddRoles{
						Roles: []statemachines.StateRole{
							statemachines.ReturnRole,
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
			input: &statemachines.UpdateInput{
				ID:      "1234",
				Version: 2,
				Actions: commercetools.UpdateActions{
					&statemachines.RemoveRoles{
						Roles: []statemachines.StateRole{
							statemachines.ReviewIncludedInStatisticsRole,
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
			svc := statemachines.New(client)

			_, err := svc.Update(tC.input)
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
	svc := statemachines.New(client)

	_, err := svc.DeleteByID("1234", 2)
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

	input := &statemachines.State{
		ID:      "7c2e2694-aefe-43d7-888e-6a99514caaca",
		Version: 1,
		Key:     "Initial",
		Type:    statemachines.LineItemStateType,
		Roles:   []statemachines.StateRole{},
		Name: commercetools.LocalizedString{
			"en": "Initial",
		},
		Description: commercetools.LocalizedString{
			"en": "Initial is the first that (custom) line item gets after it's creation",
		},
		BuiltIn:        true,
		Initial:        true,
		CreatedAt:      timestamp,
		LastModifiedAt: timestamp,
	}

	svc := statemachines.New(client)
	result, err := svc.GetByID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}
