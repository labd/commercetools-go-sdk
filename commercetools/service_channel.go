// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ChannelURLPath is the commercetools API path.
const ChannelURLPath = "channels"

// ChannelCreate creates a new instance of type Channel
func (client *Client) ChannelCreate(ctx context.Context, draft *ChannelDraft, opts ...RequestOption) (result *Channel, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ChannelURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelQuery allows querying for type Channel
func (client *Client) ChannelQuery(ctx context.Context, input *QueryInput) (result *ChannelPagedQueryResponse, err error) {
	err = client.Query(ctx, ChannelURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelDeleteWithID for type Channel
func (client *Client) ChannelDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Channel, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("channels/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelGetWithID for type Channel
func (client *Client) ChannelGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Channel, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("channels/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelUpdateWithIDInput is input for function ChannelUpdateWithID
type ChannelUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ChannelUpdateAction
}

// ChannelUpdateWithID for type Channel
func (client *Client) ChannelUpdateWithID(ctx context.Context, input *ChannelUpdateWithIDInput, opts ...RequestOption) (result *Channel, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("channels/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
