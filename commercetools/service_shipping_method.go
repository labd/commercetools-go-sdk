// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ShippingMethodURLPath is the commercetools API path.
const ShippingMethodURLPath = "shipping-methods"

// ShippingMethodCreate creates a new instance of type ShippingMethod
func (client *Client) ShippingMethodCreate(ctx context.Context, draft *ShippingMethodDraft, opts ...RequestOption) (result *ShippingMethod, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ShippingMethodURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodQuery allows querying for type ShippingMethod
func (client *Client) ShippingMethodQuery(ctx context.Context, input *QueryInput) (result *ShippingMethodPagedQueryResponse, err error) {
	err = client.Query(ctx, ShippingMethodURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodDeleteWithKey for type ShippingMethod
func (client *Client) ShippingMethodDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *ShippingMethod, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("shipping-methods/key=%s", key)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodGetWithKey for type ShippingMethod
func (client *Client) ShippingMethodGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *ShippingMethod, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("shipping-methods/key=%s", key)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodUpdateWithKeyInput is input for function ShippingMethodUpdateWithKey
type ShippingMethodUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ShippingMethodUpdateAction
}

func (input *ShippingMethodUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ShippingMethodUpdateWithKey for type ShippingMethod
func (client *Client) ShippingMethodUpdateWithKey(ctx context.Context, input *ShippingMethodUpdateWithKeyInput, opts ...RequestOption) (result *ShippingMethod, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("shipping-methods/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodDeleteWithID for type ShippingMethod
func (client *Client) ShippingMethodDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *ShippingMethod, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("shipping-methods/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodGetWithID for type ShippingMethod
func (client *Client) ShippingMethodGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *ShippingMethod, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("shipping-methods/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodUpdateWithIDInput is input for function ShippingMethodUpdateWithID
type ShippingMethodUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ShippingMethodUpdateAction
}

func (input *ShippingMethodUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ShippingMethodUpdateWithID for type ShippingMethod
func (client *Client) ShippingMethodUpdateWithID(ctx context.Context, input *ShippingMethodUpdateWithIDInput, opts ...RequestOption) (result *ShippingMethod, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("shipping-methods/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
