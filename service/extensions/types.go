package extensions

import (
	"encoding/json"
	"time"
)

// Extension extends the behavior of an API with your business logic.
type Extension struct {
	ID             string      `json:"id"`
	Version        int         `json:"version"`
	Key            string      `json:"key"`
	Destination    Destination `json:"destination"`
	Triggers       []Trigger   `json:"triggers"`
	CreatedAt      time.Time   `json:"createdAt"`
	LastModifiedAt time.Time   `json:"lastModifiedAt"`
}

// UnmarshalJSON override to map the destination to the corresponding struct.
func (e *Extension) UnmarshalJSON(data []byte) error {
	type ExtensionClone Extension
	if err := json.Unmarshal(data, (*ExtensionClone)(e)); err != nil {
		return err
	}
	e.Destination = destinationMapping(e.Destination)
	return nil
}

// ExtensionDraft is given as payload for Create Extension requests.
type ExtensionDraft struct {
	Key         string      `json:"key"`
	Destination Destination `json:"destination"`
	Triggers    []Trigger   `json:"triggers"`
}

// Destination contains all info necessary for the commercetools platform to
// call the extension. Destinations can be differentiated by the type field.
type Destination interface{}

// DestinationHTTP implementation An encrypted https connection is strongly
// recommended for production setups, but we accept unencrypted http connections
// for development purposes. HTTP redirects will not be followed. Cache headers
// will be ignored.
type DestinationHTTP struct {
	URL            string                    `json:"url"`
	Authentication DestinationAuthentication `json:"authentication"`
}

// MarshalJSON override to add the type value.
func (ed DestinationHTTP) MarshalJSON() ([]byte, error) {
	type Alias DestinationHTTP
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "HTTP",
		Alias: (*Alias)(&ed),
	})
}

// DestinationAWSLambda uses AWS Lambda as type for the Destination.
//
// It is recommended to create an IAM user with an accessKey and accessSecret
// pair specifically for each API Extension that only has the
// lambda:InvokeFunction permission on this function.
//
// Please note that AWS Lambda limits the size of the payload to 6 MB (this
// limit also applies if the Lambda function is invoked by the API Gateway). You
// should not use AWS Lambda if you anticipate that your API Extensions will
// receive JSON input exceeding 6 MB.
type DestinationAWSLambda struct {
	ARN          string `json:"arn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

// MarshalJSON override to add the type value.
func (ed DestinationAWSLambda) MarshalJSON() ([]byte, error) {
	type Alias DestinationAWSLambda
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "AWSLambda",
		Alias: (*Alias)(&ed),
	})
}

type DestinationAuthentication interface{}

type DestinationAuthenticationAzure struct {
	Key string `json:"key"`
}

func (ed DestinationAuthenticationAzure) MarshalJSON() ([]byte, error) {
	type Alias DestinationAuthenticationAzure
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "AzureFunctions",
		Alias: (*Alias)(&ed),
	})
}

type DestinationAuthenticationAuth struct {
	HeaderValue string `json:"headerValue"`
}

func (ed DestinationAuthenticationAuth) MarshalJSON() ([]byte, error) {
	type Alias DestinationAuthenticationAuth
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "AuthorizationHeader",
		Alias: (*Alias)(&ed),
	})
}

type Trigger struct {
	ResourceTypeID string   `json:"resourceTypeId"`
	Actions        []string `json:"actions"`
}
