package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	The customer verifies the email using the token value.
*	Verifying the email of the Customer produces the [CustomerEmailVerified](ctp:api:type:CustomerEmailVerifiedMessage) Message.
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email/confirm", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
