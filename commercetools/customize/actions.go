package customize

import "github.com/labd/commercetools-go-sdk/commercetools"

type SubscriptionSetKey struct {
	commercetools.UpdateAction
	Key string `json:"key,omitempty"`
}

type SubscriptionSetMessages struct {
	commercetools.UpdateAction
	Messages []MessageSubscription `json:"messages"`
}

type SubscriptionSetChanges struct {
	commercetools.UpdateAction
	Changes []ChangeSubscription `json:"changes"`
}
