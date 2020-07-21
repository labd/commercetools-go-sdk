// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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

// OrderImportOrder Create an Order by Import
func (client *Client) OrderImportOrder(ctx context.Context, value *OrderImportDraft, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "orders/import"
	err = client.Create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderDeleteWithOrderNumber for type Order
func (client *Client) OrderDeleteWithOrderNumber(ctx context.Context, ordernumber string, version int, dataErasure bool, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/order-number=%s", ordernumber)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
OrderGetWithOrderNumber In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
it should be provided in URL-encoded format.
*/
func (client *Client) OrderGetWithOrderNumber(ctx context.Context, ordernumber string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/order-number=%s", ordernumber)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderUpdateWithOrderNumberInput is input for function OrderUpdateWithOrderNumber
type OrderUpdateWithOrderNumberInput struct {
	OrderNumber string
	Version     int
	Actions     []OrderUpdateAction
}

func (input *OrderUpdateWithOrderNumberInput) Validate() error {
	if input.OrderNumber == "" {
		return fmt.Errorf("no valid value for OrderNumber given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// OrderUpdateWithOrderNumber for type Order
func (client *Client) OrderUpdateWithOrderNumber(ctx context.Context, input *OrderUpdateWithOrderNumberInput, opts ...RequestOption) (result *Order, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("orders/order-number=%s", input.OrderNumber)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderDeleteWithID for type Order
func (client *Client) OrderDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderGetWithID for type Order
func (client *Client) OrderGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
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

func (input *OrderUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// OrderUpdateWithID for type Order
func (client *Client) OrderUpdateWithID(ctx context.Context, input *OrderUpdateWithIDInput, opts ...RequestOption) (result *Order, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("orders/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
