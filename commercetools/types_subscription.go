// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type AzureEventGridDestination struct {
	Uri       string `json:"uri"`
	AccessKey string `json:"accessKey"`
}

func (obj AzureEventGridDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureEventGridDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "EventGrid", Alias: (*Alias)(&obj)})
}

type AzureServiceBusDestination struct {
	ConnectionString string `json:"connectionString"`
}

func (obj AzureServiceBusDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureServiceBusDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "AzureServiceBus", Alias: (*Alias)(&obj)})
}

type ChangeSubscription struct {
	ResourceTypeID string `json:"resourceTypeId"`
}

type DeliveryCloudEventsFormat struct {
	CloudEventsVersion string `json:"cloudEventsVersion"`
}

func (obj DeliveryCloudEventsFormat) MarshalJSON() ([]byte, error) {
	type Alias DeliveryCloudEventsFormat
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CloudEvents", Alias: (*Alias)(&obj)})
}

type DeliveryFormat interface{}
type AbstractDeliveryFormat struct{}

func AbstractDeliveryFormatDiscriminatorMapping(input DeliveryFormat) DeliveryFormat {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "CloudEvents":
		new := DeliveryCloudEventsFormat{}
		mapstructure.Decode(input, &new)
		return new
	case "Platform":
		new := DeliveryPlatformFormat{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type DeliveryPlatformFormat struct{}

func (obj DeliveryPlatformFormat) MarshalJSON() ([]byte, error) {
	type Alias DeliveryPlatformFormat
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Platform", Alias: (*Alias)(&obj)})
}

type Destination interface{}
type AbstractDestination struct{}

func AbstractDestinationDiscriminatorMapping(input Destination) Destination {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "EventGrid":
		new := AzureEventGridDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "AzureServiceBus":
		new := AzureServiceBusDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "GoogleCloudPubSub":
		new := GoogleCloudPubSubDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "IronMQ":
		new := IronMqDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "SNS":
		new := SnsDestination{}
		mapstructure.Decode(input, &new)
		return new
	case "SQS":
		new := SqsDestination{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type GoogleCloudPubSubDestination struct {
	Topic     string `json:"topic"`
	ProjectID string `json:"projectId"`
}

func (obj GoogleCloudPubSubDestination) MarshalJSON() ([]byte, error) {
	type Alias GoogleCloudPubSubDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "GoogleCloudPubSub", Alias: (*Alias)(&obj)})
}

type IronMqDestination struct {
	Uri string `json:"uri"`
}

func (obj IronMqDestination) MarshalJSON() ([]byte, error) {
	type Alias IronMqDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "IronMQ", Alias: (*Alias)(&obj)})
}

type MessageDelivery struct {
	Resource           Reference           `json:"resource"`
	ProjectKey         string              `json:"projectKey"`
	Version            int                 `json:"version"`
	SequenceNumber     int                 `json:"sequenceNumber"`
	ResourceVersion    int                 `json:"resourceVersion"`
	PayloadNotIncluded *PayloadNotIncluded `json:"payloadNotIncluded"`
	LastModifiedAt     time.Time           `json:"lastModifiedAt"`
	ID                 string              `json:"id"`
	CreatedAt          time.Time           `json:"createdAt"`
}

func (obj MessageDelivery) MarshalJSON() ([]byte, error) {
	type Alias MessageDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "Message", Alias: (*Alias)(&obj)})
}

type MessageSubscription struct {
	Types          []string `json:"types,omitempty"`
	ResourceTypeID string   `json:"resourceTypeId"`
}

type PayloadNotIncluded struct {
	Reason      string `json:"reason"`
	PayloadType string `json:"payloadType"`
}

type ResourceCreatedDelivery struct {
	Resource   Reference `json:"resource"`
	ProjectKey string    `json:"projectKey"`
	Version    int       `json:"version"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func (obj ResourceCreatedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceCreatedDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "ResourceCreated", Alias: (*Alias)(&obj)})
}

type ResourceDeletedDelivery struct {
	Resource   Reference `json:"resource"`
	ProjectKey string    `json:"projectKey"`
	Version    int       `json:"version"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func (obj ResourceDeletedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceDeletedDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "ResourceDeleted", Alias: (*Alias)(&obj)})
}

type ResourceUpdatedDelivery struct {
	Resource   Reference `json:"resource"`
	ProjectKey string    `json:"projectKey"`
	Version    int       `json:"version"`
	OldVersion int       `json:"oldVersion"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func (obj ResourceUpdatedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceUpdatedDelivery
	return json.Marshal(struct {
		NotificationType string `json:"notificationType"`
		*Alias
	}{NotificationType: "ResourceUpdated", Alias: (*Alias)(&obj)})
}

type SnsDestination struct {
	TopicArn     string `json:"topicArn"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

func (obj SnsDestination) MarshalJSON() ([]byte, error) {
	type Alias SnsDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "SNS", Alias: (*Alias)(&obj)})
}

type SqsDestination struct {
	Region       string `json:"region"`
	QueueURL     string `json:"queueUrl"`
	AccessSecret string `json:"accessSecret"`
	AccessKey    string `json:"accessKey"`
}

func (obj SqsDestination) MarshalJSON() ([]byte, error) {
	type Alias SqsDestination
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "SQS", Alias: (*Alias)(&obj)})
}

type Subscription struct {
	Version        int                   `json:"version"`
	LastModifiedAt time.Time             `json:"lastModifiedAt"`
	ID             string                `json:"id"`
	CreatedAt      time.Time             `json:"createdAt"`
	Messages       []MessageSubscription `json:"messages"`
	Key            string                `json:"key,omitempty"`
	Format         DeliveryFormat        `json:"format"`
	Destination    Destination           `json:"destination"`
	Changes        []ChangeSubscription  `json:"changes"`
}

func (obj *Subscription) UnmarshalJSON(data []byte) error {
	type Alias Subscription
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = AbstractDestinationDiscriminatorMapping(obj.Destination)
	}
	if obj.Format != nil {
		obj.Format = AbstractDeliveryFormatDiscriminatorMapping(obj.Format)
	}

	return nil
}

type SubscriptionChangeDestinationAction struct {
	Destination Destination `json:"destination"`
}

func (obj SubscriptionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}
func (obj *SubscriptionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = AbstractDestinationDiscriminatorMapping(obj.Destination)
	}

	return nil
}

type SubscriptionDelivery interface{}
type AbstractSubscriptionDelivery struct {
	Resource   Reference `json:"resource"`
	ProjectKey string    `json:"projectKey"`
}

func AbstractSubscriptionDeliveryDiscriminatorMapping(input SubscriptionDelivery) SubscriptionDelivery {
	discriminator := input.(map[string]interface{})["notificationType"]
	switch discriminator {
	case "Message":
		new := MessageDelivery{}
		mapstructure.Decode(input, &new)
		return new
	case "ResourceCreated":
		new := ResourceCreatedDelivery{}
		mapstructure.Decode(input, &new)
		return new
	case "ResourceDeleted":
		new := ResourceDeletedDelivery{}
		mapstructure.Decode(input, &new)
		return new
	case "ResourceUpdated":
		new := ResourceUpdatedDelivery{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
func (obj *AbstractSubscriptionDelivery) UnmarshalJSON(data []byte) error {
	type Alias AbstractSubscriptionDelivery
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Resource != nil {
		obj.Resource = AbstractReferenceDiscriminatorMapping(obj.Resource)
	}

	return nil
}

type SubscriptionDraft struct {
	Messages    []MessageSubscription `json:"messages,omitempty"`
	Key         string                `json:"key,omitempty"`
	Format      DeliveryFormat        `json:"format,omitempty"`
	Destination Destination           `json:"destination"`
	Changes     []ChangeSubscription  `json:"changes,omitempty"`
}

func (obj *SubscriptionDraft) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		obj.Destination = AbstractDestinationDiscriminatorMapping(obj.Destination)
	}
	if obj.Format != nil {
		obj.Format = AbstractDeliveryFormatDiscriminatorMapping(obj.Format)
	}

	return nil
}

type SubscriptionPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []Subscription `json:"results"`
}

type SubscriptionSetChangesAction struct {
	Changes []ChangeSubscription `json:"changes,omitempty"`
}

func (obj SubscriptionSetChangesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setChanges", Alias: (*Alias)(&obj)})
}

type SubscriptionSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj SubscriptionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type SubscriptionSetMessagesAction struct {
	Messages []MessageSubscription `json:"messages,omitempty"`
}

func (obj SubscriptionSetMessagesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetMessagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMessages", Alias: (*Alias)(&obj)})
}

type SubscriptionUpdate struct {
	Version int                        `json:"version"`
	Actions []SubscriptionUpdateAction `json:"actions"`
}

func (obj *SubscriptionUpdate) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractSubscriptionUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type SubscriptionUpdateAction interface{}
type AbstractSubscriptionUpdateAction struct{}

func AbstractSubscriptionUpdateActionDiscriminatorMapping(input SubscriptionUpdateAction) SubscriptionUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeDestination":
		new := SubscriptionChangeDestinationAction{}
		mapstructure.Decode(input, &new)
		if new.Destination != nil {
			new.Destination = AbstractDestinationDiscriminatorMapping(new.Destination)
		}

		return new
	case "setChanges":
		new := SubscriptionSetChangesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := SubscriptionSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMessages":
		new := SubscriptionSetMessagesAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
