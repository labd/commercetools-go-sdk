// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyOrdersEditsByIdApplyRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsByIdApplyRequestBuilder) Post(body OrderEditApply) *ByProjectKeyOrdersEditsByIdApplyRequestMethodPost {
	return &ByProjectKeyOrdersEditsByIdApplyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/%s/apply", rb.projectKey, rb.id),
		client: rb.client,
	}
}
