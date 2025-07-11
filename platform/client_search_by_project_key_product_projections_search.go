package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	For implementing funnel search on Product Listing Pages where users select multiple filters, use this POST method.
*	To avoid URL length restrictions, this method passes the same query parameters as defined in the [GET](ctp:api:endpoint:/{projectKey}/product-projections/search:GET) method, within the request body in URL-encoded format.
*
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestBuilder) Post(body string) *ByProjectKeyProductProjectionsSearchRequestMethodPost {
	return &ByProjectKeyProductProjectionsSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-projections/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	This method appends query parameters to the URL.
*	The maximum allowed URL length is 8192 characters.
*	Exceeding this limit will result in URL truncation, potentially leading to unexpected results.
*	For funnel searches on Product Listing Pages, where users select multiple filters, we recommend the [POST](ctp:api:endpoint:/{projectKey}/product-projections/search:POST) method which passes the query parameters within the request body, avoiding URL length restrictions.
*
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestBuilder) Get() *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	return &ByProjectKeyProductProjectionsSearchRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/search", rb.projectKey),
		client: rb.client,
	}
}
