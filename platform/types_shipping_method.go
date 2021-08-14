// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type PriceFunction struct {
	// The currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
	Function     string `json:"function"`
}

type ShippingMethod struct {
	// The unique ID of the shipping method.
	ID string `json:"id"`
	// The current version of the shipping method.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-specific unique identifier for the shipping method.
	Key                  *string              `json:"key,omitempty"`
	Name                 string               `json:"name"`
	LocalizedName        *LocalizedString     `json:"localizedName,omitempty"`
	Description          *string              `json:"description,omitempty"`
	LocalizedDescription *LocalizedString     `json:"localizedDescription,omitempty"`
	TaxCategory          TaxCategoryReference `json:"taxCategory"`
	ZoneRates            []ZoneRate           `json:"zoneRates"`
	// One shipping method in a project can be default.
	IsDefault bool `json:"isDefault"`
	// A Cart predicate which can be used to more precisely select a shipping method for a cart.
	Predicate *string       `json:"predicate,omitempty"`
	Custom    *CustomFields `json:"custom,omitempty"`
}

type ShippingMethodDraft struct {
	Key                  *string                       `json:"key,omitempty"`
	Name                 string                        `json:"name"`
	LocalizedName        *LocalizedString              `json:"localizedName,omitempty"`
	Description          *string                       `json:"description,omitempty"`
	LocalizedDescription *LocalizedString              `json:"localizedDescription,omitempty"`
	TaxCategory          TaxCategoryResourceIdentifier `json:"taxCategory"`
	ZoneRates            []ZoneRateDraft               `json:"zoneRates"`
	// If `true` the shipping method will be the default one in a project.
	IsDefault bool `json:"isDefault"`
	// A Cart predicate which can be used to more precisely select a shipping method for a cart.
	Predicate *string            `json:"predicate,omitempty"`
	Custom    *CustomFieldsDraft `json:"custom,omitempty"`
}

type ShippingMethodPagedQueryResponse struct {
	Limit   *int             `json:"limit,omitempty"`
	Count   int              `json:"count"`
	Total   *int             `json:"total,omitempty"`
	Offset  *int             `json:"offset,omitempty"`
	Results []ShippingMethod `json:"results"`
}

type ShippingMethodReference struct {
	// Unique ID of the referenced resource.
	ID  string          `json:"id"`
	Obj *ShippingMethod `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

type ShippingMethodResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

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

	return nil
}

type ShippingMethodUpdateAction interface{}

func mapDiscriminatorShippingMethodUpdateAction(input interface{}) (ShippingMethodUpdateAction, error) {

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
	case "addShippingRate":
		obj := ShippingMethodAddShippingRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addZone":
		obj := ShippingMethodAddZoneAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeIsDefault":
		obj := ShippingMethodChangeIsDefaultAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ShippingMethodChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxCategory":
		obj := ShippingMethodChangeTaxCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingRate":
		obj := ShippingMethodRemoveShippingRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeZone":
		obj := ShippingMethodRemoveZoneAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := ShippingMethodSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := ShippingMethodSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ShippingMethodSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ShippingMethodSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocalizedDescription":
		obj := ShippingMethodSetLocalizedDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocalizedName":
		obj := ShippingMethodSetLocalizedNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPredicate":
		obj := ShippingMethodSetPredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ShippingRate struct {
	Price TypedMoney `json:"price"`
	// The shipping is free if the order total (the sum of line item prices) exceeds the `freeAbove` value.
	// Note: `freeAbove` applies before any Cart or Product discounts, and can cause discounts to apply in invalid scenarios.
	// Use a Cart Discount to set the shipping price to 0 to avoid providing free shipping in invalid discount scenarios.
	FreeAbove TypedMoney `json:"freeAbove,omitempty"`
	// Only appears in response to requests for shipping methods by cart or location to mark this shipping rate as one that matches the cart or location.
	IsMatching *bool `json:"isMatching,omitempty"`
	// A list of shipping rate price tiers.
	Tiers []ShippingRatePriceTier `json:"tiers"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRate) UnmarshalJSON(data []byte) error {
	type Alias ShippingRate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Price != nil {
		var err error
		obj.Price, err = mapDiscriminatorTypedMoney(obj.Price)
		if err != nil {
			return err
		}
	}
	if obj.FreeAbove != nil {
		var err error
		obj.FreeAbove, err = mapDiscriminatorTypedMoney(obj.FreeAbove)
		if err != nil {
			return err
		}
	}
	return nil
}

type ShippingRateDraft struct {
	Price Money `json:"price"`
	// The shipping is free if the order total (the sum of line item prices) exceeds the freeAbove value.
	// Note: `freeAbove` applies before any Cart or Product discounts, and can cause discounts to apply in invalid scenarios.
	// Use a Cart Discount to set the shipping price to 0 to avoid providing free shipping in invalid discount scenarios.
	FreeAbove *Money `json:"freeAbove,omitempty"`
	// A list of shipping rate price tiers.
	Tiers []ShippingRatePriceTier `json:"tiers,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ShippingRatePriceTier interface{}

func mapDiscriminatorShippingRatePriceTier(input interface{}) (ShippingRatePriceTier, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "CartClassification":
		obj := CartClassificationTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CartScore":
		obj := CartScoreTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CartValue":
		obj := CartValueTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CartClassificationTier struct {
	Value      string `json:"value"`
	Price      Money  `json:"price"`
	IsMatching *bool  `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartClassification", Alias: (*Alias)(&obj)})
}

type CartScoreTier struct {
	Score         float64        `json:"score"`
	Price         *Money         `json:"price,omitempty"`
	PriceFunction *PriceFunction `json:"priceFunction,omitempty"`
	IsMatching    *bool          `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartScoreTier) MarshalJSON() ([]byte, error) {
	type Alias CartScoreTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartScore", Alias: (*Alias)(&obj)})
}

type CartValueTier struct {
	MinimumCentAmount int   `json:"minimumCentAmount"`
	Price             Money `json:"price"`
	IsMatching        *bool `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartValueTier) MarshalJSON() ([]byte, error) {
	type Alias CartValueTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartValue", Alias: (*Alias)(&obj)})
}

type ShippingRateTierType string

const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

type ZoneRate struct {
	Zone ZoneReference `json:"zone"`
	// The array does not contain two shipping rates with the same currency.
	ShippingRates []ShippingRate `json:"shippingRates"`
}

type ZoneRateDraft struct {
	Zone ZoneResourceIdentifier `json:"zone"`
	// The array must not contain two shipping rates with the same currency.
	ShippingRates []ShippingRateDraft `json:"shippingRates"`
}

type ShippingMethodAddShippingRateAction struct {
	Zone         ZoneResourceIdentifier `json:"zone"`
	ShippingRate ShippingRateDraft      `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodAddShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingRate", Alias: (*Alias)(&obj)})
}

type ShippingMethodAddZoneAction struct {
	Zone ZoneResourceIdentifier `json:"zone"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodAddZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addZone", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeIsDefaultAction struct {
	// Only one ShippingMethod in a project can be default.
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

type ShippingMethodChangeTaxCategoryAction struct {
	TaxCategory TaxCategoryResourceIdentifier `json:"taxCategory"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodChangeTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCategory", Alias: (*Alias)(&obj)})
}

type ShippingMethodRemoveShippingRateAction struct {
	Zone         ZoneResourceIdentifier `json:"zone"`
	ShippingRate ShippingRateDraft      `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodRemoveShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingRate", Alias: (*Alias)(&obj)})
}

type ShippingMethodRemoveZoneAction struct {
	Zone ZoneResourceIdentifier `json:"zone"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodRemoveZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeZone", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetCustomTypeAction struct {
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetDescriptionAction struct {
	Description *string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetKeyAction struct {
	// If `key` is absent or `null`, it is removed if it exists.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetLocalizedDescriptionAction struct {
	LocalizedDescription *LocalizedString `json:"localizedDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetLocalizedDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetLocalizedDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocalizedDescription", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetLocalizedNameAction struct {
	LocalizedName *LocalizedString `json:"localizedName,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetLocalizedNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetLocalizedNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocalizedName", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetPredicateAction struct {
	// A valid Cart predicate.
	// If `predicate` is absent or `null`, it is removed if it exists.
	Predicate *string `json:"predicate,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ShippingMethodSetPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}
