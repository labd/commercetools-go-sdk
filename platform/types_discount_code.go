// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type DiscountCode struct {
	// The unique ID of the discount code.
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy   *CreatedBy       `json:"createdBy,omitEmpty"`
	Name        *LocalizedString `json:"name,omitEmpty"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// Unique identifier of this discount code.
	// This value is added to the cart
	// to enable the related cart discounts in the cart.
	Code string `json:"code"`
	// The referenced matching cart discounts can be applied to the cart once the DiscountCode is added.
	CartDiscounts []CartDiscountReference `json:"cartDiscounts"`
	// The discount code can only be applied to carts that match this predicate.
	CartPredicate string `json:"cartPredicate,omitEmpty"`
	IsActive      bool   `json:"isActive"`
	// The platform will generate this array from the cart predicate.
	// It contains the references of all the resources that are addressed in the predicate.
	References []Reference `json:"references"`
	// The discount code can only be applied `maxApplications` times.
	MaxApplications int `json:"maxApplications,omitEmpty"`
	// The discount code can only be applied `maxApplicationsPerCustomer` times per customer.
	MaxApplicationsPerCustomer int           `json:"maxApplicationsPerCustomer,omitEmpty"`
	Custom                     *CustomFields `json:"custom,omitEmpty"`
	// The groups to which this discount code belong.
	Groups []string `json:"groups"`
	// The time from which the discount can be applied on a cart.
	// Before that time the code is invalid.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// The time until the discount can be applied on a cart.
	// After that time the code is invalid.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCode) UnmarshalJSON(data []byte) error {
	type Alias DiscountCode
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type DiscountCodeDraft struct {
	Name        *LocalizedString `json:"name,omitEmpty"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// Unique identifier of this discount code.
	// This value is added to the cart
	// to enable the related cart discounts in the cart.
	Code string `json:"code"`
	// The referenced matching cart discounts can be applied to the cart once the discount code is added.
	// The number of cart discounts in a discount code is limited to **10**.
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
	// The discount code can only be applied to carts that match this predicate.
	CartPredicate              string             `json:"cartPredicate,omitEmpty"`
	IsActive                   bool               `json:"isActive,omitEmpty"`
	MaxApplications            int                `json:"maxApplications,omitEmpty"`
	MaxApplicationsPerCustomer int                `json:"maxApplicationsPerCustomer,omitEmpty"`
	Custom                     *CustomFieldsDraft `json:"custom,omitEmpty"`
	// The groups to which this discount code shall belong to.
	Groups []string `json:"groups,omitEmpty"`
	// The time from which the discount can be applied on a cart.
	// Before that time the code is invalid.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// The time until the discount can be applied on a cart.
	// After that time the code is invalid.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

type DiscountCodePagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   int            `json:"total,omitEmpty"`
	Offset  int            `json:"offset"`
	Results []DiscountCode `json:"results"`
}

type DiscountCodeReference struct {
	Id  string        `json:"id"`
	Obj *DiscountCode `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

type DiscountCodeResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

type DiscountCodeUpdate struct {
	Version int                        `json:"version"`
	Actions []DiscountCodeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCodeUpdate) UnmarshalJSON(data []byte) error {
	type Alias DiscountCodeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type DiscountCodeUpdateAction interface{}

func mapDiscriminatorDiscountCodeUpdateAction(input interface{}) (DiscountCodeUpdateAction, error) {

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
	case "changeCartDiscounts":
		new := DiscountCodeChangeCartDiscountsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeGroups":
		new := DiscountCodeChangeGroupsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeIsActive":
		new := DiscountCodeChangeIsActiveAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCartPredicate":
		new := DiscountCodeSetCartPredicateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := DiscountCodeSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := DiscountCodeSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := DiscountCodeSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMaxApplications":
		new := DiscountCodeSetMaxApplicationsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMaxApplicationsPerCustomer":
		new := DiscountCodeSetMaxApplicationsPerCustomerAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := DiscountCodeSetNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFrom":
		new := DiscountCodeSetValidFromAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFromAndUntil":
		new := DiscountCodeSetValidFromAndUntilAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidUntil":
		new := DiscountCodeSetValidUntilAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type DiscountCodeChangeCartDiscountsAction struct {
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeChangeCartDiscountsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeCartDiscountsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartDiscounts", Alias: (*Alias)(&obj)})
}

type DiscountCodeChangeGroupsAction struct {
	// The groups to which this discount code shall belong to.
	// Use empty array to remove the code from all groups.
	Groups []string `json:"groups"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCartPredicateAction struct {
	// If the `cartPredicate` parameter is not included, the field will be emptied.
	CartPredicate string `json:"cartPredicate,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartPredicate", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetDescriptionAction struct {
	// If the `description` parameter is not included, the field will be emptied.
	Description *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsAction struct {
	// If the `maxApplications` parameter is not included, the field will be emptied.
	MaxApplications int `json:"maxApplications,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetMaxApplicationsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplications", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsPerCustomerAction struct {
	// If the `maxApplicationsPerCustomer` parameter is not included, the field will be emptied.
	MaxApplicationsPerCustomer int `json:"maxApplicationsPerCustomer,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetMaxApplicationsPerCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsPerCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplicationsPerCustomer", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetNameAction struct {
	// If the `name` parameter is not included, the field will be emptied.
	Name *LocalizedString `json:"name,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAndUntilAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// If absent, the field with the value is removed in case a value was set before.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidUntilAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj DiscountCodeSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
