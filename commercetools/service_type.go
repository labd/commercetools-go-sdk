// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// TypeURLPath is the commercetools API path.
const TypeURLPath = "types"

// TypeCreate creates a new instance of type Type
func (client *Client) TypeCreate(ctx context.Context, draft *TypeDraft) (result *Type, err error) {
	err = client.Create(ctx, TypeURLPath, nil, draft, &result)
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
func (client *Client) TypeDeleteWithKey(ctx context.Context, key string, version int) (result *Type, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("types/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeGetWithKey for type Type
func (client *Client) TypeGetWithKey(ctx context.Context, key string) (result *Type, err error) {
	err = client.Get(ctx, strings.Replace("types/key={key}", "{key}", key, 1), nil, &result)
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
func (client *Client) TypeUpdateWithKey(ctx context.Context, input *TypeUpdateWithKeyInput) (result *Type, err error) {
	err = client.Update(ctx, strings.Replace("types/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeDeleteWithID for type Type
func (client *Client) TypeDeleteWithID(ctx context.Context, ID string, version int) (result *Type, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("types/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TypeGetWithID for type Type
func (client *Client) TypeGetWithID(ctx context.Context, ID string) (result *Type, err error) {
	err = client.Get(ctx, strings.Replace("types/{ID}", "{ID}", ID, 1), nil, &result)
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
func (client *Client) TypeUpdateWithID(ctx context.Context, input *TypeUpdateWithIDInput) (result *Type, err error) {
	err = client.Update(ctx, strings.Replace("types/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
