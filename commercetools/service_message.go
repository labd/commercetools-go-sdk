// Automatically generated, do not edit

package commercetools

import (
	"context"
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
func (client *Client) AbstractMessageGetWithID(ctx context.Context, ID string) (result *Message, err error) {
	err = client.Get(ctx, strings.Replace("messages/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
