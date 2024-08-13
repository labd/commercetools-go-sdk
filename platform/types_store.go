package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ProductSelectionSetting struct {
	// Reference to a ProductSelection.
	ProductSelection ProductSelectionReference `json:"productSelection"`
	// If `true`, all Products assigned to this Product Selection are part of the Store's assortment.
	Active bool `json:"active"`
}

type ProductSelectionSettingDraft struct {
	// Resource Identifier of a ProductSelection.
	ProductSelection ProductSelectionResourceIdentifier `json:"productSelection"`
	// Set to `true` if all Products assigned to the Product Selection should become part of the Store's assortment.
	Active *bool `json:"active,omitempty"`
}

type Store struct {
	// Unique ID of the Store.
	ID string `json:"id"`
	// Current version of the Store.
	Version int `json:"version"`
	// Date and time (UTC) the Store was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Store was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the Store.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the Store.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique and immutable identifier for the Store.
	Key string `json:"key"`
	// Name of the Store.
	Name *LocalizedString `json:"name,omitempty"`
	// Languages configured for the Store.
	Languages []string `json:"languages"`
	// Countries defined for the Store.
	Countries []StoreCountry `json:"countries"`
	// Product Distribution Channels allowed for the Store.
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	// Inventory Supply Channels allowed for the Store.
	SupplyChannels []ChannelReference `json:"supplyChannels"`
	// Controls availability of Products for this Store via Product Selections:
	//
	// - Leave empty if all Products in the [Project](ctp:api:type:Project) should be available in this Store.
	// - If only `inactive` Product Selections with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode) are provided, all the Products are availlable in this Store.
	// - If all the Product Selections provided are `inactive` and there's at least a Product Selection of mode `Individual`, no Product is availlable in this Store.
	// - If at least an `active` Product Selection is provided, only `active` Product Selections are considered to compute the availlability in this Store.
	ProductSelections []ProductSelectionSetting `json:"productSelections"`
	// Custom fields for the Store.
	Custom *CustomFields `json:"custom,omitempty"`
}

type StoreDraft struct {
	// User-defined unique and immutable identifier for the Store.
	// Keys can only contain alphanumeric characters, underscores, and hyphens.
	Key string `json:"key"`
	// Name of the Store.
	Name *LocalizedString `json:"name,omitempty"`
	// Languages defined in [Project](ctp:api:type:Project). Only languages defined in the Project can be used.
	Languages []string `json:"languages"`
	// Countries defined for the Store.
	Countries []StoreCountry `json:"countries"`
	// ResourceIdentifier of a Channel with `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels"`
	// ResourceIdentifier of a Channel with `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels"`
	// Controls availability of Products for this Store via active/inactive Product Selections:
	//
	// - Leave empty if all Products in the [Project](ctp:api:type:Project) should be available in this Store.
	// - If only `inactive` Product Selections with `IndividualExclusion` [ProductSelectionMode](ctp:api:type:ProductSelectionMode) are provided, all the Products are available in this Store.
	// - If all the Product Selections provided are `inactive` and there's at least a Product Selection of mode `Individual`, no Product is available in this Store.
	// - If at least an `active` Product Selection is provided, only `active` Product Selections are considered to compute the availability in this Store.
	ProductSelections []ProductSelectionSettingDraft `json:"productSelections"`
	// Custom fields for the Store.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreDraft) MarshalJSON() ([]byte, error) {
	type Alias StoreDraft
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

	if raw["languages"] == nil {
		delete(raw, "languages")
	}

	if raw["countries"] == nil {
		delete(raw, "countries")
	}

	if raw["distributionChannels"] == nil {
		delete(raw, "distributionChannels")
	}

	if raw["supplyChannels"] == nil {
		delete(raw, "supplyChannels")
	}

	if raw["productSelections"] == nil {
		delete(raw, "productSelections")
	}

	return json.Marshal(raw)

}

/**
*	[KeyReference](ctp:api:type:KeyReference) to a [Store](ctp:api:type:Store).
*
 */
type StoreKeyReference struct {
	// Unique and immutable key of the referenced [Store](ctp:api:type:Store).
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Store](ctp:api:type:Store).
*
 */
type StorePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Stores](ctp:api:type:Store) matching the query.
	Results []Store `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Store](ctp:api:type:Store).
*
 */
type StoreReference struct {
	// Unique ID of the referenced [Store](ctp:api:type:Store).
	ID string `json:"id"`
	// Contains the representation of the expanded Store. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Stores.
	Obj *Store `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreReference) MarshalJSON() ([]byte, error) {
	type Alias StoreReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Store](ctp:api:type:Store). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type StoreResourceIdentifier struct {
	// Unique ID of the referenced [Store](ctp:api:type:Store). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced [Store](ctp:api:type:Store). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StoreResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

type StoreUpdate struct {
	// Expected version of the Store on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Store.
	Actions []StoreUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreUpdate) UnmarshalJSON(data []byte) error {
	type Alias StoreUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStoreUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type StoreUpdateAction interface{}

func mapDiscriminatorStoreUpdateAction(input interface{}) (StoreUpdateAction, error) {
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
	case "addCountry":
		obj := StoreAddCountryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addDistributionChannel":
		obj := StoreAddDistributionChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addProductSelection":
		obj := StoreAddProductSelectionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addSupplyChannel":
		obj := StoreAddSupplyChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeProductSelectionActive":
		obj := StoreChangeProductSelectionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeCountry":
		obj := StoreRemoveCountryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeDistributionChannel":
		obj := StoreRemoveDistributionChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeProductSelection":
		obj := StoreRemoveProductSelectionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeSupplyChannel":
		obj := StoreRemoveSupplyChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCountries":
		obj := StoreSetCountriesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := StoreSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := StoreSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDistributionChannels":
		obj := StoreSetDistributionChannelsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLanguages":
		obj := StoreSetLanguagesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := StoreSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductSelections":
		obj := StoreSetProductSelectionsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSupplyChannels":
		obj := StoreSetSupplyChannelsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	This update action produces the [StoreCountriesChanged](ctp:api:type:StoreCountriesChangedMessage) Message.
*	It has no effect if the given country is already present in a Store.
*
 */
type StoreAddCountryAction struct {
	// Value to append to `countries`.
	Country StoreCountry `json:"country"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreAddCountryAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCountry", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [StoreDistributionChannelsChanged](ctp:api:type:StoreDistributionChannelsChangedMessage) Message.
*	It has no effect if a given distribution channel is already present in a Store.
*
*	Adding a [Channel](ctp:api:type:Channel) without the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) returns a [MissingRoleOnChannel](ctp:api:type:MissingRoleOnChannelError) error.
*
 */
type StoreAddDistributionChannelAction struct {
	// Value to append.
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreAddDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDistributionChannel", Alias: (*Alias)(&obj)})
}

/**
*	To make all included Products available to your customers of a given Store, add the [Product Selections](/../api/projects/product-selections) to the respective Store. This action has no effect if the given Product Selection is already present in the Store and has the same `active` flag.
*
 */
type StoreAddProductSelectionAction struct {
	// Product Selection to add to the Store either activated or deactivated.
	ProductSelection ProductSelectionResourceIdentifier `json:"productSelection"`
	// Set to `true` to make all Products assigned to the referenced Product Selection available in the Store.
	Active *bool `json:"active,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreAddProductSelectionAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddProductSelectionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addProductSelection", Alias: (*Alias)(&obj)})
}

/**
*	This action has no effect if a given supply channel is already present in a Store.
*
*	Adding a supply channel produces the [StoreSupplyChannelsChanged](ctp:api:type:StoreSupplyChannelsChangedMessage) Message.
*
*	Adding a [Channel](ctp:api:type:Channel) without the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) returns a [MissingRoleOnChannel](ctp:api:type:MissingRoleOnChannelError) error.
*
 */
type StoreAddSupplyChannelAction struct {
	// Value to append.
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreAddSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addSupplyChannel", Alias: (*Alias)(&obj)})
}

/**
*	[ProductSelection](ctp:api:type:ProductSelection) in a Store can be activated or deactivated using this update action.
*
 */
type StoreChangeProductSelectionAction struct {
	// Current Product Selection of the Store to be activated or deactivated.
	ProductSelection ProductSelectionResourceIdentifier `json:"productSelection"`
	// Set to `true` if all Products assigned to the Product Selection should become part of the Store's assortment.
	Active *bool `json:"active,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreChangeProductSelectionAction) MarshalJSON() ([]byte, error) {
	type Alias StoreChangeProductSelectionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeProductSelectionActive", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [StoreCountriesChanged](ctp:api:type:StoreCountriesChangedMessage) Message.
*	It has no effect if a given country is not present in a Store.
*
 */
type StoreRemoveCountryAction struct {
	// Value to remove from `countries`.
	Country StoreCountry `json:"country"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreRemoveCountryAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCountry", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [StoreDistributionChannelsChanged](ctp:api:type:StoreDistributionChannelsChangedMessage) Message.
*
 */
type StoreRemoveDistributionChannelAction struct {
	// Value to remove. ResourceIdentifier of a Channel with the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreRemoveDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDistributionChannel", Alias: (*Alias)(&obj)})
}

/**
*	This action has no effect if the given Product Selection is not in the Store.
*
 */
type StoreRemoveProductSelectionAction struct {
	// Value to remove. The removed Product Selection is made offline.
	ProductSelection ProductSelectionResourceIdentifier `json:"productSelection"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreRemoveProductSelectionAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveProductSelectionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeProductSelection", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [StoreSupplyChannelsChanged](ctp:api:type:StoreSupplyChannelsChangedMessage) Message.
*
 */
type StoreRemoveSupplyChannelAction struct {
	// Value to remove. ResourceIdentifier of a Channel with the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreRemoveSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeSupplyChannel", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [StoreCountriesChanged](ctp:api:type:StoreCountriesChangedMessage) Message.
*
 */
type StoreSetCountriesAction struct {
	// New value to set.
	Countries []StoreCountry `json:"countries"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetCountriesAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCountriesAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountries", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["countries"] == nil {
		delete(raw, "countries")
	}

	return json.Marshal(raw)

}

type StoreSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type StoreSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Store with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Store.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Store.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [StoreDistributionChannelsChanged](ctp:api:type:StoreDistributionChannelsChangedMessage) Message.
*
*	Setting a [Channel](ctp:api:type:Channel) without the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) returns a [MissingRoleOnChannel](ctp:api:type:MissingRoleOnChannelError) error.
*
 */
type StoreSetDistributionChannelsAction struct {
	// Value to set.
	// If not defined, the Store's `distributionChannels` are unset.
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetDistributionChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetDistributionChannelsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDistributionChannels", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["distributionChannels"] == nil {
		delete(raw, "distributionChannels")
	}

	return json.Marshal(raw)

}

/**
*	This update action produces the [StoreLanguagesChanged](ctp:api:type:StoreLanguagesChangedMessage) Message.
*	Adding a language other than the ones defined in the [Project](ctp:api:type:Project) returns a [ProjectNotConfiguredForLanguages](ctp:api:type:ProjectNotConfiguredForLanguagesError) error.
*
 */
type StoreSetLanguagesAction struct {
	// Value to set.
	Languages []string `json:"languages"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetLanguagesAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLanguages", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["languages"] == nil {
		delete(raw, "languages")
	}

	return json.Marshal(raw)

}

/**
*	This update action produces the [StoreNameSet](ctp:api:type:StoreNameSetMessage) Message.
*
 */
type StoreSetNameAction struct {
	// Value to set.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Instead of adding or removing [Product Selections](/../api/projects/product-selections) individually, you can also change all the Store's Product Selections in one go using this update action. The Store will only contain the Product Selections specified in the request.
*
 */
type StoreSetProductSelectionsAction struct {
	// Value to set.
	//
	// - If provided, Product Selections for which `active` is set to `true` are available in the Store.
	// - If not provided or provided as empty array, the action removes all Product Selections from this Store, meaning all Products in the [Project](ctp:api:type:Project) are available in this Store.
	ProductSelections []ProductSelectionSettingDraft `json:"productSelections"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetProductSelectionsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetProductSelectionsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductSelections", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["productSelections"] == nil {
		delete(raw, "productSelections")
	}

	return json.Marshal(raw)

}

/**
*	Setting a supply channel produces the [StoreSupplyChannelsChanged](ctp:api:type:StoreSupplyChannelsChangedMessage) Message.
*
*	Setting a [Channel](ctp:api:type:Channel) without the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) returns a [MissingRoleOnChannel](ctp:api:type:MissingRoleOnChannelError) error.
*
 */
type StoreSetSupplyChannelsAction struct {
	// Value to set.
	// If not defined, the Store's `supplyChannels` are unset.
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetSupplyChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetSupplyChannelsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannels", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["supplyChannels"] == nil {
		delete(raw, "supplyChannels")
	}

	return json.Marshal(raw)

}
