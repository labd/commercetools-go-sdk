package extension

import (
	"time"
)

type MessageSubscription struct {
	ResourceTypeID string   `json:"resourceTypeId"`
	Types          []string `json:"types"`
}

type ChangeSubscription struct {
	resourceTypeID string
}

type SubscriptionAWSSQSDestination struct {
	Type         string `json:"type"`
	QueueURL     string `json:"queueUrl"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	Region       string `json:"region"`
}

type SubscriptionAWSSNSDestination struct {
	Type         string `json:"type"`
	TopicArn     string `json:"topicArn"`
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

type SubscriptionIronMQDestination struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SubscriptionDestination interface {
}

type Subscription struct {
	ID             string                  `json:"id"`
	Version        int                     `json:"version"`
	Key            string                  `json:"key"`
	Destination    SubscriptionDestination `json:"destination"`
	Messages       []MessageSubscription   `json:"messages"`
	Changes        []ChangeSubscription    `json:"changes"`
	CreatedAt      time.Time               `json:"createdAt"`
	LastModifiedAt time.Time               `json:"lastModifiedAt"`
}

type SubscriptionDraft struct {
	Key         string                  `json:"key"`
	Destination SubscriptionDestination `json:"destination"`
	Messages    []MessageSubscription   `json:"messages"`
	Changes     []ChangeSubscription    `json:"changes"`
}
