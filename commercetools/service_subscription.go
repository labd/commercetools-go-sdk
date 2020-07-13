// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// SubscriptionURLPath is the commercetools API path.
const SubscriptionURLPath = "subscriptions"

// SubscriptionCreate creates a new instance of type Subscription
func (client *Client) SubscriptionCreate(ctx context.Context, draft *SubscriptionDraft) (result *Subscription, err error) {
	err = client.Create(ctx, SubscriptionURLPath, nil, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionQuery allows querying for type Subscription
func (client *Client) SubscriptionQuery(ctx context.Context, input *QueryInput) (result *SubscriptionPagedQueryResponse, err error) {
	err = client.Query(ctx, SubscriptionURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionDeleteWithKey for type Subscription
func (client *Client) SubscriptionDeleteWithKey(ctx context.Context, key string, version int) (result *Subscription, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("subscriptions/key={key}", "{key}", key, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionGetWithKey Retrieves the representation of a subscription by its key.
func (client *Client) SubscriptionGetWithKey(ctx context.Context, key string) (result *Subscription, err error) {
	err = client.Get(ctx, strings.Replace("subscriptions/key={key}", "{key}", key, 1), nil, &result)
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

// SubscriptionUpdateWithKey for type Subscription
func (client *Client) SubscriptionUpdateWithKey(ctx context.Context, input *SubscriptionUpdateWithKeyInput) (result *Subscription, err error) {
	err = client.Update(ctx, strings.Replace("subscriptions/key={key}", "{key}", input.Key, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionDeleteWithID for type Subscription
func (client *Client) SubscriptionDeleteWithID(ctx context.Context, ID string, version int) (result *Subscription, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	err = client.Delete(ctx, strings.Replace("subscriptions/{ID}", "{ID}", ID, 1), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubscriptionGetWithID Retrieves the representation of a subscription by its id.
func (client *Client) SubscriptionGetWithID(ctx context.Context, ID string) (result *Subscription, err error) {
	err = client.Get(ctx, strings.Replace("subscriptions/{ID}", "{ID}", ID, 1), nil, &result)
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

// SubscriptionUpdateWithID for type Subscription
func (client *Client) SubscriptionUpdateWithID(ctx context.Context, input *SubscriptionUpdateWithIDInput) (result *Subscription, err error) {
	err = client.Update(ctx, strings.Replace("subscriptions/{ID}", "{ID}", input.ID, 1), nil, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
