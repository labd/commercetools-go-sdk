package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Defines the method of authentication for AWS SQS and SNS Destinations.
 */
type AwsAuthenticationMode string

const (
	AwsAuthenticationModeCredentials AwsAuthenticationMode = "Credentials"
	AwsAuthenticationModeIAM         AwsAuthenticationMode = "IAM"
)

/**
*	Notification about changes to a resource. The payload format differs for resource [creation](ctp:api:type:ResourceCreatedDeliveryPayload),
*	[update](ctp:api:type:ResourceUpdatedDeliveryPayload),
*	and [deletion](ctp:api:type:ResourceDeletedDeliveryPayload).
*
 */
type ChangeSubscription struct {
	// Unique identifier for the type of resource, for example, `cart`.
	ResourceTypeId ChangeSubscriptionResourceTypeId `json:"resourceTypeId"`
}

/**
*	Resource types supported by [ChangeSubscriptions](ctp:api:type:ChangeSubscription):
*
 */
type ChangeSubscriptionResourceTypeId string

const (
	ChangeSubscriptionResourceTypeIdApprovalFlow          ChangeSubscriptionResourceTypeId = "approval-flow"
	ChangeSubscriptionResourceTypeIdApprovalRule          ChangeSubscriptionResourceTypeId = "approval-rule"
	ChangeSubscriptionResourceTypeIdAssociateRole         ChangeSubscriptionResourceTypeId = "associate-role"
	ChangeSubscriptionResourceTypeIdBusinessUnit          ChangeSubscriptionResourceTypeId = "business-unit"
	ChangeSubscriptionResourceTypeIdCart                  ChangeSubscriptionResourceTypeId = "cart"
	ChangeSubscriptionResourceTypeIdCartDiscount          ChangeSubscriptionResourceTypeId = "cart-discount"
	ChangeSubscriptionResourceTypeIdCategory              ChangeSubscriptionResourceTypeId = "category"
	ChangeSubscriptionResourceTypeIdChannel               ChangeSubscriptionResourceTypeId = "channel"
	ChangeSubscriptionResourceTypeIdCustomer              ChangeSubscriptionResourceTypeId = "customer"
	ChangeSubscriptionResourceTypeIdCustomerEmailToken    ChangeSubscriptionResourceTypeId = "customer-email-token"
	ChangeSubscriptionResourceTypeIdCustomerGroup         ChangeSubscriptionResourceTypeId = "customer-group"
	ChangeSubscriptionResourceTypeIdCustomerPasswordToken ChangeSubscriptionResourceTypeId = "customer-password-token"
	ChangeSubscriptionResourceTypeIdDiscountCode          ChangeSubscriptionResourceTypeId = "discount-code"
	ChangeSubscriptionResourceTypeIdExtension             ChangeSubscriptionResourceTypeId = "extension"
	ChangeSubscriptionResourceTypeIdInventoryEntry        ChangeSubscriptionResourceTypeId = "inventory-entry"
	ChangeSubscriptionResourceTypeIdKeyValueDocument      ChangeSubscriptionResourceTypeId = "key-value-document"
	ChangeSubscriptionResourceTypeIdOrder                 ChangeSubscriptionResourceTypeId = "order"
	ChangeSubscriptionResourceTypeIdOrderEdit             ChangeSubscriptionResourceTypeId = "order-edit"
	ChangeSubscriptionResourceTypeIdPayment               ChangeSubscriptionResourceTypeId = "payment"
	ChangeSubscriptionResourceTypeIdProduct               ChangeSubscriptionResourceTypeId = "product"
	ChangeSubscriptionResourceTypeIdProductDiscount       ChangeSubscriptionResourceTypeId = "product-discount"
	ChangeSubscriptionResourceTypeIdProductPrice          ChangeSubscriptionResourceTypeId = "product-price"
	ChangeSubscriptionResourceTypeIdProductSelection      ChangeSubscriptionResourceTypeId = "product-selection"
	ChangeSubscriptionResourceTypeIdProductType           ChangeSubscriptionResourceTypeId = "product-type"
	ChangeSubscriptionResourceTypeIdQuote                 ChangeSubscriptionResourceTypeId = "quote"
	ChangeSubscriptionResourceTypeIdQuoteRequest          ChangeSubscriptionResourceTypeId = "quote-request"
	ChangeSubscriptionResourceTypeIdReview                ChangeSubscriptionResourceTypeId = "review"
	ChangeSubscriptionResourceTypeIdShippingMethod        ChangeSubscriptionResourceTypeId = "shipping-method"
	ChangeSubscriptionResourceTypeIdShoppingList          ChangeSubscriptionResourceTypeId = "shopping-list"
	ChangeSubscriptionResourceTypeIdStagedQuote           ChangeSubscriptionResourceTypeId = "staged-quote"
	ChangeSubscriptionResourceTypeIdStandalonePrice       ChangeSubscriptionResourceTypeId = "standalone-price"
	ChangeSubscriptionResourceTypeIdState                 ChangeSubscriptionResourceTypeId = "state"
	ChangeSubscriptionResourceTypeIdStore                 ChangeSubscriptionResourceTypeId = "store"
	ChangeSubscriptionResourceTypeIdSubscription          ChangeSubscriptionResourceTypeId = "subscription"
	ChangeSubscriptionResourceTypeIdTaxCategory           ChangeSubscriptionResourceTypeId = "tax-category"
	ChangeSubscriptionResourceTypeIdType                  ChangeSubscriptionResourceTypeId = "type"
	ChangeSubscriptionResourceTypeIdZone                  ChangeSubscriptionResourceTypeId = "zone"
)

/**
*	The [CloudEventsFormat](ctp:api:type:CloudEventsFormat) represents event data in a way that conforms to a common specification. The message payload can be found inside the `data` field.
*
 */
type CloudEventsPayload struct {
	// The version of the [CloudEvents](https://github.com/cloudevents/spec/blob/v1.0.2/cloudevents/spec.md) specification which the event uses.
	Specversion string `json:"specversion"`
	// Unique identifier of the event.
	ID string `json:"id"`
	// The `type` is namespaced with `com.commercetools`, followed by the [ReferenceTypeId](ctp:api:type:ReferenceTypeId), the type of Subscription (either `message` or `change`), and the message or change type.
	// For example, `com.commercetools.product.message.ProductPublished` or `com.commercetools.order.change.ResourceCreated`.
	Type string `json:"type"`
	// The default REST URI of the [ReferenceTypeId](ctp:api:type:ReferenceTypeId) that triggered this event, including the project key.
	Source string `json:"source"`
	// Unique identifier of the resource that triggered the event.
	Subject string `json:"subject"`
	// Corresponds to the `lastModifiedAt` of the resource at the time the event was triggered.
	Time time.Time `json:"time"`
	// Corresponds to the `sequenceNumber` of a [MessageSubscription](ctp:api:type:MessageSubscription). Can be used to process messages in the correct order.
	Sequence *string `json:"sequence,omitempty"`
	// `"Integer"`
	Sequencetype *string `json:"sequencetype,omitempty"`
	// The URI from which the message can be retrieved if messages are [enabled](/../api/projects/messages#enable-querying-messages-via-the-api). Only set for [MessageSubscriptions](ctp:api:type:MessageSubscription).
	Dataref *string `json:"dataref,omitempty"`
	// [MessageDeliveryPayload](ctp:api:type:MessageDeliveryPayload), [ResourceCreatedDeliveryPayload](ctp:api:type:ResourceCreatedDeliveryPayload), [ResourceUpdatedDeliveryPayload](ctp:api:type:ResourceUpdatedDeliveryPayload), or [ResourceDeletedDeliveryPayload](ctp:api:type:ResourceDeletedDeliveryPayload).
	Data DeliveryPayload `json:"data"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CloudEventsPayload) UnmarshalJSON(data []byte) error {
	type Alias CloudEventsPayload
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Data != nil {
		var err error
		obj.Data, err = mapDiscriminatorDeliveryPayload(obj.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CloudEventsPayload) MarshalJSON() ([]byte, error) {
	type Alias CloudEventsPayload
	return json.Marshal(struct {
		Action string `json:"null"`
		*Alias
	}{Action: "CloudEvents", Alias: (*Alias)(&obj)})
}

type DeliveryFormat interface{}

func mapDiscriminatorDeliveryFormat(input interface{}) (DeliveryFormat, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "CloudEvents":
		obj := CloudEventsFormat{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Platform":
		obj := PlatformFormat{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	The CloudEventsFormat can be used with any [Destination](#destination), and the payload is delivered in the `JSON Event Format`. [AzureEventGridDestination](ctp:api:type:AzureEventGridDestination) offers native support to filter and route CloudEvents.
*
 */
type CloudEventsFormat struct {
	CloudEventsVersion string `json:"cloudEventsVersion"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CloudEventsFormat) MarshalJSON() ([]byte, error) {
	type Alias CloudEventsFormat
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CloudEvents", Alias: (*Alias)(&obj)})
}

/**
*	All payloads for the [PlatformFormat](ctp:api:type:PlatformFormat) share these common fields.
*
 */
type DeliveryPayload interface{}

func mapDiscriminatorDeliveryPayload(input interface{}) (DeliveryPayload, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["notificationType"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'notificationType'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "Message":
		obj := MessageDeliveryPayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ResourceCreated":
		obj := ResourceCreatedDeliveryPayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ResourceDeleted":
		obj := ResourceDeletedDeliveryPayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "ResourceUpdated":
		obj := ResourceUpdatedDeliveryPayload{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Resource != nil {
			var err error
			obj.Resource, err = mapDiscriminatorReference(obj.Resource)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

type Destination interface{}

func mapDiscriminatorDestination(input interface{}) (Destination, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "EventGrid":
		obj := AzureEventGridDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "AzureServiceBus":
		obj := AzureServiceBusDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ConfluentCloud":
		obj := ConfluentCloudDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "EventBridge":
		obj := EventBridgeDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "GoogleCloudPubSub":
		obj := GoogleCloudPubSubDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "IronMQ":
		obj := IronMqDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SNS":
		obj := SnsDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "SQS":
		obj := SqsDestination{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	[Azure Event Grid](https://azure.microsoft.com/en-us/products/event-grid/) can be used to push messages to Azure Functions, HTTP endpoints (webhooks), and several other Azure tools. Event Grid can only be used with the [CloudEventsFormat](ctp:api:type:CloudEventsFormat).
*	To set up a Subscription with Azure Event Grid, first create a topic in the [Azure Portal](https://azure.microsoft.com/en-us/get-started/azure-portal/). To allow Composable Commerce to push messages to your topic, provide an [access key](https://docs.microsoft.com/en-us/azure/event-grid/get-access-keys).
*
 */
type AzureEventGridDestination struct {
	// URI of the topic.
	Uri string `json:"uri"`
	// Partially hidden on retrieval for security reasons.
	AccessKey string `json:"accessKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AzureEventGridDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureEventGridDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "EventGrid", Alias: (*Alias)(&obj)})
}

/**
*	[Azure Service Bus](https://azure.microsoft.com/en-us/products/service-bus/) can be used as a pull-queue with [Queues](https://learn.microsoft.com/en-us/azure/service-bus-messaging/service-bus-queues-topics-subscriptions#queues), or to fan-out messages with [Topics and Subscriptions](https://learn.microsoft.com/en-us/azure/service-bus-messaging/service-bus-queues-topics-subscriptions).
*	To set up a Subscription with Azure Service Bus, first create a queue/topic in the [Azure Portal](https://azure.microsoft.com/en-us/get-started/azure-portal/) with a Shared Access Policy including the `Send` permission.
*
 */
type AzureServiceBusDestination struct {
	// SharedAccessKey is partially hidden on retrieval for security reasons.
	ConnectionString string `json:"connectionString"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AzureServiceBusDestination) MarshalJSON() ([]byte, error) {
	type Alias AzureServiceBusDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "AzureServiceBus", Alias: (*Alias)(&obj)})
}

/**
*	This destination can be used to push events and messages to [Confluent Cloud](https://www.confluent.io/confluent-cloud/).
*	To set up a Subscription of this type, first, create a topic in Confluent Cloud.
*	Then, to allow Composable Commerce to push events and messages to your topic, generate [API keys](https://docs.confluent.io/cloud/current/access-management/authenticate/api-keys/api-keys.html) for your topic, and create the Subscription destination using the generated credentials.
*
*	The Composable Commerce producer uses the following values: `SASL_SSL` for`security.protocol`, `PLAIN` for`sasl.mechanism`, and the default value (1048576) for `max.request.size`.
*
 */
type ConfluentCloudDestination struct {
	// URL to the bootstrap server including the port number in the format `<xxxxx>.<region>.<provider>.confluent.cloud:9092`.
	BootstrapServer string `json:"bootstrapServer"`
	// Partially hidden on retrieval for security reasons.
	ApiKey string `json:"apiKey"`
	// Partially hidden on retrieval for security reasons.
	ApiSecret string `json:"apiSecret"`
	// The Kafka `acks` value.
	Acks string `json:"acks"`
	// The name of the topic.
	Topic string `json:"topic"`
	// The Kafka record key.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ConfluentCloudDestination) MarshalJSON() ([]byte, error) {
	type Alias ConfluentCloudDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ConfluentCloud", Alias: (*Alias)(&obj)})
}

/**
*	[AWS EventBridge](https://aws.amazon.com/eventbridge/) can be used to push events and messages to a serverless event bus that can forward them to AWS SQS, SNS, Lambda, and other AWS services based on forwarding rules.
*	Once the Subscription is created, an equivalent "partner event source" is created in AWS EventBridge. This event source must be associated with an event bus for the Subscription setup to be complete.
*
 */
type EventBridgeDestination struct {
	// AWS region that receives the events.
	Region string `json:"region"`
	// ID of the AWS account that receives the events.
	AccountId string `json:"accountId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EventBridgeDestination) MarshalJSON() ([]byte, error) {
	type Alias EventBridgeDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "EventBridge", Alias: (*Alias)(&obj)})
}

/**
*	Destination for [Google Cloud Pub/Sub](https://cloud.google.com/pubsub/) that can be used
*	for [Pull subscriptions](https://cloud.google.com/pubsub/docs/pull) as well as for [Push subscriptions](https://cloud.google.com/pubsub/docs/push).
*	The `topic` must give the `pubsub.topics.publish` permission to the service account `subscriptions@commercetools-platform.iam.gserviceaccount.com`.
*	If used with the [CloudEventsFormat](#cloudeventsformat), the message conforms to the [PubSub Protocol Binding](https://github.com/google/knative-gcp/blob/master/docs/spec/pubsub-protocol-binding.md) of the [Structured Content Mode](https://github.com/google/knative-gcp/blob/master/docs/spec/pubsub-protocol-binding.md#32-structured-content-mode).
*
 */
type GoogleCloudPubSubDestination struct {
	// ID of the Google Cloud project that contains the Pub/Sub topic.
	ProjectId string `json:"projectId"`
	// Name of the topic.
	Topic string `json:"topic"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj IronMqDestination) MarshalJSON() ([]byte, error) {
	type Alias IronMqDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "IronMQ", Alias: (*Alias)(&obj)})
}

/**
*	This payload is sent for a [MessageSubscription](ctp:api:type:MessageSubscription).
*
 */
type MessageDeliveryPayload struct {
	// `key` of the [Project](ctp:api:type:Project).
	// Useful in message processing if the Destination receives events from multiple Projects.
	ProjectKey string `json:"projectKey"`
	// Reference to the resource that triggered the message.
	Resource Reference `json:"resource"`
	// User-defined unique identifiers of the resource.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Unique ID of the message.
	ID string `json:"id"`
	// Last seen version of the resource.
	Version int `json:"version"`
	// Date and time (UTC) the resource was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the resource was last modified.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Used to ensure all messages of the resource are processed in correct order.
	// The `sequenceNumber` of the next message of the resource is a successor of the `sequenceNumber` of the current message.
	SequenceNumber int `json:"sequenceNumber"`
	// Version of the resource on which the change was performed.
	ResourceVersion int `json:"resourceVersion"`
	// If the payload does not fit into the size limit or its format is not accepted by the messaging service, the `payloadNotIncluded` field is present.
	PayloadNotIncluded *PayloadNotIncluded `json:"payloadNotIncluded,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MessageDeliveryPayload) UnmarshalJSON(data []byte) error {
	type Alias MessageDeliveryPayload
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MessageDeliveryPayload) MarshalJSON() ([]byte, error) {
	type Alias MessageDeliveryPayload
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "Message", Alias: (*Alias)(&obj)})
}

/**
*	For supported resources and message types, see [Message Types](/../api/projects/messages#message-types). Messages will be delivered even if the Messages Query HTTP API [is not enabled](/../api/projects/messages#enable-querying-messages-via-the-api).
*
*	For MessageSubscriptions, the format of the payload is [MessageDeliveryPayload](ctp:api:type:MessageDeliveryPayload).
*
 */
type MessageSubscription struct {
	// Unique identifier for the type of resource, for example, `order`.
	ResourceTypeId MessageSubscriptionResourceTypeId `json:"resourceTypeId"`
	// Must contain valid message types for the resource. For example, for resource type `product` the message type `ProductPublished` is valid.
	// If no `types` of messages are given, the Subscription will receive all messages for this resource.
	Types []string `json:"types"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MessageSubscription) MarshalJSON() ([]byte, error) {
	type Alias MessageSubscription
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["types"] == nil {
		delete(raw, "types")
	}

	return json.Marshal(raw)

}

/**
*	Resource types supported by [MessageSubscriptions](ctp:api:type:MessageSubscription):
*
 */
type MessageSubscriptionResourceTypeId string

const (
	MessageSubscriptionResourceTypeIdApprovalFlow          MessageSubscriptionResourceTypeId = "approval-flow"
	MessageSubscriptionResourceTypeIdApprovalRule          MessageSubscriptionResourceTypeId = "approval-rule"
	MessageSubscriptionResourceTypeIdAssociateRole         MessageSubscriptionResourceTypeId = "associate-role"
	MessageSubscriptionResourceTypeIdBusinessUnit          MessageSubscriptionResourceTypeId = "business-unit"
	MessageSubscriptionResourceTypeIdCategory              MessageSubscriptionResourceTypeId = "category"
	MessageSubscriptionResourceTypeIdCustomer              MessageSubscriptionResourceTypeId = "customer"
	MessageSubscriptionResourceTypeIdCustomerEmailToken    MessageSubscriptionResourceTypeId = "customer-email-token"
	MessageSubscriptionResourceTypeIdCustomerGroup         MessageSubscriptionResourceTypeId = "customer-group"
	MessageSubscriptionResourceTypeIdCustomerPasswordToken MessageSubscriptionResourceTypeId = "customer-password-token"
	MessageSubscriptionResourceTypeIdInventoryEntry        MessageSubscriptionResourceTypeId = "inventory-entry"
	MessageSubscriptionResourceTypeIdOrder                 MessageSubscriptionResourceTypeId = "order"
	MessageSubscriptionResourceTypeIdPayment               MessageSubscriptionResourceTypeId = "payment"
	MessageSubscriptionResourceTypeIdProduct               MessageSubscriptionResourceTypeId = "product"
	MessageSubscriptionResourceTypeIdProductSelection      MessageSubscriptionResourceTypeId = "product-selection"
	MessageSubscriptionResourceTypeIdQuote                 MessageSubscriptionResourceTypeId = "quote"
	MessageSubscriptionResourceTypeIdQuoteRequest          MessageSubscriptionResourceTypeId = "quote-request"
	MessageSubscriptionResourceTypeIdReview                MessageSubscriptionResourceTypeId = "review"
	MessageSubscriptionResourceTypeIdStagedQuote           MessageSubscriptionResourceTypeId = "staged-quote"
	MessageSubscriptionResourceTypeIdStandalonePrice       MessageSubscriptionResourceTypeId = "standalone-price"
	MessageSubscriptionResourceTypeIdStore                 MessageSubscriptionResourceTypeId = "store"
)

type PayloadNotIncluded struct {
	// Reason the payload is not included. For example, the payload is too large, or its content is not supported by the Subscription destination.
	Reason string `json:"reason"`
	// Value of the `type` field in the original payload.
	PayloadType string `json:"payloadType"`
}

/**
*	The PlatformFormat uses constructs that are similar to the ones used in the REST API, for example, on the [Messages Query HTTP API](/../api/projects/messages).
*
 */
type PlatformFormat struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PlatformFormat) MarshalJSON() ([]byte, error) {
	type Alias PlatformFormat
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Platform", Alias: (*Alias)(&obj)})
}

/**
*	This payload is sent for a [ChangeSubscription](ctp:api:type:ChangeSubscription) when a resource is created.
*
 */
type ResourceCreatedDeliveryPayload struct {
	// `key` of the [Project](ctp:api:type:Project).
	// Useful in message processing if the Destination receives events from multiple Projects.
	ProjectKey string `json:"projectKey"`
	// Reference to the resource that triggered the message.
	Resource Reference `json:"resource"`
	// User-defined unique identifiers of the resource.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Last seen version of the resource.
	Version int `json:"version"`
	// Date and time (UTC) the resource was last modified.
	ModifiedAt time.Time `json:"modifiedAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceCreatedDeliveryPayload) UnmarshalJSON(data []byte) error {
	type Alias ResourceCreatedDeliveryPayload
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceCreatedDeliveryPayload) MarshalJSON() ([]byte, error) {
	type Alias ResourceCreatedDeliveryPayload
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "ResourceCreated", Alias: (*Alias)(&obj)})
}

/**
*	This payload is sent for a [ChangeSubscription](ctp:api:type:ChangeSubscription) when a resource is deleted.
*
 */
type ResourceDeletedDeliveryPayload struct {
	// `key` of the [Project](ctp:api:type:Project).
	// Useful in message processing if the Destination receives events from multiple Projects.
	ProjectKey string `json:"projectKey"`
	// Reference to the resource that triggered the message.
	Resource Reference `json:"resource"`
	// User-defined unique identifiers of the resource.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Last seen version of the resource.
	Version int `json:"version"`
	// Date and time (UTC) the resource was last deleted.
	ModifiedAt time.Time `json:"modifiedAt"`
	// `true` if the `dataErasure` [parameter](/../api/general-concepts#data-erasure-of-personal-data) on the `DELETE` request was set to `true`.
	DataErasure *bool `json:"dataErasure,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceDeletedDeliveryPayload) UnmarshalJSON(data []byte) error {
	type Alias ResourceDeletedDeliveryPayload
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceDeletedDeliveryPayload) MarshalJSON() ([]byte, error) {
	type Alias ResourceDeletedDeliveryPayload
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "ResourceDeleted", Alias: (*Alias)(&obj)})
}

/**
*	This payload is sent for a [ChangeSubscription](ctp:api:type:ChangeSubscription) when a resource is updated. This includes updates by a background process, like a change in product availability.
*
 */
type ResourceUpdatedDeliveryPayload struct {
	// `key` of the [Project](ctp:api:type:Project).
	// Useful in message processing if the Destination receives events from multiple Projects.
	ProjectKey string `json:"projectKey"`
	// Reference to the resource that triggered the message.
	Resource Reference `json:"resource"`
	// User-defined unique identifiers of the resource.
	ResourceUserProvidedIdentifiers *UserProvidedIdentifiers `json:"resourceUserProvidedIdentifiers,omitempty"`
	// Last seen version of the resource.
	Version int `json:"version"`
	// Version of the resource before the update.
	OldVersion int `json:"oldVersion"`
	// Date and time (UTC) the resource was last updated.
	ModifiedAt time.Time `json:"modifiedAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ResourceUpdatedDeliveryPayload) UnmarshalJSON(data []byte) error {
	type Alias ResourceUpdatedDeliveryPayload
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ResourceUpdatedDeliveryPayload) MarshalJSON() ([]byte, error) {
	type Alias ResourceUpdatedDeliveryPayload
	return json.Marshal(struct {
		Action string `json:"notificationType"`
		*Alias
	}{Action: "ResourceUpdated", Alias: (*Alias)(&obj)})
}

/**
*	[AWS SNS](https://aws.amazon.com/sns/) can be used to push messages to AWS Lambda, HTTP endpoints (webhooks), or fan-out messages to SQS queues. The SQS queue must be a [Standard](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/standard-queues.html) queue type.
*
*	We recommend setting `authenticationMode` to `IAM`, to avoid unnecessary key management. For IAM authentication and before creating the Subscription, give permissions to the following user account: `arn:aws-cn:iam::417094354346:user/subscriptions` if the Project is hosted in the China (AWS, Ningxia) Region; `arn:aws:iam::362576667341:user/subscriptions` for all other [Regions](/../api/general-concepts#regions). Otherwise, a test message will not be sent.
*
*	If you prefer to use `Credentials` for authentication, we recommend [creating an IAM user](https://docs.aws.amazon.com/sns/latest/dg/sns-setting-up.html#create-iam-user) with an `accessKey` and `accessSecret` pair specifically for each Subscription.
*
*	The IAM user should only have the `sns:Publish` permission on this topic.
*
 */
type SnsDestination struct {
	// Only present if `authenticationMode` is set to `Credentials`.
	AccessKey *string `json:"accessKey,omitempty"`
	// Only present if `authenticationMode` is set to `Credentials`.
	AccessSecret *string `json:"accessSecret,omitempty"`
	// Amazon Resource Name (ARN) of the topic.
	TopicArn string `json:"topicArn"`
	// Defines the method of authentication for the SNS topic.
	AuthenticationMode *AwsAuthenticationMode `json:"authenticationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SnsDestination) MarshalJSON() ([]byte, error) {
	type Alias SnsDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SNS", Alias: (*Alias)(&obj)})
}

/**
*	[AWS SQS](https://aws.amazon.com/sqs/) is a pull-queue on AWS.
*	The queue must be a [Standard](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/standard-queues.html) queue type with a `MaximumMessageSize` of `256 KB`.
*
*	We recommend setting `authenticationMode` to `IAM`, to avoid unnecessary key management. For IAM authentication and before creating the Subscription, give permissions to the following user account: `arn:aws-cn:iam::417094354346:user/subscriptions` if the Project is hosted in the China (AWS, Ningxia) Region; `arn:aws:iam::362576667341:user/subscriptions` for all other [Regions](/../api/general-concepts#regions). Otherwise, a test message will not be sent.
*
*	If you prefer to use `Credentials` for authentication, we recommend [creating an IAM user](https://docs.aws.amazon.com/sns/latest/dg/sns-setting-up.html#create-iam-user) with an `accessKey` and `accessSecret` pair specifically for each Subscription.
*
*	The IAM user should only have the `sqs:SendMessage` permission on this queue.
*
 */
type SqsDestination struct {
	// Only present if `authenticationMode` is set to `Credentials`.
	AccessKey *string `json:"accessKey,omitempty"`
	// Only present if `authenticationMode` is set to `Credentials`.
	AccessSecret *string `json:"accessSecret,omitempty"`
	// URL of the Amazon SQS queue.
	QueueUrl string `json:"queueUrl"`
	// [AWS Region](https://docs.aws.amazon.com/general/latest/gr/rande-manage.html) the message queue is located in.
	Region string `json:"region"`
	// Defines the method of authentication for the SQS queue.
	AuthenticationMode *AwsAuthenticationMode `json:"authenticationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SqsDestination) MarshalJSON() ([]byte, error) {
	type Alias SqsDestination
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "SQS", Alias: (*Alias)(&obj)})
}

type Subscription struct {
	// Unique identifier of the Subscription.
	ID string `json:"id"`
	// Current version of the Subscription.
	Version int `json:"version"`
	// Date and time (UTC) the Subscription was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Subscription was last modified.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Change notifications subscribed to.
	Changes []ChangeSubscription `json:"changes"`
	// Messaging service to which the messages are to be sent.
	Destination Destination `json:"destination"`
	// User-defined unique identifier of the Subscription.
	Key *string `json:"key,omitempty"`
	// Messages subscribed to.
	Messages []MessageSubscription `json:"messages"`
	// Format in which the payload is delivered.
	Format DeliveryFormat `json:"format"`
	// Status of the Subscription.
	Status SubscriptionHealthStatus `json:"status"`
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

/**
*	Either `messages` or `changes` must be set.
*
 */
type SubscriptionDraft struct {
	// Change notifications to be subscribed to.
	Changes []ChangeSubscription `json:"changes"`
	// Messaging service to which the messages are sent.
	Destination Destination `json:"destination"`
	// User-defined unique identifier for the Subscription.
	Key *string `json:"key,omitempty"`
	// Messages to be subscribed to.
	Messages []MessageSubscription `json:"messages"`
	// Format in which the payload is delivered. When not provided, the [PlatformFormat](ctp:api:type:PlatformFormat) is selected by default.
	Format DeliveryFormat `json:"format,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SubscriptionDraft) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["changes"] == nil {
		delete(raw, "changes")
	}

	if raw["messages"] == nil {
		delete(raw, "messages")
	}

	return json.Marshal(raw)

}

/**
*	The health status of the Subscription that indicates whether messages are being delivered to the Destination.
*
 */
type SubscriptionHealthStatus string

const (
	SubscriptionHealthStatusHealthy                           SubscriptionHealthStatus = "Healthy"
	SubscriptionHealthStatusConfigurationError                SubscriptionHealthStatus = "ConfigurationError"
	SubscriptionHealthStatusConfigurationErrorDeliveryStopped SubscriptionHealthStatus = "ConfigurationErrorDeliveryStopped"
	SubscriptionHealthStatusTemporaryError                    SubscriptionHealthStatus = "TemporaryError"
)

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Subscription](ctp:api:type:Subscription).
*
 */
type SubscriptionPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Subscriptions matching the query.
	Results []Subscription `json:"results"`
}

type SubscriptionUpdate struct {
	// Expected version of the Subscription on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Subscription.
	Actions []SubscriptionUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SubscriptionUpdate) UnmarshalJSON(data []byte) error {
	type Alias SubscriptionUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorSubscriptionUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type SubscriptionUpdateAction interface{}

func mapDiscriminatorSubscriptionUpdateAction(input interface{}) (SubscriptionUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "changeDestination":
		obj := SubscriptionChangeDestinationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Destination != nil {
			var err error
			obj.Destination, err = mapDiscriminatorDestination(obj.Destination)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setChanges":
		obj := SubscriptionSetChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := SubscriptionSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMessages":
		obj := SubscriptionSetMessagesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	A test message is sent to ensure the correct configuration of the Destination. If the message cannot be delivered, the update will fail. The payload of the test message is a notification of type [ResourceCreated](ctp:api:type:ResourceCreatedDeliveryPayload) for the `resourceTypeId` `subscription`. The `status` will change to [Healthy](ctp:api:type:SubscriptionHealthStatus), if it isn't already.
*
 */
type SubscriptionChangeDestinationAction struct {
	// New value to set. Must not be empty.
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SubscriptionChangeDestinationAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionChangeDestinationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDestination", Alias: (*Alias)(&obj)})
}

type SubscriptionSetChangesAction struct {
	// Value to set. Can only be unset if `messages` is set.
	Changes []ChangeSubscription `json:"changes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SubscriptionSetChangesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetChangesAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setChanges", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["changes"] == nil {
		delete(raw, "changes")
	}

	return json.Marshal(raw)

}

type SubscriptionSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SubscriptionSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type SubscriptionSetMessagesAction struct {
	// Value to set. Can only be unset if `changes` is set.
	Messages []MessageSubscription `json:"messages"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SubscriptionSetMessagesAction) MarshalJSON() ([]byte, error) {
	type Alias SubscriptionSetMessagesAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMessages", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["messages"] == nil {
		delete(raw, "messages")
	}

	return json.Marshal(raw)

}
