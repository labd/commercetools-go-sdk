package customize

import "github.com/labd/commercetools-go-sdk/common"

type SubscriptionSetKey struct {
	common.UpdateAction
	Key string `json:"key,omitempty"`
}

type SubscriptionSetMessages struct {
	common.UpdateAction
	Messages []MessageSubscription `json:"messages"`
}

type SubscriptionSetChanges struct {
	common.UpdateAction
	Changes []ChangeSubscription `json:"changes"`
}
