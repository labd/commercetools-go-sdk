// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ProductURLPath is the commercetools API path.
const ProductURLPath = "products"

// ProductCreate creates a new instance of type Product
func (client *Client) ProductCreate(ctx context.Context, draft *ProductDraft, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ProductURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductQuery allows querying for type Product
func (client *Client) ProductQuery(ctx context.Context, input *QueryInput) (result *ProductPagedQueryResponse, err error) {
	err = client.Query(ctx, ProductURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDeleteWithKey for type Product
func (client *Client) ProductDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("products/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductGetWithKey Gets the full representation of a product by Key.
func (client *Client) ProductGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("products/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdateWithKeyInput is input for function ProductUpdateWithKey
type ProductUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductUpdateAction
}

// ProductUpdateWithKey for type Product
func (client *Client) ProductUpdateWithKey(ctx context.Context, input *ProductUpdateWithKeyInput, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("products/key={key}", "{key}", input.Key, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductDeleteWithID for type Product
func (client *Client) ProductDeleteWithID(ctx context.Context, ID string, version int, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("products/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductGetWithID Gets the full representation of a product by ID.
func (client *Client) ProductGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("products/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductUpdateWithIDInput is input for function ProductUpdateWithID
type ProductUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ProductUpdateAction
}

// ProductUpdateWithID for type Product
func (client *Client) ProductUpdateWithID(ctx context.Context, input *ProductUpdateWithIDInput, opts ...RequestOption) (result *Product, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("products/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
