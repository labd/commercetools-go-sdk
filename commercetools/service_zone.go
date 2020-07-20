// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ZoneURLPath is the commercetools API path.
const ZoneURLPath = "zones"

// ZoneCreate creates a new instance of type Zone
func (client *Client) ZoneCreate(ctx context.Context, draft *ZoneDraft, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Create(ctx, ZoneURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneQuery allows querying for type Zone
func (client *Client) ZoneQuery(ctx context.Context, input *QueryInput) (result *ZonePagedQueryResponse, err error) {
	err = client.Query(ctx, ZoneURLPath, input.toParams(), &result)
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
	err = client.Delete(ctx, strings.Replace("zones/key={key}", "{key}", key, 1), params, &result)
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
	err = client.Get(ctx, strings.Replace("zones/key={key}", "{key}", key, 1), params, &result)
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

// ZoneUpdateWithKey for type Zone
func (client *Client) ZoneUpdateWithKey(ctx context.Context, input *ZoneUpdateWithKeyInput, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("zones/key={key}", "{key}", input.Key, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneDeleteWithID for type Zone
func (client *Client) ZoneDeleteWithID(ctx context.Context, ID string, version int, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("zones/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ZoneGetWithID for type Zone
func (client *Client) ZoneGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("zones/{ID}", "{ID}", ID, 1), params, &result)
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

// ZoneUpdateWithID for type Zone
func (client *Client) ZoneUpdateWithID(ctx context.Context, input *ZoneUpdateWithIDInput, opts ...RequestOption) (result *Zone, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Update(ctx, strings.Replace("zones/{ID}", "{ID}", input.ID, 1), params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
