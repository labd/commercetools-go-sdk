// Automatically generated, do not edit

package commercetools

import (
	"net/url"
	"strings"
)

// APIClientURLPath is the commercetools API path.
const APIClientURLPath = "api-clients"

func (client *Client) APIClientCreate(draft *APIClientDraft) (result *APIClient, err error) {
	err = client.Create(APIClientURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) APIClientQuery(input *QueryInput) (result *APIClientPagedQueryResponse, err error) {
	err = client.Query(APIClientURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) APIClientDeleteWithID(ID string) (result *APIClient, err error) {
	params := url.Values{}

	err = client.Delete(strings.Replace("api-clients/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *Client) APIClientGetWithID(ID string) (result *APIClient, err error) {
	err = client.Get(strings.Replace("api-clients/{ID}", "{ID}", ID, 1), nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
