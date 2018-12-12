// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ShippingRateTierType is an enum type
type ShippingRateTierType string

// Enum values for ShippingRateTierType
const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

// ShippingMethodUpdateAction uses action as discriminator attribute
type ShippingMethodUpdateAction interface{}

func mapDiscriminatorShippingMethodUpdateAction(input interface{}) ShippingMethodUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addShippingRate":
		new := ShippingMethodAddShippingRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addZone":
		new := ShippingMethodAddZoneAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeIsDefault":
		new := ShippingMethodChangeIsDefaultAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ShippingMethodChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxCategory":
		new := ShippingMethodChangeTaxCategoryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeShippingRate":
		new := ShippingMethodRemoveShippingRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeZone":
		new := ShippingMethodRemoveZoneAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := ShippingMethodSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ShippingMethodSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setPredicate":
		new := ShippingMethodSetPredicateAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// ShippingRatePriceTier uses type as discriminator attribute
type ShippingRatePriceTier interface{}

func mapDiscriminatorShippingRatePriceTier(input interface{}) ShippingRatePriceTier {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "CartClassification":
		new := CartClassificationTier{}
		mapstructure.Decode(input, &new)
		return new
	case "CartScore":
		new := CartScoreTier{}
		mapstructure.Decode(input, &new)
		return new
	case "CartValue":
		new := CartValueTier{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// CartClassificationTier implements the interface ShippingRatePriceTier
type CartClassificationTier struct {
	Value      string `json:"value"`
	Price      *Money `json:"price"`
	IsMatching bool   `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartClassification", Alias: (*Alias)(&obj)})
}

// CartScoreTier implements the interface ShippingRatePriceTier
type CartScoreTier struct {
	Score         float64        `json:"score"`
	PriceFunction *PriceFunction `json:"priceFunction,omitempty"`
	Price         *Money         `json:"price,omitempty"`
	IsMatching    bool           `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartScoreTier) MarshalJSON() ([]byte, error) {
	type Alias CartScoreTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartScore", Alias: (*Alias)(&obj)})
}

// CartValueTier implements the interface ShippingRatePriceTier
type CartValueTier struct {
	Price             *Money `json:"price"`
	MinimumCentAmount int    `json:"minimumCentAmount"`
	IsMatching        bool   `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartValueTier) MarshalJSON() ([]byte, error) {
	type Alias CartValueTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartValue", Alias: (*Alias)(&obj)})
}

// PriceFunction is a standalone struct
type PriceFunction struct {
	Function     string       `json:"function"`
	CurrencyCode CurrencyCode `json:"currencyCode"`
}

// ShippingMethod is of type Resource
type ShippingMethod struct {
	Version        int                   `json:"version"`
	LastModifiedAt time.Time             `json:"lastModifiedAt"`
	ID             string                `json:"id"`
	CreatedAt      time.Time             `json:"createdAt"`
	ZoneRates      []ZoneRate            `json:"zoneRates"`
	TaxCategory    *TaxCategoryReference `json:"taxCategory"`
	Predicate      string                `json:"predicate,omitempty"`
	Name           string                `json:"name"`
	Key            string                `json:"key,omitempty"`
	IsDefault      bool                  `json:"isDefault"`
	Description    string                `json:"description,omitempty"`
}

// ShippingMethodAddShippingRateAction implements the interface ShippingMethodUpdateAction
type ShippingMethodAddShippingRateAction struct {
	Zone         *ZoneReference     `json:"zone"`
	ShippingRate *ShippingRateDraft `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodAddShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingRate", Alias: (*Alias)(&obj)})
}

// ShippingMethodAddZoneAction implements the interface ShippingMethodUpdateAction
type ShippingMethodAddZoneAction struct {
	Zone *ZoneReference `json:"zone"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodAddZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addZone", Alias: (*Alias)(&obj)})
}

// ShippingMethodChangeIsDefaultAction implements the interface ShippingMethodUpdateAction
type ShippingMethodChangeIsDefaultAction struct {
	IsDefault bool `json:"isDefault"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeIsDefaultAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeIsDefaultAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsDefault", Alias: (*Alias)(&obj)})
}

// ShippingMethodChangeNameAction implements the interface ShippingMethodUpdateAction
type ShippingMethodChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ShippingMethodChangeTaxCategoryAction implements the interface ShippingMethodUpdateAction
type ShippingMethodChangeTaxCategoryAction struct {
	TaxCategory *TaxCategoryReference `json:"taxCategory"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCategory", Alias: (*Alias)(&obj)})
}

// ShippingMethodDraft is a standalone struct
type ShippingMethodDraft struct {
	ZoneRates   []ZoneRateDraft       `json:"zoneRates"`
	TaxCategory *TaxCategoryReference `json:"taxCategory"`
	Predicate   string                `json:"predicate,omitempty"`
	Name        string                `json:"name"`
	Key         string                `json:"key,omitempty"`
	IsDefault   bool                  `json:"isDefault"`
	Description string                `json:"description,omitempty"`
}

// ShippingMethodPagedQueryResponse is of type PagedQueryResponse
type ShippingMethodPagedQueryResponse struct {
	Total   int              `json:"total,omitempty"`
	Offset  int              `json:"offset"`
	Count   int              `json:"count"`
	Results []ShippingMethod `json:"results"`
}

// ShippingMethodReference implements the interface Reference
type ShippingMethodReference struct {
	Key string          `json:"key,omitempty"`
	ID  string          `json:"id,omitempty"`
	Obj *ShippingMethod `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shipping-method", Alias: (*Alias)(&obj)})
}

// ShippingMethodRemoveShippingRateAction implements the interface ShippingMethodUpdateAction
type ShippingMethodRemoveShippingRateAction struct {
	Zone         *ZoneReference     `json:"zone"`
	ShippingRate *ShippingRateDraft `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodRemoveShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingRate", Alias: (*Alias)(&obj)})
}

// ShippingMethodRemoveZoneAction implements the interface ShippingMethodUpdateAction
type ShippingMethodRemoveZoneAction struct {
	Zone *ZoneReference `json:"zone"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodRemoveZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeZone", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetDescriptionAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetKeyAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ShippingMethodSetPredicateAction implements the interface ShippingMethodUpdateAction
type ShippingMethodSetPredicateAction struct {
	Predicate string `json:"predicate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}

// ShippingMethodUpdate is of type Update
type ShippingMethodUpdate struct {
	Version int                          `json:"version"`
	Actions []ShippingMethodUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingMethodUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShippingMethodUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorShippingMethodUpdateAction(obj.Actions[i])
	}

	return nil
}

// ShippingRate is a standalone struct
type ShippingRate struct {
	Tiers      []ShippingRatePriceTier `json:"tiers"`
	Price      TypedMoney              `json:"price"`
	IsMatching bool                    `json:"isMatching,omitempty"`
	FreeAbove  TypedMoney              `json:"freeAbove,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRate) UnmarshalJSON(data []byte) error {
	type Alias ShippingRate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.FreeAbove != nil {
		obj.FreeAbove = mapDiscriminatorTypedMoney(obj.FreeAbove)
	}
	if obj.Price != nil {
		obj.Price = mapDiscriminatorTypedMoney(obj.Price)
	}
	for i := range obj.Tiers {
		obj.Tiers[i] = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
	}

	return nil
}

// ShippingRateDraft is a standalone struct
type ShippingRateDraft struct {
	Tiers     []ShippingRatePriceTier `json:"tiers,omitempty"`
	Price     *Money                  `json:"price"`
	FreeAbove *Money                  `json:"freeAbove,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Tiers {
		obj.Tiers[i] = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
	}

	return nil
}

// ZoneRate is a standalone struct
type ZoneRate struct {
	Zone          *ZoneReference `json:"zone"`
	ShippingRates []ShippingRate `json:"shippingRates"`
}

// ZoneRateDraft is a standalone struct
type ZoneRateDraft struct {
	Zone          *ZoneReference      `json:"zone"`
	ShippingRates []ShippingRateDraft `json:"shippingRates"`
}
