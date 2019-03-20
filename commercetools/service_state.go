package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// StateDeleteInput provides the data required to delete a state machine.
type StateDeleteInput struct {
	ID      string
	Version int
}

// StateUpdateInput provides the data required to update a state machine.
type StateUpdateInput struct {
	ID string

	// The expected version of the state on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict
	// will be returned.
	Version int

	// The list of update actions to be performed on the state.
	Actions []StateUpdateAction
}

// StateURLPath is the commercetools API state path.
const StateURLPath = "states"

// StateGetByID will return a state matching the provided ID. OAuth2 Scopes:
// view_states:{projectKey} (or, deprecated: view_orders:{projectKey})
func (client *Client) StateGetByID(id string) (result *State, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", StateURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateCreate will create a new state from a draft, and return the newly
// created state. OAuth2 Scopes: manage_states:{projectKey} (or, deprecated:
// manage_orders:{projectKey})
func (client *Client) StateCreate(draft *StateDraft) (result *State, err error) {
	err = client.Create(StateURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateUpdate will update a state matching the provided ID with the
// defined StateUpdateActions. OAuth2 Scopes: manage_states:{projectKey}
// (or, deprecated: manage_orders:{projectKey})
func (client *Client) StateUpdate(input *StateUpdateInput) (result *State, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid state id passed")
	}

	endpoint := fmt.Sprintf("%s/%s", StateURLPath, input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateDeleteByID will delete a state matching the provided ID. OAuth2
// Scopes: manage_states:{projectKey} (or, deprecated:
// manage_orders:{projectKey})
func (client *Client) StateDeleteByID(id string, version int) (result *State, err error) {
	endpoint := fmt.Sprintf("%s/%s", StateURLPath, id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateQuery will query state objects.
// OAuth2 Scopes: view_states:{projectKey}
func (client *Client) StateQuery(input *QueryInput) (result *StatePagedQueryResponse, err error) {
	err = client.Query(StateURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
