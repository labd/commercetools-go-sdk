package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// DeleteInput provides the data required to delete a state machine.
type StateDeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a state machine.
type StateUpdateInput struct {
	ID string

	// The expected version of the state on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the state.
	Actions []StateUpdateAction
}

// GetByID will return a state matching the provided ID. OAuth2 Scopes:
// view_states:{projectKey} (or, deprecated: view_orders:{projectKey})
func (client *Client) StateMachineGetByID(id string) (result *State, err error) {
	err = client.Get(fmt.Sprintf("states/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create will create a new state from a draft, and return the newly created
// state. OAuth2 Scopes: manage_states:{projectKey} (or, deprecated:
// manage_orders:{projectKey})
func (client *Client) StateMachineCreate(draft *StateDraft) (result *State, err error) {
	err = client.Create("states", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update will update a state matching the provided ID with the defined
// UpdateActions. OAuth2 Scopes: manage_states:{projectKey} (or, deprecated:
// manage_orders:{projectKey})
func (client *Client) StateMachineUpdate(input *StateUpdateInput) (result *State, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid state id passed")
	}

	endpoint := fmt.Sprintf("states/%s", input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteByID will delete a state matching the provided ID. OAuth2 Scopes:
// manage_states:{projectKey} (or, deprecated: manage_orders:{projectKey})
func (client *Client) StateMachineDeleteByID(id string, version int) (result *State, err error) {
	endpoint := fmt.Sprintf("states/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
