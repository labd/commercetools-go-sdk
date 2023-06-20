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
*	Associates with the `UpdateAssociates` [Permission](ctp:api:type:Permission) can use this endpoint to create a new Customer and associate it with the Business Unit.
*	If the required [Permission](/projects/associate-roles#permission) is missing, an [AssociateMissingPermission](/errors#associatemissingpermission) error is returned.
*	The new Associate is created with an empty set of roles.
*
 */
func (rb *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestBuilder) Post(body MyBusinessUnitAssociateDraft) *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost {
	return &ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-business-unit/key=%s/me/customers", rb.projectKey, rb.businessUnitKey),
		client: rb.client,
	}
}
