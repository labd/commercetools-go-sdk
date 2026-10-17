package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTransactionsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Returns a Transaction with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyTransactionsByIdRequestBuilder) Get() *ByProjectKeyTransactionsByIdRequestMethodGet {
	return &ByProjectKeyTransactionsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/transactions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
