// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type DiscountCode struct {
	Version                    int                     `json:"version"`
	LastModifiedAt             time.Time               `json:"lastModifiedAt"`
	ID                         string                  `json:"id"`
	CreatedAt                  time.Time               `json:"createdAt"`
	ValidUntil                 time.Time               `json:"validUntil,omitempty"`
	ValidFrom                  time.Time               `json:"validFrom,omitempty"`
	References                 []Reference             `json:"references"`
	Name                       *LocalizedString        `json:"name,omitempty"`
	MaxApplicationsPerCustomer int                     `json:"maxApplicationsPerCustomer,omitempty"`
	MaxApplications            int                     `json:"maxApplications,omitempty"`
	IsActive                   bool                    `json:"isActive"`
	Groups                     []string                `json:"groups"`
	Description                *LocalizedString        `json:"description,omitempty"`
	Custom                     *CustomFields           `json:"custom,omitempty"`
	Code                       string                  `json:"code"`
	CartPredicate              string                  `json:"cartPredicate,omitempty"`
	CartDiscounts              []CartDiscountReference `json:"cartDiscounts"`
}

func (obj *DiscountCode) UnmarshalJSON(data []byte) error {
	type Alias DiscountCode
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		obj.References[i] = AbstractReferenceDiscriminatorMapping(obj.References[i])
	}

	return nil
}

type DiscountCodeChangeCartDiscountsAction struct {
	CartDiscounts []CartDiscountReference `json:"cartDiscounts"`
}

func (obj DiscountCodeChangeCartDiscountsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeCartDiscountsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartDiscounts", Alias: (*Alias)(&obj)})
}

type DiscountCodeChangeGroupsAction struct {
	Groups []string `json:"groups"`
}

func (obj DiscountCodeChangeGroupsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeGroupsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeGroups", Alias: (*Alias)(&obj)})
}

type DiscountCodeChangeIsActiveAction struct {
	IsActive bool `json:"isActive"`
}

func (obj DiscountCodeChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type DiscountCodeDraft struct {
	ValidUntil                 time.Time               `json:"validUntil,omitempty"`
	ValidFrom                  time.Time               `json:"validFrom,omitempty"`
	Name                       *LocalizedString        `json:"name,omitempty"`
	MaxApplicationsPerCustomer int                     `json:"maxApplicationsPerCustomer,omitempty"`
	MaxApplications            int                     `json:"maxApplications,omitempty"`
	IsActive                   bool                    `json:"isActive,omitempty"`
	Groups                     []string                `json:"groups,omitempty"`
	Description                *LocalizedString        `json:"description,omitempty"`
	Custom                     *CustomFieldsDraft      `json:"custom,omitempty"`
	Code                       string                  `json:"code"`
	CartPredicate              string                  `json:"cartPredicate,omitempty"`
	CartDiscounts              []CartDiscountReference `json:"cartDiscounts"`
}

type DiscountCodePagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []DiscountCode `json:"results"`
}

type DiscountCodeReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *DiscountCode `json:"obj,omitempty"`
}

func (obj DiscountCodeReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "discount-code", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCartPredicateAction struct {
	CartPredicate string `json:"cartPredicate,omitempty"`
}

func (obj DiscountCodeSetCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartPredicate", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj DiscountCodeSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj DiscountCodeSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

func (obj DiscountCodeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsAction struct {
	MaxApplications int `json:"maxApplications,omitempty"`
}

func (obj DiscountCodeSetMaxApplicationsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplications", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsPerCustomerAction struct {
	MaxApplicationsPerCustomer int `json:"maxApplicationsPerCustomer,omitempty"`
}

func (obj DiscountCodeSetMaxApplicationsPerCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsPerCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplicationsPerCustomer", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

func (obj DiscountCodeSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAction struct {
	ValidFrom time.Time `json:"validFrom,omitempty"`
}

func (obj DiscountCodeSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAndUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
	ValidFrom  time.Time `json:"validFrom,omitempty"`
}

func (obj DiscountCodeSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
}

func (obj DiscountCodeSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

type DiscountCodeUpdate struct {
	Version int                        `json:"version"`
	Actions []DiscountCodeUpdateAction `json:"actions"`
}

func (obj *DiscountCodeUpdate) UnmarshalJSON(data []byte) error {
	type Alias DiscountCodeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractDiscountCodeUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type DiscountCodeUpdateAction interface{}
type AbstractDiscountCodeUpdateAction struct{}

func AbstractDiscountCodeUpdateActionDiscriminatorMapping(input DiscountCodeUpdateAction) DiscountCodeUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeCartDiscounts":
		new := DiscountCodeChangeCartDiscountsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeGroups":
		new := DiscountCodeChangeGroupsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeIsActive":
		new := DiscountCodeChangeIsActiveAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCartPredicate":
		new := DiscountCodeSetCartPredicateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := DiscountCodeSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := DiscountCodeSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := DiscountCodeSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMaxApplications":
		new := DiscountCodeSetMaxApplicationsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMaxApplicationsPerCustomer":
		new := DiscountCodeSetMaxApplicationsPerCustomerAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setName":
		new := DiscountCodeSetNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidFrom":
		new := DiscountCodeSetValidFromAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidFromAndUntil":
		new := DiscountCodeSetValidFromAndUntilAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidUntil":
		new := DiscountCodeSetValidUntilAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
