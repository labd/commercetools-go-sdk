// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
)

// MessageQuery allows querying for type Message
func (client *Client) MessageQuery(ctx context.Context, input *QueryInput) (result *MessagePagedQueryResponse, err error) {
	endpoint := "messages"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MessageGetWithID for type Message
func (client *Client) MessageGetWithID(ctx context.Context, id string, opts ...RequestOption) (result Message, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("messages/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return mapDiscriminatorMessage(result)
}
