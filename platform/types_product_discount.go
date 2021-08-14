// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type ProductDiscount struct {
	// The unique ID of the product discount
	Id string `json:"id"`
	// The current version of the product discount.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy      `json:"createdBy,omitEmpty"`
	Name      LocalizedString `json:"name"`
	// User-specific unique identifier for a product discount.
	// Must be unique across a project.
	Key         string               `json:"key,omitEmpty"`
	Description *LocalizedString     `json:"description,omitEmpty"`
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
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// The time from which the discount should be ineffective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
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
	return nil
}

type ProductDiscountDraft struct {
	Name LocalizedString `json:"name"`
	// User-specific unique identifier for a product discount.
	// Must be unique across a project.
	// The field can be reset using the Set Key UpdateAction
	Key         string                    `json:"key,omitEmpty"`
	Description *LocalizedString          `json:"description,omitEmpty"`
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
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// The time from which the discount should be effective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
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
	Total   int               `json:"total,omitEmpty"`
	Offset  int               `json:"offset"`
	Results []ProductDiscount `json:"results"`
}

type ProductDiscountReference struct {
	Id  string           `json:"id"`
	Obj *ProductDiscount `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

type ProductDiscountResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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

	return nil
}

type ProductDiscountUpdateAction interface{}

func mapDiscriminatorProductDiscountUpdateAction(input interface{}) (ProductDiscountUpdateAction, error) {

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
	case "changeIsActive":
		new := ProductDiscountChangeIsActiveAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ProductDiscountChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePredicate":
		new := ProductDiscountChangePredicateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeSortOrder":
		new := ProductDiscountChangeSortOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeValue":
		new := ProductDiscountChangeValueAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ProductDiscountSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ProductDiscountSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFrom":
		new := ProductDiscountSetValidFromAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFromAndUntil":
		new := ProductDiscountSetValidFromAndUntilAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidUntil":
		new := ProductDiscountSetValidUntilAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ProductDiscountValue interface{}

func mapDiscriminatorProductDiscountValue(input interface{}) (ProductDiscountValue, error) {

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
	case "absolute":
		new := ProductDiscountValueAbsolute{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "external":
		new := ProductDiscountValueExternal{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "relative":
		new := ProductDiscountValueRelative{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
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

	return nil
}

// MarshalJSON override to set the discriminator value
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "absolute":
		new := ProductDiscountValueAbsoluteDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "external":
		new := ProductDiscountValueExternalDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "relative":
		new := ProductDiscountValueRelativeDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ProductDiscountValueAbsoluteDraft struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueAbsoluteDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsoluteDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueExternal struct {
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueExternal) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternal
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueExternalDraft struct {
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAndUntilAction struct {
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// The timeframe for which the discount should be effective.
	// Please take Eventual Consistency into account for calculated undiscounted values.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
