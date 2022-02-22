// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type CartsConfiguration struct {
	// if country - no state tax rate fallback should be used when a shipping address state is not explicitly covered in the rates lists of all tax categories of a cart line items. Default value 'false'
	CountryTaxRateFallbackEnabled *bool `json:"countryTaxRateFallbackEnabled,omitempty"`
	// The default value for the deleteDaysAfterLastModification parameter of the CartDraft. Initially set to 90 for projects created after December 2019.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

type ExternalOAuth struct {
	Url                 string `json:"url"`
	AuthorizationHeader string `json:"authorizationHeader"`
}

/**
*	Activated indicates that the Order Search feature is active. Deactivated means that the namely feature is currently configured to be inactive.
 */
type OrderSearchStatus string

const (
	OrderSearchStatusActivated   OrderSearchStatus = "Activated"
	OrderSearchStatusDeactivated OrderSearchStatus = "Deactivated"
)

type Project struct {
	// The current version of the project.
	Version int `json:"version"`
	// The unique key of the project.
	Key string `json:"key"`
	// The name of the project.
	Name string `json:"name"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Countries []string `json:"countries"`
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currencies []string  `json:"currencies"`
	Languages  []string  `json:"languages"`
	CreatedAt  time.Time `json:"createdAt"`
	// The time is in the format Year-Month `YYYY-MM`.
	TrialUntil            *string                      `json:"trialUntil,omitempty"`
	Messages              MessagesConfiguration        `json:"messages"`
	ShippingRateInputType ShippingRateInputType        `json:"shippingRateInputType,omitempty"`
	ExternalOAuth         *ExternalOAuth               `json:"externalOAuth,omitempty"`
	Carts                 CartsConfiguration           `json:"carts"`
	SearchIndexing        *SearchIndexingConfiguration `json:"searchIndexing,omitempty"`
	ShoppingLists         *ShoppingListsConfiguration  `json:"shoppingLists,omitempty"`
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
	Version int                   `json:"version"`
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

type SearchIndexingConfiguration struct {
	// Configuration for endpoints serving indexed [Product](ctp:api:type:Product) information.
	Products *SearchIndexingConfigurationValues `json:"products,omitempty"`
	// Configuration for the [Order Search](/../api/projects/order-search) feature.
	Orders *SearchIndexingConfigurationValues `json:"orders,omitempty"`
}

/**
*	Can be one of the following or absent. "Activated" or absent means that the search and suggest endpoints for the specified resource type are active. "Deactivated" means that the search and suggest endpoints for the specified resource type cannot be used. "Indexing" indicates that the search and suggest endpoints can _temporally_ not be used because the search index is being re-built.
 */
type SearchIndexingConfigurationStatus string

const (
	SearchIndexingConfigurationStatusActivated   SearchIndexingConfigurationStatus = "Activated"
	SearchIndexingConfigurationStatusDeactivated SearchIndexingConfigurationStatus = "Deactivated"
	SearchIndexingConfigurationStatusIndexing    SearchIndexingConfigurationStatus = "Indexing"
)

type SearchIndexingConfigurationValues struct {
	// Can be one of the following or absent. "Activated" or absent means that the search and suggest endpoints for the specified resource type are active. "Deactivated" means that the search and suggest endpoints for the specified resource type cannot be used. "Indexing" indicates that the search and suggest endpoints can _temporally_ not be used because the search index is being re-built.
	Status         *SearchIndexingConfigurationStatus `json:"status,omitempty"`
	LastModifiedAt *time.Time                         `json:"lastModifiedAt,omitempty"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
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

type CartClassificationType struct {
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
	// The default value for the deleteDaysAfterLastModification parameter of the ShoppingListDraft. Initially set to 360 for projects created after December 2019.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

type ProjectChangeCartsConfigurationAction struct {
	CartsConfiguration *CartsConfiguration `json:"cartsConfiguration,omitempty"`
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
	// A two-digit country code as per country code.
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
	// default value is `false`
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
	// A three-digit currency code as per currency code.
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

type ProjectChangeLanguagesAction struct {
	//  .
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
	// Activated indicates that the Order Search feature is active. Deactivated means that the namely feature is currently configured to be inactive.
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
	ShoppingListsConfiguration *ShoppingListsConfiguration `json:"shoppingListsConfiguration,omitempty"`
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
	// If you do not provide the `externalOAuth` field or provide a value
	// of `null`, the update action unsets the External OAuth provider.
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
	// If not set, removes existing shippingRateInputType.
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
