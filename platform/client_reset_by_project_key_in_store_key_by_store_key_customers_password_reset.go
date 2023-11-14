package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Resetting the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder) Post(body CustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
