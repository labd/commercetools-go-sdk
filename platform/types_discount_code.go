// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type DiscountCode struct {
	// The unique ID of the discount code.
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy   *CreatedBy       `json:"createdBy,omitempty"`
	Name        *LocalizedString `json:"name,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
	// Unique identifier of this discount code.
	// This value is added to the cart
	// to enable the related cart discounts in the cart.
	Code string `json:"code"`
	// The referenced matching cart discounts can be applied to the cart once the DiscountCode is added.
	CartDiscounts []CartDiscountReference `json:"cartDiscounts"`
	// The discount code can only be applied to carts that match this predicate.
	CartPredicate *string `json:"cartPredicate,omitempty"`
	IsActive      bool    `json:"isActive"`
	// The platform will generate this array from the cart predicate.
	// It contains the references of all the resources that are addressed in the predicate.
	References []Reference `json:"references"`
	// The discount code can only be applied `maxApplications` times.
	MaxApplications *int `json:"maxApplications,omitempty"`
	// The discount code can only be applied `maxApplicationsPerCustomer` times per customer.
	MaxApplicationsPerCustomer *int          `json:"maxApplicationsPerCustomer,omitempty"`
	Custom                     *CustomFields `json:"custom,omitempty"`
	// The groups to which this discount code belong.
	Groups []string `json:"groups"`
	// The time from which the discount can be applied on a cart.
	// Before that time the code is invalid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// The time until the discount can be applied on a cart.
	// After that time the code is invalid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Used for the internal platform only and registers the reservation of use of a discount code.
	// Its value is managed by the platform.
	// It can change at any time due to internal and external factors.
	// It should not be used in customer logic.
	ApplicationVersion *int `json:"applicationVersion,omitempty"`
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
	Name        *LocalizedString `json:"name,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
	// Unique identifier of this discount code.
	// This value is added to the cart
	// to enable the related cart discounts in the cart.
	Code string `json:"code"`
	// The referenced matching cart discounts can be applied to the cart once the discount code is added.
	// The number of cart discounts in a discount code is limited to **10**.
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
	// The discount code can only be applied to carts that match this predicate.
	CartPredicate              *string            `json:"cartPredicate,omitempty"`
	IsActive                   *bool              `json:"isActive,omitempty"`
	MaxApplications            *int               `json:"maxApplications,omitempty"`
	MaxApplicationsPerCustomer *int               `json:"maxApplicationsPerCustomer,omitempty"`
	Custom                     *CustomFieldsDraft `json:"custom,omitempty"`
	// The groups to which this discount code shall belong to.
	Groups []string `json:"groups"`
	// The time from which the discount can be applied on a cart.
	// Before that time the code is invalid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// The time until the discount can be applied on a cart.
	// After that time the code is invalid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeDraft) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["groups"] == nil {
		delete(target, "groups")
	}

	return json.Marshal(target)
}

type DiscountCodePagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   *int           `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Results []DiscountCode `json:"results"`
}

type DiscountCodeReference struct {
	// Unique ID of the referenced resource.
	ID  string        `json:"id"`
	Obj *DiscountCode `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

type DiscountCodeResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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
		obj := DiscountCodeChangeCartDiscountsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeGroups":
		obj := DiscountCodeChangeGroupsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeIsActive":
		obj := DiscountCodeChangeIsActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCartPredicate":
		obj := DiscountCodeSetCartPredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := DiscountCodeSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := DiscountCodeSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := DiscountCodeSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMaxApplications":
		obj := DiscountCodeSetMaxApplicationsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMaxApplicationsPerCustomer":
		obj := DiscountCodeSetMaxApplicationsPerCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := DiscountCodeSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFrom":
		obj := DiscountCodeSetValidFromAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFromAndUntil":
		obj := DiscountCodeSetValidFromAndUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidUntil":
		obj := DiscountCodeSetValidUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type DiscountCodeChangeCartDiscountsAction struct {
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCartPredicateAction struct {
	// If the `cartPredicate` parameter is not included, the field will be emptied.
	CartPredicate *string `json:"cartPredicate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartPredicate", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomFieldAction struct {
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetDescriptionAction struct {
	// If the `description` parameter is not included, the field will be emptied.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsAction struct {
	// If the `maxApplications` parameter is not included, the field will be emptied.
	MaxApplications *int `json:"maxApplications,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetMaxApplicationsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplications", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsPerCustomerAction struct {
	// If the `maxApplicationsPerCustomer` parameter is not included, the field will be emptied.
	MaxApplicationsPerCustomer *int `json:"maxApplicationsPerCustomer,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetMaxApplicationsPerCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsPerCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplicationsPerCustomer", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetNameAction struct {
	// If the `name` parameter is not included, the field will be emptied.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAndUntilAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// If absent, the field with the value is removed in case a value was set before.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidUntilAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
