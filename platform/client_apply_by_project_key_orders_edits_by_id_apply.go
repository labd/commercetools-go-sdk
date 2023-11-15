package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersEditsByIDApplyRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Applying an OrderEdit produces the [OrderEditApplied](ctp:api:type:OrderEditAppliedMessage) Message.
*
 */
func (rb *ByProjectKeyOrdersEditsByIDApplyRequestBuilder) Post(body OrderEditApply) *ByProjectKeyOrdersEditsByIDApplyRequestMethodPost {
	return &ByProjectKeyOrdersEditsByIDApplyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/%s/apply", rb.projectKey, rb.id),
		client: rb.client,
	}
}
