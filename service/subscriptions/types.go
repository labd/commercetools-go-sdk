package subscriptions

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Subscription allows you to be notified of new messages or changes via a
// Message Queue of your choice.
type Subscription struct {
	ID             string                `json:"id"`
	Version        int                   `json:"version"`
	Key            string                `json:"key"`
	Destination    Destination           `json:"destination"`
	Messages       []MessageSubscription `json:"messages,omitempty"`
	Changes        []ChangeSubscription  `json:"changes,omitempty"`
	CreatedAt      time.Time             `json:"createdAt"`
	LastModifiedAt time.Time             `json:"lastModifiedAt"`
}

// UnmarshalJSON override to map the destination to the corresponding struct.
func (s *Subscription) UnmarshalJSON(data []byte) error {
	type Alias Subscription
	if err := json.Unmarshal(data, (*Alias)(s)); err != nil {
		return err
	}
	s.Destination = destinationMapping(s.Destination)
	return nil
}

// SubscriptionDraft is used to create new subscriptions. It is a stripped
// version of the regular Subscription struct.
type SubscriptionDraft struct {
	Key         string                `json:"key"`
	Destination Destination           `json:"destination"`
	Messages    []MessageSubscription `json:"messages"`
	Changes     []ChangeSubscription  `json:"changes,omitempty"`
}

// MessageSubscription allows you to subscribe to messages on a specific
// resource type and type combination. You can for example subscribe on the
// resourceType 'product' to the type 'ProductPublished'.
type MessageSubscription struct {
	ResourceTypeID string   `json:"resourceTypeId"`
	Types          []string `json:"types,omitempty"`
}

// ChangeSubscription allows you to subscribe to changes on a specific resource
// type.
type ChangeSubscription struct {
	ResourceTypeID string `json:"resourceTypeId"`
}

// DestinationAWSSQS is for the AWS SQS as a destination. AWS SQS is a
// pull-queue on AWS. We support the Standard queue type (the FIFO queue type is
// not supported).
//
// The queue needs to exist beforehand. It is recommended to create an IAM user
// with an accessKey and accessSecret pair specifically for each subscription
// that only has the sqs:SendMessage permission on this queue.
type DestinationAWSSQS struct {
	QueueURL     string `json:"queueUrl"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	Region       string `json:"region"`
}

// MarshalJSON override to add the type value
func (sd DestinationAWSSQS) MarshalJSON() ([]byte, error) {
	type Alias DestinationAWSSQS

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "SQS",
		Alias: (*Alias)(&sd),
	})
}

// DestinationAWSSNS is for the AWS SNS destination. AWS SNS can be used to push
// messages to AWS Lambda, HTTP endpoints (webhooks) or fan-out messages to SQS
// queues.
//
// The topic needs to exist beforehand. It is recommended to create an IAM user
// with an accessKey and accessSecret pair specifically for each subscription
// that only has the sns:Publish permission on this topic.
type DestinationAWSSNS struct {
	TopicArn     string `json:"topicArn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

// MarshalJSON override to add the type value
func (sd DestinationAWSSNS) MarshalJSON() ([]byte, error) {
	type Alias DestinationAWSSNS

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "SNS",
		Alias: (*Alias)(&sd),
	})
}

// DestinationIronMQ is for the IronMQ destination. IronMQ can be used both as a
// pull-queue, and to push messages to IronWorkers or HTTP endpoints (webhooks)
// or fan-out messages to other IronMQ queues.
//
// The webhook URI must contain a valid OAuth token. It roughly looks like this:
// https://...iron.io/.../queues/.../webhook?oauth=....
type DestinationIronMQ struct {
	URI string `json:"uri"`
}

// MarshalJSON override to add the Type() value
func (sd DestinationIronMQ) MarshalJSON() ([]byte, error) {
	type Alias DestinationIronMQ

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "IronMQ",
		Alias: (*Alias)(&sd),
	})
}

// DestinationGooglePubSub Google Cloud Pub/Sub can be used both as a pull-queue,
// and to push messages to e.g. Google Cloud Functions or HTTP endpoints
// (webhooks).
type DestinationGooglePubSub struct {
	ProjectID string `json:"projectId"`
	Topic     string `json:"topic"`
}

// MarshalJSON override to add the Type() value
func (sd DestinationGooglePubSub) MarshalJSON() ([]byte, error) {
	type Alias DestinationGooglePubSub

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "GoogleCloudPubSub",
		Alias: (*Alias)(&sd),
	})
}

// Destination contains all info necessary for the commercetools platform to
// deliver a message onto your Message Queue. Message Queues can be
// differentiated by the type field.
type Destination interface{}

func destinationMapping(input Destination) Destination {
	DestinationType := input.(map[string]interface{})["type"]
	switch DestinationType {
	case "IronMQ":
		new := DestinationIronMQ{}
		mapstructure.Decode(input, &new)
		return new
	case "SNS":
		new := DestinationAWSSNS{}
		mapstructure.Decode(input, &new)
		return new
	case "SQS":
		new := DestinationAWSSQS{}
		mapstructure.Decode(input, &new)
		return new
	case "GoogleCloudPubSub":
		new := DestinationGooglePubSub{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
