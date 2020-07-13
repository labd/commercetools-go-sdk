// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ShippingMethodURLPath is the commercetools API path.
const ShippingMethodURLPath = "shipping-methods"

// ShippingMethodCreate creates a new instance of type ShippingMethod
func (client *Client) ShippingMethodCreate(ctx context.Context, draft *ShippingMethodDraft) (result *ShippingMethod, err error) {
	err = client.Create(ctx, ShippingMethodURLPath, nil, draft, &result)
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
func (client *Client) ShippingMethodDeleteWithKey(ctx context.Context, key string, version int) (result *ShippingMethod, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("shipping-methods/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodGetWithKey for type ShippingMethod
func (client *Client) ShippingMethodGetWithKey(ctx context.Context, key string) (result *ShippingMethod, err error) {
	err = client.Get(ctx, strings.Replace("shipping-methods/key={key}", "{key}", key, 1), nil, &result)
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

// ShippingMethodUpdateWithKey for type ShippingMethod
func (client *Client) ShippingMethodUpdateWithKey(ctx context.Context, input *ShippingMethodUpdateWithKeyInput) (result *ShippingMethod, err error) {
	err = client.Update(ctx, strings.Replace("shipping-methods/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodDeleteWithID for type ShippingMethod
func (client *Client) ShippingMethodDeleteWithID(ctx context.Context, ID string, version int) (result *ShippingMethod, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("shipping-methods/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShippingMethodGetWithID for type ShippingMethod
func (client *Client) ShippingMethodGetWithID(ctx context.Context, ID string) (result *ShippingMethod, err error) {
	err = client.Get(ctx, strings.Replace("shipping-methods/{ID}", "{ID}", ID, 1), nil, &result)
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

// ShippingMethodUpdateWithID for type ShippingMethod
func (client *Client) ShippingMethodUpdateWithID(ctx context.Context, input *ShippingMethodUpdateWithIDInput) (result *ShippingMethod, err error) {
	err = client.Update(ctx, strings.Replace("shipping-methods/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
