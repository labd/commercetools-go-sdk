// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strings"
)

// APIClientURLPath is the commercetools API path.
const APIClientURLPath = "api-clients"

// APIClientCreate creates a new instance of type APIClient
func (client *Client) APIClientCreate(ctx context.Context, draft *APIClientDraft, opts ...RequestOption) (result *APIClient, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Create(ctx, APIClientURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientQuery allows querying for type APIClient
func (client *Client) APIClientQuery(ctx context.Context, input *QueryInput) (result *APIClientPagedQueryResponse, err error) {
	err = client.Query(ctx, APIClientURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientDeleteWithID Delete ApiClient by ID
func (client *Client) APIClientDeleteWithID(ctx context.Context, ID string, opts ...RequestOption) (result *APIClient, err error) {
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}
	err = client.Delete(ctx, strings.Replace("api-clients/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientGetWithID Get ApiClient by ID
func (client *Client) APIClientGetWithID(ctx context.Context, ID string, opts ...RequestOption) (result *APIClient, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	err = client.Get(ctx, strings.Replace("api-clients/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
