package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	If omitted in the request body, the [Customer](ctp:api:type:Customer) `stores` field is set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder) Post(body MyCustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/signup", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
