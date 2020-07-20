// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ProductDiscountURLPath is the commercetools API path.
const ProductDiscountURLPath = "product-discounts"

// ProductDiscountCreate creates a new instance of type ProductDiscount
func (client *Client) ProductDiscountCreate(ctx context.Context, draft *ProductDiscountDraft, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ProductDiscountURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDiscountQuery allows querying for type ProductDiscount
func (client *Client) ProductDiscountQuery(ctx context.Context, input *QueryInput) (result *ProductDiscountPagedQueryResponse, err error) {
	err = client.Query(ctx, ProductDiscountURLPath, input.toParams(), &result)
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
	err = client.Delete(ctx, endpoint, params, &result)
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
	err = client.Get(ctx, endpoint, params, &result)
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

// ProductDiscountUpdateWithKey for type ProductDiscount
func (client *Client) ProductDiscountUpdateWithKey(ctx context.Context, input *ProductDiscountUpdateWithKeyInput, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-discounts/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
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
	err = client.Delete(ctx, endpoint, params, &result)
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
	err = client.Get(ctx, endpoint, params, &result)
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

// ProductDiscountUpdateWithID for type ProductDiscount
func (client *Client) ProductDiscountUpdateWithID(ctx context.Context, input *ProductDiscountUpdateWithIDInput, opts ...RequestOption) (result *ProductDiscount, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-discounts/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
