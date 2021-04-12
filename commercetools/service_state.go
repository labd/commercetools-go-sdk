// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// StateCreate creates a new instance of type State
func (client *Client) StateCreate(ctx context.Context, draft *StateDraft, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "states"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateQuery allows querying for type State
func (client *Client) StateQuery(ctx context.Context, input *QueryInput) (result *StatePagedQueryResponse, err error) {
	endpoint := "states"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateDeleteWithID for type State
func (client *Client) StateDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("states/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateDeleteWithKey for type State
func (client *Client) StateDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("states/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateGetWithID for type State
func (client *Client) StateGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("states/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateGetWithKey for type State
func (client *Client) StateGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *State, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("states/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
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

func (input *StateUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StateUpdateWithID for type State
func (client *Client) StateUpdateWithID(ctx context.Context, input *StateUpdateWithIDInput, opts ...RequestOption) (result *State, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("states/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StateUpdateWithKeyInput is input for function StateUpdateWithKey
type StateUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []StateUpdateAction
}

func (input *StateUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StateUpdateWithKey for type State
func (client *Client) StateUpdateWithKey(ctx context.Context, input *StateUpdateWithKeyInput, opts ...RequestOption) (result *State, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("states/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
