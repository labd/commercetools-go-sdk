package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder) Reset() *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Changing the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=false`.  Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Customer exists with the `id` specified in the [customer:{id}](/scopes#customer_idid) scope.
*	- If the Customer exists but is associated with a different Store than what is specified in the `manage_my_profile:{projectKey}:{storeKey}` scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder) Post(body MyCustomerChangePassword) *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/password", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
