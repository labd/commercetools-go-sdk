// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

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
	Languages []string         `json:"languages,omitempty"`
	// Set of References to a Channel with `ProductDistribution` role
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	// Set of ResourceIdentifiers of Channels with `InventorySupply` role
	SupplyChannels []ChannelReference `json:"supplyChannels,omitempty"`
	Custom         *CustomFields      `json:"custom,omitempty"`
}

type StoreDraft struct {
	// User-specific unique identifier for the store.
	// The `key` is mandatory and immutable.
	// It is used to reference the store.
	Key string `json:"key"`
	// The name of the store
	Name      *LocalizedString `json:"name,omitempty"`
	Languages []string         `json:"languages,omitempty"`
	// Set of ResourceIdentifiers to a Channel with `ProductDistribution` role
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitempty"`
	// Set of ResourceIdentifiers of Channels with `InventorySupply` role
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels,omitempty"`
	Custom         *CustomFieldsDraft          `json:"custom,omitempty"`
}

type StoreKeyReference struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreKeyReference) MarshalJSON() ([]byte, error) {
	type Alias StoreKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

type StorePagedQueryResponse struct {
	Limit   int     `json:"limit"`
	Count   int     `json:"count"`
	Total   *int    `json:"total,omitempty"`
	Offset  int     `json:"offset"`
	Results []Store `json:"results"`
}

type StoreReference struct {
	// Unique ID of the referenced resource.
	ID  string `json:"id"`
	Obj *Store `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreReference) MarshalJSON() ([]byte, error) {
	type Alias StoreReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "store", Alias: (*Alias)(&obj)})
}

type StoreResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
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

	return nil
}

type StoreUpdateAction interface{}

func mapDiscriminatorStoreUpdateAction(input interface{}) (StoreUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "addDistributionChannel":
		obj := StoreAddDistributionChannelAction{}
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
	case "removeDistributionChannel":
		obj := StoreRemoveDistributionChannelAction{}
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

// MarshalJSON override to set the discriminator value
func (obj StoreAddDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDistributionChannel", Alias: (*Alias)(&obj)})
}

type StoreAddSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreAddSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreAddSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addSupplyChannel", Alias: (*Alias)(&obj)})
}

type StoreRemoveDistributionChannelAction struct {
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreRemoveDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDistributionChannel", Alias: (*Alias)(&obj)})
}

type StoreRemoveSupplyChannelAction struct {
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreRemoveSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StoreRemoveSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeSupplyChannel", Alias: (*Alias)(&obj)})
}

type StoreSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type StoreSetCustomTypeAction struct {
	// If set, the custom type is reset to this value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// A valid JSON object, based on the FieldDefinitions of the Type
	// Sets the custom field to this value.
	Fields *interface{} `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type StoreSetDistributionChannelsAction struct {
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetDistributionChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetDistributionChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDistributionChannels", Alias: (*Alias)(&obj)})
}

type StoreSetLanguagesAction struct {
	Languages []string `json:"languages,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLanguages", Alias: (*Alias)(&obj)})
}

type StoreSetNameAction struct {
	// The updated name of the store
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type StoreSetSupplyChannelsAction struct {
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetSupplyChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetSupplyChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannels", Alias: (*Alias)(&obj)})
}
