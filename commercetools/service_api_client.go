package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// APIClientDeleteInput provides the data required to delete a channel.
type APIClientDeleteInput struct {
	ID      string
	Version int
}

// APIClientGetByID will return a channel matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) APIClientGetByID(id string) (result *APIClient, err error) {
	err = client.Get(fmt.Sprintf("api-clients/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientCreate will create a new channel from a draft, and return the newly created
// channel. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) APIClientCreate(draft *APIClientDraft) (result *APIClient, err error) {
	err = client.Create("api-clients", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientDelete will delete a type matching the provided ID. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) APIClientDelete(id string, version int) (result *APIClient, err error) {
	endpoint := fmt.Sprintf("api-clients/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err = client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
