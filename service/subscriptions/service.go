package subscriptions

import "github.com/labd/commercetools-go-sdk/commercetools"

// Service contains client information and bundles all actions.
type Service struct {
	client *commercetools.Client
}

// New returns an instance of the subscriptions service.
func New(client *commercetools.Client) *Service {
	return &Service{client}
}
