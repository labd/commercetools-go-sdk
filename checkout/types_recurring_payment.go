package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Paginated result containing RecurringPayments.
*
 */
type PaginatedRecurringPayment struct {
	// Number of results requested.
	Limit int `json:"limit"`
	// Number of elements skipped.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query. Only present when the `withTotal` query parameter is `true`.
	Total *int `json:"total,omitempty"`
	// RecurringPayments matching the query.
	Results []RecurringPayment `json:"results"`
}

/**
*	The [PaymentMethod](ctp:api:type:PaymentMethod) used to pay a [RecurringOrder](ctp:api:type:RecurringOrder) and the [Connector](/checkout/connectors-and-applications#payment-connectors) that processes it.
*
*	Checkout only supports processing a [RecurringPayment](ctp:checkout:type:RecurringPayment) with one PaymentMethodConfiguration.
*
 */
type PaymentMethodConfiguration struct {
	// [PaymentMethod](ctp:api:type:PaymentMethod) used to pay the [RecurringOrder](ctp:api:type:RecurringOrder).
	PaymentMethod PaymentMethodReference `json:"paymentMethod"`
	// [Connector](/checkout/connectors-and-applications#payment-connectors) Deployment that processes the future payments of the [RecurringOrder](ctp:api:type:RecurringOrder).
	ConnectorDeployment ConnectorDeploymentReference `json:"connectorDeployment"`
}

/**
*	Reference to a [RecurringOrder](ctp:api:type:RecurringOrder).
*
 */
type RecurringOrderReference struct {
	// Unique identifier of the referenced [RecurringOrder](ctp:api:type:RecurringOrder).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderReference) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurring-order", Alias: (*Alias)(&obj)})
}

/**
*	Maps a [RecurringOrder](ctp:api:type:RecurringOrder) to the [PaymentMethod](ctp:api:type:PaymentMethod) and [Connector](/checkout/connectors-and-applications#payment-connectors) that process its future payments.
*
 */
type RecurringPayment struct {
	// Unique identifier of the RecurringPayment.
	ID string `json:"id"`
	// Current version of the RecurringPayment.
	Version int `json:"version"`
	// User-defined unique identifier of the RecurringPayment.
	Key *string `json:"key,omitempty"`
	// [RecurringOrder](ctp:api:type:RecurringOrder) whose future payments are processed using Checkout.
	RecurringOrder RecurringOrderReference `json:"recurringOrder"`
	// PaymentMethod and Connector used to pay the [RecurringOrder](ctp:api:type:RecurringOrder). Checkout only supports processing this array with one PaymentMethodConfiguration.
	PaymentMethodConfigurations []PaymentMethodConfiguration `json:"paymentMethodConfigurations"`
	// Date and time (UTC) the RecurringPayment was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the RecurringPayment was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

/**
*	Draft type to create a [RecurringPayment](ctp:checkout:type:RecurringPayment).
*
 */
type RecurringPaymentDraft struct {
	// User-defined unique identifier of the RecurringPayment.
	Key *string `json:"key,omitempty"`
	// [RecurringOrder](ctp:api:type:RecurringOrder) whose future payments must be processed using Checkout.
	RecurringOrder RecurringOrderReference `json:"recurringOrder"`
	// PaymentMethod and Connector to use to pay the [RecurringOrder](ctp:api:type:RecurringOrder). Checkout only supports processing this array with one PaymentMethodConfiguration.
	PaymentMethodConfigurations []PaymentMethodConfiguration `json:"paymentMethodConfigurations"`
}

/**
*	Reference to a [RecurringPayment](ctp:checkout:type:RecurringPayment).
*
 */
type RecurringPaymentReference struct {
	// Unique identifier of the referenced [RecurringPayment](ctp:checkout:type:RecurringPayment).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringPaymentReference) MarshalJSON() ([]byte, error) {
	type Alias RecurringPaymentReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurring-payment", Alias: (*Alias)(&obj)})
}

/**
*	Base type for all RecurringPayment update actions.
*
 */
type RecurringPaymentUpdateAction interface{}

func mapDiscriminatorRecurringPaymentUpdateAction(input interface{}) (RecurringPaymentUpdateAction, error) {
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
	case "addPaymentMethodConfiguration":
		obj := RecurringPaymentAddPaymentMethodConfigurationUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := RecurringPaymentSetKeyUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPaymentMethodConfiguration":
		obj := RecurringPaymentSetPaymentMethodConfigurationUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRecurringOrder":
		obj := RecurringPaymentSetRecurringOrderUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Adds a PaymentMethodConfiguration to a RecurringPayment. Checkout only supports processing a RecurringPayment with one PaymentMethodConfiguration.
*
 */
type RecurringPaymentAddPaymentMethodConfigurationUpdateAction struct {
	// PaymentMethod and Connector to add.
	PaymentMethodConfiguration PaymentMethodConfiguration `json:"paymentMethodConfiguration"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringPaymentAddPaymentMethodConfigurationUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringPaymentAddPaymentMethodConfigurationUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPaymentMethodConfiguration", Alias: (*Alias)(&obj)})
}

/**
*	Sets or unsets the key of a RecurringPayment.
*
 */
type RecurringPaymentSetKeyUpdateAction struct {
	// Key to set. If omitted, any existing value is removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringPaymentSetKeyUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringPaymentSetKeyUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	Sets the PaymentMethodConfigurations of a RecurringPayment, replacing any existing ones. Checkout only supports processing this array with one PaymentMethodConfiguration.
*
 */
type RecurringPaymentSetPaymentMethodConfigurationUpdateAction struct {
	// PaymentMethod and Connector to set.
	PaymentMethodConfigurations []PaymentMethodConfiguration `json:"paymentMethodConfigurations"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringPaymentSetPaymentMethodConfigurationUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringPaymentSetPaymentMethodConfigurationUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPaymentMethodConfiguration", Alias: (*Alias)(&obj)})
}

/**
*	Sets the [RecurringOrder](ctp:api:type:RecurringOrder) of a RecurringPayment.
*
 */
type RecurringPaymentSetRecurringOrderUpdateAction struct {
	// [RecurringOrder](ctp:api:type:RecurringOrder) to set.
	RecurringOrder RecurringOrderReference `json:"recurringOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringPaymentSetRecurringOrderUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringPaymentSetRecurringOrderUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRecurringOrder", Alias: (*Alias)(&obj)})
}

/**
*	Request body to update a [RecurringPayment](ctp:checkout:type:RecurringPayment).
*
 */
type RecurringPaymentUpdateActions struct {
	// Expected version of the RecurringPayment on which the changes should be applied. If the expected version does not match the actual version, a [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the RecurringPayment.
	Actions []RecurringPaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurringPaymentUpdateActions) UnmarshalJSON(data []byte) error {
	type Alias RecurringPaymentUpdateActions
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorRecurringPaymentUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
