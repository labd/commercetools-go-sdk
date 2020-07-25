// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
)

// APIClientCreate creates a new instance of type APIClient
func (client *Client) APIClientCreate(ctx context.Context, draft *APIClientDraft, opts ...RequestOption) (result *APIClient, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "api-clients"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientQuery allows querying for type ApiClient
func (client *Client) APIClientQuery(ctx context.Context, input *QueryInput) (result *APIClientPagedQueryResponse, err error) {
	endpoint := "api-clients"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientDeleteWithID Delete ApiClient by ID
func (client *Client) APIClientDeleteWithID(ctx context.Context, id string, opts ...RequestOption) (result *APIClient, err error) {
	params := url.Values{}

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("api-clients/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// APIClientGetWithID Get ApiClient by ID
func (client *Client) APIClientGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *APIClient, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("api-clients/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
