// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// TypeCreate creates a new instance of type Type
func (client *Client) TypeCreate(ctx context.Context, draft *TypeDraft, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "types"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeQuery allows querying for type Type
func (client *Client) TypeQuery(ctx context.Context, input *QueryInput) (result *TypePagedQueryResponse, err error) {
	endpoint := "types"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeDeleteWithID for type Type
func (client *Client) TypeDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("types/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeDeleteWithKey for type Type
func (client *Client) TypeDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("types/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeGetWithID for type Type
func (client *Client) TypeGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("types/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeGetWithKey for type Type
func (client *Client) TypeGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("types/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeUpdateWithIDInput is input for function TypeUpdateWithID
type TypeUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []TypeUpdateAction
}

func (input *TypeUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// TypeUpdateWithID for type Type
func (client *Client) TypeUpdateWithID(ctx context.Context, input *TypeUpdateWithIDInput, opts ...RequestOption) (result *Type, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("types/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeUpdateWithKeyInput is input for function TypeUpdateWithKey
type TypeUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []TypeUpdateAction
}

func (input *TypeUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// TypeUpdateWithKey for type Type
func (client *Client) TypeUpdateWithKey(ctx context.Context, input *TypeUpdateWithKeyInput, opts ...RequestOption) (result *Type, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("types/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
