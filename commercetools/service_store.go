// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// StoreCreate creates a new instance of type Store
func (client *Client) StoreCreate(ctx context.Context, draft *StoreDraft, opts ...RequestOption) (result *Store, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "stores"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreQuery allows querying for type Store
func (client *Client) StoreQuery(ctx context.Context, input *QueryInput) (result *StorePagedQueryResponse, err error) {
	endpoint := "stores"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDeleteWithID for type Store
func (client *Client) StoreDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Store, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("stores/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDeleteWithKey for type Store
func (client *Client) StoreDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Store, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("stores/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreGetWithID for type Store
func (client *Client) StoreGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Store, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("stores/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreGetWithKey for type Store
func (client *Client) StoreGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Store, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("stores/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUpdateWithIDInput is input for function StoreUpdateWithID
type StoreUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []StoreUpdateAction
}

func (input *StoreUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreUpdateWithID for type Store
func (client *Client) StoreUpdateWithID(ctx context.Context, input *StoreUpdateWithIDInput, opts ...RequestOption) (result *Store, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("stores/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreUpdateWithKeyInput is input for function StoreUpdateWithKey
type StoreUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []StoreUpdateAction
}

func (input *StoreUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// StoreUpdateWithKey for type Store
func (client *Client) StoreUpdateWithKey(ctx context.Context, input *StoreUpdateWithKeyInput, opts ...RequestOption) (result *Store, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("stores/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
