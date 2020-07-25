// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ShoppingListCreate creates a new instance of type ShoppingList
func (client *Client) ShoppingListCreate(ctx context.Context, draft *ShoppingListDraft, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "shopping-lists"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListQuery allows querying for type ShoppingList
func (client *Client) ShoppingListQuery(ctx context.Context, input *QueryInput) (result *ShoppingListPagedQueryResponse, err error) {
	endpoint := "shopping-lists"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListDeleteWithID for type ShoppingList
func (client *Client) ShoppingListDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("shopping-lists/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
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
	endpoint := fmt.Sprintf("shopping-lists/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingListGetWithID Gets a shopping list by ID.
func (client *Client) ShoppingListGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *ShoppingList, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("shopping-lists/%s", id)
	err = client.get(ctx, endpoint, params, &result)
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
	endpoint := fmt.Sprintf("shopping-lists/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
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

func (input *ShoppingListUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ShoppingListUpdateWithID for type ShoppingList
func (client *Client) ShoppingListUpdateWithID(ctx context.Context, input *ShoppingListUpdateWithIDInput, opts ...RequestOption) (result *ShoppingList, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("shopping-lists/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
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

func (input *ShoppingListUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ShoppingListUpdateWithKey for type ShoppingList
func (client *Client) ShoppingListUpdateWithKey(ctx context.Context, input *ShoppingListUpdateWithKeyInput, opts ...RequestOption) (result *ShoppingList, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("shopping-lists/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
