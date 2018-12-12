// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ProductDiscountUpdateAction uses action as discriminator attribute
type ProductDiscountUpdateAction interface{}

func mapDiscriminatorProductDiscountUpdateAction(input ProductDiscountUpdateAction) ProductDiscountUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeIsActive":
		new := ProductDiscountChangeIsActiveAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ProductDiscountChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changePredicate":
		new := ProductDiscountChangePredicateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeSortOrder":
		new := ProductDiscountChangeSortOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeValue":
		new := ProductDiscountChangeValueAction{}
		mapstructure.Decode(input, &new)
		if new.Value != nil {
			new.Value = mapDiscriminatorProductDiscountValue(new.Value)
		}

		return new
	case "setDescription":
		new := ProductDiscountSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidFrom":
		new := ProductDiscountSetValidFromAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidFromAndUntil":
		new := ProductDiscountSetValidFromAndUntilAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidUntil":
		new := ProductDiscountSetValidUntilAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// ProductDiscountValue uses type as discriminator attribute
type ProductDiscountValue interface{}

func mapDiscriminatorProductDiscountValue(input ProductDiscountValue) ProductDiscountValue {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "absolute":
		new := ProductDiscountValueAbsolute{}
		mapstructure.Decode(input, &new)
		return new
	case "external":
		new := ProductDiscountValueExternal{}
		mapstructure.Decode(input, &new)
		return new
	case "relative":
		new := ProductDiscountValueRelative{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// ProductDiscount is of type Resource
type ProductDiscount struct {
	Version        int                  `json:"version"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
	ID             string               `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	Value          ProductDiscountValue `json:"value"`
	ValidUntil     time.Time            `json:"validUntil,omitempty"`
	ValidFrom      time.Time            `json:"validFrom,omitempty"`
	SortOrder      string               `json:"sortOrder"`
	References     []Reference          `json:"references"`
	Predicate      string               `json:"predicate"`
	Name           *LocalizedString     `json:"name"`
	IsActive       bool                 `json:"isActive"`
	Description    *LocalizedString     `json:"description,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscount) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		obj.References[i] = mapDiscriminatorReference(obj.References[i])
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorProductDiscountValue(obj.Value)
	}

	return nil
}

// ProductDiscountChangeIsActiveAction implements the interface ProductDiscountUpdateAction
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

// ProductDiscountChangeNameAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ProductDiscountChangePredicateAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangePredicateAction struct {
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

// ProductDiscountChangeSortOrderAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeSortOrderAction struct {
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

// ProductDiscountChangeValueAction implements the interface ProductDiscountUpdateAction
type ProductDiscountChangeValueAction struct {
	Value ProductDiscountValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorProductDiscountValue(obj.Value)
	}

	return nil
}

// ProductDiscountDraft is a standalone struct
type ProductDiscountDraft struct {
	Value       ProductDiscountValue `json:"value"`
	ValidUntil  time.Time            `json:"validUntil,omitempty"`
	ValidFrom   time.Time            `json:"validFrom,omitempty"`
	SortOrder   string               `json:"sortOrder"`
	Predicate   string               `json:"predicate"`
	Name        *LocalizedString     `json:"name"`
	IsActive    bool                 `json:"isActive"`
	Description *LocalizedString     `json:"description,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorProductDiscountValue(obj.Value)
	}

	return nil
}

// ProductDiscountMatchQuery is a standalone struct
type ProductDiscountMatchQuery struct {
	VariantID float64 `json:"variantId"`
	Staged    bool    `json:"staged"`
	ProductID string  `json:"productId"`
	Price     *Price  `json:"price"`
}

// ProductDiscountPagedQueryResponse is of type PagedQueryResponse
type ProductDiscountPagedQueryResponse struct {
	Total   int               `json:"total,omitempty"`
	Offset  int               `json:"offset"`
	Count   int               `json:"count"`
	Results []ProductDiscount `json:"results"`
}

// ProductDiscountReference implements the interface Reference
type ProductDiscountReference struct {
	Key string           `json:"key,omitempty"`
	ID  string           `json:"id,omitempty"`
	Obj *ProductDiscount `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product-discount", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetDescriptionAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetValidFromAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetValidFromAction struct {
	ValidFrom time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetValidFromAndUntilAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetValidFromAndUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
	ValidFrom  time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

// ProductDiscountSetValidUntilAction implements the interface ProductDiscountUpdateAction
type ProductDiscountSetValidUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

// ProductDiscountUpdate is of type Update
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
		obj.Actions[i] = mapDiscriminatorProductDiscountUpdateAction(obj.Actions[i])
	}

	return nil
}

// ProductDiscountValueAbsolute implements the interface ProductDiscountValue
type ProductDiscountValueAbsolute struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsolute
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "absolute", Alias: (*Alias)(&obj)})
}

// ProductDiscountValueExternal implements the interface ProductDiscountValue
type ProductDiscountValueExternal struct{}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueExternal) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternal
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "external", Alias: (*Alias)(&obj)})
}

// ProductDiscountValueRelative implements the interface ProductDiscountValue
type ProductDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelative
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "relative", Alias: (*Alias)(&obj)})
}
