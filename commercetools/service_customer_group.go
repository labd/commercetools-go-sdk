// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CustomerGroupCreate creates a new instance of type CustomerGroup
func (client *Client) CustomerGroupCreate(ctx context.Context, draft *CustomerGroupDraft, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "customer-groups"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupQuery allows querying for type CustomerGroup
func (client *Client) CustomerGroupQuery(ctx context.Context, input *QueryInput) (result *CustomerGroupPagedQueryResponse, err error) {
	endpoint := "customer-groups"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupDeleteWithID for type CustomerGroup
func (client *Client) CustomerGroupDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customer-groups/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupDeleteWithKey for type CustomerGroup
func (client *Client) CustomerGroupDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customer-groups/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupGetWithID for type CustomerGroup
func (client *Client) CustomerGroupGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customer-groups/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupGetWithKey Gets a customer group by Key.
func (client *Client) CustomerGroupGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *CustomerGroup, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("customer-groups/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupUpdateWithIDInput is input for function CustomerGroupUpdateWithID
type CustomerGroupUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CustomerGroupUpdateAction
}

func (input *CustomerGroupUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CustomerGroupUpdateWithID for type CustomerGroup
func (client *Client) CustomerGroupUpdateWithID(ctx context.Context, input *CustomerGroupUpdateWithIDInput, opts ...RequestOption) (result *CustomerGroup, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("customer-groups/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomerGroupUpdateWithKeyInput is input for function CustomerGroupUpdateWithKey
type CustomerGroupUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CustomerGroupUpdateAction
}

func (input *CustomerGroupUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CustomerGroupUpdateWithKey for type CustomerGroup
func (client *Client) CustomerGroupUpdateWithKey(ctx context.Context, input *CustomerGroupUpdateWithKeyInput, opts ...RequestOption) (result *CustomerGroup, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("customer-groups/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
