// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ProductTypeURLPath is the commercetools API path.
const ProductTypeURLPath = "product-types"

// ProductTypeCreate creates a new instance of type ProductType
func (client *Client) ProductTypeCreate(ctx context.Context, draft *ProductTypeDraft, opts ...RequestOption) (result *ProductType, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ProductTypeURLPath, params, draft, &result)
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
func (client *Client) ProductTypeDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *ProductType, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-types/key=%s", key)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeGetWithKey for type ProductType
func (client *Client) ProductTypeGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *ProductType, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-types/key=%s", key)
	err = client.Get(ctx, endpoint, params, &result)
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

func (input *ProductTypeUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ProductTypeUpdateWithKey for type ProductType
func (client *Client) ProductTypeUpdateWithKey(ctx context.Context, input *ProductTypeUpdateWithKeyInput, opts ...RequestOption) (result *ProductType, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("product-types/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeDeleteWithID for type ProductType
func (client *Client) ProductTypeDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *ProductType, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-types/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductTypeGetWithID for type ProductType
func (client *Client) ProductTypeGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *ProductType, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("product-types/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
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

func (input *ProductTypeUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ProductTypeUpdateWithID for type ProductType
func (client *Client) ProductTypeUpdateWithID(ctx context.Context, input *ProductTypeUpdateWithIDInput, opts ...RequestOption) (result *ProductType, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("product-types/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
