package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersRequestBuilder) WithPasswordToken(passwordToken string) *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder {
	return &ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder{
		passwordToken: passwordToken,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) WithEmailToken(emailToken string) *ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder {
	return &ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder{
		emailToken: emailToken,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	To verify a customer's email, an email token can be created. This should be embedded in a link and sent to the
*	customer via email. When the customer clicks on the link, the "verify customer's email" endpoint should be called,
*	which sets customer's isVerifiedEmail field to true.
*
 */
func (rb *ByProjectKeyCustomersRequestBuilder) EmailToken() *ByProjectKeyCustomersEmailTokenRequestBuilder {
	return &ByProjectKeyCustomersEmailTokenRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) EmailConfirm() *ByProjectKeyCustomersEmailConfirmRequestBuilder {
	return &ByProjectKeyCustomersEmailConfirmRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) Password() *ByProjectKeyCustomersPasswordRequestBuilder {
	return &ByProjectKeyCustomersPasswordRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) PasswordReset() *ByProjectKeyCustomersPasswordResetRequestBuilder {
	return &ByProjectKeyCustomersPasswordResetRequestBuilder{
		projectKey: rb.projectKey,
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
func (rb *ByProjectKeyCustomersRequestBuilder) PasswordToken() *ByProjectKeyCustomersPasswordTokenRequestBuilder {
	return &ByProjectKeyCustomersPasswordTokenRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) WithKey(key string) *ByProjectKeyCustomersKeyByKeyRequestBuilder {
	return &ByProjectKeyCustomersKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) WithId(id string) *ByProjectKeyCustomersByIDRequestBuilder {
	return &ByProjectKeyCustomersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) Get() *ByProjectKeyCustomersRequestMethodGet {
	return &ByProjectKeyCustomersRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a customer. If an anonymous cart is passed in,
*	then the cart is assigned to the created customer and the version number of the Cart will increase.
*	If the ID of an anonymous session is given, all carts and orders will be assigned to the created customer.
*
 */
func (rb *ByProjectKeyCustomersRequestBuilder) Post(body CustomerDraft) *ByProjectKeyCustomersRequestMethodPost {
	return &ByProjectKeyCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers", rb.projectKey),
		client: rb.client,
	}
}
