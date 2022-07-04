package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type CartsConfiguration struct {
	// Default value for the `deleteDaysAfterLastModification` parameter of the [CartDraft](ctp:api:type:CartDraft). This field may not be present on Projects created before January 2020.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Indicates if country _- no state_ Tax Rate fallback should be used when a shipping address state is not explicitly covered in the rates lists of all Tax Categories of a Cart Line Items. This field may not be present on Projects created before June 2020.
	CountryTaxRateFallbackEnabled *bool `json:"countryTaxRateFallbackEnabled,omitempty"`
}

/**
*	Represents a RFC 7662 compliant [OAuth 2.0 Token Introspection](https://datatracker.ietf.org/doc/html/rfc7662) endpoint. For more information, see [Requesting an access token using an external OAuth 2.0 server](/../api/authorization#requesting-an-access-token-using-an-external-oauth-server).
*
*	You can only configure **one** external OAuth 2.0 endpoint per Project. To authenticate using multiple external services (such as social network logins), use a middle layer authentication service.
*
 */
type ExternalOAuth struct {
	// URL with authorization header.
	Url string `json:"url"`
	// Partially hidden on retrieval.
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
	// Expected version of the Project on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
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
	// Configuration for the [Product Projection Search](/../api/projects/products-search) and [Product Suggestions](/../api/projects/products-suggestions) endpoints.
	Products *SearchIndexingConfigurationValues `json:"products,omitempty"`
	// Configuration for the [Order Search](/../api/projects/order-search) feature.
	Orders *SearchIndexingConfigurationValues `json:"orders,omitempty"`
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
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
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
*	Used when the ShippingRate maps to an abstract Cart categorization expressed by strings (for example, `Light`, `Medium`, or `Heavy`).
*	Only keys defined in the `values` array can be used to create a tier or to set a value of the `shippingRateInput` on the [Cart](ctp:api:type:Cart).
*	Keys must be unique.
*
 */
type CartClassificationType struct {
	// The classification items that can be used for specifiying any [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
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
*	Used when the ShippingRate maps to an abstract Cart categorization expressed by integers (such as shipping scores or weight ranges).
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
*	Used when the ShippingRate maps to the sum of [LineItem](ctp:api:type:LineItem) Prices.
*	The value of the Cart is used to select a tier.
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

/**
*	If a language is used by a [Store](ctp:api:type:Store), it cannot be deleted. Attempts to delete such language will lead to [LanguageUsedInStores](/../api/errors#projects-400-language-used-in-stores) errors.
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

type ProjectChangeProductSearchIndexingEnabledAction struct {
	// If `false`, the indexing of [Product](ctp:api:type:Product) information will stop and the [Product Projection Search](/../api/projects/products-search) as well as the [Product Suggestions](/../api/projects/products-suggestions) endpoint will not be available anymore for this Project. The Project's [SearchIndexingConfiguration](ctp:api:type:SearchIndexingConfiguration) `status` for `products` will be changed to `"Deactivated"`.
	//
	// If `true`, the indexing of [Product](ctp:api:type:Product) information will start and the [Product Projection Search](/../api/projects/products-search) as well as the [Product Suggestions](/../api/projects/products-suggestions) endpoint will become available soon after for this Project. Proportional to the amount of information being indexed, the Project's [SearchIndexingConfiguration](ctp:api:type:SearchIndexingConfiguration) `status` for `products` will be shown as `"Indexing"` during this time. As soon as the indexing has finished, the configuration status will be changed to `"Activated"` making the aforementioned endpoints fully available for this Project.
	Enabled bool `json:"enabled"`
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
