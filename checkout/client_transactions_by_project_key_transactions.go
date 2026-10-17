package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTransactionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTransactionsRequestBuilder) WithId(id string) *ByProjectKeyTransactionsByIdRequestBuilder {
	return &ByProjectKeyTransactionsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTransactionsRequestBuilder) WithKey(key string) *ByProjectKeyTransactionsKeyByKeyRequestBuilder {
	return &ByProjectKeyTransactionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Creates a Transaction on Checkout. Requests to the payment [Connector](/checkout/connectors-and-applications#payment-connectors) time out after 30 seconds.
*
*	Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [InternalConstraintViolated](ctp:checkout:type:InternalConstraintViolatedError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConnectorFailed](ctp:checkout:type:ConnectorFailedError)
*	- [ConnectorTimeout](ctp:checkout:type:ConnectorTimeoutError)
*	- [PaymentFailure](ctp:checkout:type:PaymentFailureError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*
 */
func (rb *ByProjectKeyTransactionsRequestBuilder) Post(body TransactionDraft) *ByProjectKeyTransactionsRequestMethodPost {
	return &ByProjectKeyTransactionsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/transactions", rb.projectKey),
		client: rb.client,
	}
}
