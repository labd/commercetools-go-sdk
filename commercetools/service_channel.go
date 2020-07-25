// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ChannelCreate creates a new instance of type Channel
func (client *Client) ChannelCreate(ctx context.Context, draft *ChannelDraft, opts ...RequestOption) (result *Channel, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "channels"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelQuery allows querying for type Channel
func (client *Client) ChannelQuery(ctx context.Context, input *QueryInput) (result *ChannelPagedQueryResponse, err error) {
	endpoint := "channels"
	err = client.query(ctx, endpoint, input.toParams(), &result)
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
	err = client.delete(ctx, endpoint, params, &result)
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
	err = client.get(ctx, endpoint, params, &result)
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

func (input *ChannelUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// ChannelUpdateWithID for type Channel
func (client *Client) ChannelUpdateWithID(ctx context.Context, input *ChannelUpdateWithIDInput, opts ...RequestOption) (result *Channel, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("channels/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
