package customize

import "github.com/labd/commercetools-go-sdk/commercetools"

type Service struct {
	client *commercetools.Client
}

func New(client *commercetools.Client) *Service {
	return &Service{client}
}
