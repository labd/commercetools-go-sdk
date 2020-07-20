// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// TypeURLPath is the commercetools API path.
const TypeURLPath = "types"

// TypeCreate creates a new instance of type Type
func (client *Client) TypeCreate(ctx context.Context, draft *TypeDraft, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, TypeURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeQuery allows querying for type Type
func (client *Client) TypeQuery(ctx context.Context, input *QueryInput) (result *TypePagedQueryResponse, err error) {
	err = client.Query(ctx, TypeURLPath, input.toParams(), &result)
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
	err = client.Delete(ctx, endpoint, params, &result)
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
	err = client.Get(ctx, endpoint, params, &result)
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

// TypeUpdateWithKey for type Type
func (client *Client) TypeUpdateWithKey(ctx context.Context, input *TypeUpdateWithKeyInput, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("types/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
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
	err = client.Delete(ctx, endpoint, params, &result)
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
	err = client.Get(ctx, endpoint, params, &result)
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

// TypeUpdateWithID for type Type
func (client *Client) TypeUpdateWithID(ctx context.Context, input *TypeUpdateWithIDInput, opts ...RequestOption) (result *Type, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("types/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
