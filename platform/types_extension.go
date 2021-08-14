// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Extension struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy   *CreatedBy           `json:"createdBy,omitempty"`
	Key         *string              `json:"key,omitempty"`
	Destination ExtensionDestination `json:"destination"`
	Triggers    []ExtensionTrigger   `json:"triggers"`
	// The maximum time the commercetools platform waits for a response from the extension.
	// If not present, `2000` (2 seconds) is used.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Extension) UnmarshalJSON(data []byte) error {
	type Alias Extension
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	return nil
}

type ExtensionAction string

const (
	ExtensionActionCreate ExtensionAction = "Create"
	ExtensionActionUpdate ExtensionAction = "Update"
)

type ExtensionDestination interface{}

func mapDiscriminatorExtensionDestination(input interface{}) (ExtensionDestination, error) {

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
	case "AWSLambda":
		obj := ExtensionAWSLambdaDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "HTTP":
		obj := ExtensionHttpDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Authentication != nil {
			var err error
			obj.Authentication, err = mapDiscriminatorExtensionHttpDestinationAuthentication(obj.Authentication)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

type ExtensionAWSLambdaDestination struct {
	Arn          string `json:"arn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAWSLambdaDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAWSLambdaDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AWSLambda", Alias: (*Alias)(&obj)})
}

type ExtensionDraft struct {
	// User-specific unique identifier for the extension
	Key *string `json:"key,omitempty"`
	// Details where the extension can be reached
	Destination ExtensionDestination `json:"destination"`
	// Describes what triggers the extension
	Triggers []ExtensionTrigger `json:"triggers"`
	// The maximum time the commercetools platform waits for a response from the extension.
	// The maximum value is 2000 ms (2 seconds).
	// This limit can be increased per project after we review the performance impact.
	// Please contact Support via the [Support Portal](https://support.commercetools.com) and provide the region, project key and use case.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionDraft) UnmarshalJSON(data []byte) error {
	type Alias ExtensionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	return nil
}

type ExtensionHttpDestination struct {
	Url            string                                 `json:"url"`
	Authentication ExtensionHttpDestinationAuthentication `json:"authentication,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionHttpDestination) UnmarshalJSON(data []byte) error {
	type Alias ExtensionHttpDestination
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Authentication != nil {
		var err error
		obj.Authentication, err = mapDiscriminatorExtensionHttpDestinationAuthentication(obj.Authentication)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionHttpDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionHttpDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "HTTP", Alias: (*Alias)(&obj)})
}

type ExtensionHttpDestinationAuthentication interface{}

func mapDiscriminatorExtensionHttpDestinationAuthentication(input interface{}) (ExtensionHttpDestinationAuthentication, error) {

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
	case "AuthorizationHeader":
		obj := ExtensionAuthorizationHeaderAuthentication{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AzureFunctions":
		obj := ExtensionAzureFunctionsAuthentication{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ExtensionAuthorizationHeaderAuthentication struct {
	HeaderValue string `json:"headerValue"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAuthorizationHeaderAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAuthorizationHeaderAuthentication
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AuthorizationHeader", Alias: (*Alias)(&obj)})
}

type ExtensionAzureFunctionsAuthentication struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAzureFunctionsAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAzureFunctionsAuthentication
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AzureFunctions", Alias: (*Alias)(&obj)})
}

type ExtensionInput struct {
	Action   ExtensionAction `json:"action"`
	Resource Reference       `json:"resource"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionInput) UnmarshalJSON(data []byte) error {
	type Alias ExtensionInput
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		var err error
		obj.Resource, err = mapDiscriminatorReference(obj.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

type ExtensionPagedQueryResponse struct {
	Limit   int         `json:"limit"`
	Count   int         `json:"count"`
	Total   *int        `json:"total,omitempty"`
	Offset  int         `json:"offset"`
	Results []Extension `json:"results"`
}

type ExtensionResourceTypeId string

const (
	ExtensionResourceTypeIdCart     ExtensionResourceTypeId = "cart"
	ExtensionResourceTypeIdOrder    ExtensionResourceTypeId = "order"
	ExtensionResourceTypeIdPayment  ExtensionResourceTypeId = "payment"
	ExtensionResourceTypeIdCustomer ExtensionResourceTypeId = "customer"
)

type ExtensionTrigger struct {
	ResourceTypeId ExtensionResourceTypeId `json:"resourceTypeId"`
	Actions        []ExtensionAction       `json:"actions"`
}

type ExtensionUpdate struct {
	Version int                     `json:"version"`
	Actions []ExtensionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionUpdate) UnmarshalJSON(data []byte) error {
	type Alias ExtensionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ExtensionUpdateAction interface{}

func mapDiscriminatorExtensionUpdateAction(input interface{}) (ExtensionUpdateAction, error) {

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
	case "changeDestination":
		obj := ExtensionChangeDestinationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Destination != nil {
			var err error
			obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "changeTriggers":
		obj := ExtensionChangeTriggersAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ExtensionSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTimeoutInMs":
		obj := ExtensionSetTimeoutInMsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ExtensionChangeDestinationAction struct {
	Destination ExtensionDestination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias ExtensionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorExtensionDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

type ExtensionChangeTriggersAction struct {
	Triggers []ExtensionTrigger `json:"triggers"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionChangeTriggersAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeTriggersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTriggers", Alias: (*Alias)(&obj)})
}

type ExtensionSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ExtensionSetTimeoutInMsAction struct {
	// The maximum time the commercetools platform waits for a response from the extension.
	// The maximum value is 2000 ms (2 seconds).
	// This limit can be increased per project after we review the performance impact.
	// Please contact Support via the support and provide the region, project key and use case.
	TimeoutInMs *int `json:"timeoutInMs,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionSetTimeoutInMsAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetTimeoutInMsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTimeoutInMs", Alias: (*Alias)(&obj)})
}
