// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strconv"
	"strings"
)

// ChannelURLPath is the commercetools API path.
const ChannelURLPath = "channels"

func (client *Client) ChannelCreate(draft *ChannelDraft) (result *Channel, err error) {
	err = client.Create(ChannelURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ChannelQuery(input *QueryInput) (result *ChannelPagedQueryResponse, err error) {
	err = client.Query(ChannelURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ChannelDeleteWithId(ID string, version int) (result *Channel, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(strings.Replace("channels/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) ChannelGetWithId(ID string) (result *Channel, err error) {
	err = client.Get(strings.Replace("channels/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ChannelUpdateWithIdInput struct {
	ID      string
	Version int
	Actions []ChannelUpdateAction
}

func (client *Client) ChannelUpdateWithId(input *ChannelUpdateWithIdInput) (result *Channel, err error) {
	err = client.Update(strings.Replace("channels/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
