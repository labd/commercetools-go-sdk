// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ExtensionURLPath is the commercetools API path.
const ExtensionURLPath = "extensions"

// ExtensionCreate creates a new instance of type Extension
func (client *Client) ExtensionCreate(ctx context.Context, draft *ExtensionDraft) (result *Extension, err error) {
	err = client.Create(ctx, ExtensionURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionQuery allows querying for type Extension
func (client *Client) ExtensionQuery(ctx context.Context, input *QueryInput) (result *ExtensionPagedQueryResponse, err error) {
	err = client.Query(ctx, ExtensionURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteWithKey for type Extension
func (client *Client) ExtensionDeleteWithKey(ctx context.Context, key string, version int) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("extensions/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionGetWithKey Retrieves the representation of an extension by its key.
func (client *Client) ExtensionGetWithKey(ctx context.Context, key string) (result *Extension, err error) {
	err = client.Get(ctx, strings.Replace("extensions/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionUpdateWithKeyInput is input for function ExtensionUpdateWithKey
type ExtensionUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ExtensionUpdateAction
}

// ExtensionUpdateWithKey for type Extension
func (client *Client) ExtensionUpdateWithKey(ctx context.Context, input *ExtensionUpdateWithKeyInput) (result *Extension, err error) {
	err = client.Update(ctx, strings.Replace("extensions/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteWithID for type Extension
func (client *Client) ExtensionDeleteWithID(ctx context.Context, ID string, version int) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("extensions/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionGetWithID Retrieves the representation of an extension by its id.
func (client *Client) ExtensionGetWithID(ctx context.Context, ID string) (result *Extension, err error) {
	err = client.Get(ctx, strings.Replace("extensions/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionUpdateWithIDInput is input for function ExtensionUpdateWithID
type ExtensionUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ExtensionUpdateAction
}

// ExtensionUpdateWithID for type Extension
func (client *Client) ExtensionUpdateWithID(ctx context.Context, input *ExtensionUpdateWithIDInput) (result *Extension, err error) {
	err = client.Update(ctx, strings.Replace("extensions/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
