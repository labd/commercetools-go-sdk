package extensions

import (
	"encoding/json"
	"time"
)

type Extension struct {
	ID             string      `json:"id"`
	Version        int         `json:"version"`
	Key            string      `json:"key"`
	Destination    Destination `json:"destination"`
	Triggers       []Trigger   `json:"messages"`
	CreatedAt      time.Time   `json:"createdAt"`
	LastModifiedAt time.Time   `json:"lastModifiedAt"`
}

func (e *Extension) UnmarshalJSON(data []byte) error {
	type ExtensionClone Extension
	if err := json.Unmarshal(data, (*ExtensionClone)(e)); err != nil {
		return err
	}
	e.Destination = destinationMapping(e.Destination)
	return nil
}

type ExtensionDraft struct {
	Key         string      `json:"key"`
	Destination Destination `json:"destination"`
	Triggers    []Trigger   `json:"messages"`
}

type Destination interface {
}

// DestinationHTTP implementation
// An encrypted https connection is strongly recommended for production setups,
// but we accept unencrypted http connections for development purposes. HTTP
// redirects will not be followed. Cache headers will be ignored.
type DestinationHTTP struct {
	URL string `json:"url"`
}

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

type DestinationAuthentication struct{}

type DestinationAuthenticationAzure struct{}

func (ed *DestinationAuthenticationAzure) Type() string {
	return "AzureFunctions"
}

type DestinationAuthenticationAuth struct{}

func (ed *DestinationAuthenticationAuth) Type() string {
	return "AuthorizationHeader"
}

type Trigger struct {
	ResourceTypeID string   `json:"resourceTypeId"`
	Actions        []string `json:"actions"`
}
