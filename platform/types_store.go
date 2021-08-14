// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Store struct {
	Id             string          `json:"id"`
	Version        int             `json:"version"`
	CreatedAt      time.Time       `json:"createdAt"`
	LastModifiedAt time.Time       `json:"lastModifiedAt"`
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	CreatedBy      *CreatedBy      `json:"createdBy,omitEmpty"`
	// User-specific unique identifier for the store.
	// The `key` is mandatory and immutable.
	// It is used to reference the store.
	Key string `json:"key"`
	// The name of the store
	Name      *LocalizedString `json:"name,omitEmpty"`
	Languages []string         `json:"languages,omitEmpty"`
	// Set of References to a Channel with `ProductDistribution` role
	DistributionChannels []ChannelReference `json:"distributionChannels"`
	// Set of ResourceIdentifiers of Channels with `InventorySupply` role
	SupplyChannels []ChannelReference `json:"supplyChannels,omitEmpty"`
	Custom         *CustomFields      `json:"custom,omitEmpty"`
}

type StoreDraft struct {
	// User-specific unique identifier for the store.
	// The `key` is mandatory and immutable.
	// It is used to reference the store.
	Key string `json:"key"`
	// The name of the store
	Name      LocalizedString `json:"name"`
	Languages []string        `json:"languages,omitEmpty"`
	// Set of ResourceIdentifiers to a Channel with `ProductDistribution` role
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitEmpty"`
	// Set of ResourceIdentifiers of Channels with `InventorySupply` role
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels,omitEmpty"`
	Custom         *CustomFieldsDraft          `json:"custom,omitEmpty"`
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
	Total   int     `json:"total,omitEmpty"`
	Offset  int     `json:"offset"`
	Results []Store `json:"results"`
}

type StoreReference struct {
	Id  string `json:"id"`
	Obj *Store `json:"obj,omitEmpty"`
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
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
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
		new := StoreAddDistributionChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addSupplyChannel":
		new := StoreAddSupplyChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDistributionChannel":
		new := StoreRemoveDistributionChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeSupplyChannel":
		new := StoreRemoveSupplyChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := StoreSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := StoreSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDistributionChannels":
		new := StoreSetDistributionChannelsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLanguages":
		new := StoreSetLanguagesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := StoreSetNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSupplyChannels":
		new := StoreSetSupplyChannelsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
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
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel"`
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
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel"`
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
	Value interface{} `json:"value,omitEmpty"`
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
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type
	// Sets the custom field to this value.
	Fields *interface{} `json:"fields,omitEmpty"`
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
	DistributionChannels []ChannelResourceIdentifier `json:"distributionChannels,omitEmpty"`
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
	Languages []string `json:"languages,omitEmpty"`
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
	Name *LocalizedString `json:"name,omitEmpty"`
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
	SupplyChannels []ChannelResourceIdentifier `json:"supplyChannels,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj StoreSetSupplyChannelsAction) MarshalJSON() ([]byte, error) {
	type Alias StoreSetSupplyChannelsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannels", Alias: (*Alias)(&obj)})
}
