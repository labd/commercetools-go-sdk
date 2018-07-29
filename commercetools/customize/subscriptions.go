package customize

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/mitchellh/mapstructure"
)

type SubscriptionDeleteInput struct {
	ID      string
	Version int
}

type SubscriptionUpdateInput struct {
	ID      string
	Version int
	Actions commercetools.UpdateActions
}

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

func (s *Subscription) UnmarshalJSON(data []byte) error {
	type SubscriptionClone Subscription
	if err := json.Unmarshal(data, (*SubscriptionClone)(s)); err != nil {
		return err
	}
	s.Destination = convertSubscriptionDestination(s.Destination)
	return nil
}

type SubscriptionDraft struct {
	Key         string                  `json:"key"`
	Destination SubscriptionDestination `json:"destination"`
	Messages    []MessageSubscription   `json:"messages"`
	Changes     []ChangeSubscription    `json:"changes"`
}

func (svc *Service) SubscriptionGetByID(id string) (*Subscription, error) {
	var result Subscription
	err := svc.client.Get(fmt.Sprintf("subscriptions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) SubscriptionCreate(draft *SubscriptionDraft) (*Subscription, error) {
	var result Subscription
	err := svc.client.Create("subscriptions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) SubscriptionUpdate(input *SubscriptionUpdateInput) (*Subscription, error) {
	var result Subscription

	endpoint := fmt.Sprintf("products/%s", input.ID)
	err := svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) SubscriptionDeleteByID(id string, version int) (*Subscription, error) {
	var result Subscription
	endpoint := fmt.Sprintf("subscriptions/%s", id)
	params := url.Values{}
	params.Set("version", string(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc *Service) SubscriptionDeleteByKey(key string, version int) (*Subscription, error) {
	var result Subscription
	endpoint := fmt.Sprintf("subscriptions/key=%s", key)
	params := url.Values{}
	params.Set("version", string(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func convertSubscriptionDestination(input SubscriptionDestination) SubscriptionDestination {
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
