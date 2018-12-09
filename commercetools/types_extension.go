// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Extension struct {
	Version        int                  `json:"version"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
	ID             string               `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	Triggers       []ExtensionTrigger   `json:"triggers"`
	Key            string               `json:"key,omitempty"`
	Destination    ExtensionDestination `json:"destination"`
}

func (obj *Extension) UnmarshalJSON(data []byte) error {
	type Alias Extension
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = AbstractExtensionDestinationDiscriminatorMapping(obj.Destination)
	}

	return nil
}

type ExtensionAWSLambdaDestination struct {
	Arn          string `json:"arn"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

func (obj ExtensionAWSLambdaDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAWSLambdaDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AWSLambda", Alias: (*Alias)(&obj)})
}

type ExtensionAction string

const (
	ExtensionActionCreate ExtensionAction = "Create"
	ExtensionActionUpdate ExtensionAction = "Update"
)

type ExtensionAuthorizationHeaderAuthentication struct {
	HeaderValue string `json:"headerValue"`
}

func (obj ExtensionAuthorizationHeaderAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAuthorizationHeaderAuthentication
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AuthorizationHeader", Alias: (*Alias)(&obj)})
}

type ExtensionAzureFunctionsAuthentication struct {
	Key string `json:"key"`
}

func (obj ExtensionAzureFunctionsAuthentication) MarshalJSON() ([]byte, error) {
	type Alias ExtensionAzureFunctionsAuthentication
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AzureFunctions", Alias: (*Alias)(&obj)})
}

type ExtensionChangeDestinationAction struct {
	Destination ExtensionDestination `json:"destination"`
}

func (obj ExtensionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}
func (obj *ExtensionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias ExtensionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = AbstractExtensionDestinationDiscriminatorMapping(obj.Destination)
	}

	return nil
}

type ExtensionChangeTriggersAction struct {
	Triggers []ExtensionTrigger `json:"triggers"`
}

func (obj ExtensionChangeTriggersAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeTriggersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTriggers", Alias: (*Alias)(&obj)})
}

type ExtensionDestination interface{}
type AbstractExtensionDestination struct{}

func AbstractExtensionDestinationDiscriminatorMapping(input ExtensionDestination) ExtensionDestination {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "AWSLambda":
		new := ExtensionAWSLambdaDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "HTTP":
		new := ExtensionHttpDestination{}
		mapstructure.Decode(input, &new)
		if new.Authentication != nil {
			new.Authentication = AbstractExtensionHttpDestinationAuthenticationDiscriminatorMapping(new.Authentication)
		}

		return new
	}
	return nil
}

type ExtensionDraft struct {
	Triggers    []ExtensionTrigger   `json:"triggers"`
	Key         string               `json:"key,omitempty"`
	Destination ExtensionDestination `json:"destination"`
}

func (obj *ExtensionDraft) UnmarshalJSON(data []byte) error {
	type Alias ExtensionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = AbstractExtensionDestinationDiscriminatorMapping(obj.Destination)
	}

	return nil
}

type ExtensionHttpDestination struct {
	URL            string                                 `json:"url"`
	Authentication ExtensionHttpDestinationAuthentication `json:"authentication,omitempty"`
}

func (obj ExtensionHttpDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionHttpDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "HTTP", Alias: (*Alias)(&obj)})
}
func (obj *ExtensionHttpDestination) UnmarshalJSON(data []byte) error {
	type Alias ExtensionHttpDestination
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Authentication != nil {
		obj.Authentication = AbstractExtensionHttpDestinationAuthenticationDiscriminatorMapping(obj.Authentication)
	}

	return nil
}

type ExtensionHttpDestinationAuthentication interface{}
type AbstractExtensionHttpDestinationAuthentication struct{}

func AbstractExtensionHttpDestinationAuthenticationDiscriminatorMapping(input ExtensionHttpDestinationAuthentication) ExtensionHttpDestinationAuthentication {
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

type ExtensionInput struct {
	Resource Reference       `json:"resource"`
	Action   ExtensionAction `json:"action"`
}

func (obj *ExtensionInput) UnmarshalJSON(data []byte) error {
	type Alias ExtensionInput
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		obj.Resource = AbstractReferenceDiscriminatorMapping(obj.Resource)
	}

	return nil
}

type ExtensionPagedQueryResponse struct {
	Total   int         `json:"total,omitempty"`
	Offset  int         `json:"offset"`
	Count   int         `json:"count"`
	Results []Extension `json:"results"`
}
type ExtensionResourceTypeID string

const (
	ExtensionResourceTypeIDCart     ExtensionResourceTypeID = "cart"
	ExtensionResourceTypeIDOrder    ExtensionResourceTypeID = "order"
	ExtensionResourceTypeIDPayment  ExtensionResourceTypeID = "payment"
	ExtensionResourceTypeIDCustomer ExtensionResourceTypeID = "customer"
)

type ExtensionSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj ExtensionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ExtensionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ExtensionTrigger struct {
	ResourceTypeID ExtensionResourceTypeID `json:"resourceTypeId"`
	Actions        []ExtensionAction       `json:"actions"`
}

type ExtensionUpdate struct {
	Version int                     `json:"version"`
	Actions []ExtensionUpdateAction `json:"actions"`
}

func (obj *ExtensionUpdate) UnmarshalJSON(data []byte) error {
	type Alias ExtensionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractExtensionUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ExtensionUpdateAction interface{}
type AbstractExtensionUpdateAction struct{}

func AbstractExtensionUpdateActionDiscriminatorMapping(input ExtensionUpdateAction) ExtensionUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeDestination":
		new := ExtensionChangeDestinationAction{}
		mapstructure.Decode(input, &new)
		if new.Destination != nil {
			new.Destination = AbstractExtensionDestinationDiscriminatorMapping(new.Destination)
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
