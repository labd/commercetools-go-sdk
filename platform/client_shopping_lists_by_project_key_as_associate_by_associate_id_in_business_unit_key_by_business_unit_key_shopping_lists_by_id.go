package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	id              string
	client          *Client
}

/**
*	Retrieves a ShoppingList with the provided `key` in a [BusinessUnit](ctp:api:type:BusinessUnit).
*
*	If the ShoppingList exists in the Project but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists with the provided `id` in a [BusinessUnit](ctp:api:type:BusinessUnit). Returns a `200 OK` if the ShoppingList exists; otherwise, returns [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList in a [BusinessUnit](ctp:api:type:BusinessUnit) using one or more [update actions](/../api/projects/shoppingLists#update-actions).
*	If the ShoppingList exists in the Project but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a ShoppingList in a [BusinessUnit](ctp:api:type:BusinessUnit).
*
*	If the ShoppingList exists in the Project but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}
