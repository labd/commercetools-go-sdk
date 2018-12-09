// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

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

func (obj *ProductDiscount) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		obj.References[i] = AbstractReferenceDiscriminatorMapping(obj.References[i])
	}
	if obj.Value != nil {
		obj.Value = AbstractProductDiscountValueDiscriminatorMapping(obj.Value)
	}

	return nil
}

type ProductDiscountChangeIsActiveAction struct {
	IsActive bool `json:"isActive"`
}

func (obj ProductDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

func (obj ProductDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangePredicateAction struct {
	Predicate string `json:"predicate"`
}

func (obj ProductDiscountChangePredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangePredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePredicate", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeSortOrderAction struct {
	SortOrder string `json:"sortOrder"`
}

func (obj ProductDiscountChangeSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSortOrder", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeValueAction struct {
	Value ProductDiscountValue `json:"value"`
}

func (obj ProductDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}
func (obj *ProductDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = AbstractProductDiscountValueDiscriminatorMapping(obj.Value)
	}

	return nil
}

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

func (obj *ProductDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = AbstractProductDiscountValueDiscriminatorMapping(obj.Value)
	}

	return nil
}

type ProductDiscountMatchQuery struct {
	VariantID float64 `json:"variantId"`
	Staged    bool    `json:"staged"`
	ProductID string  `json:"productId"`
	Price     *Price  `json:"price"`
}

type ProductDiscountPagedQueryResponse struct {
	Total   int               `json:"total,omitempty"`
	Offset  int               `json:"offset"`
	Count   int               `json:"count"`
	Results []ProductDiscount `json:"results"`
}

type ProductDiscountReference struct {
	Key string           `json:"key,omitempty"`
	ID  string           `json:"id,omitempty"`
	Obj *ProductDiscount `json:"obj,omitempty"`
}

func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product-discount", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

func (obj ProductDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAction struct {
	ValidFrom time.Time `json:"validFrom,omitempty"`
}

func (obj ProductDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAndUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
	ValidFrom  time.Time `json:"validFrom,omitempty"`
}

func (obj ProductDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
}

func (obj ProductDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

type ProductDiscountUpdate struct {
	Version int                           `json:"version"`
	Actions []ProductDiscountUpdateAction `json:"actions"`
}

func (obj *ProductDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractProductDiscountUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ProductDiscountUpdateAction interface{}
type AbstractProductDiscountUpdateAction struct{}

func AbstractProductDiscountUpdateActionDiscriminatorMapping(input ProductDiscountUpdateAction) ProductDiscountUpdateAction {
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
			new.Value = AbstractProductDiscountValueDiscriminatorMapping(new.Value)
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

type ProductDiscountValue interface{}
type AbstractProductDiscountValue struct{}

func AbstractProductDiscountValueDiscriminatorMapping(input ProductDiscountValue) ProductDiscountValue {
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

type ProductDiscountValueAbsolute struct {
	Money []Money `json:"money"`
}

func (obj ProductDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsolute
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "absolute", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueExternal struct{}

func (obj ProductDiscountValueExternal) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternal
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "external", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

func (obj ProductDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelative
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "relative", Alias: (*Alias)(&obj)})
}
