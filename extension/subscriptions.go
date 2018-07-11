package extension

import "fmt"

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
