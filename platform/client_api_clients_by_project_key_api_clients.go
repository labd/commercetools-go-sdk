package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyApiClientsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyApiClientsRequestBuilder) WithId(id string) *ByProjectKeyApiClientsByIDRequestBuilder {
	return &ByProjectKeyApiClientsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyApiClientsRequestBuilder) Get() *ByProjectKeyApiClientsRequestMethodGet {
	return &ByProjectKeyApiClientsRequestMethodGet{
		url:    fmt.Sprintf("/%s/api-clients", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyApiClientsRequestBuilder) Post(body ApiClientDraft) *ByProjectKeyApiClientsRequestMethodPost {
	return &ByProjectKeyApiClientsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/api-clients", rb.projectKey),
		client: rb.client,
	}
}
