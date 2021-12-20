// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyOrdersEditsByIDApplyRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsByIDApplyRequestBuilder) Post(body OrderEditApply) *ByProjectKeyOrdersEditsByIDApplyRequestMethodPost {
	return &ByProjectKeyOrdersEditsByIDApplyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/%s/apply", rb.projectKey, rb.id),
		client: rb.client,
	}
}
