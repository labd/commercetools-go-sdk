// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ProductDiscountCreate creates a new instance of type ProductDiscount
func (client *Client) ProductDiscountCreate(ctx context.Context, draft *ProductDiscountDraft, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "product-discounts"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountQuery allows querying for type ProductDiscount
func (client *Client) ProductDiscountQuery(ctx context.Context, input *QueryInput) (result *ProductDiscountPagedQueryResponse, err error) {
	endpoint := "product-discounts"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountDeleteWithID for type ProductDiscount
func (client *Client) ProductDiscountDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-discounts/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountDeleteWithKey for type ProductDiscount
func (client *Client) ProductDiscountDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-discounts/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountGetWithID for type ProductDiscount
func (client *Client) ProductDiscountGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-discounts/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountGetWithKey for type ProductDiscount
func (client *Client) ProductDiscountGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-discounts/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountUpdateWithIDInput is input for function ProductDiscountUpdateWithID
type ProductDiscountUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ProductDiscountUpdateAction
}

func (input *ProductDiscountUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ProductDiscountUpdateWithID for type ProductDiscount
func (client *Client) ProductDiscountUpdateWithID(ctx context.Context, input *ProductDiscountUpdateWithIDInput, opts ...RequestOption) (result *ProductDiscount, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("product-discounts/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountUpdateWithKeyInput is input for function ProductDiscountUpdateWithKey
type ProductDiscountUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductDiscountUpdateAction
}

func (input *ProductDiscountUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ProductDiscountUpdateWithKey for type ProductDiscount
func (client *Client) ProductDiscountUpdateWithKey(ctx context.Context, input *ProductDiscountUpdateWithKeyInput, opts ...RequestOption) (result *ProductDiscount, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("product-discounts/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountMatching for type ProductDiscountMatchQuery
func (client *Client) ProductDiscountMatching(ctx context.Context, value *ProductDiscountMatchQuery, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "product-discounts/matching"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
