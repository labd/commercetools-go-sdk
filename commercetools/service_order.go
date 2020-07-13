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
func (client *Client) OrderCreate(ctx context.Context, draft *OrderFromCartDraft) (result *Order, err error) {
	err = client.Create(ctx, OrderURLPath, nil, draft, &result)
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
func (client *Client) OrderDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	err = client.Delete(ctx, strings.Replace("orders/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderGetWithID for type Order
func (client *Client) OrderGetWithID(ctx context.Context, ID string) (result *Order, err error) {
	err = client.Get(ctx, strings.Replace("orders/{ID}", "{ID}", ID, 1), nil, &result)
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
func (client *Client) OrderUpdateWithID(ctx context.Context, input *OrderUpdateWithIDInput) (result *Order, err error) {
	err = client.Update(ctx, strings.Replace("orders/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
