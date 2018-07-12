package extension

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type SubscriptionDeleteInput struct {
	ID      string
	Version int
}

func (svc *Service) SubscriptionGetByID(id string) (*Subscription, error) {
	var result Subscription
	err := svc.client.Get(fmt.Sprintf("subscriptions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}

	result.Destination = convertSubscriptionDestination(result.Destination)
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

func (svc *Service) SubscriptionDelete(input *SubscriptionDeleteInput) (*Subscription, error) {
	return nil, nil
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
