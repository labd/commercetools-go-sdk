// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ProductTypeURLPath is the commercetools API path.
const ProductTypeURLPath = "product-types"

// ProductTypeCreate creates a new instance of type ProductType
func (client *Client) ProductTypeCreate(ctx context.Context, draft *ProductTypeDraft) (result *ProductType, err error) {
	err = client.Create(ctx, ProductTypeURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeQuery allows querying for type ProductType
func (client *Client) ProductTypeQuery(ctx context.Context, input *QueryInput) (result *ProductTypePagedQueryResponse, err error) {
	err = client.Query(ctx, ProductTypeURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeDeleteWithKey for type ProductType
func (client *Client) ProductTypeDeleteWithKey(ctx context.Context, key string, version int) (result *ProductType, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("product-types/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeGetWithKey for type ProductType
func (client *Client) ProductTypeGetWithKey(ctx context.Context, key string) (result *ProductType, err error) {
	err = client.Get(ctx, strings.Replace("product-types/key={key}", "{key}", key, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeUpdateWithKeyInput is input for function ProductTypeUpdateWithKey
type ProductTypeUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ProductTypeUpdateAction
}

// ProductTypeUpdateWithKey for type ProductType
func (client *Client) ProductTypeUpdateWithKey(ctx context.Context, input *ProductTypeUpdateWithKeyInput) (result *ProductType, err error) {
	err = client.Update(ctx, strings.Replace("product-types/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeDeleteWithID for type ProductType
func (client *Client) ProductTypeDeleteWithID(ctx context.Context, ID string, version int) (result *ProductType, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("product-types/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeGetWithID for type ProductType
func (client *Client) ProductTypeGetWithID(ctx context.Context, ID string) (result *ProductType, err error) {
	err = client.Get(ctx, strings.Replace("product-types/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeUpdateWithIDInput is input for function ProductTypeUpdateWithID
type ProductTypeUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ProductTypeUpdateAction
}

// ProductTypeUpdateWithID for type ProductType
func (client *Client) ProductTypeUpdateWithID(ctx context.Context, input *ProductTypeUpdateWithIDInput) (result *ProductType, err error) {
	err = client.Update(ctx, strings.Replace("product-types/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
