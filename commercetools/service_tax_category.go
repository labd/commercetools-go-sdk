// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// TaxCategoryURLPath is the commercetools API path.
const TaxCategoryURLPath = "tax-categories"

// TaxCategoryCreate creates a new instance of type TaxCategory
func (client *Client) TaxCategoryCreate(ctx context.Context, draft *TaxCategoryDraft, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, TaxCategoryURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryQuery allows querying for type TaxCategory
func (client *Client) TaxCategoryQuery(ctx context.Context, input *QueryInput) (result *TaxCategoryPagedQueryResponse, err error) {
	err = client.Query(ctx, TaxCategoryURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryDeleteWithKey for type TaxCategory
func (client *Client) TaxCategoryDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("tax-categories/key=%s", key)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryGetWithKey for type TaxCategory
func (client *Client) TaxCategoryGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("tax-categories/key=%s", key)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryUpdateWithKeyInput is input for function TaxCategoryUpdateWithKey
type TaxCategoryUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []TaxCategoryUpdateAction
}

// TaxCategoryUpdateWithKey for type TaxCategory
func (client *Client) TaxCategoryUpdateWithKey(ctx context.Context, input *TaxCategoryUpdateWithKeyInput, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("tax-categories/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryDeleteWithID for type TaxCategory
func (client *Client) TaxCategoryDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("tax-categories/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryGetWithID for type TaxCategory
func (client *Client) TaxCategoryGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("tax-categories/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryUpdateWithIDInput is input for function TaxCategoryUpdateWithID
type TaxCategoryUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []TaxCategoryUpdateAction
}

// TaxCategoryUpdateWithID for type TaxCategory
func (client *Client) TaxCategoryUpdateWithID(ctx context.Context, input *TaxCategoryUpdateWithIDInput, opts ...RequestOption) (result *TaxCategory, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("tax-categories/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
