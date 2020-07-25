// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CartDiscountCreate creates a new instance of type CartDiscount
func (client *Client) CartDiscountCreate(ctx context.Context, draft *CartDiscountDraft, opts ...RequestOption) (result *CartDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "cart-discounts"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountQuery allows querying for type CartDiscount
func (client *Client) CartDiscountQuery(ctx context.Context, input *QueryInput) (result *CartDiscountPagedQueryResponse, err error) {
	endpoint := "cart-discounts"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountDeleteWithID for type CartDiscount
func (client *Client) CartDiscountDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *CartDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("cart-discounts/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountDeleteWithKey for type CartDiscount
func (client *Client) CartDiscountDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *CartDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("cart-discounts/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountGetWithID for type CartDiscount
func (client *Client) CartDiscountGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *CartDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("cart-discounts/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountGetWithKey for type CartDiscount
func (client *Client) CartDiscountGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *CartDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("cart-discounts/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountUpdateWithIDInput is input for function CartDiscountUpdateWithID
type CartDiscountUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CartDiscountUpdateAction
}

func (input *CartDiscountUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CartDiscountUpdateWithID for type CartDiscount
func (client *Client) CartDiscountUpdateWithID(ctx context.Context, input *CartDiscountUpdateWithIDInput, opts ...RequestOption) (result *CartDiscount, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("cart-discounts/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CartDiscountUpdateWithKeyInput is input for function CartDiscountUpdateWithKey
type CartDiscountUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CartDiscountUpdateAction
}

func (input *CartDiscountUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CartDiscountUpdateWithKey for type CartDiscount
func (client *Client) CartDiscountUpdateWithKey(ctx context.Context, input *CartDiscountUpdateWithKeyInput, opts ...RequestOption) (result *CartDiscount, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("cart-discounts/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
