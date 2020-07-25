// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// SubscriptionCreate creates a new instance of type Subscription
func (client *Client) SubscriptionCreate(ctx context.Context, draft *SubscriptionDraft, opts ...RequestOption) (result *Subscription, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "subscriptions"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionQuery allows querying for type Subscription
func (client *Client) SubscriptionQuery(ctx context.Context, input *QueryInput) (result *SubscriptionPagedQueryResponse, err error) {
	endpoint := "subscriptions"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionDeleteWithID for type Subscription
func (client *Client) SubscriptionDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Subscription, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("subscriptions/%s", id)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionDeleteWithKey for type Subscription
func (client *Client) SubscriptionDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Subscription, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("subscriptions/key=%s", key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionGetWithID Retrieves the representation of a subscription by its id.
func (client *Client) SubscriptionGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Subscription, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("subscriptions/%s", id)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionGetWithKey Retrieves the representation of a subscription by its key.
func (client *Client) SubscriptionGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Subscription, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("subscriptions/key=%s", key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionUpdateWithIDInput is input for function SubscriptionUpdateWithID
type SubscriptionUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []SubscriptionUpdateAction
}

func (input *SubscriptionUpdateWithIDInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("no valid value for ID given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// SubscriptionUpdateWithID for type Subscription
func (client *Client) SubscriptionUpdateWithID(ctx context.Context, input *SubscriptionUpdateWithIDInput, opts ...RequestOption) (result *Subscription, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("subscriptions/%s", input.ID)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionUpdateWithKeyInput is input for function SubscriptionUpdateWithKey
type SubscriptionUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []SubscriptionUpdateAction
}

func (input *SubscriptionUpdateWithKeyInput) Validate() error {
	if input.Key == "" {
		return fmt.Errorf("no valid value for Key given")
	}
	if len(input.Actions) == 0 {
		return fmt.Errorf("no update actions specified")
	}
	return nil
}

// SubscriptionUpdateWithKey for type Subscription
func (client *Client) SubscriptionUpdateWithKey(ctx context.Context, input *SubscriptionUpdateWithKeyInput, opts ...RequestOption) (result *Subscription, err error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := fmt.Sprintf("subscriptions/key=%s", input.Key)
	err = client.update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
