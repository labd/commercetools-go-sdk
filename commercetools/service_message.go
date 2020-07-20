// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strings"
)

// AbstractMessageURLPath is the commercetools API path.
const AbstractMessageURLPath = "messages"

// AbstractMessageQuery allows querying for type Message
func (client *Client) AbstractMessageQuery(ctx context.Context, input *QueryInput) (result *MessagePagedQueryResponse, err error) {
	err = client.Query(ctx, AbstractMessageURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AbstractMessageGetWithID for type Message
func (client *Client) AbstractMessageGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *Message, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("messages/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
