package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithPasswordToken(passwordToken string) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder{
		passwordToken: passwordToken,
		projectKey:    rb.projectKey,
		storeKey:      rb.storeKey,
		client:        rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithEmailToken(emailToken string) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder{
		emailToken: emailToken,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	To verify a customer's email, an email token can be created. This should be embedded in a link and sent to the
*	customer via email. When the customer clicks on the link,
*	the "verify customer's email" endpoint should be called,
*	which sets customer's isVerifiedEmail field to true.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) EmailToken() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) EmailConfirm() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Password() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) PasswordReset() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	The following workflow can be used to reset the customer's password:
*
*	* Create a password reset token and send it embedded in a link to the customer.
*	* When the customer clicks on the link, the customer is retrieved with the token.
*	* The customer enters a new password and the "reset customer's password" endpoint is called.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) PasswordToken() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a customer in a specific Store.
*	When using this endpoint, if omitted,
*	the customer's stores field is set to the store specified in the path parameter.
*	If an anonymous cart is passed in as when using this method,
*	then the cart is assigned to the created customer and the version number of the Cart increases.
*	If the ID of an anonymous session is given, all carts and orders will be assigned to the created customer and
*	the store specified. If you pass in a cart with a store field specified,
*	the store field must reference the same store specified in the {storeKey} path parameter.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Post(body CustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
