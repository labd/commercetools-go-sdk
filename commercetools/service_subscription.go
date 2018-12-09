package commercetools

import (
	"fmt"
	"net/url"
	"strconv"
)

// DeleteInput provides the data required to delete a subscription.
type SubscriptionDeleteInput struct {
	ID      string
	Version int
}

// UpdateInput provides the data required to update a subscription.
type SubscriptionUpdateInput struct {
	ID      string
	Version int
	Actions []SubscriptionUpdateAction
}

// Service contains client information and bundles all actions.
type SubscriptionService struct {
	client *Client
}

// GetByID will return a subscription matching the provided ID.
func (svc *SubscriptionService) GetByID(id string) (*Subscription, error) {
	var result Subscription
	err := svc.client.Get(fmt.Sprintf("subscriptions/%s", id), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new Subscription. It is eventually consistent, it may take
// up to a minute before it becomes fully active.
//
// In order to test that the destination is correctly configured, a test message
// will be put into the queue. If the message could not be delivered, the
// subscription will not be created. The payload of the test message is a
// notification of type ResourceCreated for the resourceTypeId subscription.
//
// Currently, a maximum of 25 subscriptions can be created per project.
func (svc *SubscriptionService) Create(draft *SubscriptionDraft) (*Subscription, error) {
	var result Subscription
	err := svc.client.Create("subscriptions", nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a Subscription. It is eventually consistent, it may take up to
// a minute before changes becomes fully active.
func (svc *SubscriptionService) Update(input *SubscriptionUpdateInput) (*Subscription, error) {
	var result Subscription

	if input.ID == "" {
		return nil, fmt.Errorf("No valid subscription id passed")
	}

	endpoint := fmt.Sprintf("subscriptions/%s", input.ID)
	err := svc.client.Update(endpoint, nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByID will delete a subscription matching the provided ID.
func (svc *SubscriptionService) DeleteByID(id string, version int) (*Subscription, error) {
	var result Subscription
	endpoint := fmt.Sprintf("subscriptions/%s", id)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteByKey will delete a subscription matching the provided key.
func (svc *SubscriptionService) DeleteByKey(key string, version int) (*Subscription, error) {
	var result Subscription
	endpoint := fmt.Sprintf("subscriptions/key=%s", key)
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	err := svc.client.Delete(endpoint, params, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
