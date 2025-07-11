package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	key             string
	client          *Client
}

/**
*	Retrieves a ShoppingList with the provided `key` in a [BusinessUnit](ctp:api:type:BusinessUnit).
*	If the ShoppingList exists in the Project but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists with the provided `key` in a [BusinessUnit](ctp:api:type:BusinessUnit). Returns a `200 OK` if the ShoppingList exists; otherwise, returns [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList in a [BusinessUnit](ctp:api:type:BusinessUnit) using one or more [update actions](/../api/projects/shoppingLists#update-actions).
*	If the ShoppingList exists in the [Project](ctp:api:type:Project) but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a ShoppingList in a [BusinessUnit](ctp:api:type:BusinessUnit).
*
*	If the ShoppingList exists in the Project but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/shopping-lists/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}
