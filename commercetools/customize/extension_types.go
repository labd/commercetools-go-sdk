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
// An encrypted https connection is strongly recommended for production setups,
// but we accept unencrypted http connections for development purposes. HTTP
// redirects will not be followed. Cache headers will be ignored.
type ExtensionDestinationHTTP struct {
	URL string `json:"url"`
}

func (ed ExtensionDestinationHTTP) MarshalJSON() ([]byte, error) {
	type Alias ExtensionDestinationHTTP
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "HTTP",
		Alias: (*Alias)(&ed),
	})
}

// ExtensionDestinationAWSLambda uses AWS Lambda as type for the
// ExtensionDestination.
//
// It is recommended to create an IAM user with an accessKey and accessSecret
// pair specifically for each API Extension that only has the
// lambda:InvokeFunction permission on this function.
//
// Please note that AWS Lambda limits the size of the payload to 6 MB (this
// limit also applies if the Lambda function is invoked by the API Gateway). You
// should not use AWS Lambda if you anticipate that your API Extensions will
// receive JSON input exceeding 6 MB.
type ExtensionDestinationAWSLambda struct {
	ARN          string `json:"arn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

func (ed ExtensionDestinationAWSLambda) MarshalJSON() ([]byte, error) {
	type Alias ExtensionDestinationAWSLambda
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "AWSLambda",
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

// ExtensionChangeTriggers is used to update an existing API Extension with
// a new triggers.
type ExtensionChangeTriggers struct {
	Messages []ExtensionTrigger `json:"messages"`
}

func (ua ExtensionChangeTriggers) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeTriggers

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeTriggers",
		Alias:  (*Alias)(&ua),
	})
}

// ExtensionChangeDestination is used to update an existing API Extension with
// a new destination.
type ExtensionChangeDestination struct {
	Destination ExtensionDestination `json:"destination"`
}

func (ua ExtensionChangeDestination) MarshalJSON() ([]byte, error) {
	type Alias ExtensionChangeDestination

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeDestination",
		Alias:  (*Alias)(&ua),
	})
}
