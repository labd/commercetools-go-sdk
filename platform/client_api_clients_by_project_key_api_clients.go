// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyApiClientsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyApiClientsRequestBuilder) WithId(id string) *ByProjectKeyApiClientsByIdRequestBuilder {
	return &ByProjectKeyApiClientsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query api-clients
 */
func (rb *ByProjectKeyApiClientsRequestBuilder) Get() *ByProjectKeyApiClientsRequestMethodGet {
	return &ByProjectKeyApiClientsRequestMethodGet{
		url:    fmt.Sprintf("/%s/api-clients", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create ApiClient
 */
func (rb *ByProjectKeyApiClientsRequestBuilder) Post(body ApiClientDraft) *ByProjectKeyApiClientsRequestMethodPost {
	return &ByProjectKeyApiClientsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/api-clients", rb.projectKey),
		client: rb.client,
	}
}
