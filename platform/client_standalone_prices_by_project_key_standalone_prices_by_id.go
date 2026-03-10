package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStandalonePricesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Get() *ByProjectKeyStandalonePricesByIDRequestMethodGet {
	return &ByProjectKeyStandalonePricesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a StandalonePrice exists with the provided `id`. Returns a `200` status if the StandalonePrice exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Head() *ByProjectKeyStandalonePricesByIDRequestMethodHead {
	return &ByProjectKeyStandalonePricesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	If a modification is already in progress for the exact combination of SKU and price scope fields, an [ExactLockConflict](ctp:api:type:ExactLockConflictError) or [ValidityLockConflict](ctp:api:type:ValidityLockConflictError) error is returned.
*
 */
func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Post(body StandalonePriceUpdate) *ByProjectKeyStandalonePricesByIDRequestMethodPost {
	return &ByProjectKeyStandalonePricesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Produces the [StandalonePriceDeleted](ctp:api:type:StandalonePriceDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Delete() *ByProjectKeyStandalonePricesByIDRequestMethodDelete {
	return &ByProjectKeyStandalonePricesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
