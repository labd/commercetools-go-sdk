package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestBuilder struct {
	projectKey      string
	businessUnitKey string
	client          *Client
}

/**
*	The My Business Unit endpoint does not support assigning existing Customers to a Business Unit.
*	Use this endpoint to create a new Customer and associate it with the Business Unit.
*	The user must have the `Admin` role within the Business Unit to perform this request.
*	The new Customer is created with an empty set of roles.
*
 */
func (rb *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestBuilder) Post(body MyBusinessUnitAssociateDraft) *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost {
	return &ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-business-unit/key=%s/me/customers", rb.projectKey, rb.businessUnitKey),
		client: rb.client,
	}
}
