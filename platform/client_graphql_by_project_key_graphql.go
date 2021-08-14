// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyGraphqlRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Execute a GraphQL query
 */
func (rb *ByProjectKeyGraphqlRequestBuilder) Post(body GraphQLRequest) *ByProjectKeyGraphqlRequestMethodPost {
	return &ByProjectKeyGraphqlRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/graphql", rb.projectKey),
		client: rb.client,
	}
}
