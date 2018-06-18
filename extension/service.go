package extension

import "github.com/labd/commercetools-go-sdk/common"

type Service struct {
	client *common.Client
}

func New(client *common.Client) *Service {
	return &Service{client}
}
