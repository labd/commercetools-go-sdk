package history

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyGraphqlRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Execute a GraphQL request.
 */
func (rb *ByProjectKeyGraphqlRequestBuilder) Post(body GraphQLRequest) *ByProjectKeyGraphqlRequestMethodPost {
	return &ByProjectKeyGraphqlRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/graphql", rb.projectKey),
		client: rb.client,
	}
}
