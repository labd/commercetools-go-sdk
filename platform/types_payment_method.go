package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type PaymentMethod struct {
	// Unique identifier of the PaymentMethod.
	ID string `json:"id"`
	// Current version of the PaymentMethod.
	Version int `json:"version"`
	// Date and time (UTC) the PaymentMethod was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the PaymentMethod was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the PaymentMethod.
	Key *string `json:"key,omitempty"`
	// Name of the PaymentMethod.
	Name *LocalizedString `json:"name,omitempty"`
	// Reference to a Customer associated with the PaymentMethod.
	//
	// If `businessUnit` is set, the Customer is an [Associate](ctp:api:type:Associate) of the Business Unit.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Reference to a BusinessUnit associated with the PaymentMethod.
	//
	// Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// Payment Method used for the Payment—for example, a credit card or direct debit.
	Method *string `json:"method,omitempty"`
	// Payment service that processes the Payment—for example, a PSP.
	PaymentInterface *string `json:"paymentInterface,omitempty"`
	// Account or instance of the payment interface when multiple accounts are used (per interface).
	InterfaceAccount *string `json:"interfaceAccount,omitempty"`
	// Tokenized representation of the PaymentMethod used by the payment interface.
	Token *PaymentMethodToken `json:"token,omitempty"`
	// Status of the PaymentMethod.
	PaymentMethodStatus PaymentMethodStatus `json:"paymentMethodStatus"`
	// Indicates if the PaymentMethod is the default.
	//
	// The default applies per Customer, Business Unit, or the combination of both (Associate).
	Default bool `json:"default"`
	// Custom Fields of the PaymentMethod.
	Custom *CustomFields `json:"custom,omitempty"`
	// IDs and references that last modified the PaymentMethod.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the PaymentMethod.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

type PaymentMethodDraft struct {
	// User-defined unique identifier for the PaymentMethod.
	Key *string `json:"key,omitempty"`
	// Name of the PaymentMethod.
	Name *LocalizedString `json:"name,omitempty"`
	// Reference to a Customer the PaymentMethod should belong to.
	//
	// If `businessUnit` is set, the Customer must be an [Associate](ctp:api:type:Associate) of the Business Unit.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	// Reference to a BusinessUnit the PaymentMethod should belong to.
	//
	// Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// Payment method to use for the Payment—for example, a credit card or direct debit.
	Method *string `json:"method,omitempty"`
	// Payment service that processes the Payment—for example, a PSP.
	PaymentInterface *string `json:"paymentInterface,omitempty"`
	// Account or instance of the payment interface when multiple accounts are used (per interface).
	InterfaceAccount *string `json:"interfaceAccount,omitempty"`
	// Tokenized representation of the PaymentMethod used by the payment interface.
	Token *PaymentMethodToken `json:"token,omitempty"`
	// Status of the PaymentMethod.
	PaymentMethodStatus *PaymentMethodStatus `json:"paymentMethodStatus,omitempty"`
	// Set to `true` if the PaymentMethod should be the default.
	//
	// The default applies per Customer, Business Unit, or the combination of both (Associate).
	Default *bool `json:"default,omitempty"`
	// Custom Fields for the PaymentMethod.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [PaymentMethod](ctp:api:type:PaymentMethod).
*
 */
type PaymentMethodPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [PaymentMethods](ctp:api:type:PaymentMethod) matching the query.
	Results []PaymentMethod `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [PaymentMethod](ctp:api:type:PaymentMethod).
*
 */
type PaymentMethodReference struct {
	// Unique identifier of the referenced [PaymentMethod](ctp:api:type:PaymentMethod).
	ID string `json:"id"`
	// Contains the representation of the expanded PaymentMethod. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for PaymentMethods.
	Obj *PaymentMethod `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment-method", Alias: (*Alias)(&obj)})
}

type PaymentMethodStatus string

const (
	PaymentMethodStatusActive   PaymentMethodStatus = "Active"
	PaymentMethodStatusInactive PaymentMethodStatus = "Inactive"
)

type PaymentMethodToken struct {
	// Tokenized representation of the Payment Method.
	//
	// It is _displayed_ in the following instances:
	//
	// - in the payload of an API Extension for Payments and PaymentMethods
	// - when querying Payments and PaymentMethods
	//
	// It is _masked_ in the following instances:
	//
	// - in the payload of Payment and PaymentMethod messages
	// - when querying MyPayments
	// - in referenced Payments and PaymentMethods embedded through [Reference Expansion](/general-concepts#reference-expansion)
	Value string `json:"value"`
}

type PaymentMethodUpdate struct {
	// Expected version of the PaymentMethod on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the PaymentMethod.
	Actions []PaymentMethodUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentMethodUpdate) UnmarshalJSON(data []byte) error {
	type Alias PaymentMethodUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorPaymentMethodUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type PaymentMethodUpdateAction interface{}

func mapDiscriminatorPaymentMethodUpdateAction(input interface{}) (PaymentMethodUpdateAction, error) {
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
	case "setCustomField":
		obj := PaymentMethodSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := PaymentMethodSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDefault":
		obj := PaymentMethodSetDefaultAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setInterfaceAccount":
		obj := PaymentMethodSetInterfaceAccountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := PaymentMethodSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethod":
		obj := PaymentMethodSetMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := PaymentMethodSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPaymentInterface":
		obj := PaymentMethodSetPaymentInterfaceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPaymentMethodStatus":
		obj := PaymentMethodSetPaymentMethodStatusAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Adding a Custom Field to a PaymentMethod generates the [PaymentMethodCustomFieldAdded](ctp:api:type:PaymentMethodCustomFieldAddedMessage) Message, removing one generates the [PaymentMethodCustomFieldRemoved](ctp:api:type:PaymentMethodCustomFieldRemovedMessage) Message, and updating an existing one generates the [PaymentMethodCustomFieldChanged](ctp:api:type:PaymentMethodCustomFieldChangedMessage) Message.
*
 */
type PaymentMethodSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields) to add, update, or remove.
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Adding or updating a Custom Type on a PaymentMethod generates the [PaymentMethodCustomTypeSet](ctp:api:type:PaymentMethodCustomTypeSetMessage) Message, removing one generates the [PaymentMethodCustomTypeRemoved](ctp:api:type:PaymentMethodCustomTypeRemovedMessage) Message.
*
 */
type PaymentMethodSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the PaymentMethod with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the PaymentMethod.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the PaymentMethod.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	This action generates the [PaymentMethodDefaultSet](ctp:api:type:PaymentMethodDefaultSetMessage) Message.
*
*	An inactive Payment Method cannot be set as the default, and the action will return an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
type PaymentMethodSetDefaultAction struct {
	// Value to set.
	Default bool `json:"default"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetDefaultAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetDefaultAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDefault", Alias: (*Alias)(&obj)})
}

/**
*	This action generates the [PaymentMethodInterfaceAccountSet](ctp:api:type:PaymentMethodInterfaceAccountSetMessage) Message.
*
 */
type PaymentMethodSetInterfaceAccountAction struct {
	// New account or instance of the payment interface.
	// If empty, any existing value will be removed.
	InterfaceAccount *string `json:"interfaceAccount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetInterfaceAccountAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetInterfaceAccountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceAccount", Alias: (*Alias)(&obj)})
}

/**
*	This action generates the [PaymentMethodKeySet](ctp:api:type:PaymentMethodKeySetMessage) Message.
*
 */
type PaymentMethodSetKeyAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	This action generates the [PaymentMethodMethodSet](ctp:api:type:PaymentMethodMethodSetMessage) Message.
*
 */
type PaymentMethodSetMethodAction struct {
	// New payment method—for example, a credit card or direct debit.
	// If empty, any existing value will be removed.
	Method *string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethod", Alias: (*Alias)(&obj)})
}

/**
*	This update action generates the [PaymentMethodNameSet](ctp:api:type:PaymentMethodNameSetMessage) Message.
*
 */
type PaymentMethodSetNameAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	This action generates the [PaymentMethodPaymentInterfaceSet](ctp:api:type:PaymentMethodPaymentInterfaceSetMessage) Message.
*
 */
type PaymentMethodSetPaymentInterfaceAction struct {
	// New payment service that processes the Payment—for example, a PSP.
	// If empty, any existing value will be removed.
	PaymentInterface *string `json:"paymentInterface,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetPaymentInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetPaymentInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPaymentInterface", Alias: (*Alias)(&obj)})
}

/**
*	This action generates the [PaymentMethodPaymentMethodStatusSet](ctp:api:type:PaymentMethodPaymentMethodStatusSetMessage) Message.
*
*	A default Payment Method cannot be set as inactive, and the action will return an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
type PaymentMethodSetPaymentMethodStatusAction struct {
	// Value to set.
	PaymentMethodStatus PaymentMethodStatus `json:"paymentMethodStatus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentMethodSetPaymentMethodStatusAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentMethodSetPaymentMethodStatusAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPaymentMethodStatus", Alias: (*Alias)(&obj)})
}
