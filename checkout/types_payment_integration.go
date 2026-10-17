package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Information regarding IDs and references that created the Payment Integration.
*
 */
type CreatedBy struct {
	// ID of the API client that created the PaymentIntegration.
	ClientId *string `json:"clientId,omitempty"`
	// External user ID that created the PaymentIntegration.
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Customer reference if created by a customer.
	Customer Reference `json:"customer,omitempty"`
	// Anonymous ID if created by an anonymous user.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CreatedBy) UnmarshalJSON(data []byte) error {
	type Alias CreatedBy
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Customer != nil {
		var err error
		obj.Customer, err = mapDiscriminatorReference(obj.Customer)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Information regarding IDs and references that last modified the Payment Integration.
*
 */
type LastModifiedBy struct {
	// ID of the API client that last modified the PaymentIntegration.
	ClientId *string `json:"clientId,omitempty"`
	// Session ID associated with the modification of the PaymentIntegration.
	SessionId *string `json:"sessionId,omitempty"`
	// External user ID that last modified the PaymentIntegration.
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Anonymous ID if the resource was last modified by an anonymous user.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Customer reference or customer ID if the resource was last modified by a customer.
	Customer Reference `json:"customer,omitempty"`
	// Associate (customer) reference if the resource was last modified by an associate.
	Associate Reference `json:"associate,omitempty"`
	// User reference if the resource was last modified by a platform user.
	User Reference `json:"user,omitempty"`
	// Indicates whether the resource was last modified by a platform client.
	IsPlatformClient *bool `json:"isPlatformClient,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LastModifiedBy) UnmarshalJSON(data []byte) error {
	type Alias LastModifiedBy
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Customer != nil {
		var err error
		obj.Customer, err = mapDiscriminatorReference(obj.Customer)
		if err != nil {
			return err
		}
	}
	if obj.Associate != nil {
		var err error
		obj.Associate, err = mapDiscriminatorReference(obj.Associate)
		if err != nil {
			return err
		}
	}
	if obj.User != nil {
		var err error
		obj.User, err = mapDiscriminatorReference(obj.User)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Configuration for automated reversal of payments.
*
 */
type AutomatedReversalConfiguration struct {
	// Indicates whether the automated reversal configuration is active or not.
	Status string `json:"status"`
	// Cart-based predicate in JSONata format to determine when automated reversal should be applied.
	Predicate *string `json:"predicate,omitempty"`
}

/**
*	Reference to a ConnectorDeployment for the payment integration.
*
 */
type ConnectorDeploymentReference struct {
	// Unique identifier of the referenced ConnectorDeployment.
	ID string `json:"id"`
	// Type identifier, always `deployment` for ConnectorDeployment references.
	TypeId string `json:"typeId"`
}

/**
*	Display information for the payment integration in the UI.
*
 */
type DisplayInfo struct {
	// Label of the payment method displayed in the UI.
	Label *LocalizedString `json:"label,omitempty"`
	// Logo of the payment method displayed in the UI.
	LogoUrl *LocalizedUrl `json:"logoUrl,omitempty"`
	// Description of the payment method displayed in the UI.
	Description *LocalizedString `json:"description,omitempty"`
	// Text of the Pay button displayed in the UI when paying with the payment method.
	PayButtonText *LocalizedString `json:"payButtonText,omitempty"`
}

/**
*	Paginated result containing PaymentIntegrations.
*
 */
type PaginatedPaymentIntegration struct {
	// Number of results requested.
	Limit int `json:"limit"`
	// Number of elements skipped.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	Total int `json:"total"`
	// PaymentIntegrations matching the query.
	Results []PaymentIntegration `json:"results"`
}

type PaymentComponentType string

const (
	PaymentComponentTypeComponent PaymentComponentType = "Component"
	PaymentComponentTypeDropIn    PaymentComponentType = "DropIn"
	PaymentComponentTypeExpress   PaymentComponentType = "Express"
)

type PaymentIntegration struct {
	// Unique identifier of the PaymentIntegration.
	ID string `json:"id"`
	// Current version of the PaymentIntegration.
	Version int `json:"version"`
	// User-defined unique identifier of the PaymentIntegration. `MinLength: 2 ​MaxLength: 256 ​Pattern: ^[A-Za-z0-9_-]+$`
	Key *string `json:"key,omitempty"`
	// Reference to an Application associated with the PaymentIntegration.
	Application ApplicationReference `json:"application"`
	// Indicates whether the Payment Integration is active or not.
	Status *PaymentIntegrationStatus `json:"status,omitempty"`
	// Type of the PaymentIntegration, for example, card, Paypal, or Applepay. This field must be same as used by the payment service provider (PSP) and must be supported by the Connector.
	//
	// Use `drop-in` for drop-in Payment Integrations.
	Type string `json:"type"`
	// Name of the PaymentIntegration in the Merchant Center.
	Name string `json:"name"`
	// Predicate in JSONata format. Applies only when `componentType` is `Component`.
	Predicate *string `json:"predicate,omitempty"`
	// Indicates the component type of Payment Integration.
	ComponentType PaymentComponentType `json:"componentType"`
	// Connector deployment info for the Payment Integration.
	ConnectorDeployment ConnectorDeploymentReference `json:"connectorDeployment"`
	// Customization for the information that is displayed in the UI. Applies only when `componentType` is `Component`
	DisplayInfo *DisplayInfo `json:"displayInfo,omitempty"`
	// Sorting information of the Payment Integration.
	SortingInfo *SortingInfo `json:"sortingInfo,omitempty"`
	// Configuration of the Payment Integration automatic reversal.
	AutomatedReversalConfiguration *AutomatedReversalConfiguration `json:"automatedReversalConfiguration,omitempty"`
	// Date and time (UTC) the PaymentIntegration was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// IDs and references that created the PaymentIntegration.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Date and time (UTC) the PaymentIntegration was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the PaymentIntegration.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
}

/**
*	Reference to a [Payment Integration](ctp:checkout:type:PaymentIntegration).
*
 */
type PaymentIntegrationReference struct {
	// Unique identifier of the referenced Payment Integration.
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment-integration", Alias: (*Alias)(&obj)})
}

/**
*	Resource identifier to a [Payment Integration](ctp:checkout:type:PaymentIntegration). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](ctp:api:type:InvalidJsonInputError) error is returned.
*
 */
type PaymentIntegrationResourceIdentifier struct {
	// Unique identifier of the referenced Payment Integration. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced Payment Integration. Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntegrationResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntegrationResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment-integration", Alias: (*Alias)(&obj)})
}

type PaymentIntegrationStatus string

const (
	PaymentIntegrationStatusActive   PaymentIntegrationStatus = "Active"
	PaymentIntegrationStatusInactive PaymentIntegrationStatus = "Inactive"
)

type PaymentIntegrationType string

const (
	PaymentIntegrationTypePsp      PaymentIntegrationType = "Psp"
	PaymentIntegrationTypeGiftCard PaymentIntegrationType = "GiftCard"
)

type SortingInfo struct {
	// Sorting priority of the Payment Integration.
	Priority int `json:"priority"`
}

/**
*	Base type for all PaymentIntegration update actions.
*
 */
type PaymentIntegrationUpdateAction interface{}

func mapDiscriminatorPaymentIntegrationUpdateAction(input interface{}) (PaymentIntegrationUpdateAction, error) {
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
	case "setAutomatedReversalConfigurationPredicate":
		obj := SetAutomatedReversalConfigurationPredicateUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAutomatedReversalConfigurationStatus":
		obj := SetAutomatedReversalConfigurationStatusUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAutomatedReversalConfiguration":
		obj := SetAutomatedReversalConfigurationUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDisplayInfoDescription":
		obj := SetDisplayInfoDescriptionUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDisplayInfoLabel":
		obj := SetDisplayInfoLabelUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDisplayInfoLogoUrl":
		obj := SetDisplayInfoLogoUrlUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDisplayInfoPayButtonText":
		obj := SetDisplayInfoPayButtonTextUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDisplayInfo":
		obj := SetDisplayInfoUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := SetKeyUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := SetNameUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPredicate":
		obj := SetPredicateUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSortingInfo":
		obj := SetSortingInfoUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatus":
		obj := SetStatusUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setType":
		obj := SetTypeUpdateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type PaymentIntegrationUpdateActions struct {
	// Expected version of the PaymentIntegration on which the changes should be applied. If the expected version does not match the actual version, a [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the PaymentIntegration.
	Actions []PaymentIntegrationUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentIntegrationUpdateActions) UnmarshalJSON(data []byte) error {
	type Alias PaymentIntegrationUpdateActions
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorPaymentIntegrationUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Sets the automated reversal configuration predicate of a PaymentIntegration.
*
 */
type SetAutomatedReversalConfigurationPredicateUpdateAction struct {
	// If omitted, any existing value is removed.
	Predicate *string `json:"predicate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAutomatedReversalConfigurationPredicateUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetAutomatedReversalConfigurationPredicateUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAutomatedReversalConfigurationPredicate", Alias: (*Alias)(&obj)})
}

/**
*	Sets the automated reversal configuration status of a PaymentIntegration.
*
 */
type SetAutomatedReversalConfigurationStatusUpdateAction struct {
	// Value to set as the automated reversal configuration status of the PaymentIntegration.
	Status string `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAutomatedReversalConfigurationStatusUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetAutomatedReversalConfigurationStatusUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAutomatedReversalConfigurationStatus", Alias: (*Alias)(&obj)})
}

/**
*	Sets the automated reversal configuration of a PaymentIntegration.
*
 */
type SetAutomatedReversalConfigurationUpdateAction struct {
	// Value to set as the automated reversal configuration of the PaymentIntegration.
	AutomatedReversalConfiguration *AutomatedReversalConfiguration `json:"automatedReversalConfiguration,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetAutomatedReversalConfigurationUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetAutomatedReversalConfigurationUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAutomatedReversalConfiguration", Alias: (*Alias)(&obj)})
}

/**
*	Sets the display info description of a PaymentIntegration.
*
 */
type SetDisplayInfoDescriptionUpdateAction struct {
	// If omitted, any existing value is removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDisplayInfoDescriptionUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDisplayInfoDescriptionUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDisplayInfoDescription", Alias: (*Alias)(&obj)})
}

/**
*	Sets the display info label of a PaymentIntegration.
*
 */
type SetDisplayInfoLabelUpdateAction struct {
	// If omitted, any existing value is removed.
	Label *LocalizedString `json:"label,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDisplayInfoLabelUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDisplayInfoLabelUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDisplayInfoLabel", Alias: (*Alias)(&obj)})
}

/**
*	Sets the display info logo URL of a PaymentIntegration.
*
 */
type SetDisplayInfoLogoUrlUpdateAction struct {
	// If omitted, any existing value is removed.
	LogoUrl *LocalizedString `json:"logoUrl,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDisplayInfoLogoUrlUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDisplayInfoLogoUrlUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDisplayInfoLogoUrl", Alias: (*Alias)(&obj)})
}

/**
*	Sets the display info pay button text of a PaymentIntegration.
*
 */
type SetDisplayInfoPayButtonTextUpdateAction struct {
	// If omitted, any existing value is removed.
	PayButtonText *LocalizedString `json:"payButtonText,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDisplayInfoPayButtonTextUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDisplayInfoPayButtonTextUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDisplayInfoPayButtonText", Alias: (*Alias)(&obj)})
}

type SetDisplayInfoUpdateAction struct {
	// If omitted, any existing value is removed.
	DisplayInfo *DisplayInfo `json:"displayInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetDisplayInfoUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetDisplayInfoUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDisplayInfo", Alias: (*Alias)(&obj)})
}

/**
*	Sets or unsets the key of a PaymentIntegration.
*
 */
type SetKeyUpdateAction struct {
	// If omitted, any existing value is removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetKeyUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetKeyUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	Sets the name of a PaymentIntegration.
*
 */
type SetNameUpdateAction struct {
	// Value to set as the name of the PaymentIntegration.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetNameUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetNameUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Sets the predicate of a PaymentIntegration.
*
 */
type SetPredicateUpdateAction struct {
	// If omitted, any existing value is removed.
	Predicate *string `json:"predicate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetPredicateUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetPredicateUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}

/**
*	Sets the sorting info of a PaymentIntegration.
*
 */
type SetSortingInfoUpdateAction struct {
	// If omitted, any existing value is removed.
	SortingInfo *SortingInfo `json:"sortingInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetSortingInfoUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetSortingInfoUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSortingInfo", Alias: (*Alias)(&obj)})
}

/**
*	Sets the status of a PaymentIntegration.
*
 */
type SetStatusUpdateAction struct {
	// Value to set as the status of the PaymentIntegration.
	Status PaymentIntegrationStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetStatusUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetStatusUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatus", Alias: (*Alias)(&obj)})
}

/**
*	Sets the type of a PaymentIntegration.
*
 */
type SetTypeUpdateAction struct {
	// Value to set as the type of the PaymentIntegration.
	Type string `json:"type"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SetTypeUpdateAction) MarshalJSON() ([]byte, error) {
	type Alias SetTypeUpdateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setType", Alias: (*Alias)(&obj)})
}
