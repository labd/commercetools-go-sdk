// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// StoreLogin Authenticate Customer (Sign In) in store
func (client *Client) StoreLogin(ctx context.Context, storeKey string, value *CustomerSignin, opts ...RequestOption) (result *CustomerSignInResult, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/login", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShippingMethodsForMatchingCartInput is input for function StoreShippingMethodsForMatchingCart
type StoreShippingMethodsForMatchingCartInput struct {
	CartID string `url:"cartId"`
}

// StoreShippingMethodsForMatchingCart for type StoreShippingMethodsForMatchingCartInput
func (client *Client) StoreShippingMethodsForMatchingCart(ctx context.Context, storeKey string, value *StoreShippingMethodsForMatchingCartInput, opts ...RequestOption) (result *ShippingMethodPagedQueryResponse, err error) {
	params := serializeQueryParams(value)
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/shipping-methods/matching-cart", storeKey)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartCreate creates a new instance of type Cart
func (client *Client) StoreCartCreate(ctx context.Context, storeKey string, draft *CartDraft, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/carts", storeKey)
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartQuery allows querying for type
// Queries carts in a specific Store. The {storeKey} path parameter maps to a Store's key.
func (client *Client) StoreCartQuery(ctx context.Context, storeKey string, input *QueryInput) (result *CartPagedQueryResponse, err error) {
	endpoint := fmt.Sprintf("in-store/key=%s/carts", storeKey)
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartDeleteWithID for type Cart
func (client *Client) StoreCartDeleteWithID(ctx context.Context, storeKey string, id string, version int, dataErasure bool, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/carts/%s", storeKey, id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartDeleteWithKey for type Cart
func (client *Client) StoreCartDeleteWithKey(ctx context.Context, storeKey string, key string, version int, dataErasure bool, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/carts/key=%s", storeKey, key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCartGetWithCustomerID Retrieves the active cart of the customer that has been modified most recently in a specific Store.
The {storeKey} path parameter maps to a Store's key.

If the cart exists in the commercetools project but does not have the store field, or the store field
references a different store, this method returns a ResourceNotFound error.

The cart may not contain up-to-date prices, discounts etc. If you want to ensure they're up-to-date,
send an Update request with the Recalculate update action instead.
*/
func (client *Client) StoreCartGetWithCustomerID(ctx context.Context, storeKey string, customerId string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/carts/customer-id=%s", storeKey, customerId)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCartGetWithID Returns a cart by its ID from a specific Store. The {storeKey} path parameter maps to a Store's key.
If the cart exists in the commercetools project but does not have the store field,
or the store field references a different store, this method returns a ResourceNotFound error.
The cart may not contain up-to-date prices, discounts etc.
If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*/
func (client *Client) StoreCartGetWithID(ctx context.Context, storeKey string, id string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/carts/%s", storeKey, id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCartGetWithKey Returns a cart by its key from a specific Store. The {storeKey} path parameter maps to a Store's key.
If the cart exists in the commercetools project but does not have the store field,
or the store field references a different store, this method returns a ResourceNotFound error.
The cart may not contain up-to-date prices, discounts etc.
If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*/
func (client *Client) StoreCartGetWithKey(ctx context.Context, storeKey string, key string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/carts/key=%s", storeKey, key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartUpdateWithIDInput is input for function StoreCartUpdateWithID
type StoreCartUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartUpdateAction
}

func (input *StoreCartUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreCartUpdateWithID for type Cart
func (client *Client) StoreCartUpdateWithID(ctx context.Context, storeKey string, input *StoreCartUpdateWithIDInput, opts ...RequestOption) (result *Cart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/carts/%s", storeKey, input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartUpdateWithKeyInput is input for function StoreCartUpdateWithKey
type StoreCartUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CartUpdateAction
}

func (input *StoreCartUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreCartUpdateWithKey for type Cart
func (client *Client) StoreCartUpdateWithKey(ctx context.Context, storeKey string, input *StoreCartUpdateWithKeyInput, opts ...RequestOption) (result *Cart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/carts/key=%s", storeKey, input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCartReplicate for type ReplicaCartDraft
func (client *Client) StoreCartReplicate(ctx context.Context, storeKey string, value *ReplicaCartDraft, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/carts/replicate", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerCreate creates a new instance of type Customer
func (client *Client) StoreCustomerCreate(ctx context.Context, storeKey string, draft *CustomerDraft, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers", storeKey)
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerQuery allows querying for type
func (client *Client) StoreCustomerQuery(ctx context.Context, storeKey string, input *QueryInput) (result *CustomerPagedQueryResponse, err error) {
	endpoint := fmt.Sprintf("in-store/key=%s/customers", storeKey)
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerDeleteWithID for type Customer
func (client *Client) StoreCustomerDeleteWithID(ctx context.Context, storeKey string, id string, version int, dataErasure bool, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/customers/%s", storeKey, id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerDeleteWithKey for type Customer
func (client *Client) StoreCustomerDeleteWithKey(ctx context.Context, storeKey string, key string, version int, dataErasure bool, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/customers/key=%s", storeKey, key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerGetWithEmailToken for type Customer
func (client *Client) StoreCustomerGetWithEmailToken(ctx context.Context, storeKey string, emailToken string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/customers/email-token=%s", storeKey, emailToken)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCustomerGetWithID Returns a customer by its ID from a specific Store. The {storeKey} path parameter maps to a Store's key.
It also considers customers that do not have the stores field.
If the customer exists in the commercetools project but the stores field references different stores,
this method returns a ResourceNotFound error.
*/
func (client *Client) StoreCustomerGetWithID(ctx context.Context, storeKey string, id string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/customers/%s", storeKey, id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCustomerGetWithKey Returns a customer by its Key from a specific Store. The {storeKey} path parameter maps to a Store's key.
It also considers customers that do not have the stores field.
If the customer exists in the commercetools project but the stores field references different stores,
this method returns a ResourceNotFound error.
*/
func (client *Client) StoreCustomerGetWithKey(ctx context.Context, storeKey string, key string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/customers/key=%s", storeKey, key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerGetWithPasswordToken for type Customer
func (client *Client) StoreCustomerGetWithPasswordToken(ctx context.Context, storeKey string, passwordToken string, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/customers/password-token=%s", storeKey, passwordToken)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerUpdateWithIDInput is input for function StoreCustomerUpdateWithID
type StoreCustomerUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CustomerUpdateAction
}

func (input *StoreCustomerUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreCustomerUpdateWithID for type Customer
func (client *Client) StoreCustomerUpdateWithID(ctx context.Context, storeKey string, input *StoreCustomerUpdateWithIDInput, opts ...RequestOption) (result *Customer, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/%s", storeKey, input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerUpdateWithKeyInput is input for function StoreCustomerUpdateWithKey
type StoreCustomerUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CustomerUpdateAction
}

func (input *StoreCustomerUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreCustomerUpdateWithKey for type Customer
func (client *Client) StoreCustomerUpdateWithKey(ctx context.Context, storeKey string, input *StoreCustomerUpdateWithKeyInput, opts ...RequestOption) (result *Customer, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/key=%s", storeKey, input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCustomerEmailToken To verify a customer's email, an email token can be created. This should be embedded in a link and sent to the
customer via email. When the customer clicks on the link,
the "verify customer's email" endpoint should be called,
which sets customer's isVerifiedEmail field to true.
*/
func (client *Client) StoreCustomerEmailToken(ctx context.Context, storeKey string, value *CustomerCreateEmailToken, opts ...RequestOption) (result *CustomerToken, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/email-token", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerEmailconfirm for type CustomerEmailVerify
func (client *Client) StoreCustomerEmailconfirm(ctx context.Context, storeKey string, value *CustomerEmailVerify, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/email/confirm", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerPassword for type CustomerChangePassword
func (client *Client) StoreCustomerPassword(ctx context.Context, storeKey string, value *CustomerChangePassword, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/password", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreCustomerPasswordToken The following workflow can be used to reset the customer's password:

* Create a password reset token and send it embedded in a link to the customer.
* When the customer clicks on the link, the customer is retrieved with the token.
* The customer enters a new password and the "reset customer's password" endpoint is called.
*/
func (client *Client) StoreCustomerPasswordToken(ctx context.Context, storeKey string, value *CustomerCreatePasswordResetToken, opts ...RequestOption) (result *CustomerToken, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/password-token", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCustomerPasswordreset for type CustomerResetPassword
func (client *Client) StoreCustomerPasswordreset(ctx context.Context, storeKey string, value *CustomerResetPassword, opts ...RequestOption) (result *Customer, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/customers/password/reset", storeKey)
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyActiveCart for type
func (client *Client) StoreMyActiveCart(ctx context.Context, storeKey string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/me/active-cart", storeKey)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyCartCreate creates a new instance of type Cart
func (client *Client) StoreMyCartCreate(ctx context.Context, storeKey string, draft *MyCartDraft, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/me/carts", storeKey)
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyCartQuery allows querying for type
func (client *Client) StoreMyCartQuery(ctx context.Context, storeKey string, input *QueryInput) (result *CartPagedQueryResponse, err error) {
	endpoint := fmt.Sprintf("in-store/key=%s/me/carts", storeKey)
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyCartDeleteWithID for type Cart
func (client *Client) StoreMyCartDeleteWithID(ctx context.Context, storeKey string, id string, version int, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/me/carts/%s", storeKey, id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyCartGetWithID for type Cart
func (client *Client) StoreMyCartGetWithID(ctx context.Context, storeKey string, id string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/me/carts/%s", storeKey, id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyCartUpdateWithIDInput is input for function StoreMyCartUpdateWithID
type StoreMyCartUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartUpdateAction
}

func (input *StoreMyCartUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreMyCartUpdateWithID for type Cart
func (client *Client) StoreMyCartUpdateWithID(ctx context.Context, storeKey string, input *StoreMyCartUpdateWithIDInput, opts ...RequestOption) (result *Cart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/me/carts/%s", storeKey, input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyOrderCreate creates a new instance of type Order
func (client *Client) StoreMyOrderCreate(ctx context.Context, storeKey string, draft *MyOrderFromCartDraft, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/me/orders", storeKey)
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyOrderQuery allows querying for type
func (client *Client) StoreMyOrderQuery(ctx context.Context, storeKey string, input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	endpoint := fmt.Sprintf("in-store/key=%s/me/orders", storeKey)
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreMyOrderGetWithID for type Order
func (client *Client) StoreMyOrderGetWithID(ctx context.Context, storeKey string, id string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/me/orders/%s", storeKey, id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreOrderCreate creates a new instance of type Order
func (client *Client) StoreOrderCreate(ctx context.Context, storeKey string, draft *OrderFromCartDraft, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/orders", storeKey)
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreOrderQuery allows querying for type
// Queries orders in a specific Store. The {storeKey} path parameter maps to a Store's key.
func (client *Client) StoreOrderQuery(ctx context.Context, storeKey string, input *QueryInput) (result *OrderPagedQueryResponse, err error) {
	endpoint := fmt.Sprintf("in-store/key=%s/orders", storeKey)
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreOrderDeleteWithID for type Order
func (client *Client) StoreOrderDeleteWithID(ctx context.Context, storeKey string, id string, version int, dataErasure bool, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/orders/%s", storeKey, id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreOrderDeleteWithOrderNumber for type Order
func (client *Client) StoreOrderDeleteWithOrderNumber(ctx context.Context, storeKey string, orderNumber string, version int, dataErasure bool, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/orders/order-number=%s", storeKey, orderNumber)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreOrderGetWithID Returns an order by its ID from a specific Store. The {storeKey} path parameter maps to a Store's key.
If the order exists in the commercetools project but does not have the store field,
or the store field references a different store, this method returns a ResourceNotFound error.
*/
func (client *Client) StoreOrderGetWithID(ctx context.Context, storeKey string, id string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/orders/%s", storeKey, id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
StoreOrderGetWithOrderNumber Returns an order by its order number from a specific Store.
The {storeKey} path parameter maps to a Store's key.
If the order exists in the commercetools project but does not have the store field,
or the store field references a different store, this method returns a ResourceNotFound error.
In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
it should be provided in URL-encoded format.
*/
func (client *Client) StoreOrderGetWithOrderNumber(ctx context.Context, storeKey string, orderNumber string, opts ...RequestOption) (result *Order, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/orders/order-number=%s", storeKey, orderNumber)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreOrderUpdateWithIDInput is input for function StoreOrderUpdateWithID
type StoreOrderUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []OrderUpdateAction
}

func (input *StoreOrderUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreOrderUpdateWithID for type Order
func (client *Client) StoreOrderUpdateWithID(ctx context.Context, storeKey string, input *StoreOrderUpdateWithIDInput, opts ...RequestOption) (result *Order, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/orders/%s", storeKey, input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreOrderUpdateWithOrderNumberInput is input for function StoreOrderUpdateWithOrderNumber
type StoreOrderUpdateWithOrderNumberInput struct {
	OrderNumber string
	Version     int
	Actions     []OrderUpdateAction
}

func (input *StoreOrderUpdateWithOrderNumberInput) Validate() error {
	if input.OrderNumber == "" {
		return fmt.Errorf("no valid value for OrderNumber given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreOrderUpdateWithOrderNumber for type Order
func (client *Client) StoreOrderUpdateWithOrderNumber(ctx context.Context, storeKey string, input *StoreOrderUpdateWithOrderNumberInput, opts ...RequestOption) (result *Order, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/orders/order-number=%s", storeKey, input.OrderNumber)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListCreate creates a new instance of type ShoppingList
func (client *Client) StoreShoppingListCreate(ctx context.Context, storeKey string, draft *ShoppingListDraft, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists", storeKey)
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListQuery allows querying for type
func (client *Client) StoreShoppingListQuery(ctx context.Context, storeKey string, input *QueryInput) (result *ShoppingListPagedQueryResponse, err error) {
	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists", storeKey)
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListDeleteWithID for type ShoppingList
func (client *Client) StoreShoppingListDeleteWithID(ctx context.Context, storeKey string, id string, version int, dataErasure bool, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists/%s", storeKey, id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListDeleteWithKey for type ShoppingList
func (client *Client) StoreShoppingListDeleteWithKey(ctx context.Context, storeKey string, key string, version int, dataErasure bool, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists/key=%s", storeKey, key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListGetWithID Gets a shopping list by ID.
func (client *Client) StoreShoppingListGetWithID(ctx context.Context, storeKey string, id string, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists/%s", storeKey, id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListGetWithKey Gets a shopping list by Key.
func (client *Client) StoreShoppingListGetWithKey(ctx context.Context, storeKey string, key string, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists/key=%s", storeKey, key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListUpdateWithIDInput is input for function StoreShoppingListUpdateWithID
type StoreShoppingListUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ShoppingListUpdateAction
}

func (input *StoreShoppingListUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreShoppingListUpdateWithID for type ShoppingList
func (client *Client) StoreShoppingListUpdateWithID(ctx context.Context, storeKey string, input *StoreShoppingListUpdateWithIDInput, opts ...RequestOption) (result *ShoppingList, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists/%s", storeKey, input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreShoppingListUpdateWithKeyInput is input for function StoreShoppingListUpdateWithKey
type StoreShoppingListUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ShoppingListUpdateAction
}

func (input *StoreShoppingListUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreShoppingListUpdateWithKey for type ShoppingList
func (client *Client) StoreShoppingListUpdateWithKey(ctx context.Context, storeKey string, input *StoreShoppingListUpdateWithKeyInput, opts ...RequestOption) (result *ShoppingList, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("in-store/key=%s/shopping-lists/key=%s", storeKey, input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
