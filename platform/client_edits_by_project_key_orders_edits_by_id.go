package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersEditsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Apply() *ByProjectKeyOrdersEditsByIDApplyRequestBuilder {
	return &ByProjectKeyOrdersEditsByIDApplyRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Get() *ByProjectKeyOrdersEditsByIDRequestMethodGet {
	return &ByProjectKeyOrdersEditsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Post(body OrderEditUpdate) *ByProjectKeyOrdersEditsByIDRequestMethodPost {
	return &ByProjectKeyOrdersEditsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Delete() *ByProjectKeyOrdersEditsByIDRequestMethodDelete {
	return &ByProjectKeyOrdersEditsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
