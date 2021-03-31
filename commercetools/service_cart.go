// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CartCreate creates a new instance of type Cart
func (client *Client) CartCreate(ctx context.Context, draft *CartDraft, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "carts"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartQuery allows querying for type Cart
func (client *Client) CartQuery(ctx context.Context, input *QueryInput) (result *CartPagedQueryResponse, err error) {
	endpoint := "carts"
	err = client.query(ctx, endpoint, input.toParams(), &result)
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
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDeleteWithKey for type Cart
func (client *Client) CartDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CartGetWithCustomerID Retrieves the active cart of the customer that has been modified most recently.
It does not consider carts with CartOrigin Merchant. If no active cart exists, a 404 Not Found error is returned.

The cart may not contain up-to-date prices, discounts etc. If you want to ensure they're up-to-date,
send an Update request with the Recalculate update action instead.
*/
func (client *Client) CartGetWithCustomerID(ctx context.Context, customerId string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/customer-id=%s", customerId)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CartGetWithID The cart may not contain up-to-date prices, discounts etc.
If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*/
func (client *Client) CartGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
CartGetWithKey The cart may not contain up-to-date prices, discounts etc.
If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*/
func (client *Client) CartGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("carts/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
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

func (input *CartUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CartUpdateWithID for type Cart
func (client *Client) CartUpdateWithID(ctx context.Context, input *CartUpdateWithIDInput, opts ...RequestOption) (result *Cart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("carts/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartUpdateWithKeyInput is input for function CartUpdateWithKey
type CartUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CartUpdateAction
}

func (input *CartUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CartUpdateWithKey for type Cart
func (client *Client) CartUpdateWithKey(ctx context.Context, input *CartUpdateWithKeyInput, opts ...RequestOption) (result *Cart, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("carts/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartReplicate for type ReplicaCartDraft
func (client *Client) CartReplicate(ctx context.Context, value *ReplicaCartDraft, opts ...RequestOption) (result *Cart, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "carts/replicate"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
