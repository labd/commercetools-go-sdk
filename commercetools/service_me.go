// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// MyActiveCart for type
func (client *Client) MyActiveCart(ctx context.Context, opts ...RequestOption) (result *MyCart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/active-cart"
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyLogin for type
func (client *Client) MyLogin(ctx context.Context, opts ...RequestOption) (result *CustomerSignInResult, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/login"
	err = client.create(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyReset for type
func (client *Client) MyReset(ctx context.Context, opts ...RequestOption) (result *MyCustomer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/password/reset"
	err = client.create(ctx, endpoint, params, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MySignup for type MyCustomerDraft
func (client *Client) MySignup(ctx context.Context, value *MyCustomerDraft, opts ...RequestOption) (result *CustomerSignInResult, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/signup"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartCreate creates a new instance of type MyCart
func (client *Client) MyCartCreate(ctx context.Context, draft *MyCartDraft, opts ...RequestOption) (result *MyCart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/carts"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartQuery allows querying for type
func (client *Client) MyCartQuery(ctx context.Context, input *QueryInput) (result *CartPagedQueryResponse, err error) {
	endpoint := "me/carts"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartDeleteWithID for type MyCart
func (client *Client) MyCartDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *MyCart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/carts/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartDeleteWithKey for type MyCart
func (client *Client) MyCartDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *MyCart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/carts/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartGetWithID for type MyCart
func (client *Client) MyCartGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *MyCart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/carts/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartGetWithKey for type MyCart
func (client *Client) MyCartGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *MyCart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/carts/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartUpdateWithIDInput is input for function MyCartUpdateWithID
type MyCartUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyCartUpdateAction
}

func (input *MyCartUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// MyCartUpdateWithID for type MyCart
func (client *Client) MyCartUpdateWithID(ctx context.Context, input *MyCartUpdateWithIDInput, opts ...RequestOption) (result *MyCart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("me/carts/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyCartUpdateWithKeyInput is input for function MyCartUpdateWithKey
type MyCartUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []MyCartUpdateAction
}

func (input *MyCartUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// MyCartUpdateWithKey for type MyCart
func (client *Client) MyCartUpdateWithKey(ctx context.Context, input *MyCartUpdateWithKeyInput, opts ...RequestOption) (result *MyCart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("me/carts/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyOrderCreate creates a new instance of type MyOrder
func (client *Client) MyOrderCreate(ctx context.Context, draft *MyOrderFromCartDraft, opts ...RequestOption) (result *MyOrder, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/orders"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyOrderQuery allows querying for type
func (client *Client) MyOrderQuery(ctx context.Context, input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	endpoint := "me/orders"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyOrderGetWithID for type MyOrder
func (client *Client) MyOrderGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *MyOrder, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/orders/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentCreate creates a new instance of type MyPayment
func (client *Client) MyPaymentCreate(ctx context.Context, draft *MyPaymentDraft, opts ...RequestOption) (result *MyPayment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/payments"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentQuery allows querying for type
func (client *Client) MyPaymentQuery(ctx context.Context, input *QueryInput) (result *MyPaymentPagedQueryResponse, err error) {
	endpoint := "me/payments"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentDeleteWithID for type MyPayment
func (client *Client) MyPaymentDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *MyPayment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/payments/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentDeleteWithKey for type MyPayment
func (client *Client) MyPaymentDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *MyPayment, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/payments/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentGetWithID for type MyPayment
func (client *Client) MyPaymentGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *MyPayment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/payments/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentGetWithKey for type MyPayment
func (client *Client) MyPaymentGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *MyPayment, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/payments/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentUpdateWithIDInput is input for function MyPaymentUpdateWithID
type MyPaymentUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyPaymentUpdateAction
}

func (input *MyPaymentUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// MyPaymentUpdateWithID for type MyPayment
func (client *Client) MyPaymentUpdateWithID(ctx context.Context, input *MyPaymentUpdateWithIDInput, opts ...RequestOption) (result *MyPayment, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("me/payments/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyPaymentUpdateWithKeyInput is input for function MyPaymentUpdateWithKey
type MyPaymentUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []MyPaymentUpdateAction
}

func (input *MyPaymentUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// MyPaymentUpdateWithKey for type MyPayment
func (client *Client) MyPaymentUpdateWithKey(ctx context.Context, input *MyPaymentUpdateWithKeyInput, opts ...RequestOption) (result *MyPayment, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("me/payments/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListCreate creates a new instance of type MyShoppingList
func (client *Client) MyShoppingListCreate(ctx context.Context, draft *MyShoppingListDraft, opts ...RequestOption) (result *MyShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "me/shopping-lists"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListQuery allows querying for type
func (client *Client) MyShoppingListQuery(ctx context.Context, input *QueryInput) (result *ShoppingListPagedQueryResponse, err error) {
	endpoint := "me/shopping-lists"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListDeleteWithID for type MyShoppingList
func (client *Client) MyShoppingListDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *MyShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/shopping-lists/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListDeleteWithKey for type MyShoppingList
func (client *Client) MyShoppingListDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *MyShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/shopping-lists/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListGetWithID for type MyShoppingList
func (client *Client) MyShoppingListGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *MyShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/shopping-lists/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListGetWithKey for type MyShoppingList
func (client *Client) MyShoppingListGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *MyShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("me/shopping-lists/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListUpdateWithIDInput is input for function MyShoppingListUpdateWithID
type MyShoppingListUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []MyShoppingListUpdateAction
}

func (input *MyShoppingListUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// MyShoppingListUpdateWithID for type MyShoppingList
func (client *Client) MyShoppingListUpdateWithID(ctx context.Context, input *MyShoppingListUpdateWithIDInput, opts ...RequestOption) (result *MyShoppingList, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("me/shopping-lists/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MyShoppingListUpdateWithKeyInput is input for function MyShoppingListUpdateWithKey
type MyShoppingListUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []MyShoppingListUpdateAction
}

func (input *MyShoppingListUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// MyShoppingListUpdateWithKey for type MyShoppingList
func (client *Client) MyShoppingListUpdateWithKey(ctx context.Context, input *MyShoppingListUpdateWithKeyInput, opts ...RequestOption) (result *MyShoppingList, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("me/shopping-lists/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
