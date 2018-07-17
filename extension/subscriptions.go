package extension

import (
	"fmt"
	"net/url"

	"github.com/labd/commercetools-go-sdk/common"
)

type SubscriptionDeleteInput struct {
	ID      string
	Version int
}

type SubscriptionUpdateInput struct {
	ID      string
	Version int
	Actions common.UpdateActions
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
