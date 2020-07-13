// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// ChannelURLPath is the commercetools API path.
const ChannelURLPath = "channels"

// ChannelCreate creates a new instance of type Channel
func (client *Client) ChannelCreate(ctx context.Context, draft *ChannelDraft) (result *Channel, err error) {
	err = client.Create(ctx, ChannelURLPath, nil, draft, &result)
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
func (client *Client) ChannelDeleteWithID(ctx context.Context, ID string, version int) (result *Channel, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("channels/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelGetWithID for type Channel
func (client *Client) ChannelGetWithID(ctx context.Context, ID string) (result *Channel, err error) {
	err = client.Get(ctx, strings.Replace("channels/{ID}", "{ID}", ID, 1), nil, &result)
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
func (client *Client) ChannelUpdateWithID(ctx context.Context, input *ChannelUpdateWithIDInput) (result *Channel, err error) {
	err = client.Update(ctx, strings.Replace("channels/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
