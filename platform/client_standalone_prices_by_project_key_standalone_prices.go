package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStandalonePricesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesRequestBuilder) WithKey(key string) *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStandalonePricesRequestBuilder) WithId(id string) *ByProjectKeyStandalonePricesByIDRequestBuilder {
	return &ByProjectKeyStandalonePricesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStandalonePricesRequestBuilder) Get() *ByProjectKeyStandalonePricesRequestMethodGet {
	return &ByProjectKeyStandalonePricesRequestMethodGet{
		url:    fmt.Sprintf("/%s/standalone-prices", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a StandalonePrice exists for a given Query Predicate. Returns a `200 OK` status if any StandalonePrices match the Query Predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStandalonePricesRequestBuilder) Head() *ByProjectKeyStandalonePricesRequestMethodHead {
	return &ByProjectKeyStandalonePricesRequestMethodHead{
		url:    fmt.Sprintf("/%s/standalone-prices", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a Standalone Price produces the [StandalonePriceCreated](ctp:api:type:StandalonePriceCreatedMessage) Message.
*
*	- If the Standalone Price has the same price scope as an existing Standalone Price, a [DuplicateStandalonePriceScope](ctp:api:type:DuplicateStandalonePriceScopeError) error is returned.
*	- If the Standalone Price has overlapping validity periods within the same price scope, a [OverlappingStandalonePriceValidity](ctp:api:type:OverlappingStandalonePriceValidityError) error is returned. A Price without validity period does not conflict with a Price defined for a time period.
*
 */
func (rb *ByProjectKeyStandalonePricesRequestBuilder) Post(body StandalonePriceDraft) *ByProjectKeyStandalonePricesRequestMethodPost {
	return &ByProjectKeyStandalonePricesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices", rb.projectKey),
		client: rb.client,
	}
}
