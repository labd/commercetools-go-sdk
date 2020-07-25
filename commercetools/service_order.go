// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// OrderCreate creates a new instance of type Order
func (client *Client) OrderCreate(ctx context.Context, draft *OrderFromCartDraft, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "orders"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderQuery allows querying for type Order
func (client *Client) OrderQuery(ctx context.Context, input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	endpoint := "orders"
	err = client.query(ctx, endpoint, input.toParams(), &result)
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
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderDeleteWithOrderNumber for type Order
func (client *Client) OrderDeleteWithOrderNumber(ctx context.Context, orderNumber string, version int, dataErasure bool, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/order-number=%s", orderNumber)
	err = client.delete(ctx, endpoint, params, &result)
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
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
OrderGetWithOrderNumber In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
it should be provided in URL-encoded format.
*/
func (client *Client) OrderGetWithOrderNumber(ctx context.Context, orderNumber string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/order-number=%s", orderNumber)
	err = client.get(ctx, endpoint, params, &result)
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
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
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
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderImport for type OrderImportDraft
func (client *Client) OrderImport(ctx context.Context, value *OrderImportDraft, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "orders/import"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditCreate creates a new instance of type OrderEdit
func (client *Client) OrderEditCreate(ctx context.Context, draft *OrderEditDraft, opts ...RequestOption) (result *OrderEdit, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "orders/edits"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditQuery allows querying for type Order
func (client *Client) OrderEditQuery(ctx context.Context, input *QueryInput) (result *OrderEditPagedQueryResponse, err error) {
	endpoint := "orders/edits"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditDeleteWithID for type OrderEdit
func (client *Client) OrderEditDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *OrderEdit, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/edits/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditDeleteWithKey for type OrderEdit
func (client *Client) OrderEditDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *OrderEdit, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/edits/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditGetWithID for type OrderEdit
func (client *Client) OrderEditGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *OrderEdit, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/edits/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditGetWithKey for type OrderEdit
func (client *Client) OrderEditGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *OrderEdit, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("orders/edits/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditUpdateWithIDInput is input for function OrderEditUpdateWithID
type OrderEditUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []OrderEditUpdateAction
}

func (input *OrderEditUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// OrderEditUpdateWithID for type OrderEdit
func (client *Client) OrderEditUpdateWithID(ctx context.Context, input *OrderEditUpdateWithIDInput, opts ...RequestOption) (result *OrderEdit, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("orders/edits/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditUpdateWithKeyInput is input for function OrderEditUpdateWithKey
type OrderEditUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []OrderEditUpdateAction
}

func (input *OrderEditUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// OrderEditUpdateWithKey for type OrderEdit
func (client *Client) OrderEditUpdateWithKey(ctx context.Context, input *OrderEditUpdateWithKeyInput, opts ...RequestOption) (result *OrderEdit, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("orders/edits/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OrderEditApply for type OrderEditApply
func (client *Client) OrderEditApply(ctx context.Context, value *OrderEditApply, opts ...RequestOption) (result *OrderEdit, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "orders/edits/apply"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
