package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ProductSelectionSetting struct {
	// Reference to a Product Selection
	ProductSelection ProductSelectionReference `json:"productSelection"`
	// If `true` all Products assigned to this Product Selection are part of the Store's assortment.
	Active bool `json:"active"`
}

type ProductSelectionSettingDraft struct {
	// Resource Identifier of a Product Selection
	ProductSelection ProductSelectionResourceIdentifier `json:"productSelection"`
	// If `true` all Products assigned to this Product Selection become part of the Store's assortment.
	Active *bool `json:"active,omitempty"`
}

type Store struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-specific unique identifier for the store.
	// The `key` is mandatory and immutable.
	// It is used to reference the store.
	Key string `json:"key"`
	// The name of the store
	Name      *LocalizedString `json:"name,omitempty"`
	Languages []string         `json:"languages"`
	// Set of References to a Channel with `ProductDistribution` role
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	// Set of ResourceIdentifiers of Channels with `InventorySupply` role
	SupplyChannels []ChannelReference `json:"supplyChannels"`
	// Set of References to Product Selections along with settings.
	// If `productSelections` is empty all products in the project are available in this Store.
	// If `productSelections` is not empty but there exists no `active` Product Selection then no Product is available in this Store.
	ProductSelections []ProductSelectionSetting `json:"productSelections"`
	Custom            *CustomFields             `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Store) MarshalJSON() ([]byte, error) {
	type Alias Store
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

	if target["languages"] == nil {
		delete(target, "languages")
	}

	if target["supplyChannels"] == nil {
		delete(target, "supplyChannels")
	}

	if target["productSelections"] == nil {
		delete(target, "productSelections")
	}

	return json.Marshal(target)
}

type StoreDraft struct {
	// User-specific unique identifier for the store.
	// The `key` is mandatory and immutable.
	// It is used to reference the store.
	Key string `json:"key"`
	// The name of the store
	Name      *LocalizedString `json:"name,omitempty"`
	Languages []string         `json:"languages"`
	// Set of ResourceIdentifiers to a Channel with `ProductDistribution` role
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels"`
	// Set of ResourceIdentifiers of Channels with `InventorySupply` role
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels"`
	// Set of ResourceIdentifiers of Product Selections along with settings.
	// If `productSelections` is empty all products in the project are available in this Store.
	// If `productSelections` is not empty but there exists no `active` Product Selection then no Product is available in this Store.
	ProductSelections []ProductSelectionSettingDraft `json:"productSelections"`
	Custom            *CustomFieldsDraft             `json:"custom,omitempty"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["languages"] == nil {
		delete(target, "languages")
	}

	if target["distributionChannels"] == nil {
		delete(target, "distributionChannels")
	}

	if target["supplyChannels"] == nil {
		delete(target, "supplyChannels")
	}

	if target["productSelections"] == nil {
		delete(target, "productSelections")
	}

	return json.Marshal(target)
}

type StoreKeyReference struct {
	// User-defined unique and immutable key of the referenced resource.
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

type StorePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int     `json:"offset"`
	Results []Store `json:"results"`
}

type StoreReference struct {
	// Unique ID of the referenced resource.
	ID  string `json:"id"`
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

type StoreResourceIdentifier struct {
	// Platform-generated unique identifier of the referenced resource. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource. Required if `id` is absent.
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
	Version int                 `json:"version"`
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
		if obj.ProductSelection != nil {
			var err error
			obj.ProductSelection, err = mapDiscriminatorResourceIdentifier(obj.ProductSelection)
			if err != nil {
				return nil, err
			}
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
		if obj.ProductSelection != nil {
			var err error
			obj.ProductSelection, err = mapDiscriminatorResourceIdentifier(obj.ProductSelection)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "removeSupplyChannel":
		obj := StoreRemoveSupplyChannelAction{}
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

type StoreAddDistributionChannelAction struct {
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

type StoreAddProductSelectionAction struct {
	// Resource Identifier of a Product Selection
	ProductSelection ProductSelectionResourceIdentifier `json:"productSelection"`
	// If `true` all Products assigned to this Product Selection become part of the Store's assortment.
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

type StoreAddSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
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

type StoreChangeProductSelectionAction struct {
	// A current Product Selection of this Store that is to be activated or deactivated.
	ProductSelection ResourceIdentifier `json:"productSelection"`
	// If `true` all Products assigned to the Product Selection become part of the Store's assortment.
	Active *bool `json:"active,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreChangeProductSelectionAction) UnmarshalJSON(data []byte) error {
	type Alias StoreChangeProductSelectionAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ProductSelection != nil {
		var err error
		obj.ProductSelection, err = mapDiscriminatorResourceIdentifier(obj.ProductSelection)
		if err != nil {
			return err
		}
	}
	return nil
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

type StoreRemoveDistributionChannelAction struct {
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

type StoreRemoveProductSelectionAction struct {
	// A Product Selection to be removed from the current Product Selections of this Store.
	ProductSelection ResourceIdentifier `json:"productSelection"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StoreRemoveProductSelectionAction) UnmarshalJSON(data []byte) error {
	type Alias StoreRemoveProductSelectionAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ProductSelection != nil {
		var err error
		obj.ProductSelection, err = mapDiscriminatorResourceIdentifier(obj.ProductSelection)
		if err != nil {
			return err
		}
	}
	return nil
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

type StoreRemoveSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
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

type StoreSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
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

type StoreSetDistributionChannelsAction struct {
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["distributionChannels"] == nil {
		delete(target, "distributionChannels")
	}

	return json.Marshal(target)
}

type StoreSetLanguagesAction struct {
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["languages"] == nil {
		delete(target, "languages")
	}

	return json.Marshal(target)
}

type StoreSetNameAction struct {
	// The updated name of the store
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

type StoreSetProductSelectionsAction struct {
	// The total of Product Selections to be set for this Store.
	ProductSelections []ProductSelectionSettingDraft `json:"productSelections"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StoreSetProductSelectionsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetProductSelectionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductSelections", Alias: (*Alias)(&obj)})
}

type StoreSetSupplyChannelsAction struct {
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["supplyChannels"] == nil {
		delete(target, "supplyChannels")
	}

	return json.Marshal(target)
}
