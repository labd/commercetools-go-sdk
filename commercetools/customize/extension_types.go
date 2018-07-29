package customize

import (
	"encoding/json"
	"time"
)

type Extension struct {
	ID             string               `json:"id"`
	Version        int                  `json:"version"`
	Key            string               `json:"key"`
	Destination    ExtensionDestination `json:"destination"`
	Triggers       []ExtensionTrigger   `json:"messages"`
	CreatedAt      time.Time            `json:"createdAt"`
	LastModifiedAt time.Time            `json:"lastModifiedAt"`
}

func (e *Extension) UnmarshalJSON(data []byte) error {
	type ExtensionClone Extension
	if err := json.Unmarshal(data, (*ExtensionClone)(e)); err != nil {
		return err
	}
	e.Destination = ExtensionDestinationMapping(e.Destination)
	return nil
}

type ExtensionDraft struct {
	Key         string               `json:"key"`
	Destination ExtensionDestination `json:"destination"`
	Triggers    []ExtensionTrigger   `json:"messages"`
}

type ExtensionDestination interface {
}

// ExtensionDestinationHTTP implementation

type ExtensionDestinationHTTP struct {
	URL string `json:"url"`
}

func (ed *ExtensionDestinationHTTP) Type() string {
	return "HTTP"
}

func (ed ExtensionDestinationHTTP) MarshalJSON() ([]byte, error) {
	type Alias ExtensionDestinationHTTP
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  ed.Type(),
		Alias: (*Alias)(&ed),
	})
}

// ExtensionDestinationAWSLambda implementation

type ExtensionDestinationAWSLambda struct {
	ARN          string `json:"arn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

func (ed *ExtensionDestinationAWSLambda) Type() string {
	return "AWSLambda"
}

func (ed ExtensionDestinationAWSLambda) MarshalJSON() ([]byte, error) {
	type Alias ExtensionDestinationAWSLambda
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  ed.Type(),
		Alias: (*Alias)(&ed),
	})
}

type ExtensionDestinationAuthentication struct {
}

type ExtensionDestinationAuthenticationAzure struct{}

func (ed *ExtensionDestinationAuthenticationAzure) Type() string {
	return "AzureFunctions"
}

type ExtensionDestinationAuthenticationAuth struct {
}

func (ed *ExtensionDestinationAuthenticationAuth) Type() string {
	return "AuthorizationHeader"
}

type ExtensionTrigger struct {
	ResourceTypeID string   `json:"resourceTypeId"`
	Actions        []string `json:"actions"`
}

type ExtensionChangeTriggers struct {
	Messages []ExtensionTrigger `json:"messages"`
}

func (ua ExtensionChangeTriggers) Action() string {
	return "changeTriggers"
}

func (ua ExtensionChangeTriggers) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeTriggers

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

type ExtensionChangeDestination struct {
	Destination ExtensionDestination `json:"destination"`
}

func (ua ExtensionChangeDestination) Action() string {
	return "changeDestination"
}

func (ua ExtensionChangeDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestination

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}
