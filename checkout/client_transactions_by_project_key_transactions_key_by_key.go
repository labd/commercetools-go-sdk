package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTransactionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Returns a Transaction with a given `key`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyTransactionsKeyByKeyRequestBuilder) Get() *ByProjectKeyTransactionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyTransactionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/transactions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
