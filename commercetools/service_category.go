// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CategoryURLPath is the commercetools API path.
const CategoryURLPath = "categories"

// CategoryCreate creates a new instance of type Category
func (client *Client) CategoryCreate(ctx context.Context, draft *CategoryDraft, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, CategoryURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryQuery allows querying for type Category
func (client *Client) CategoryQuery(ctx context.Context, input *QueryInput) (result *CategoryPagedQueryResponse, err error) {
	err = client.Query(ctx, CategoryURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryDeleteWithKey for type Category
func (client *Client) CategoryDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("categories/key=%s", key)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryGetWithKey for type Category
func (client *Client) CategoryGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("categories/key=%s", key)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryUpdateWithKeyInput is input for function CategoryUpdateWithKey
type CategoryUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []CategoryUpdateAction
}

// CategoryUpdateWithKey for type Category
func (client *Client) CategoryUpdateWithKey(ctx context.Context, input *CategoryUpdateWithKeyInput, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("categories/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryDeleteWithID for type Category
func (client *Client) CategoryDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("categories/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryGetWithID for type Category
func (client *Client) CategoryGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("categories/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryUpdateWithIDInput is input for function CategoryUpdateWithID
type CategoryUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []CategoryUpdateAction
}

// CategoryUpdateWithID for type Category
func (client *Client) CategoryUpdateWithID(ctx context.Context, input *CategoryUpdateWithIDInput, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("categories/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
