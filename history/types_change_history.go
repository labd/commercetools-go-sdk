package history

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Captures the differences between the previous and next version of a resource.
*
*	The maximum number of Records that can be stored and their retention period are subject to a [limit](/../api/limits#records).
*
 */
type Record struct {
	// Version of the resource after the change.
	//
	// For more information on how the version is incremented, see [Optimistic Concurrency Control](/../api/general-concepts#optimistic-concurrency-control).
	Version int `json:"version"`
	// Version of the resource before the change.
	PreviousVersion int `json:"previousVersion"`
	// Indicates the type of change.
	// For creation, update, or deletion, the value is `"ResourceCreated"`, `"ResourceUpdated"`, or `"ResourceDeleted"` respectively.
	Type string `json:"type"`
	// Information about the user or API Client who performed the change.
	ModifiedBy ModifiedBy `json:"modifiedBy"`
	// Date and time (UTC) when the change was made.
	ModifiedAt string `json:"modifiedAt"`
	// Information that describes the resource after the change.
	Label Label `json:"label"`
	// Information that describes the resource before the change.
	PreviousLabel Label `json:"previousLabel"`
	// Shows the differences in the resource between `previousVersion` and `version`.
	//
	// The value is not identical to the actual array of update actions sent and is not limited to update actions (see, for example, [Optimistic  Concurrency Control](/general-concepts#optimistic-concurrency-control)).
	Changes []Change `json:"changes"`
	// ResourceIdentifier of the changed resource.
	Resource ResourceIdentifier `json:"resource"`
	// References to the [Stores](ctp:api:type:Store) associated with the [Change](ctp:history:type:Change).
	Stores []KeyReference `json:"stores"`
	// Reference to the [Business Unit](ctp:api:type:BusinessUnit) associated with the [Change](ctp:history:type:Change).
	BusinessUnit *KeyReference `json:"businessUnit,omitempty"`
	// `true` if no change was detected.
	//
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
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Record](ctp:history:type:Record).
*
 */
type RecordPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation and not [strongly consistent](/../api/general-concepts#strong-consistency).
	Total int `json:"total"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Records matching the query.
	Results []Record `json:"results"`
}

/**
*	This data type represents the supported resource types.
*	The value must be one of the following:
*
 */
type ChangeHistoryResourceType string

const (
	ChangeHistoryResourceTypeAssociateRole    ChangeHistoryResourceType = "associate-role"
	ChangeHistoryResourceTypeBusinessUnit     ChangeHistoryResourceType = "business-unit"
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
	ChangeHistoryResourceTypeProductSelection ChangeHistoryResourceType = "product-selection"
	ChangeHistoryResourceTypeProductType      ChangeHistoryResourceType = "product-type"
	ChangeHistoryResourceTypeQuoteRequest     ChangeHistoryResourceType = "quote-request"
	ChangeHistoryResourceTypeQuote            ChangeHistoryResourceType = "quote"
	ChangeHistoryResourceTypeReview           ChangeHistoryResourceType = "review"
	ChangeHistoryResourceTypeShoppingList     ChangeHistoryResourceType = "shopping-list"
	ChangeHistoryResourceTypeStagedQuote      ChangeHistoryResourceType = "staged-quote"
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
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ErrorObject: failed to parse error response"
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}

func (obj ErrorResponse) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ErrorResponse: failed to parse error response"
}

/**
*	Information about the user or API Client who performed the change. This is a variant of [LastModifiedBy](ctp:api:type:LastModifiedBy).
*
 */
type ModifiedBy struct {
	// [ID](/general-concepts#identifier) of the Merchant Center user who made the change.
	//
	// Present only if the change was made in the Merchant Center.
	ID string `json:"id"`
	// Indicates who performed the change.
	//
	// - If the change was made by a user, the value is `"user"`.
	// - If the change was made by an API Client with or without an [external user ID](/general-concepts#external-user-ids), the value is `"external-user"`.
	// - If the change was made by an [Associate](ctp:api:type:Associate), the value is `"associate"`.
	Type string `json:"type"`
	// [ID](/general-concepts#identifier) of the [API Client](ctp:api:type:ApiClient) that made the change.
	//
	// Present only if the change was made using an API Client.
	ClientId *string `json:"clientId,omitempty"`
	// Present only if the change was made using a token from an [anonymous session](/authorization#tokens-for-anonymous-sessions).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// The [Customer](ctp:api:type:Customer) who made the change.
	//
	// Present only if the change was made using a token from the [password flow](/authorization#password-flow).
	Customer *Reference `json:"customer,omitempty"`
	// The [Associate](ctp:api:type:Associate) who made the change in the context of a [Business Unit](ctp:api:type:BusinessUnit). Present only if the Associate acts on behalf of a company using the [associate endpoints](/associates-overview#on-the-associate-endpoints).
	Associate *Reference `json:"associate,omitempty"`
	// `true` if the change was made using the Merchant Center or [ImpEx](https://impex.europe-west1.gcp.commercetools.com/).
	IsPlatformClient bool `json:"isPlatformClient"`
}

/**
*	Updates that are triggered automatically as a result of a user-initiated change.
 */
type PlatformInitiatedChange string

const (
	PlatformInitiatedChangeExcludeAll                   PlatformInitiatedChange = "excludeAll"
	PlatformInitiatedChangeChangeLineItemName           PlatformInitiatedChange = "changeLineItemName"
	PlatformInitiatedChangeChangeReviewRatingStatistics PlatformInitiatedChange = "changeReviewRatingStatistics"
	PlatformInitiatedChangeSetApplicationVersion        PlatformInitiatedChange = "setApplicationVersion"
	PlatformInitiatedChangeSetIsValid                   PlatformInitiatedChange = "setIsValid"
	PlatformInitiatedChangeSetVariantAvailability       PlatformInitiatedChange = "setVariantAvailability"
)

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
	UpdateTypeAddAddress                            UpdateType = "addAddress"
	UpdateTypeAddAsset                              UpdateType = "addAsset"
	UpdateTypeAddAssociate                          UpdateType = "addAssociate"
	UpdateTypeAddAttributeDefinition                UpdateType = "addAttributeDefinition"
	UpdateTypeAddBillingAddressId                   UpdateType = "addBillingAddressId"
	UpdateTypeAddCustomLineItem                     UpdateType = "addCustomLineItem"
	UpdateTypeAddDelivery                           UpdateType = "addDelivery"
	UpdateTypeAddDiscountCode                       UpdateType = "addDiscountCode"
	UpdateTypeAddEnumValue                          UpdateType = "addEnumValue"
	UpdateTypeAddExternalImage                      UpdateType = "addExternalImage"
	UpdateTypeAddFieldDefinition                    UpdateType = "addFieldDefinition"
	UpdateTypeAddInterfaceInteraction               UpdateType = "addInterfaceInteraction"
	UpdateTypeAddItemShippingAddress                UpdateType = "addItemShippingAddress"
	UpdateTypeAddLineItem                           UpdateType = "addLineItem"
	UpdateTypeAddLocalizedEnumValue                 UpdateType = "addLocalizedEnumValue"
	UpdateTypeAddLocation                           UpdateType = "addLocation"
	UpdateTypeAddParcelToDelivery                   UpdateType = "addParcelToDelivery"
	UpdateTypeAddPayment                            UpdateType = "addPayment"
	UpdateTypeAddPlainEnumValue                     UpdateType = "addPlainEnumValue"
	UpdateTypeAddPrice                              UpdateType = "addPrice"
	UpdateTypeAddProduct                            UpdateType = "addProduct"
	UpdateTypeAddProductSelection                   UpdateType = "addProductSelection"
	UpdateTypeAddProperty                           UpdateType = "addProperty"
	UpdateTypeAddReturnInfo                         UpdateType = "addReturnInfo"
	UpdateTypeAddRoles                              UpdateType = "addRoles"
	UpdateTypeAddShippingAddressId                  UpdateType = "addShippingAddressId"
	UpdateTypeAddTaxRate                            UpdateType = "addTaxRate"
	UpdateTypeAddTextLineItem                       UpdateType = "addTextLineItem"
	UpdateTypeAddToCategory                         UpdateType = "addToCategory"
	UpdateTypeAddTransaction                        UpdateType = "addTransaction"
	UpdateTypeAddVariant                            UpdateType = "addVariant"
	UpdateTypeChangeAddress                         UpdateType = "changeAddress"
	UpdateTypeChangeAmountAuthorized                UpdateType = "changeAmountAuthorized"
	UpdateTypeChangeAmountPlanned                   UpdateType = "changeAmountPlanned"
	UpdateTypeChangeAssetName                       UpdateType = "changeAssetName"
	UpdateTypeChangeAssetOrder                      UpdateType = "changeAssetOrder"
	UpdateTypeChangeAssociate                       UpdateType = "changeAssociate"
	UpdateTypeChangeAssociateMode                   UpdateType = "changeAssociateMode"
	UpdateTypeChangeAttributeConstraint             UpdateType = "changeAttributeConstraint"
	UpdateTypeChangeAttributeName                   UpdateType = "changeAttributeName"
	UpdateTypeChangeAttributeOrderByName            UpdateType = "changeAttributeOrderByName"
	UpdateTypeChangeCartDiscounts                   UpdateType = "changeCartDiscounts"
	UpdateTypeChangeCartPredicate                   UpdateType = "changeCartPredicate"
	UpdateTypeChangeCustomLineItemQuantity          UpdateType = "changeCustomLineItemQuantity"
	UpdateTypeChangeDescription                     UpdateType = "changeDescription"
	UpdateTypeChangeEmail                           UpdateType = "changeEmail"
	UpdateTypeChangeEnumKey                         UpdateType = "changeEnumKey"
	UpdateTypeChangeEnumValueLabel                  UpdateType = "changeEnumValueLabel"
	UpdateTypeChangeEnumValueOrder                  UpdateType = "changeEnumValueOrder"
	UpdateTypeChangeFieldDefinitionOrder            UpdateType = "changeFieldDefinitionOrder"
	UpdateTypeChangeGroups                          UpdateType = "changeGroups"
	UpdateTypeChangeInitial                         UpdateType = "changeInitial"
	UpdateTypeChangeInputHint                       UpdateType = "changeInputHint"
	UpdateTypeChangeIsActive                        UpdateType = "changeIsActive"
	UpdateTypeChangeIsSearchable                    UpdateType = "changeIsSearchable"
	UpdateTypeChangeKey                             UpdateType = "changeKey"
	UpdateTypeChangeLabel                           UpdateType = "changeLabel"
	UpdateTypeChangeLineItemName                    UpdateType = "changeLineItemName"
	UpdateTypeChangeLineItemQuantity                UpdateType = "changeLineItemQuantity"
	UpdateTypeChangeLineItemsOrder                  UpdateType = "changeLineItemsOrder"
	UpdateTypeChangeLocalizedEnumValueLabel         UpdateType = "changeLocalizedEnumValueLabel"
	UpdateTypeChangeLocalizedEnumValueOrder         UpdateType = "changeLocalizedEnumValueOrder"
	UpdateTypeChangeMasterVariant                   UpdateType = "changeMasterVariant"
	UpdateTypeChangeName                            UpdateType = "changeName"
	UpdateTypeChangeOrderHint                       UpdateType = "changeOrderHint"
	UpdateTypeChangeOrderState                      UpdateType = "changeOrderState"
	UpdateTypeChangeParent                          UpdateType = "changeParent"
	UpdateTypeChangeParentUnit                      UpdateType = "changeParentUnit"
	UpdateTypeChangePaymentState                    UpdateType = "changePaymentState"
	UpdateTypeChangePlainEnumValueLabel             UpdateType = "changePlainEnumValueLabel"
	UpdateTypeChangePredicate                       UpdateType = "changePredicate"
	UpdateTypeChangePrice                           UpdateType = "changePrice"
	UpdateTypeChangeProductSelectionActive          UpdateType = "changeProductSelectionActive"
	UpdateTypeChangeQuantity                        UpdateType = "changeQuantity"
	UpdateTypeChangeQuoteRequestState               UpdateType = "changeQuoteRequestState"
	UpdateTypeChangeQuoteState                      UpdateType = "changeQuoteState"
	UpdateTypeChangeRequiresDiscountCode            UpdateType = "changeRequiresDiscountCode"
	UpdateTypeChangeReviewRatingStatistics          UpdateType = "changeReviewRatingStatistics"
	UpdateTypeChangeShipmentState                   UpdateType = "changeShipmentState"
	UpdateTypeChangeSlug                            UpdateType = "changeSlug"
	UpdateTypeChangeSortOrder                       UpdateType = "changeSortOrder"
	UpdateTypeChangeStackingMode                    UpdateType = "changeStackingMode"
	UpdateTypeChangeStagedQuoteState                UpdateType = "changeStagedQuoteState"
	UpdateTypeChangeStatus                          UpdateType = "changeStatus"
	UpdateTypeChangeTarget                          UpdateType = "changeTarget"
	UpdateTypeChangeTaxCalculationMode              UpdateType = "changeTaxCalculationMode"
	UpdateTypeChangeTaxMode                         UpdateType = "changeTaxMode"
	UpdateTypeChangeTaxRoundingMode                 UpdateType = "changeTaxRoundingMode"
	UpdateTypeChangeTextLineItemName                UpdateType = "changeTextLineItemName"
	UpdateTypeChangeTextLineItemQuantity            UpdateType = "changeTextLineItemQuantity"
	UpdateTypeChangeTextLineItemsOrder              UpdateType = "changeTextLineItemsOrder"
	UpdateTypeChangeTransactionInteractionId        UpdateType = "changeTransactionInteractionId"
	UpdateTypeChangeTransactionState                UpdateType = "changeTransactionState"
	UpdateTypeChangeTransactionTimestamp            UpdateType = "changeTransactionTimestamp"
	UpdateTypeChangeType                            UpdateType = "changeType"
	UpdateTypeChangeValue                           UpdateType = "changeValue"
	UpdateTypeMoveImageToPosition                   UpdateType = "moveImageToPosition"
	UpdateTypePublish                               UpdateType = "publish"
	UpdateTypeRemoveAddress                         UpdateType = "removeAddress"
	UpdateTypeRemoveAsset                           UpdateType = "removeAsset"
	UpdateTypeRemoveAssociate                       UpdateType = "removeAssociate"
	UpdateTypeRemoveAttributeDefinition             UpdateType = "removeAttributeDefinition"
	UpdateTypeRemoveBillingAddressId                UpdateType = "removeBillingAddressId"
	UpdateTypeRemoveCustomLineItem                  UpdateType = "removeCustomLineItem"
	UpdateTypeRemoveDelivery                        UpdateType = "removeDelivery"
	UpdateTypeRemoveDiscountCode                    UpdateType = "removeDiscountCode"
	UpdateTypeRemoveEnumValues                      UpdateType = "removeEnumValues"
	UpdateTypeRemoveFieldDefinition                 UpdateType = "removeFieldDefinition"
	UpdateTypeRemoveFromCategory                    UpdateType = "removeFromCategory"
	UpdateTypeRemoveImage                           UpdateType = "removeImage"
	UpdateTypeRemoveItemShippingAddress             UpdateType = "removeItemShippingAddress"
	UpdateTypeRemoveLineItem                        UpdateType = "removeLineItem"
	UpdateTypeRemoveLocation                        UpdateType = "removeLocation"
	UpdateTypeRemoveParcelFromDelivery              UpdateType = "removeParcelFromDelivery"
	UpdateTypeRemovePayment                         UpdateType = "removePayment"
	UpdateTypeRemovePrice                           UpdateType = "removePrice"
	UpdateTypeRemoveProduct                         UpdateType = "removeProduct"
	UpdateTypeRemoveProductSelection                UpdateType = "removeProductSelection"
	UpdateTypeRemoveProperty                        UpdateType = "removeProperty"
	UpdateTypeRemoveRoles                           UpdateType = "removeRoles"
	UpdateTypeRemoveShippingAddressId               UpdateType = "removeShippingAddressId"
	UpdateTypeRemoveTaxRate                         UpdateType = "removeTaxRate"
	UpdateTypeRemoveTextLineItem                    UpdateType = "removeTextLineItem"
	UpdateTypeRemoveVariant                         UpdateType = "removeVariant"
	UpdateTypeRequestQuoteRenegotiation             UpdateType = "requestQuoteRenegotiation"
	UpdateTypeSetAddress                            UpdateType = "setAddress"
	UpdateTypeSetAddressCustomField                 UpdateType = "setAddressCustomField"
	UpdateTypeSetAddressCustomType                  UpdateType = "setAddressCustomType"
	UpdateTypeSetAnonymousId                        UpdateType = "setAnonymousId"
	UpdateTypeSetApplicationVersion                 UpdateType = "setApplicationVersion"
	UpdateTypeSetAssetCustomField                   UpdateType = "setAssetCustomField"
	UpdateTypeSetAssetCustomType                    UpdateType = "setAssetCustomType"
	UpdateTypeSetAssetDescription                   UpdateType = "setAssetDescription"
	UpdateTypeSetAssetSources                       UpdateType = "setAssetSources"
	UpdateTypeSetAssetTags                          UpdateType = "setAssetTags"
	UpdateTypeSetAsssetKey                          UpdateType = "setAsssetKey"
	UpdateTypeSetAttribute                          UpdateType = "setAttribute"
	UpdateTypeSetAuthenticationMode                 UpdateType = "setAuthenticationMode"
	UpdateTypeSetAuthorName                         UpdateType = "setAuthorName"
	UpdateTypeSetBillingAddress                     UpdateType = "setBillingAddress"
	UpdateTypeSetCartPredicate                      UpdateType = "setCartPredicate"
	UpdateTypeSetCategoryOrderHint                  UpdateType = "setCategoryOrderHint"
	UpdateTypeSetCompanyName                        UpdateType = "setCompanyName"
	UpdateTypeSetContactEmail                       UpdateType = "setContactEmail"
	UpdateTypeSetCountries                          UpdateType = "setCountries"
	UpdateTypeSetCountry                            UpdateType = "setCountry"
	UpdateTypeSetCustomField                        UpdateType = "setCustomField"
	UpdateTypeSetCustomLineItemCustomField          UpdateType = "setCustomLineItemCustomField"
	UpdateTypeSetCustomLineItemCustomType           UpdateType = "setCustomLineItemCustomType"
	UpdateTypeSetCustomLineItemMoney                UpdateType = "setCustomLineItemMoney"
	UpdateTypeSetCustomLineItemShippingDetails      UpdateType = "setCustomLineItemShippingDetails"
	UpdateTypeSetCustomLineItemTaxAmount            UpdateType = "setCustomLineItemTaxAmount"
	UpdateTypeSetCustomLineItemTaxCategory          UpdateType = "setCustomLineItemTaxCategory"
	UpdateTypeSetCustomLineItemTaxRate              UpdateType = "setCustomLineItemTaxRate"
	UpdateTypeSetCustomLineItemTaxedPrice           UpdateType = "setCustomLineItemTaxedPrice"
	UpdateTypeSetCustomLineItemTotalPrice           UpdateType = "setCustomLineItemTotalPrice"
	UpdateTypeSetCustomShippingMethod               UpdateType = "setCustomShippingMethod"
	UpdateTypeSetCustomType                         UpdateType = "setCustomType"
	UpdateTypeSetCustomer                           UpdateType = "setCustomer"
	UpdateTypeSetCustomerEmail                      UpdateType = "setCustomerEmail"
	UpdateTypeSetCustomerGroup                      UpdateType = "setCustomerGroup"
	UpdateTypeSetCustomerId                         UpdateType = "setCustomerId"
	UpdateTypeSetCustomerNumber                     UpdateType = "setCustomerNumber"
	UpdateTypeSetDateOfBirth                        UpdateType = "setDateOfBirth"
	UpdateTypeSetDefaultBillingAddress              UpdateType = "setDefaultBillingAddress"
	UpdateTypeSetDefaultShippingAddress             UpdateType = "setDefaultShippingAddress"
	UpdateTypeSetDeleteDaysAfterLastModification    UpdateType = "setDeleteDaysAfterLastModification"
	UpdateTypeSetDeliveryAddress                    UpdateType = "setDeliveryAddress"
	UpdateTypeSetDeliveryItems                      UpdateType = "setDeliveryItems"
	UpdateTypeSetDescription                        UpdateType = "setDescription"
	UpdateTypeSetDiscountedPrice                    UpdateType = "setDiscountedPrice"
	UpdateTypeSetDistributionChannels               UpdateType = "setDistributionChannels"
	UpdateTypeSetExpectedDelivery                   UpdateType = "setExpectedDelivery"
	UpdateTypeSetExternalId                         UpdateType = "setExternalId"
	UpdateTypeSetFirstName                          UpdateType = "setFirstName"
	UpdateTypeSetGeoLocation                        UpdateType = "setGeoLocation"
	UpdateTypeSetImageLabel                         UpdateType = "setImageLabel"
	UpdateTypeSetInputTip                           UpdateType = "setInputTip"
	UpdateTypeSetInterfaceId                        UpdateType = "setInterfaceId"
	UpdateTypeSetIsValid                            UpdateType = "setIsValid"
	UpdateTypeSetKey                                UpdateType = "setKey"
	UpdateTypeSetLanguages                          UpdateType = "setLanguages"
	UpdateTypeSetLastName                           UpdateType = "setLastName"
	UpdateTypeSetLineItemCustomField                UpdateType = "setLineItemCustomField"
	UpdateTypeSetLineItemCustomType                 UpdateType = "setLineItemCustomType"
	UpdateTypeSetLineItemDeactivatedAt              UpdateType = "setLineItemDeactivatedAt"
	UpdateTypeSetLineItemDiscountedPrice            UpdateType = "setLineItemDiscountedPrice"
	UpdateTypeSetLineItemDiscountedPricePerQuantity UpdateType = "setLineItemDiscountedPricePerQuantity"
	UpdateTypeSetLineItemDistributionChannel        UpdateType = "setLineItemDistributionChannel"
	UpdateTypeSetLineItemPrice                      UpdateType = "setLineItemPrice"
	UpdateTypeSetLineItemProductKey                 UpdateType = "setLineItemProductKey"
	UpdateTypeSetLineItemProductSlug                UpdateType = "setLineItemProductSlug"
	UpdateTypeSetLineItemShippingDetails            UpdateType = "setLineItemShippingDetails"
	UpdateTypeSetLineItemTaxAmount                  UpdateType = "setLineItemTaxAmount"
	UpdateTypeSetLineItemTaxRate                    UpdateType = "setLineItemTaxRate"
	UpdateTypeSetLineItemTaxedPrice                 UpdateType = "setLineItemTaxedPrice"
	UpdateTypeSetLineItemTotalPrice                 UpdateType = "setLineItemTotalPrice"
	UpdateTypeSetLocale                             UpdateType = "setLocale"
	UpdateTypeSetMaxApplications                    UpdateType = "setMaxApplications"
	UpdateTypeSetMaxApplicationsPerCustomer         UpdateType = "setMaxApplicationsPerCustomer"
	UpdateTypeSetMetaDescription                    UpdateType = "setMetaDescription"
	UpdateTypeSetMetaKeywords                       UpdateType = "setMetaKeywords"
	UpdateTypeSetMetaTitle                          UpdateType = "setMetaTitle"
	UpdateTypeSetMethodInfoInterface                UpdateType = "setMethodInfoInterface"
	UpdateTypeSetMethodInfoMethod                   UpdateType = "setMethodInfoMethod"
	UpdateTypeSetMethodInfoName                     UpdateType = "setMethodInfoName"
	UpdateTypeSetMiddleName                         UpdateType = "setMiddleName"
	UpdateTypeSetName                               UpdateType = "setName"
	UpdateTypeSetOrderNumber                        UpdateType = "setOrderNumber"
	UpdateTypeSetOrderTaxedPrice                    UpdateType = "setOrderTaxedPrice"
	UpdateTypeSetOrderTotalPrice                    UpdateType = "setOrderTotalPrice"
	UpdateTypeSetOrderTotalTax                      UpdateType = "setOrderTotalTax"
	UpdateTypeSetParcelItems                        UpdateType = "setParcelItems"
	UpdateTypeSetParcelMeasurements                 UpdateType = "setParcelMeasurements"
	UpdateTypeSetParcelTrackingData                 UpdateType = "setParcelTrackingData"
	UpdateTypeSetPassword                           UpdateType = "setPassword"
	UpdateTypeSetPrices                             UpdateType = "setPrices"
	UpdateTypeSetProductCount                       UpdateType = "setProductCount"
	UpdateTypeSetProductPriceCustomField            UpdateType = "setProductPriceCustomField"
	UpdateTypeSetProductPriceCustomType             UpdateType = "setProductPriceCustomType"
	UpdateTypeSetProductSelections                  UpdateType = "setProductSelections"
	UpdateTypeSetProductVariantKey                  UpdateType = "setProductVariantKey"
	UpdateTypeSetProperty                           UpdateType = "setProperty"
	UpdateTypeSetPurchaseOrderNumber                UpdateType = "setPurchaseOrderNumber"
	UpdateTypeSetRating                             UpdateType = "setRating"
	UpdateTypeSetReservations                       UpdateType = "setReservations"
	UpdateTypeSetRestockableInDays                  UpdateType = "setRestockableInDays"
	UpdateTypeSetReturnPaymentState                 UpdateType = "setReturnPaymentState"
	UpdateTypeSetReturnShipmentState                UpdateType = "setReturnShipmentState"
	UpdateTypeSetRoles                              UpdateType = "setRoles"
	UpdateTypeSetSalutation                         UpdateType = "setSalutation"
	UpdateTypeSetSearchKeywords                     UpdateType = "setSearchKeywords"
	UpdateTypeSetSellerComment                      UpdateType = "setSellerComment"
	UpdateTypeSetShippingAddress                    UpdateType = "setShippingAddress"
	UpdateTypeSetShippingInfoPrice                  UpdateType = "setShippingInfoPrice"
	UpdateTypeSetShippingInfoTaxedPrice             UpdateType = "setShippingInfoTaxedPrice"
	UpdateTypeSetShippingMethod                     UpdateType = "setShippingMethod"
	UpdateTypeSetShippingMethodTaxAmount            UpdateType = "setShippingMethodTaxAmount"
	UpdateTypeSetShippingMethodTaxRate              UpdateType = "setShippingMethodTaxRate"
	UpdateTypeSetShippingRate                       UpdateType = "setShippingRate"
	UpdateTypeSetShippingRateInput                  UpdateType = "setShippingRateInput"
	UpdateTypeSetSku                                UpdateType = "setSku"
	UpdateTypeSetSlug                               UpdateType = "setSlug"
	UpdateTypeSetStatusInterfaceCode                UpdateType = "setStatusInterfaceCode"
	UpdateTypeSetStatusInterfaceText                UpdateType = "setStatusInterfaceText"
	UpdateTypeSetStore                              UpdateType = "setStore"
	UpdateTypeSetStoreMode                          UpdateType = "setStoreMode"
	UpdateTypeSetStores                             UpdateType = "setStores"
	UpdateTypeSetSupplyChannel                      UpdateType = "setSupplyChannel"
	UpdateTypeSetSupplyChannels                     UpdateType = "setSupplyChannels"
	UpdateTypeSetTarget                             UpdateType = "setTarget"
	UpdateTypeSetTaxCategory                        UpdateType = "setTaxCategory"
	UpdateTypeSetText                               UpdateType = "setText"
	UpdateTypeSetTextLineItemCustomField            UpdateType = "setTextLineItemCustomField"
	UpdateTypeSetTextLineItemCustomType             UpdateType = "setTextLineItemCustomType"
	UpdateTypeSetTextLineItemDescription            UpdateType = "setTextLineItemDescription"
	UpdateTypeSetTitle                              UpdateType = "setTitle"
	UpdateTypeSetTransitions                        UpdateType = "setTransitions"
	UpdateTypeSetValidFrom                          UpdateType = "setValidFrom"
	UpdateTypeSetValidFromAndUntil                  UpdateType = "setValidFromAndUntil"
	UpdateTypeSetValidTo                            UpdateType = "setValidTo"
	UpdateTypeSetValidUntil                         UpdateType = "setValidUntil"
	UpdateTypeSetValue                              UpdateType = "setValue"
	UpdateTypeSetVariantAvailability                UpdateType = "setVariantAvailability"
	UpdateTypeSetVariantSelection                   UpdateType = "setVariantSelection"
	UpdateTypeSetVatId                              UpdateType = "setVatId"
	UpdateTypeTransitionCustomLineItemState         UpdateType = "transitionCustomLineItemState"
	UpdateTypeTransitionLineItemState               UpdateType = "transitionLineItemState"
	UpdateTypeTransitionState                       UpdateType = "transitionState"
	UpdateTypeUnpublish                             UpdateType = "unpublish"
	UpdateTypeUpdateItemShippingAddress             UpdateType = "updateItemShippingAddress"
	UpdateTypeUpdateSyncInfo                        UpdateType = "updateSyncInfo"
	UpdateTypeVerifyEmail                           UpdateType = "verifyEmail"
)
