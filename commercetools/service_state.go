// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// StateURLPath is the commercetools API path.
const StateURLPath = "states"

// StateCreate creates a new instance of type State
func (client *Client) StateCreate(ctx context.Context, draft *StateDraft, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, StateURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateQuery allows querying for type State
func (client *Client) StateQuery(ctx context.Context, input *QueryInput) (result *StatePagedQueryResponse, err error) {
	err = client.Query(ctx, StateURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateDeleteWithID for type State
func (client *Client) StateDeleteWithID(ctx context.Context, ID string, version int, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("states/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateGetWithID for type State
func (client *Client) StateGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("states/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateUpdateWithIDInput is input for function StateUpdateWithID
type StateUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []StateUpdateAction
}

// StateUpdateWithID for type State
func (client *Client) StateUpdateWithID(ctx context.Context, input *StateUpdateWithIDInput, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("states/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
