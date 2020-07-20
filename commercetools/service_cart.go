// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CartURLPath is the commercetools API path.
const CartURLPath = "carts"

// CartCreate creates a new instance of type Cart
func (client *Client) CartCreate(ctx context.Context, draft *CartDraft, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, CartURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartQuery allows querying for type Cart
func (client *Client) CartQuery(ctx context.Context, input *QueryInput) (result *CartPagedQueryResponse, err error) {
	err = client.Query(ctx, CartURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDeleteWithID for type Cart
func (client *Client) CartDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CartGetWithID The cart may not contain up-to-date prices, discounts etc.
If you want to ensure they’re up-to-date, send an Update request with the Recalculate update action instead.
*/
func (client *Client) CartGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartUpdateWithIDInput is input for function CartUpdateWithID
type CartUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartUpdateAction
}

// CartUpdateWithID for type Cart
func (client *Client) CartUpdateWithID(ctx context.Context, input *CartUpdateWithIDInput, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
