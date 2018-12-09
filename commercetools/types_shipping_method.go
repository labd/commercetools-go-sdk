// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type CartClassificationTier struct {
	Value      string `json:"value"`
	Price      *Money `json:"price"`
	IsMatching bool   `json:"isMatching,omitempty"`
}

func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartClassification", Alias: (*Alias)(&obj)})
}

type CartScoreTier struct {
	Score         float64        `json:"score"`
	PriceFunction *PriceFunction `json:"priceFunction,omitempty"`
	Price         *Money         `json:"price,omitempty"`
	IsMatching    bool           `json:"isMatching,omitempty"`
}

func (obj CartScoreTier) MarshalJSON() ([]byte, error) {
	type Alias CartScoreTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartScore", Alias: (*Alias)(&obj)})
}

type CartValueTier struct {
	Price             *Money `json:"price"`
	MinimumCentAmount int    `json:"minimumCentAmount"`
	IsMatching        bool   `json:"isMatching,omitempty"`
}

func (obj CartValueTier) MarshalJSON() ([]byte, error) {
	type Alias CartValueTier
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartValue", Alias: (*Alias)(&obj)})
}

type PriceFunction struct {
	Function     string       `json:"function"`
	CurrencyCode CurrencyCode `json:"currencyCode"`
}

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

type ShippingMethodAddShippingRateAction struct {
	Zone         *ZoneReference     `json:"zone"`
	ShippingRate *ShippingRateDraft `json:"shippingRate"`
}

func (obj ShippingMethodAddShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingRate", Alias: (*Alias)(&obj)})
}

type ShippingMethodAddZoneAction struct {
	Zone *ZoneReference `json:"zone"`
}

func (obj ShippingMethodAddZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addZone", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeIsDefaultAction struct {
	IsDefault bool `json:"isDefault"`
}

func (obj ShippingMethodChangeIsDefaultAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeIsDefaultAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsDefault", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeNameAction struct {
	Name string `json:"name"`
}

func (obj ShippingMethodChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeTaxCategoryAction struct {
	TaxCategory *TaxCategoryReference `json:"taxCategory"`
}

func (obj ShippingMethodChangeTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCategory", Alias: (*Alias)(&obj)})
}

type ShippingMethodDraft struct {
	ZoneRates   []ZoneRateDraft       `json:"zoneRates"`
	TaxCategory *TaxCategoryReference `json:"taxCategory"`
	Predicate   string                `json:"predicate,omitempty"`
	Name        string                `json:"name"`
	Key         string                `json:"key,omitempty"`
	IsDefault   bool                  `json:"isDefault"`
	Description string                `json:"description,omitempty"`
}

type ShippingMethodPagedQueryResponse struct {
	Total   int              `json:"total,omitempty"`
	Offset  int              `json:"offset"`
	Count   int              `json:"count"`
	Results []ShippingMethod `json:"results"`
}

type ShippingMethodReference struct {
	Key string          `json:"key,omitempty"`
	ID  string          `json:"id,omitempty"`
	Obj *ShippingMethod `json:"obj,omitempty"`
}

func (obj ShippingMethodReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "shipping-method", Alias: (*Alias)(&obj)})
}

type ShippingMethodRemoveShippingRateAction struct {
	Zone         *ZoneReference     `json:"zone"`
	ShippingRate *ShippingRateDraft `json:"shippingRate"`
}

func (obj ShippingMethodRemoveShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingRate", Alias: (*Alias)(&obj)})
}

type ShippingMethodRemoveZoneAction struct {
	Zone *ZoneReference `json:"zone"`
}

func (obj ShippingMethodRemoveZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeZone", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

func (obj ShippingMethodSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj ShippingMethodSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetPredicateAction struct {
	Predicate string `json:"predicate,omitempty"`
}

func (obj ShippingMethodSetPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}

type ShippingMethodUpdate struct {
	Version int                          `json:"version"`
	Actions []ShippingMethodUpdateAction `json:"actions"`
}

func (obj *ShippingMethodUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShippingMethodUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractShippingMethodUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ShippingMethodUpdateAction interface{}
type AbstractShippingMethodUpdateAction struct{}

func AbstractShippingMethodUpdateActionDiscriminatorMapping(input ShippingMethodUpdateAction) ShippingMethodUpdateAction {
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

type ShippingRate struct {
	Tiers      []ShippingRatePriceTier `json:"tiers"`
	Price      TypedMoney              `json:"price"`
	IsMatching bool                    `json:"isMatching,omitempty"`
	FreeAbove  TypedMoney              `json:"freeAbove,omitempty"`
}

func (obj *ShippingRate) UnmarshalJSON(data []byte) error {
	type Alias ShippingRate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.FreeAbove != nil {
		obj.FreeAbove = AbstractTypedMoneyDiscriminatorMapping(obj.FreeAbove)
	}
	if obj.Price != nil {
		obj.Price = AbstractTypedMoneyDiscriminatorMapping(obj.Price)
	}
	for i := range obj.Tiers {
		obj.Tiers[i] = AbstractShippingRatePriceTierDiscriminatorMapping(obj.Tiers[i])
	}

	return nil
}

type ShippingRateDraft struct {
	Tiers     []ShippingRatePriceTier `json:"tiers,omitempty"`
	Price     *Money                  `json:"price"`
	FreeAbove *Money                  `json:"freeAbove,omitempty"`
}

func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Tiers {
		obj.Tiers[i] = AbstractShippingRatePriceTierDiscriminatorMapping(obj.Tiers[i])
	}

	return nil
}

type ShippingRatePriceTier interface{}
type AbstractShippingRatePriceTier struct{}

func AbstractShippingRatePriceTierDiscriminatorMapping(input ShippingRatePriceTier) ShippingRatePriceTier {
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

type ShippingRateTierType string

const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

type ZoneRate struct {
	Zone          *ZoneReference `json:"zone"`
	ShippingRates []ShippingRate `json:"shippingRates"`
}

type ZoneRateDraft struct {
	Zone          *ZoneReference      `json:"zone"`
	ShippingRates []ShippingRateDraft `json:"shippingRates"`
}
