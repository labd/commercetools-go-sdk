// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CategoryCreate creates a new instance of type Category
func (client *Client) CategoryCreate(ctx context.Context, draft *CategoryDraft, opts ...RequestOption) (result *Category, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "categories"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CategoryQuery allows querying for type Category
func (client *Client) CategoryQuery(ctx context.Context, input *QueryInput) (result *CategoryPagedQueryResponse, err error) {
	endpoint := "categories"
	err = client.query(ctx, endpoint, input.toParams(), &result)
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
	err = client.delete(ctx, endpoint, params, &result)
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
	err = client.delete(ctx, endpoint, params, &result)
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
	err = client.get(ctx, endpoint, params, &result)
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
	err = client.get(ctx, endpoint, params, &result)
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

func (input *CategoryUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CategoryUpdateWithID for type Category
func (client *Client) CategoryUpdateWithID(ctx context.Context, input *CategoryUpdateWithIDInput, opts ...RequestOption) (result *Category, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("categories/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
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

func (input *CategoryUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// CategoryUpdateWithKey for type Category
func (client *Client) CategoryUpdateWithKey(ctx context.Context, input *CategoryUpdateWithKeyInput, opts ...RequestOption) (result *Category, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("categories/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
