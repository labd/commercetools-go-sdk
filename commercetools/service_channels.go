package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// ChannelDeleteInput provides the data required to delete a channel.
type ChannelDeleteInput struct {
	ID      string
	Version int
}

// ChannelUpdateInput provides the data required to update a channel.
type ChannelUpdateInput struct {
	ID string

	// The expected version of the channel on which the changes should be
	// applied. If the expected version does not match the actual version, a 409
	// Conflict will be returned.
	Version int

	// The list of update actions to be performed on the channel.
	Actions []ChannelUpdateAction
}

// ChannelGetByID will return a channel matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) ChannelGetByID(id string) (result *Channel, err error) {
	err = client.Get(fmt.Sprintf("channels/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelCreate will create a new channel from a draft, and return the newly created
// channel. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ChannelCreate(draft *ChannelDraft) (result *Channel, err error) {
	err = client.Create("channels", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelUpdate will update a channel matching the provided ID with the defined
// ChannelUpdateActions. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) ChannelUpdate(input *ChannelUpdateInput) (result *Channel, err error) {
	if input.ID == "" {
		return nil, fmt.Errorf("no valid type id passed")
	}

	endpoint := fmt.Sprintf("channels/%s", input.ID)
	err = client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ChannelDelete will delete a type matching the provided ID. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) ChannelDelete(id string, version int) (result *Channel, err error) {
	endpoint := fmt.Sprintf("channels/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
