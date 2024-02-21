package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyLoginRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Authenticates a global Customer not associated with a Store.
*	For more information, see [Global versus Store-specific Customers](/../api/customers-overview#global-versus-store-specific-customers).
*	If the Customer is registered in a Store, use the [Authenticate (sign in) Customer in Store](ctp:api:endpoint:/{projectKey}/in-store/key={storeKey}/login:POST) method.
*
*	If an account with the given credentials is not found, an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error is returned.
*
 */
func (rb *ByProjectKeyLoginRequestBuilder) Post(body CustomerSignin) *ByProjectKeyLoginRequestMethodPost {
	return &ByProjectKeyLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/login", rb.projectKey),
		client: rb.client,
	}
}
