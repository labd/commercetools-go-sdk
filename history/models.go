// Generated file, please do not change!!!
package history

import (
	"encoding/json"
)

/**
*	A Record captures the differences in a resource between one version and the next.
*	(Recall that the version number is not always incremented by one; see [Optimistic Concurrency Control](/general-concepts#optimistic-concurrency-control).)
*
 */
type Record struct {
	// Version of the resource after the change.
	Version int `json:"version"`
	// Version of the resource before the change.
	PreviousVersion int `json:"previousVersion"`
	// Type of the change (creation, update or deletion).
	Type string `json:"type"`
	// Information about the user or the API client who performed the change.
	ModifiedBy ModifiedBy `json:"modifiedBy"`
	// Date and time when the change was made.
	ModifiedAt string `json:"modifiedAt"`
	// Information that describes the resource after the change.
	Label Label `json:"label"`
	// Information that describes the resource before the change.
	PreviousLabel Label `json:"previousLabel"`
	// Shows the differences in the resource between `previousVersion` and `version`.
	// The value is not identical to the actual array of update actions that was sent to the platform and is not limited to update actions (see, for example, [Optimistic  Concurrency Control](/general-concepts#optimistic-concurrency-control)).
	Changes []Change `json:"changes"`
	// [Reference](/types#reference) to the changed resource.
	Resource Reference `json:"resource"`
	// `true` if no change was detected.
	// The version number of the resource can be increased even without any change in the resource.
	WithoutChanges bool `json:"withoutChanges"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Record) UnmarshalJSON(data []byte) error {
	type Alias Record
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Label != nil {
		var err error
		obj.Label, err = mapDiscriminatorLabel(obj.Label)
		if err != nil {
			return err
		}
	}
	if obj.PreviousLabel != nil {
		var err error
		obj.PreviousLabel, err = mapDiscriminatorLabel(obj.PreviousLabel)
		if err != nil {
			return err
		}
	}
	for i := range obj.Changes {
		var err error
		obj.Changes[i], err = mapDiscriminatorChange(obj.Changes[i])
		if err != nil {
			return err
		}
	}
	return nil
}

/**
*	Response to a query request for [Record](#record).
*
 */
type RecordPagedQueryResponse struct {
	// Maximum number of results requested in the query request.
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation and not [strongly consistent](/general-concepts#strong-consistency).
	Total int `json:"total"`
	// The number of elements skipped, not a page number. Supplied by the client or the server default.
	Offset  int      `json:"offset"`
	Results []Record `json:"results"`
}

/**
*	This data type represents the supported resource types.
*	The value must be one of the following:
*
 */
type ChangeHistoryResourceType string

const (
	ChangeHistoryResourceTypeCartDiscount     ChangeHistoryResourceType = "cart-discount"
	ChangeHistoryResourceTypeCategory         ChangeHistoryResourceType = "category"
	ChangeHistoryResourceTypeChannel          ChangeHistoryResourceType = "channel"
	ChangeHistoryResourceTypeCustomer         ChangeHistoryResourceType = "customer"
	ChangeHistoryResourceTypeCustomerGroup    ChangeHistoryResourceType = "customer-group"
	ChangeHistoryResourceTypeDiscountCode     ChangeHistoryResourceType = "discount-code"
	ChangeHistoryResourceTypeInventoryEntry   ChangeHistoryResourceType = "inventory-entry"
	ChangeHistoryResourceTypeKeyValueDocument ChangeHistoryResourceType = "key-value-document"
	ChangeHistoryResourceTypeOrder            ChangeHistoryResourceType = "order"
	ChangeHistoryResourceTypePayment          ChangeHistoryResourceType = "payment"
	ChangeHistoryResourceTypeProduct          ChangeHistoryResourceType = "product"
	ChangeHistoryResourceTypeProductDiscount  ChangeHistoryResourceType = "product-discount"
	ChangeHistoryResourceTypeProductType      ChangeHistoryResourceType = "product-type"
	ChangeHistoryResourceTypeReview           ChangeHistoryResourceType = "review"
	ChangeHistoryResourceTypeShoppingList     ChangeHistoryResourceType = "shopping-list"
	ChangeHistoryResourceTypeState            ChangeHistoryResourceType = "state"
	ChangeHistoryResourceTypeStore            ChangeHistoryResourceType = "store"
	ChangeHistoryResourceTypeTaxCategory      ChangeHistoryResourceType = "tax-category"
	ChangeHistoryResourceTypeType             ChangeHistoryResourceType = "type"
	ChangeHistoryResourceTypeZone             ChangeHistoryResourceType = "zone"
)

/**
*	This type consists of one enum value:
*
 */
type DateStringFilter string

const (
	DateStringFilterNow DateStringFilter = "now"
)

type ErrorObject struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (obj ErrorObject) Error() string {
	return obj.Message
}

type ErrorResponse struct {
	StatusCode       int           `json:"statusCode"`
	Message          string        `json:"message"`
	ErrorMessage     *string       `json:"error,omitempty"`
	ErrorDescription *string       `json:"error_description,omitempty"`
	Errors           []ErrorObject `json:"errors"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ErrorResponse) MarshalJSON() ([]byte, error) {
	type Alias ErrorResponse
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["errors"] == nil {
		delete(target, "errors")
	}

	return json.Marshal(target)
}
func (obj ErrorResponse) Error() string {
	return obj.Message
}

/**
*	Information about the user or the API client who performed the change. This is a variant of
*	[LastModifiedBy](/types#lastmodifiedby).
*
 */
type ModifiedBy struct {
	// [ID](/general-concepts#identifier) of the Merchant Center user who made the change.
	// Present only if the change was made in the Merchant Center.
	ID string `json:"id"`
	// Indicates whether the change was made by a user or the API client with or without an
	// [External user ID](/client-logging#external-user-ids).
	Type string `json:"type"`
	// [Reference](/types#reference) to the
	// [Customer](/projects/customers#customer) who made the change. Present only if
	// the change was made using a token from the [Password
	// Flow](/authorization#password-flow).
	Customer *Reference `json:"customer,omitempty"`
	// Present only if the change was made using a token from an [Anonymous
	// Session](/authorization#tokens-for-anonymous-sessions).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// [ID](/general-concepts#identifier) of the [API
	// Client](/projects/api-clients#apiclient) that made the change. Present only if
	// the change was made using an API Client.
	ClientId *string `json:"clientId,omitempty"`
	// `true` if the change was made via Merchant Center or [ImpEx](https://impex.europe-west1.gcp.commercetools.com/).
	IsPlatformClient bool `json:"isPlatformClient"`
}

/**
*	Values for the Source enumeration.
 */
type Source string

const (
	SourceMerchantCenter Source = "MerchantCenter"
	SourceImpEx          Source = "ImpEx"
	SourceApiClient      Source = "ApiClient"
)

type UpdateType string

const (
	UpdateTypeAddAddress                         UpdateType = "addAddress"
	UpdateTypeAddAsset                           UpdateType = "addAsset"
	UpdateTypeAddAttributeDefinition             UpdateType = "addAttributeDefinition"
	UpdateTypeAddBillingAddressId                UpdateType = "addBillingAddressId"
	UpdateTypeAddDelivery                        UpdateType = "addDelivery"
	UpdateTypeAddEnumValue                       UpdateType = "addEnumValue"
	UpdateTypeAddExternalImage                   UpdateType = "addExternalImage"
	UpdateTypeAddFieldDefinition                 UpdateType = "addFieldDefinition"
	UpdateTypeAddInterfaceInteraction            UpdateType = "addInterfaceInteraction"
	UpdateTypeAddItemShippingAddress             UpdateType = "addItemShippingAddress"
	UpdateTypeAddLineItem                        UpdateType = "addLineItem"
	UpdateTypeAddLocalizedEnumValue              UpdateType = "addLocalizedEnumValue"
	UpdateTypeAddLocation                        UpdateType = "addLocation"
	UpdateTypeAddParcelToDelivery                UpdateType = "addParcelToDelivery"
	UpdateTypeAddPayment                         UpdateType = "addPayment"
	UpdateTypeAddPlainEnumValue                  UpdateType = "addPlainEnumValue"
	UpdateTypeAddPrice                           UpdateType = "addPrice"
	UpdateTypeAddReturnInfo                      UpdateType = "addReturnInfo"
	UpdateTypeAddRoles                           UpdateType = "addRoles"
	UpdateTypeAddShippingAddressId               UpdateType = "addShippingAddressId"
	UpdateTypeAddTaxRate                         UpdateType = "addTaxRate"
	UpdateTypeAddTextLineItem                    UpdateType = "addTextLineItem"
	UpdateTypeAddToCategory                      UpdateType = "addToCategory"
	UpdateTypeAddTransaction                     UpdateType = "addTransaction"
	UpdateTypeAddVariant                         UpdateType = "addVariant"
	UpdateTypeChangeAddress                      UpdateType = "changeAddress"
	UpdateTypeChangeAmountPlanned                UpdateType = "changeAmountPlanned"
	UpdateTypeChangeAssetName                    UpdateType = "changeAssetName"
	UpdateTypeChangeAssetOrder                   UpdateType = "changeAssetOrder"
	UpdateTypeChangeAttributeConstraint          UpdateType = "changeAttributeConstraint"
	UpdateTypeChangeAttributeName                UpdateType = "changeAttributeName"
	UpdateTypeChangeAttributeOrderByName         UpdateType = "changeAttributeOrderByName"
	UpdateTypeChangeCartDiscounts                UpdateType = "changeCartDiscounts"
	UpdateTypeChangeCartPredicate                UpdateType = "changeCartPredicate"
	UpdateTypeChangeDescription                  UpdateType = "changeDescription"
	UpdateTypeChangeEmail                        UpdateType = "changeEmail"
	UpdateTypeChangeEnumKey                      UpdateType = "changeEnumKey"
	UpdateTypeChangeEnumValueLabel               UpdateType = "changeEnumValueLabel"
	UpdateTypeChangeEnumValueOrder               UpdateType = "changeEnumValueOrder"
	UpdateTypeChangeFieldDefinitionOrder         UpdateType = "changeFieldDefinitionOrder"
	UpdateTypeChangeGroups                       UpdateType = "changeGroups"
	UpdateTypeChangeInitial                      UpdateType = "changeInitial"
	UpdateTypeChangeInputHint                    UpdateType = "changeInputHint"
	UpdateTypeChangeIsActive                     UpdateType = "changeIsActive"
	UpdateTypeChangeIsSearchable                 UpdateType = "changeIsSearchable"
	UpdateTypeChangeKey                          UpdateType = "changeKey"
	UpdateTypeChangeLabel                        UpdateType = "changeLabel"
	UpdateTypeChangeLineItemQuantity             UpdateType = "changeLineItemQuantity"
	UpdateTypeChangeLineItemsOrder               UpdateType = "changeLineItemsOrder"
	UpdateTypeChangeLocalizedEnumValueLabel      UpdateType = "changeLocalizedEnumValueLabel"
	UpdateTypeChangeLocalizedEnumValueOrder      UpdateType = "changeLocalizedEnumValueOrder"
	UpdateTypeChangeMasterVariant                UpdateType = "changeMasterVariant"
	UpdateTypeChangeName                         UpdateType = "changeName"
	UpdateTypeChangeOrderHint                    UpdateType = "changeOrderHint"
	UpdateTypeChangeOrderState                   UpdateType = "changeOrderState"
	UpdateTypeChangeParent                       UpdateType = "changeParent"
	UpdateTypeChangePaymentState                 UpdateType = "changePaymentState"
	UpdateTypeChangePlainEnumValueLabel          UpdateType = "changePlainEnumValueLabel"
	UpdateTypeChangePredicate                    UpdateType = "changePredicate"
	UpdateTypeChangePrice                        UpdateType = "changePrice"
	UpdateTypeChangeQuantity                     UpdateType = "changeQuantity"
	UpdateTypeChangeRequiresDiscountCode         UpdateType = "changeRequiresDiscountCode"
	UpdateTypeChangeReviewRatingStatistics       UpdateType = "changeReviewRatingStatistics"
	UpdateTypeChangeShipmentState                UpdateType = "changeShipmentState"
	UpdateTypeChangeSlug                         UpdateType = "changeSlug"
	UpdateTypeChangeSortOrder                    UpdateType = "changeSortOrder"
	UpdateTypeChangeStackingMode                 UpdateType = "changeStackingMode"
	UpdateTypeChangeTarget                       UpdateType = "changeTarget"
	UpdateTypeChangeTextLineItemName             UpdateType = "changeTextLineItemName"
	UpdateTypeChangeTextLineItemQuantity         UpdateType = "changeTextLineItemQuantity"
	UpdateTypeChangeTextLineItemsOrder           UpdateType = "changeTextLineItemsOrder"
	UpdateTypeChangeTransactionInteractionId     UpdateType = "changeTransactionInteractionId"
	UpdateTypeChangeTransactionState             UpdateType = "changeTransactionState"
	UpdateTypeChangeTransactionTimestamp         UpdateType = "changeTransactionTimestamp"
	UpdateTypeChangeType                         UpdateType = "changeType"
	UpdateTypeChangeValue                        UpdateType = "changeValue"
	UpdateTypePublish                            UpdateType = "publish"
	UpdateTypeRemoveAddress                      UpdateType = "removeAddress"
	UpdateTypeRemoveAsset                        UpdateType = "removeAsset"
	UpdateTypeRemoveAttributeDefinition          UpdateType = "removeAttributeDefinition"
	UpdateTypeRemoveBillingAddressId             UpdateType = "removeBillingAddressId"
	UpdateTypeRemoveDelivery                     UpdateType = "removeDelivery"
	UpdateTypeRemoveEnumValues                   UpdateType = "removeEnumValues"
	UpdateTypeRemoveFieldDefinition              UpdateType = "removeFieldDefinition"
	UpdateTypeRemoveFromCategory                 UpdateType = "removeFromCategory"
	UpdateTypeRemoveImage                        UpdateType = "removeImage"
	UpdateTypeRemoveItemShippingAddress          UpdateType = "removeItemShippingAddress"
	UpdateTypeRemoveLineItem                     UpdateType = "removeLineItem"
	UpdateTypeRemoveLocation                     UpdateType = "removeLocation"
	UpdateTypeRemoveParcelFromDelivery           UpdateType = "removeParcelFromDelivery"
	UpdateTypeRemovePayment                      UpdateType = "removePayment"
	UpdateTypeRemovePrice                        UpdateType = "removePrice"
	UpdateTypeRemoveRoles                        UpdateType = "removeRoles"
	UpdateTypeRemoveShippingAddressId            UpdateType = "removeShippingAddressId"
	UpdateTypeRemoveTaxRate                      UpdateType = "removeTaxRate"
	UpdateTypeRemoveTextLineItem                 UpdateType = "removeTextLineItem"
	UpdateTypeRemoveVariant                      UpdateType = "removeVariant"
	UpdateTypeSetAddress                         UpdateType = "setAddress"
	UpdateTypeSetAnonymousId                     UpdateType = "setAnonymousId"
	UpdateTypeSetAssetCustomField                UpdateType = "setAssetCustomField"
	UpdateTypeSetAssetCustomType                 UpdateType = "setAssetCustomType"
	UpdateTypeSetAssetDescription                UpdateType = "setAssetDescription"
	UpdateTypeSetAssetSources                    UpdateType = "setAssetSources"
	UpdateTypeSetAssetTags                       UpdateType = "setAssetTags"
	UpdateTypeSetAsssetKey                       UpdateType = "setAsssetKey"
	UpdateTypeSetAttribute                       UpdateType = "setAttribute"
	UpdateTypeSetAuthorName                      UpdateType = "setAuthorName"
	UpdateTypeSetBillingAddress                  UpdateType = "setBillingAddress"
	UpdateTypeSetCartPredicate                   UpdateType = "setCartPredicate"
	UpdateTypeSetCategoryOrderHint               UpdateType = "setCategoryOrderHint"
	UpdateTypeSetCompanyName                     UpdateType = "setCompanyName"
	UpdateTypeSetCustomField                     UpdateType = "setCustomField"
	UpdateTypeSetCustomLineItemCustomField       UpdateType = "setCustomLineItemCustomField"
	UpdateTypeSetCustomLineItemCustomType        UpdateType = "setCustomLineItemCustomType"
	UpdateTypeSetCustomLineItemShippingDetails   UpdateType = "setCustomLineItemShippingDetails"
	UpdateTypeSetCustomType                      UpdateType = "setCustomType"
	UpdateTypeSetCustomer                        UpdateType = "setCustomer"
	UpdateTypeSetCustomerEmail                   UpdateType = "setCustomerEmail"
	UpdateTypeSetCustomerGroup                   UpdateType = "setCustomerGroup"
	UpdateTypeSetCustomerId                      UpdateType = "setCustomerId"
	UpdateTypeSetCustomerNumber                  UpdateType = "setCustomerNumber"
	UpdateTypeSetDateOfBirth                     UpdateType = "setDateOfBirth"
	UpdateTypeSetDefaultBillingAddress           UpdateType = "setDefaultBillingAddress"
	UpdateTypeSetDefaultShippingAddress          UpdateType = "setDefaultShippingAddress"
	UpdateTypeSetDeleteDaysAfterLastModification UpdateType = "setDeleteDaysAfterLastModification"
	UpdateTypeSetDeliveryAddress                 UpdateType = "setDeliveryAddress"
	UpdateTypeSetDeliveryItems                   UpdateType = "setDeliveryItems"
	UpdateTypeSetDescription                     UpdateType = "setDescription"
	UpdateTypeSetDiscountedPrice                 UpdateType = "setDiscountedPrice"
	UpdateTypeSetDistributionChannels            UpdateType = "setDistributionChannels"
	UpdateTypeSetExpectedDelivery                UpdateType = "setExpectedDelivery"
	UpdateTypeSetExternalId                      UpdateType = "setExternalId"
	UpdateTypeSetFirstName                       UpdateType = "setFirstName"
	UpdateTypeSetGeoLocation                     UpdateType = "setGeoLocation"
	UpdateTypeSetImageLabel                      UpdateType = "setImageLabel"
	UpdateTypeSetInputTip                        UpdateType = "setInputTip"
	UpdateTypeSetInterfaceId                     UpdateType = "setInterfaceId"
	UpdateTypeSetKey                             UpdateType = "setKey"
	UpdateTypeSetLanguages                       UpdateType = "setLanguages"
	UpdateTypeSetLastName                        UpdateType = "setLastName"
	UpdateTypeSetLineItemCustomField             UpdateType = "setLineItemCustomField"
	UpdateTypeSetLineItemCustomType              UpdateType = "setLineItemCustomType"
	UpdateTypeSetLineItemShippingDetails         UpdateType = "setLineItemShippingDetails"
	UpdateTypeSetLocale                          UpdateType = "setLocale"
	UpdateTypeSetMaxApplications                 UpdateType = "setMaxApplications"
	UpdateTypeSetMaxApplicationsPerCustomer      UpdateType = "setMaxApplicationsPerCustomer"
	UpdateTypeSetMetaDescription                 UpdateType = "setMetaDescription"
	UpdateTypeSetMetaKeywords                    UpdateType = "setMetaKeywords"
	UpdateTypeSetMetaTitle                       UpdateType = "setMetaTitle"
	UpdateTypeSetMethodInfoInterface             UpdateType = "setMethodInfoInterface"
	UpdateTypeSetMethodInfoMethod                UpdateType = "setMethodInfoMethod"
	UpdateTypeSetMethodInfoName                  UpdateType = "setMethodInfoName"
	UpdateTypeSetMiddleName                      UpdateType = "setMiddleName"
	UpdateTypeSetName                            UpdateType = "setName"
	UpdateTypeSetOrderNumber                     UpdateType = "setOrderNumber"
	UpdateTypeSetParcelItems                     UpdateType = "setParcelItems"
	UpdateTypeSetParcelMeasurements              UpdateType = "setParcelMeasurements"
	UpdateTypeSetParcelTrackingData              UpdateType = "setParcelTrackingData"
	UpdateTypeSetPassword                        UpdateType = "setPassword"
	UpdateTypeSetProductPriceCustomField         UpdateType = "setProductPriceCustomField"
	UpdateTypeSetProductPriceCustomType          UpdateType = "setProductPriceCustomType"
	UpdateTypeSetProductVariantKey               UpdateType = "setProductVariantKey"
	UpdateTypeSetRating                          UpdateType = "setRating"
	UpdateTypeSetRestockableInDays               UpdateType = "setRestockableInDays"
	UpdateTypeSetReturnPaymentState              UpdateType = "setReturnPaymentState"
	UpdateTypeSetReturnShipmentState             UpdateType = "setReturnShipmentState"
	UpdateTypeSetRoles                           UpdateType = "setRoles"
	UpdateTypeSetSalutation                      UpdateType = "setSalutation"
	UpdateTypeSetSearchKeywords                  UpdateType = "setSearchKeywords"
	UpdateTypeSetShippingAddress                 UpdateType = "setShippingAddress"
	UpdateTypeSetSku                             UpdateType = "setSku"
	UpdateTypeSetSlug                            UpdateType = "setSlug"
	UpdateTypeSetStatusInterfaceCode             UpdateType = "setStatusInterfaceCode"
	UpdateTypeSetStatusInterfaceText             UpdateType = "setStatusInterfaceText"
	UpdateTypeSetStore                           UpdateType = "setStore"
	UpdateTypeSetStores                          UpdateType = "setStores"
	UpdateTypeSetSupplyChannel                   UpdateType = "setSupplyChannel"
	UpdateTypeSetTarget                          UpdateType = "setTarget"
	UpdateTypeSetTaxCategory                     UpdateType = "setTaxCategory"
	UpdateTypeSetText                            UpdateType = "setText"
	UpdateTypeSetTextLineItemCustomField         UpdateType = "setTextLineItemCustomField"
	UpdateTypeSetTextLineItemCustomType          UpdateType = "setTextLineItemCustomType"
	UpdateTypeSetTextLineItemDescription         UpdateType = "setTextLineItemDescription"
	UpdateTypeSetTitle                           UpdateType = "setTitle"
	UpdateTypeSetTransitions                     UpdateType = "setTransitions"
	UpdateTypeSetValidFrom                       UpdateType = "setValidFrom"
	UpdateTypeSetValidFromAndUntil               UpdateType = "setValidFromAndUntil"
	UpdateTypeSetValidUntil                      UpdateType = "setValidUntil"
	UpdateTypeSetVariantAvailability             UpdateType = "setVariantAvailability"
	UpdateTypeSetVatId                           UpdateType = "setVatId"
	UpdateTypeTransitionCustomLineItemState      UpdateType = "transitionCustomLineItemState"
	UpdateTypeTransitionLineItemState            UpdateType = "transitionLineItemState"
	UpdateTypeTransitionState                    UpdateType = "transitionState"
	UpdateTypeUnpublish                          UpdateType = "unpublish"
	UpdateTypeUpdateItemShippingAddress          UpdateType = "updateItemShippingAddress"
	UpdateTypeUpdateSyncInfo                     UpdateType = "updateSyncInfo"
	UpdateTypeVerifyEmail                        UpdateType = "verifyEmail"
)

type PlatformInitiatedChange string

const (
	PlatformInitiatedChangeExcludeAll                   PlatformInitiatedChange = "excludeAll"
	PlatformInitiatedChangeChangeLineItemName           PlatformInitiatedChange = "changeLineItemName"
	PlatformInitiatedChangeChangeReviewRatingStatistics PlatformInitiatedChange = "changeReviewRatingStatistics"
	PlatformInitiatedChangeSetApplicationVersion        PlatformInitiatedChange = "setApplicationVersion"
	PlatformInitiatedChangeSetIsValid                   PlatformInitiatedChange = "setIsValid"
	PlatformInitiatedChangeSetVariantAvailability       PlatformInitiatedChange = "setVariantAvailability"
)
