package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	The current indexing status of Business Unit Search.
*
 */
type BusinessUnitIndexingStatus string

const (
	BusinessUnitIndexingStatusScheduled BusinessUnitIndexingStatus = "Scheduled"
	BusinessUnitIndexingStatusIndexing  BusinessUnitIndexingStatus = "Indexing"
	BusinessUnitIndexingStatusReady     BusinessUnitIndexingStatus = "Ready"
	BusinessUnitIndexingStatusFailed    BusinessUnitIndexingStatus = "Failed"
)

/**
*	The current indexing status of Customer Search.
*
 */
type CustomerIndexingStatus string

const (
	CustomerIndexingStatusScheduled CustomerIndexingStatus = "Scheduled"
	CustomerIndexingStatusIndexing  CustomerIndexingStatus = "Indexing"
	CustomerIndexingStatusReady     CustomerIndexingStatus = "Ready"
	CustomerIndexingStatusFailed    CustomerIndexingStatus = "Failed"
)

type BusinessUnitConfiguration struct {
	// Status of Business Units created using the [My Business Unit endpoint](ctp:api:endpoint:/{projectKey}/me/business-units:POST).
	MyBusinessUnitStatusOnCreation BusinessUnitConfigurationStatus `json:"myBusinessUnitStatusOnCreation"`
	// Default [Associate Role](ctp:api:type:AssociateRole) assigned to the Associate creating a Business Unit using the [My Business Unit endpoint](ctp:api:endpoint:/{projectKey}/me/business-units:POST).
	MyBusinessUnitAssociateRoleOnCreation *AssociateRoleKeyReference `json:"myBusinessUnitAssociateRoleOnCreation,omitempty"`
}

/**
*	Default value for [Business Unit Status](ctp:api:type:BusinessUnitStatus) configured though [Project settings](/../api/projects/project#change-my-business-unit-status-on-creation).
 */
type BusinessUnitConfigurationStatus string

const (
	BusinessUnitConfigurationStatusActive   BusinessUnitConfigurationStatus = "Active"
	BusinessUnitConfigurationStatusInactive BusinessUnitConfigurationStatus = "Inactive"
)

/**
*	Specifies the status of the [Business Unit Search](/../api/projects/business-unit-search) index.
*	You can change the status using the [Change Business Unit Search Status](ctp:api:type:ProjectChangeBusinessUnitSearchStatusAction) update action.
*
 */
type BusinessUnitSearchStatus string

const (
	BusinessUnitSearchStatusActivated   BusinessUnitSearchStatus = "Activated"
	BusinessUnitSearchStatusDeactivated BusinessUnitSearchStatus = "Deactivated"
)

type CartsConfiguration struct {
	// Default value for the `deleteDaysAfterLastModification` parameter of the [CartDraft](ctp:api:type:CartDraft) and [MyCartDraft](ctp:api:type:MyCartDraft).
	// If a [ChangeSubscription](ctp:api:type:ChangeSubscription) for Carts exists, a [ResourceDeletedDeliveryPayload](ctp:api:type:ResourceDeletedDeliveryPayload) is sent upon deletion of a Cart.
	//
	// This field may not be present on Projects created before January 2020.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Indicates if country _- no state_ Tax Rate fallback should be used when a shipping address state is not explicitly covered in the rates lists of all Tax Categories of a Cart Line Items. This field may not be present on Projects created before June 2020.
	CountryTaxRateFallbackEnabled *bool `json:"countryTaxRateFallbackEnabled,omitempty"`
	// Default value for the `priceRoundingMode` parameter of the [CartDraft](ctp:api:type:CartDraft).
	// Indicates how the total prices on [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem) are rounded when calculated.
	PriceRoundingMode *RoundingMode `json:"priceRoundingMode,omitempty"`
	// Default value for the `taxRoundingMode` parameter of the [CartDraft](ctp:api:type:CartDraft).
	// Indicates how monetary values are rounded when calculating taxes for `taxedPrice`.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
}

/**
*	Specifies the status of the [Customer Search](/../api/projects/customer-search) index.
*	You can change the status using the [Change Customer Search Status](ctp:api:type:ProjectChangeCustomerSearchStatusAction) update action.
*
 */
type CustomerSearchStatus string

const (
	CustomerSearchStatusActivated   CustomerSearchStatus = "Activated"
	CustomerSearchStatusDeactivated CustomerSearchStatus = "Deactivated"
)

/**
*	Represents a RFC 7662 compliant [OAuth 2.0 Token Introspection](https://datatracker.ietf.org/doc/html/rfc7662) endpoint. For more information, see [Requesting an access token using an external OAuth 2.0 server](/../api/authorization#request-an-access-token-using-an-external-oauth-server).
*
*	You can only configure **one** external OAuth 2.0 endpoint per Project. To authenticate using multiple external services (such as social network logins), use a middle layer authentication service.
*
 */
type ExternalOAuth struct {
	// URL with authorization header.
	Url string `json:"url"`
	// Must not contain any leading or trailing whitespaces. Partially hidden on retrieval.
	AuthorizationHeader string `json:"authorizationHeader"`
}

/**
*	Specifies the status of the [Order Search](/../api/projects/order-search) index.
 */
type OrderSearchStatus string

const (
	OrderSearchStatusActivated   OrderSearchStatus = "Activated"
	OrderSearchStatusDeactivated OrderSearchStatus = "Deactivated"
)

type ProductSearchIndexingMode string

const (
	ProductSearchIndexingModeProductProjectionsSearch ProductSearchIndexingMode = "ProductProjectionsSearch"
	ProductSearchIndexingModeProductsSearch           ProductSearchIndexingMode = "ProductsSearch"
)

type Project struct {
	// Current version of the Project.
	Version int `json:"version"`
	// User-defined unique identifier of the Project.
	Key string `json:"key"`
	// Name of the Project.
	Name string `json:"name"`
	// Country code of the geographic location.
	Countries []string `json:"countries"`
	// Currency code of the country. A Project must have at least one currency.
	Currencies []string `json:"currencies"`
	// Language of the country. A Project must have at least one language.
	Languages []string `json:"languages"`
	// Date and time (UTC) the Project was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date in YYYY-MM format specifying when the trial period for the Project ends. Only present on Projects in trial period.
	TrialUntil *string `json:"trialUntil,omitempty"`
	// Holds the configuration for the [Messages Query](/../api/projects/messages) feature.
	Messages MessagesConfiguration `json:"messages"`
	// Holds the configuration for the [Carts](/../api/projects/carts) feature.
	Carts CartsConfiguration `json:"carts"`
	// Holds the configuration for the [Shopping Lists](/../api/projects/shoppingLists) feature. This field may not be present on Projects created before January 2020.
	ShoppingLists *ShoppingListsConfiguration `json:"shoppingLists,omitempty"`
	// Holds the configuration for the [tiered shipping rates](ctp:api:type:ShippingRatePriceTier) feature.
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
	// Represents a RFC 7662 compliant [OAuth 2.0 Token Introspection](https://datatracker.ietf.org/doc/html/rfc7662) endpoint.
	ExternalOAuth *ExternalOAuth `json:"externalOAuth,omitempty"`
	// Controls indexing of resources to be provided on high performance read-only search endpoints.
	SearchIndexing *SearchIndexingConfiguration `json:"searchIndexing,omitempty"`
	// Holds configuration specific to [Business Units](ctp:api:type:BusinessUnit).
	BusinessUnits *BusinessUnitConfiguration `json:"businessUnits,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		var err error
		obj.ShippingRateInputType, err = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
		if err != nil {
			return err
		}
	}

	return nil
}

type ProjectUpdate struct {
	// Expected version of the Project on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Project.
	Actions []ProjectUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProjectUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProjectUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ProjectUpdateAction interface{}

func mapDiscriminatorProjectUpdateAction(input interface{}) (ProjectUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "changeBusinessUnitSearchStatus":
		obj := ProjectChangeBusinessUnitSearchStatusAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeMyBusinessUnitStatusOnCreation":
		obj := ProjectChangeBusinessUnitStatusOnCreationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCartsConfiguration":
		obj := ProjectChangeCartsConfigurationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCountries":
		obj := ProjectChangeCountriesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCountryTaxRateFallbackEnabled":
		obj := ProjectChangeCountryTaxRateFallbackEnabledAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCurrencies":
		obj := ProjectChangeCurrenciesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCustomerSearchStatus":
		obj := ProjectChangeCustomerSearchStatusAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLanguages":
		obj := ProjectChangeLanguagesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeMessagesConfiguration":
		obj := ProjectChangeMessagesConfigurationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeMessagesEnabled":
		obj := ProjectChangeMessagesEnabledAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProjectChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeOrderSearchStatus":
		obj := ProjectChangeOrderSearchStatusAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePriceRoundingMode":
		obj := ProjectChangePriceRoundingModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeProductSearchIndexingEnabled":
		obj := ProjectChangeProductSearchIndexingEnabledAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeShoppingListsConfiguration":
		obj := ProjectChangeShoppingListsConfigurationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxRoundingMode":
		obj := ProjectChangeTaxRoundingModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMyBusinessUnitAssociateRoleOnCreation":
		obj := ProjectSetBusinessUnitAssociateRoleOnCreationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExternalOAuth":
		obj := ProjectSetExternalOAuthAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingRateInputType":
		obj := ProjectSetShippingRateInputTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ShippingRateInputType != nil {
			var err error
			obj.ShippingRateInputType, err = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Controls indexing of resources to be provided on high performance read-only search endpoints.
*
 */
type SearchIndexingConfiguration struct {
	// Configuration for the [Product Projection Search](/../api/projects/product-projection-search) and [Search Term Suggestions](/../api/projects/search-term-suggestions) APIs.
	Products *SearchIndexingConfigurationValues `json:"products,omitempty"`
	// Configuration for the [Product Search](/../api/projects/product-search) feature.
	ProductsSearch *SearchIndexingConfigurationValues `json:"productsSearch,omitempty"`
	// Configuration for the [Order Search](/../api/projects/order-search) feature.
	Orders *SearchIndexingConfigurationValues `json:"orders,omitempty"`
	// Configuration for the [Customer Search](/../api/projects/customer-search) feature.
	Customers *SearchIndexingConfigurationValues `json:"customers,omitempty"`
	// Configuration for the [Business Unit Search](/../api/projects/business-unit-search) feature.
	BusinessUnits *SearchIndexingConfigurationValues `json:"businessUnits,omitempty"`
}

/**
*	Status of resource indexing.
 */
type SearchIndexingConfigurationStatus string

const (
	SearchIndexingConfigurationStatusActivated   SearchIndexingConfigurationStatus = "Activated"
	SearchIndexingConfigurationStatusDeactivated SearchIndexingConfigurationStatus = "Deactivated"
	SearchIndexingConfigurationStatusIndexing    SearchIndexingConfigurationStatus = "Indexing"
)

type SearchIndexingConfigurationValues struct {
	// Current status of resource indexing. Present on Projects from 1 February 2019.
	Status *SearchIndexingConfigurationStatus `json:"status,omitempty"`
	// Date and time (UTC) the Project was last updated. Only present on Projects last modified after 1 February 2019.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// IDs and references that last modified the SearchIndexingConfigurationValues.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
}

type ShippingRateInputType interface{}

func mapDiscriminatorShippingRateInputType(input interface{}) (ShippingRateInputType, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "CartClassification":
		obj := CartClassificationType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CartScore":
		obj := CartScoreType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CartValue":
		obj := CartValueType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	The [ShippingRate](ctp:api:type:ShippingRate) maps to an abstract Cart categorization expressed by strings (for example, `Light`, `Medium`, or `Heavy`).
*	Only keys defined in the `values` array can be used to create a tier or to set a value of the `shippingRateInput` on the [Cart](ctp:api:type:Cart).
*	Keys must be unique.
*
 */
type CartClassificationType struct {
	// The classification items that can be used for specifying any [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartClassificationType) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationType
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartClassification", Alias: (*Alias)(&obj)})
}

/**
*	The [ShippingRate](ctp:api:type:ShippingRate) maps to an abstract [Cart](ctp:api:type:Cart) categorization expressed by integers (such as shipping scores or weight ranges).
*
 */
type CartScoreType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartScoreType) MarshalJSON() ([]byte, error) {
	type Alias CartScoreType
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartScore", Alias: (*Alias)(&obj)})
}

/**
*	The [ShippingRate](ctp:api:type:ShippingRate) maps to the value of the Cart and is used to select a tier.
*	The value of the [Cart](ctp:api:type:Cart) is the sum of all Line Item totals and Custom Line Item totals (via the `totalPrice` field) after any Product Discounts and Cart Discounts have been applied.
*	If chosen, it is not possible to set a value for the `shippingRateInput` on the [Cart](ctp:api:type:Cart).
*
 */
type CartValueType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartValueType) MarshalJSON() ([]byte, error) {
	type Alias CartValueType
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartValue", Alias: (*Alias)(&obj)})
}

type ShoppingListsConfiguration struct {
	// Default value for the `deleteDaysAfterLastModification` parameter of the [ShoppingListDraft](ctp:api:type:ShoppingListDraft).
	// This field may not be present on Projects created before January 2020.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

type ProjectChangeBusinessUnitSearchStatusAction struct {
	// Activates or deactivates the [Search Business Units](ctp:api:endpoint:/{projectKey}/business-units/search:POST) feature. Activation will trigger building a search index for the Business Units in the Project.
	Status BusinessUnitSearchStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeBusinessUnitSearchStatusAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeBusinessUnitSearchStatusAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeBusinessUnitSearchStatus", Alias: (*Alias)(&obj)})
}

type ProjectChangeBusinessUnitStatusOnCreationAction struct {
	// Status for Business Units created using the [My Business Unit endpoint](ctp:api:endpoint:/{projectKey}/me/business-units:POST).
	Status BusinessUnitConfigurationStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeBusinessUnitStatusOnCreationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeBusinessUnitStatusOnCreationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMyBusinessUnitStatusOnCreation", Alias: (*Alias)(&obj)})
}

type ProjectChangeCartsConfigurationAction struct {
	// Configuration for the [Carts](/../api/projects/carts) feature.
	CartsConfiguration CartsConfiguration `json:"cartsConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeCartsConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCartsConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartsConfiguration", Alias: (*Alias)(&obj)})
}

type ProjectChangeCountriesAction struct {
	// New value to set. Must not be empty.
	Countries []string `json:"countries"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeCountriesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCountriesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCountries", Alias: (*Alias)(&obj)})
}

type ProjectChangeCountryTaxRateFallbackEnabledAction struct {
	// When `true`, country _- no state_ Tax Rate is used as fallback. See [CartsConfiguration](ctp:api:type:CartsConfiguration).
	CountryTaxRateFallbackEnabled bool `json:"countryTaxRateFallbackEnabled"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeCountryTaxRateFallbackEnabledAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCountryTaxRateFallbackEnabledAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCountryTaxRateFallbackEnabled", Alias: (*Alias)(&obj)})
}

type ProjectChangeCurrenciesAction struct {
	// New value to set. Must not be empty.
	Currencies []string `json:"currencies"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeCurrenciesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCurrenciesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCurrencies", Alias: (*Alias)(&obj)})
}

type ProjectChangeCustomerSearchStatusAction struct {
	// Activates or deactivates the [Customer Search](/../api/projects/customer-search) feature. Activation will trigger building a search index for the Customers in the Project.
	Status CustomerSearchStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeCustomerSearchStatusAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCustomerSearchStatusAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomerSearchStatus", Alias: (*Alias)(&obj)})
}

/**
*	Removing a language used by a [Store](ctp:api:type:Store) returns a [LanguageUsedInStores](ctp:api:type:LanguageUsedInStoresError) error.
*
 */
type ProjectChangeLanguagesAction struct {
	// New value to set. Must not be empty.
	Languages []string `json:"languages"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLanguages", Alias: (*Alias)(&obj)})
}

type ProjectChangeMessagesConfigurationAction struct {
	// Configuration for the [Messages Query](/../api/projects/messages) feature.
	MessagesConfiguration MessagesConfigurationDraft `json:"messagesConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeMessagesConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesConfiguration", Alias: (*Alias)(&obj)})
}

type ProjectChangeMessagesEnabledAction struct {
	MessagesEnabled bool `json:"messagesEnabled"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeMessagesEnabledAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesEnabledAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesEnabled", Alias: (*Alias)(&obj)})
}

type ProjectChangeNameAction struct {
	// New value to set. Must not be empty.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProjectChangeOrderSearchStatusAction struct {
	// Activates or deactivates the [Order Search](/../api/projects/order-search) feature. Activation will trigger building a search index for the Orders in the Project.
	Status OrderSearchStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeOrderSearchStatusAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeOrderSearchStatusAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderSearchStatus", Alias: (*Alias)(&obj)})
}

type ProjectChangePriceRoundingModeAction struct {
	// Project-level default rounding mode for calculating the total prices on [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem). See [CartsConfiguration](ctp:api:type:CartsConfiguration).
	PriceRoundingMode RoundingMode `json:"priceRoundingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangePriceRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangePriceRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePriceRoundingMode", Alias: (*Alias)(&obj)})
}

type ProjectChangeProductSearchIndexingEnabledAction struct {
	// - If `false`, the indexing of [Product](ctp:api:type:Product) information will stop and the [Product Projection Search](/../api/projects/product-projection-search) as well as the [Search Term Suggestions](/../api/projects/search-term-suggestions) API will no longer be available for this Project. The Project's [SearchIndexingConfiguration](ctp:api:type:SearchIndexingConfiguration) `status` for `products` will be changed to `"Deactivated"`.
	// - If `true`, the indexing of [Product](ctp:api:type:Product) information will start and the [Product Projection Search](/../api/projects/product-projection-search) as well as the [Search Term Suggestions](/../api/projects/search-term-suggestions) API will become available soon after for this Project. Proportional to the amount of information being indexed, the Project's [SearchIndexingConfiguration](ctp:api:type:SearchIndexingConfiguration) `status` for `products` will be shown as `"Indexing"` during this time. As soon as the indexing has finished, the configuration status will be changed to `"Activated"` making the aforementioned APIs fully available for this Project.
	Enabled bool `json:"enabled"`
	// Controls whether the action should apply to [Product Projection Search](/../api/projects/product-projection-search) or to [Product Search](/../api/projects/product-search).
	Mode *ProductSearchIndexingMode `json:"mode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeProductSearchIndexingEnabledAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeProductSearchIndexingEnabledAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeProductSearchIndexingEnabled", Alias: (*Alias)(&obj)})
}

type ProjectChangeShoppingListsConfigurationAction struct {
	// Configuration for the [Shopping Lists](/../api/projects/shoppingLists) feature.
	ShoppingListsConfiguration ShoppingListsConfiguration `json:"shoppingListsConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeShoppingListsConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeShoppingListsConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShoppingListsConfiguration", Alias: (*Alias)(&obj)})
}

type ProjectChangeTaxRoundingModeAction struct {
	// Project-level default rounding mode for tax calculation. See [CartsConfiguration](ctp:api:type:CartsConfiguration).
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectChangeTaxRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeTaxRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxRoundingMode", Alias: (*Alias)(&obj)})
}

type ProjectSetBusinessUnitAssociateRoleOnCreationAction struct {
	// Default [Associate Role](ctp:api:type:AssociateRole) assigned to the Associate creating a Business Unit using the [My Business Unit endpoint](ctp:api:endpoint:/{projectKey}/me/business-units:POST).
	AssociateRole AssociateRoleResourceIdentifier `json:"associateRole"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectSetBusinessUnitAssociateRoleOnCreationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetBusinessUnitAssociateRoleOnCreationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMyBusinessUnitAssociateRoleOnCreation", Alias: (*Alias)(&obj)})
}

type ProjectSetExternalOAuthAction struct {
	// Value to set. If empty, any existing value will be removed.
	ExternalOAuth *ExternalOAuth `json:"externalOAuth,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectSetExternalOAuthAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetExternalOAuthAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalOAuth", Alias: (*Alias)(&obj)})
}

type ProjectSetShippingRateInputTypeAction struct {
	// Value to set. If empty, any existing value will be removed.
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectSetShippingRateInputTypeAction) UnmarshalJSON(data []byte) error {
	type Alias ProjectSetShippingRateInputTypeAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		var err error
		obj.ShippingRateInputType, err = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProjectSetShippingRateInputTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetShippingRateInputTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInputType", Alias: (*Alias)(&obj)})
}
