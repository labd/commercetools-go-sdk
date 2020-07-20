// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// OrderURLPath is the commercetools API path.
const OrderURLPath = "orders"

// OrderCreate creates a new instance of type Order
func (client *Client) OrderCreate(ctx context.Context, draft *OrderFromCartDraft, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Create(ctx, OrderURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderQuery allows querying for type Order
func (client *Client) OrderQuery(ctx context.Context, input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	err = client.Query(ctx, OrderURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderDeleteWithID for type Order
func (client *Client) OrderDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("orders/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderGetWithID for type Order
func (client *Client) OrderGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("orders/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderUpdateWithIDInput is input for function OrderUpdateWithID
type OrderUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []OrderUpdateAction
}

// OrderUpdateWithID for type Order
func (client *Client) OrderUpdateWithID(ctx context.Context, input *OrderUpdateWithIDInput, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("orders/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
