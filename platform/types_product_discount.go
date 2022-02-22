// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type ProductDiscount struct {
	// The unique ID of the product discount
	ID string `json:"id"`
	// The current version of the product discount.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy      `json:"createdBy,omitempty"`
	Name      LocalizedString `json:"name"`
	// User-specific unique identifier for a product discount.
	// Must be unique across a project.
	Key         *string              `json:"key,omitempty"`
	Description *LocalizedString     `json:"description,omitempty"`
	Value       ProductDiscountValue `json:"value"`
	// A valid ProductDiscount Predicate.
	Predicate string `json:"predicate"`
	// The string contains a number between 0 and 1.
	// A discount with greater sortOrder is prioritized higher than a discount with lower sortOrder.
	// A sortOrder must be unambiguous.
	SortOrder string `json:"sortOrder"`
	// Only active discount will be applied to product prices.
	IsActive bool `json:"isActive"`
	// The platform will generate this array from the predicate.
	// It contains the references of all the resources that are addressed in the predicate.
	References []Reference `json:"references"`
	// The time from which the discount should be effective.
	// Please take Eventual Consistency into account for calculated product discount values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// The time from which the discount should be ineffective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscount) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}
	for i := range obj.References {
		var err error
		obj.References[i], err = mapDiscriminatorReference(obj.References[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type ProductDiscountDraft struct {
	Name LocalizedString `json:"name"`
	// User-specific unique identifier for a product discount.
	// Must be unique across a project.
	// The field can be reset using the Set Key UpdateAction
	Key         *string                   `json:"key,omitempty"`
	Description *LocalizedString          `json:"description,omitempty"`
	Value       ProductDiscountValueDraft `json:"value"`
	// A valid ProductDiscount Predicate.
	Predicate string `json:"predicate"`
	// The string must contain a decimal number between 0 and 1.
	// A discount with greater sortOrder is prioritized higher than a discount with lower sortOrder.
	SortOrder string `json:"sortOrder"`
	// If set to `true` the discount will be applied to product prices.
	IsActive bool `json:"isActive"`
	// The time from which the discount should be effective.
	// Please take Eventual Consistency into account for calculated product discount values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// The time from which the discount should be effective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValueDraft(obj.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

type ProductDiscountMatchQuery struct {
	ProductId string     `json:"productId"`
	VariantId int        `json:"variantId"`
	Staged    bool       `json:"staged"`
	Price     QueryPrice `json:"price"`
}

type ProductDiscountPagedQueryResponse struct {
	Limit   int               `json:"limit"`
	Count   int               `json:"count"`
	Total   *int              `json:"total,omitempty"`
	Offset  int               `json:"offset"`
	Results []ProductDiscount `json:"results"`
}

type ProductDiscountReference struct {
	// Unique ID of the referenced resource.
	ID  string           `json:"id"`
	Obj *ProductDiscount `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

type ProductDiscountResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

type ProductDiscountUpdate struct {
	Version int                           `json:"version"`
	Actions []ProductDiscountUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductDiscountUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type ProductDiscountUpdateAction interface{}

func mapDiscriminatorProductDiscountUpdateAction(input interface{}) (ProductDiscountUpdateAction, error) {
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
	case "changeIsActive":
		obj := ProductDiscountChangeIsActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProductDiscountChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePredicate":
		obj := ProductDiscountChangePredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeSortOrder":
		obj := ProductDiscountChangeSortOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeValue":
		obj := ProductDiscountChangeValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Value != nil {
			var err error
			obj.Value, err = mapDiscriminatorProductDiscountValueDraft(obj.Value)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setDescription":
		obj := ProductDiscountSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ProductDiscountSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFrom":
		obj := ProductDiscountSetValidFromAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFromAndUntil":
		obj := ProductDiscountSetValidFromAndUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidUntil":
		obj := ProductDiscountSetValidUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ProductDiscountValue interface{}

func mapDiscriminatorProductDiscountValue(input interface{}) (ProductDiscountValue, error) {
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
	case "absolute":
		obj := ProductDiscountValueAbsolute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Money {
			var err error
			obj.Money[i], err = mapDiscriminatorTypedMoney(obj.Money[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "external":
		obj := ProductDiscountValueExternal{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := ProductDiscountValueRelative{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ProductDiscountValueAbsolute struct {
	Money []TypedMoney `json:"money"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountValueAbsolute) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountValueAbsolute
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Money {
		var err error
		obj.Money[i], err = mapDiscriminatorTypedMoney(obj.Money[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsolute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueDraft interface{}

func mapDiscriminatorProductDiscountValueDraft(input interface{}) (ProductDiscountValueDraft, error) {
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
	case "absolute":
		obj := ProductDiscountValueAbsoluteDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "external":
		obj := ProductDiscountValueExternalDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := ProductDiscountValueRelativeDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ProductDiscountValueAbsoluteDraft struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueAbsoluteDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsoluteDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueExternal struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueExternal) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternal
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueExternalDraft struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueExternalDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternalDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelative
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueRelativeDraft struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueRelativeDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelativeDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeIsActiveAction struct {
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangePredicateAction struct {
	// A valid ProductDiscount Predicate.
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangePredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangePredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePredicate", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeSortOrderAction struct {
	// The string must contain a number between 0 and 1.
	// A discount with greater sortOrder is prioritized higher than a discount with lower sortOrder.
	SortOrder string `json:"sortOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSortOrder", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeValueAction struct {
	Value ProductDiscountValueDraft `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValueDraft(obj.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetKeyAction struct {
	// The key to set.
	// If you provide a `null` value or do not set this field at all, the existing `key` field is removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAction struct {
	// The time from which the discount should be effective.
	// Please take Eventual Consistency into account for calculated product discount values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAndUntilAction struct {
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// The timeframe for which the discount should be effective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidUntilAction struct {
	// The time from which the discount should be ineffective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
