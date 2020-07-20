// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ShoppingListURLPath is the commercetools API path.
const ShoppingListURLPath = "shopping-lists"

// ShoppingListCreate creates a new instance of type ShoppingList
func (client *Client) ShoppingListCreate(ctx context.Context, draft *ShoppingListDraft, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ShoppingListURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListQuery allows querying for type ShoppingList
func (client *Client) ShoppingListQuery(ctx context.Context, input *QueryInput) (result *ShoppingListPagedQueryResponse, err error) {
	err = client.Query(ctx, ShoppingListURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListDeleteWithKey for type ShoppingList
func (client *Client) ShoppingListDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("shopping-lists/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListGetWithKey Gets a shopping list by Key.
func (client *Client) ShoppingListGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("shopping-lists/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListUpdateWithKeyInput is input for function ShoppingListUpdateWithKey
type ShoppingListUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ShoppingListUpdateAction
}

// ShoppingListUpdateWithKey Update a shopping list found by its Key.
func (client *Client) ShoppingListUpdateWithKey(ctx context.Context, input *ShoppingListUpdateWithKeyInput, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("shopping-lists/key={key}", "{key}", input.Key, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListDeleteWithID for type ShoppingList
func (client *Client) ShoppingListDeleteWithID(ctx context.Context, ID string, version int, dataErasure bool, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("shopping-lists/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListGetWithID Gets a shopping list by ID.
func (client *Client) ShoppingListGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("shopping-lists/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListUpdateWithIDInput is input for function ShoppingListUpdateWithID
type ShoppingListUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ShoppingListUpdateAction
}

// ShoppingListUpdateWithID for type ShoppingList
func (client *Client) ShoppingListUpdateWithID(ctx context.Context, input *ShoppingListUpdateWithIDInput, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("shopping-lists/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
