// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ZoneCreate creates a new instance of type Zone
func (client *Client) ZoneCreate(ctx context.Context, draft *ZoneDraft, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "zones"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneQuery allows querying for type Zone
func (client *Client) ZoneQuery(ctx context.Context, input *QueryInput) (result *ZonePagedQueryResponse, err error) {
	endpoint := "zones"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneDeleteWithID for type Zone
func (client *Client) ZoneDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("zones/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneDeleteWithKey for type Zone
func (client *Client) ZoneDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("zones/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneGetWithID for type Zone
func (client *Client) ZoneGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("zones/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneGetWithKey for type Zone
func (client *Client) ZoneGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("zones/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneUpdateWithIDInput is input for function ZoneUpdateWithID
type ZoneUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ZoneUpdateAction
}

func (input *ZoneUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ZoneUpdateWithID for type Zone
func (client *Client) ZoneUpdateWithID(ctx context.Context, input *ZoneUpdateWithIDInput, opts ...RequestOption) (result *Zone, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("zones/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneUpdateWithKeyInput is input for function ZoneUpdateWithKey
type ZoneUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ZoneUpdateAction
}

func (input *ZoneUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ZoneUpdateWithKey for type Zone
func (client *Client) ZoneUpdateWithKey(ctx context.Context, input *ZoneUpdateWithKeyInput, opts ...RequestOption) (result *Zone, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("zones/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
