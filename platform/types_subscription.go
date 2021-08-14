// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type ChangeSubscription struct {
	ResourceTypeId string `json:"resourceTypeId"`
}

type DeliveryFormat interface{}

func mapDiscriminatorDeliveryFormat(input interface{}) (DeliveryFormat, error) {

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
	case "CloudEvents":
		new := DeliveryCloudEventsFormat{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Platform":
		new := DeliveryPlatformFormat{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type DeliveryCloudEventsFormat struct {
	CloudEventsVersion string `json:"cloudEventsVersion"`
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryCloudEventsFormat) MarshalJSON() ([]byte, error) {
	type Alias DeliveryCloudEventsFormat
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CloudEvents", Alias: (*Alias)(&obj)})
}

type DeliveryPlatformFormat struct {
}

// MarshalJSON override to set the discriminator value
func (obj DeliveryPlatformFormat) MarshalJSON() ([]byte, error) {
	type Alias DeliveryPlatformFormat
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Platform", Alias: (*Alias)(&obj)})
}

type Destination interface{}

func mapDiscriminatorDestination(input interface{}) (Destination, error) {

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
	case "EventGrid":
		new := AzureEventGridDestination{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "AzureServiceBus":
		new := AzureServiceBusDestination{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "GoogleCloudPubSub":
		new := GoogleCloudPubSubDestination{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "IronMQ":
		new := IronMqDestination{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SNS":
		new := SnsDestination{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "SQS":
		new := SqsDestination{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type AzureEventGridDestination struct {
	Uri       string `json:"uri"`
	AccessKey string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value
func (obj AzureEventGridDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureEventGridDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "EventGrid", Alias: (*Alias)(&obj)})
}

type AzureServiceBusDestination struct {
	ConnectionString string `json:"connectionString"`
}

// MarshalJSON override to set the discriminator value
func (obj AzureServiceBusDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureServiceBusDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AzureServiceBus", Alias: (*Alias)(&obj)})
}

type GoogleCloudPubSubDestination struct {
	ProjectId string `json:"projectId"`
	Topic     string `json:"topic"`
}

// MarshalJSON override to set the discriminator value
func (obj GoogleCloudPubSubDestination) MarshalJSON() ([]byte, error) {
	type Alias GoogleCloudPubSubDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "GoogleCloudPubSub", Alias: (*Alias)(&obj)})
}

type IronMqDestination struct {
	Uri string `json:"uri"`
}

// MarshalJSON override to set the discriminator value
func (obj IronMqDestination) MarshalJSON() ([]byte, error) {
	type Alias IronMqDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "IronMQ", Alias: (*Alias)(&obj)})
}

type MessageSubscription struct {
	ResourceTypeId string   `json:"resourceTypeId"`
	Types          []string `json:"types,omitEmpty"`
}

type PayloadNotIncluded struct {
	Reason      string `json:"reason"`
	PayloadType string `json:"payloadType"`
}

type SnsDestination struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	TopicArn     string `json:"topicArn"`
}

// MarshalJSON override to set the discriminator value
func (obj SnsDestination) MarshalJSON() ([]byte, error) {
	type Alias SnsDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SNS", Alias: (*Alias)(&obj)})
}

type SqsDestination struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	QueueUrl     string `json:"queueUrl"`
	Region       string `json:"region"`
}

// MarshalJSON override to set the discriminator value
func (obj SqsDestination) MarshalJSON() ([]byte, error) {
	type Alias SqsDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SQS", Alias: (*Alias)(&obj)})
}

type Subscription struct {
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy   *CreatedBy               `json:"createdBy,omitEmpty"`
	Changes     []ChangeSubscription     `json:"changes"`
	Destination Destination              `json:"destination"`
	Key         string                   `json:"key,omitEmpty"`
	Messages    []MessageSubscription    `json:"messages"`
	Format      DeliveryFormat           `json:"format"`
	Status      SubscriptionHealthStatus `json:"status"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Subscription) UnmarshalJSON(data []byte) error {
	type Alias Subscription
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	if obj.Format != nil {
		var err error
		obj.Format, err = mapDiscriminatorDeliveryFormat(obj.Format)
		if err != nil {
			return err
		}
	}
	return nil
}

type SubscriptionDelivery interface{}

func mapDiscriminatorSubscriptionDelivery(input interface{}) (SubscriptionDelivery, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["notificationType"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'notificationType'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "Message":
		new := MessageDelivery{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceCreated":
		new := ResourceCreatedDelivery{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceDeleted":
		new := ResourceDeletedDelivery{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ResourceUpdated":
		new := ResourceUpdatedDelivery{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type MessageDelivery struct {
	ProjectKey                      string                   `json:"projectKey"`
	Resource                        Reference                `json:"resource"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Id                              string                   `json:"id"`
	Version                         int                      `json:"version"`
	CreatedAt                       time.Time                `json:"createdAt"`
	LastModifiedAt                  time.Time                `json:"lastModifiedAt"`
	SequenceNumber                  int                      `json:"sequenceNumber"`
	ResourceVersion                 int                      `json:"resourceVersion"`
	PayloadNotIncluded              PayloadNotIncluded       `json:"payloadNotIncluded"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MessageDelivery) UnmarshalJSON(data []byte) error {
	type Alias MessageDelivery
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

// MarshalJSON override to set the discriminator value
func (obj MessageDelivery) MarshalJSON() ([]byte, error) {
	type Alias MessageDelivery
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "Message", Alias: (*Alias)(&obj)})
}

type ResourceCreatedDelivery struct {
	ProjectKey                      string                   `json:"projectKey"`
	Resource                        Reference                `json:"resource"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Version                         int                      `json:"version"`
	ModifiedAt                      time.Time                `json:"modifiedAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceCreatedDelivery) UnmarshalJSON(data []byte) error {
	type Alias ResourceCreatedDelivery
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

// MarshalJSON override to set the discriminator value
func (obj ResourceCreatedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceCreatedDelivery
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "ResourceCreated", Alias: (*Alias)(&obj)})
}

type ResourceDeletedDelivery struct {
	ProjectKey                      string                   `json:"projectKey"`
	Resource                        Reference                `json:"resource"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Version                         int                      `json:"version"`
	ModifiedAt                      time.Time                `json:"modifiedAt"`
	DataErasure                     bool                     `json:"dataErasure,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceDeletedDelivery) UnmarshalJSON(data []byte) error {
	type Alias ResourceDeletedDelivery
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

// MarshalJSON override to set the discriminator value
func (obj ResourceDeletedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceDeletedDelivery
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "ResourceDeleted", Alias: (*Alias)(&obj)})
}

type ResourceUpdatedDelivery struct {
	ProjectKey                      string                   `json:"projectKey"`
	Resource                        Reference                `json:"resource"`
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitEmpty"`
	Version                         int                      `json:"version"`
	OldVersion                      int                      `json:"oldVersion"`
	ModifiedAt                      time.Time                `json:"modifiedAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceUpdatedDelivery) UnmarshalJSON(data []byte) error {
	type Alias ResourceUpdatedDelivery
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

// MarshalJSON override to set the discriminator value
func (obj ResourceUpdatedDelivery) MarshalJSON() ([]byte, error) {
	type Alias ResourceUpdatedDelivery
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "ResourceUpdated", Alias: (*Alias)(&obj)})
}

type SubscriptionDraft struct {
	Changes     []ChangeSubscription  `json:"changes,omitEmpty"`
	Destination Destination           `json:"destination"`
	Key         string                `json:"key,omitEmpty"`
	Messages    []MessageSubscription `json:"messages,omitEmpty"`
	Format      DeliveryFormat        `json:"format,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionDraft) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	if obj.Format != nil {
		var err error
		obj.Format, err = mapDiscriminatorDeliveryFormat(obj.Format)
		if err != nil {
			return err
		}
	}
	return nil
}

type SubscriptionHealthStatus string

const (
	SubscriptionHealthStatusHealthy                           SubscriptionHealthStatus = "Healthy"
	SubscriptionHealthStatusConfigurationError                SubscriptionHealthStatus = "ConfigurationError"
	SubscriptionHealthStatusConfigurationErrorDeliveryStopped SubscriptionHealthStatus = "ConfigurationErrorDeliveryStopped"
	SubscriptionHealthStatusTemporaryError                    SubscriptionHealthStatus = "TemporaryError"
)

type SubscriptionPagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   int            `json:"total,omitEmpty"`
	Offset  int            `json:"offset"`
	Results []Subscription `json:"results"`
}

type SubscriptionUpdate struct {
	Version int                        `json:"version"`
	Actions []SubscriptionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionUpdate) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type SubscriptionUpdateAction interface{}

func mapDiscriminatorSubscriptionUpdateAction(input interface{}) (SubscriptionUpdateAction, error) {

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
		new := SubscriptionChangeDestinationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setChanges":
		new := SubscriptionSetChangesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := SubscriptionSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMessages":
		new := SubscriptionSetMessagesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type SubscriptionChangeDestinationAction struct {
	Destination Destination `json:"destination"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionChangeDestinationAction) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionChangeDestinationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Destination != nil {
		var err error
		obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

type SubscriptionSetChangesAction struct {
	Changes []ChangeSubscription `json:"changes,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionSetChangesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setChanges", Alias: (*Alias)(&obj)})
}

type SubscriptionSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type SubscriptionSetMessagesAction struct {
	Messages []MessageSubscription `json:"messages,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj SubscriptionSetMessagesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetMessagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMessages", Alias: (*Alias)(&obj)})
}
