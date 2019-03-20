package commercetools

import (
	"fmt"
)

// APIClientDeleteInput provides the data required to delete a channel.
type APIClientDeleteInput struct {
	ID      string
	Version int
}

// APIClientURLPath is the commercetools API client path.
const APIClientURLPath = "api-clients"

// APIClientGetByID will return a channel matching the provided ID. OAuth2 Scopes:
// view_products:{projectKey}
func (client *Client) APIClientGetByID(id string) (result *APIClient, err error) {
	err = client.Get(fmt.Sprintf("%s/%s", APIClientURLPath, id), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientCreate will create a new channel from a draft, and return the newly created
// channel. OAuth2 Scopes: manage_products:{projectKey}
func (client *Client) APIClientCreate(draft *APIClientDraft) (result *APIClient, err error) {
	err = client.Create(APIClientURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientDelete will delete a type matching the provided ID. These requests
// delete a type only if it’s not referenced by other entities. OAuth2 Scopes:
// manage_types:{projectKey}
func (client *Client) APIClientDelete(id string) (result *APIClient, err error) {
	endpoint := fmt.Sprintf("%s/%s", APIClientURLPath, id)
	err = client.Delete(endpoint, nil, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientQuery will query API clients.
// OAuth2 Scopes: view_api_clients:{projectKey}
func (client *Client) APIClientQuery(input *QueryInput) (result *APIClientPagedQueryResponse, err error) {
	err = client.Query(APIClientURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
