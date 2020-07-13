// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// TaxCategoryURLPath is the commercetools API path.
const TaxCategoryURLPath = "tax-categories"

// TaxCategoryCreate creates a new instance of type TaxCategory
func (client *Client) TaxCategoryCreate(ctx context.Context, draft *TaxCategoryDraft) (result *TaxCategory, err error) {
	err = client.Create(ctx, TaxCategoryURLPath, nil, draft, &result)
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
func (client *Client) TaxCategoryDeleteWithKey(ctx context.Context, key string, version int) (result *TaxCategory, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("tax-categories/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryGetWithKey for type TaxCategory
func (client *Client) TaxCategoryGetWithKey(ctx context.Context, key string) (result *TaxCategory, err error) {
	err = client.Get(ctx, strings.Replace("tax-categories/key={key}", "{key}", key, 1), nil, &result)
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
func (client *Client) TaxCategoryUpdateWithKey(ctx context.Context, input *TaxCategoryUpdateWithKeyInput) (result *TaxCategory, err error) {
	err = client.Update(ctx, strings.Replace("tax-categories/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryDeleteWithID for type TaxCategory
func (client *Client) TaxCategoryDeleteWithID(ctx context.Context, ID string, version int) (result *TaxCategory, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("tax-categories/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxCategoryGetWithID for type TaxCategory
func (client *Client) TaxCategoryGetWithID(ctx context.Context, ID string) (result *TaxCategory, err error) {
	err = client.Get(ctx, strings.Replace("tax-categories/{ID}", "{ID}", ID, 1), nil, &result)
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
func (client *Client) TaxCategoryUpdateWithID(ctx context.Context, input *TaxCategoryUpdateWithIDInput) (result *TaxCategory, err error) {
	err = client.Update(ctx, strings.Replace("tax-categories/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
