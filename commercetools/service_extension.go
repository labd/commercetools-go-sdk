// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ExtensionCreate creates a new instance of type Extension
func (client *Client) ExtensionCreate(ctx context.Context, draft *ExtensionDraft, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "extensions"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionQuery allows querying for type Extension
func (client *Client) ExtensionQuery(ctx context.Context, input *QueryInput) (result *ExtensionPagedQueryResponse, err error) {
	endpoint := "extensions"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteWithID for type Extension
func (client *Client) ExtensionDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteWithKey for type Extension
func (client *Client) ExtensionDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionGetWithID Retrieves the representation of an extension by its id.
func (client *Client) ExtensionGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionGetWithKey Retrieves the representation of an extension by its key.
func (client *Client) ExtensionGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
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

func (input *ExtensionUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ExtensionUpdateWithID for type Extension
func (client *Client) ExtensionUpdateWithID(ctx context.Context, input *ExtensionUpdateWithIDInput, opts ...RequestOption) (result *Extension, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("extensions/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
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

func (input *ExtensionUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ExtensionUpdateWithKey for type Extension
func (client *Client) ExtensionUpdateWithKey(ctx context.Context, input *ExtensionUpdateWithKeyInput, opts ...RequestOption) (result *Extension, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("extensions/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
