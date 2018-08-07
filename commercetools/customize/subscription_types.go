package customize

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Subscription struct {
	ID             string                  `json:"id"`
	Version        int                     `json:"version"`
	Key            string                  `json:"key"`
	Destination    SubscriptionDestination `json:"destination"`
	Messages       []MessageSubscription   `json:"messages,omitempty"`
	Changes        []ChangeSubscription    `json:"changes,omitempty"`
	CreatedAt      time.Time               `json:"createdAt"`
	LastModifiedAt time.Time               `json:"lastModifiedAt"`
}

func (s *Subscription) UnmarshalJSON(data []byte) error {
	type Alias Subscription
	if err := json.Unmarshal(data, (*Alias)(s)); err != nil {
		return err
	}
	s.Destination = subscriptionDestinationMapping(s.Destination)
	return nil
}

// SubscriptionDraft is used to create new subscriptions. It is a stripped
// version of the regular Subscription struct.
type SubscriptionDraft struct {
	Key         string                  `json:"key"`
	Destination SubscriptionDestination `json:"destination"`
	Messages    []MessageSubscription   `json:"messages"`
	Changes     []ChangeSubscription    `json:"changes,omitempty"`
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

// SubscriptionAWSSQSDestination is for the AWS SQS as a destination. AWS SQS is
// a pull-queue on AWS. We support the Standard queue type (the FIFO queue type
// is not supported).
//
// The queue needs to exist beforehand. It is recommended to create an IAM user
// with an accessKey and accessSecret pair specifically for each subscription
// that only has the sqs:SendMessage permission on this queue.
type SubscriptionAWSSQSDestination struct {
	QueueURL     string `json:"queueUrl"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	Region       string `json:"region"`
}

func (sd SubscriptionAWSSQSDestination) Type() string {
	return "SQS"
}

// MarshalJSON override to add the Type() value
func (sd SubscriptionAWSSQSDestination) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionAWSSQSDestination

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  sd.Type(),
		Alias: (*Alias)(&sd),
	})
}

// SubscriptionAWSSNSDestination is for the AWS SNS destination. AWS SNS can be
// used to push messages to AWS Lambda, HTTP endpoints (webhooks) or fan-out
// messages to SQS queues.
//
// The topic needs to exist beforehand. It is recommended to create an IAM user
// with an accessKey and accessSecret pair specifically for each subscription
// that only has the sns:Publish permission on this topic.
type SubscriptionAWSSNSDestination struct {
	TopicArn     string `json:"topicArn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

func (sd SubscriptionAWSSNSDestination) Type() string {
	return "SNS"
}

// MarshalJSON override to add the Type() value
func (sd SubscriptionAWSSNSDestination) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionAWSSNSDestination

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  sd.Type(),
		Alias: (*Alias)(&sd),
	})
}

// SubscriptionIronMQDestination is for the IronMQ destination. IronMQ can be
// used both as a pull-queue, and to push messages to IronWorkers or HTTP
// endpoints (webhooks) or fan-out messages to other IronMQ queues.
//
// The webhook URI must contain a valid OAuth token. It roughly looks like this:
// https://...iron.io/.../queues/.../webhook?oauth=....
type SubscriptionIronMQDestination struct {
	URI string `json:"uri"`
}

func (sd SubscriptionIronMQDestination) Type() string {
	return "IronMQ"
}

// MarshalJSON override to add the Type() value
func (sd SubscriptionIronMQDestination) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionIronMQDestination

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  sd.Type(),
		Alias: (*Alias)(&sd),
	})
}

type SubscriptionDestination interface{}

func subscriptionDestinationMapping(input SubscriptionDestination) SubscriptionDestination {
	DestinationType := input.(map[string]interface{})["type"]
	switch DestinationType {
	case "IronMQ":
		new := SubscriptionIronMQDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "SNS":
		new := SubscriptionAWSSNSDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "SQS":
		new := SubscriptionAWSSQSDestination{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type SubscriptionSetKey struct {
	Key string `json:"key,omitempty"`
}

func (ua SubscriptionSetKey) Action() string {
	return "setKey"
}

// MarshalJSON override to add the Action() value
func (ua SubscriptionSetKey) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

type SubscriptionSetMessages struct {
	Messages []MessageSubscription `json:"messages"`
}

func (ua SubscriptionSetMessages) Action() string {
	return "setMessages"
}

// MarshalJSON override to add the Action() value
func (ua SubscriptionSetMessages) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetMessages

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

type SubscriptionSetChanges struct {
	Changes []ChangeSubscription `json:"changes"`
}

func (ua SubscriptionSetChanges) Action() string {
	return "setChanges"
}

// MarshalJSON override to add the Action() value
func (ua SubscriptionSetChanges) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetChanges

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}
