// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ExtensionAction is an enum type
type ExtensionAction string

// Enum values for ExtensionAction
const (
	ExtensionActionCreate ExtensionAction = "Create"
	ExtensionActionUpdate ExtensionAction = "Update"
)

// ExtensionResourceTypeID is an enum type
type ExtensionResourceTypeID string

// Enum values for ExtensionResourceTypeID
const (
	ExtensionResourceTypeIDCart     ExtensionResourceTypeID = "cart"
	ExtensionResourceTypeIDOrder    ExtensionResourceTypeID = "order"
	ExtensionResourceTypeIDPayment  ExtensionResourceTypeID = "payment"
	ExtensionResourceTypeIDCustomer ExtensionResourceTypeID = "customer"
)

// ExtensionDestination uses type as discriminator attribute
type ExtensionDestination interface{}

func mapDiscriminatorExtensionDestination(input ExtensionDestination) ExtensionDestination {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "AWSLambda":
		new := ExtensionAWSLambdaDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "HTTP":
		new := ExtensionHTTPDestination{}
		mapstructure.Decode(input, &new)
		if new.Authentication != nil {
			new.Authentication = mapDiscriminatorExtensionHTTPDestinationAuthentication(new.Authentication)
		}

		return new
	}
	return nil
}

// ExtensionHTTPDestinationAuthentication uses type as discriminator attribute
type ExtensionHTTPDestinationAuthentication interface{}

func mapDiscriminatorExtensionHTTPDestinationAuthentication(input ExtensionHTTPDestinationAuthentication) ExtensionHTTPDestinationAuthentication {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "AuthorizationHeader":
		new := ExtensionAuthorizationHeaderAuthentication{}
		mapstructure.Decode(input, &new)
		return new
	case "AzureFunctions":
		new := ExtensionAzureFunctionsAuthentication{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// ExtensionUpdateAction uses action as discriminator attribute
type ExtensionUpdateAction interface{}

func mapDiscriminatorExtensionUpdateAction(input ExtensionUpdateAction) ExtensionUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeDestination":
		new := ExtensionChangeDestinationAction{}
		mapstructure.Decode(input, &new)
		if new.Destination != nil {
			new.Destination = mapDiscriminatorExtensionDestination(new.Destination)
		}

		return new
	case "changeTriggers":
		new := ExtensionChangeTriggersAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ExtensionSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// Extension is of type Resource
type Extension struct {
	Version        int                  `json:"version"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
	ID             string               `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	Triggers       []ExtensionTrigger   `json:"triggers"`
	Key            string               `json:"key,omitempty"`
	Destination    ExtensionDestination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Extension) UnmarshalJSON(data []byte) error {
	type Alias Extension
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = mapDiscriminatorExtensionDestination(obj.Destination)
	}

	return nil
}

// ExtensionAWSLambdaDestination implements the interface ExtensionDestination
type ExtensionAWSLambdaDestination struct {
	Arn          string `json:"arn"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAWSLambdaDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAWSLambdaDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AWSLambda", Alias: (*Alias)(&obj)})
}

// ExtensionAuthorizationHeaderAuthentication implements the interface ExtensionHTTPDestinationAuthentication
type ExtensionAuthorizationHeaderAuthentication struct {
	HeaderValue string `json:"headerValue"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAuthorizationHeaderAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAuthorizationHeaderAuthentication
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AuthorizationHeader", Alias: (*Alias)(&obj)})
}

// ExtensionAzureFunctionsAuthentication implements the interface ExtensionHTTPDestinationAuthentication
type ExtensionAzureFunctionsAuthentication struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionAzureFunctionsAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAzureFunctionsAuthentication
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AzureFunctions", Alias: (*Alias)(&obj)})
}

// ExtensionChangeDestinationAction implements the interface ExtensionUpdateAction
type ExtensionChangeDestinationAction struct {
	Destination ExtensionDestination `json:"destination"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias ExtensionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = mapDiscriminatorExtensionDestination(obj.Destination)
	}

	return nil
}

// ExtensionChangeTriggersAction implements the interface ExtensionUpdateAction
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

// ExtensionDraft is a standalone struct
type ExtensionDraft struct {
	Triggers    []ExtensionTrigger   `json:"triggers"`
	Key         string               `json:"key,omitempty"`
	Destination ExtensionDestination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionDraft) UnmarshalJSON(data []byte) error {
	type Alias ExtensionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = mapDiscriminatorExtensionDestination(obj.Destination)
	}

	return nil
}

// ExtensionHTTPDestination implements the interface ExtensionDestination
type ExtensionHTTPDestination struct {
	URL            string                                 `json:"url"`
	Authentication ExtensionHTTPDestinationAuthentication `json:"authentication,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionHTTPDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionHTTPDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "HTTP", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionHTTPDestination) UnmarshalJSON(data []byte) error {
	type Alias ExtensionHTTPDestination
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Authentication != nil {
		obj.Authentication = mapDiscriminatorExtensionHTTPDestinationAuthentication(obj.Authentication)
	}

	return nil
}

// ExtensionInput is a standalone struct
type ExtensionInput struct {
	Resource Reference       `json:"resource"`
	Action   ExtensionAction `json:"action"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ExtensionInput) UnmarshalJSON(data []byte) error {
	type Alias ExtensionInput
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		obj.Resource = mapDiscriminatorReference(obj.Resource)
	}

	return nil
}

// ExtensionPagedQueryResponse is of type PagedQueryResponse
type ExtensionPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
	Results []Extension `json:"results"`
}

// ExtensionSetKeyAction implements the interface ExtensionUpdateAction
type ExtensionSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ExtensionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ExtensionTrigger is a standalone struct
type ExtensionTrigger struct {
	ResourceTypeID ExtensionResourceTypeID `json:"resourceTypeId"`
	Actions        []ExtensionAction       `json:"actions"`
}

// ExtensionUpdate is of type Update
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
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorExtensionUpdateAction(obj.Actions[i])
	}

	return nil
}
