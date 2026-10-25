package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyApplicationsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves an [Application](ctp:checkout:type:Application) by a given `id`.
*
 */
func (rb *ByProjectKeyApplicationsByIdRequestBuilder) Get() *ByProjectKeyApplicationsByIdRequestMethodGet {
	return &ByProjectKeyApplicationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/applications/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an [Application](/applications-api#application) with a given `id` exists. Returns `200 OK` if the Application exists, or `404 Not Found` if it doesn't. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyApplicationsByIdRequestBuilder) Head() *ByProjectKeyApplicationsByIdRequestMethodHead {
	return &ByProjectKeyApplicationsByIdRequestMethodHead{
		url:    fmt.Sprintf("/%s/applications/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyApplicationsByIdRequestBuilder) Post(body ApplicationUpdateActions) *ByProjectKeyApplicationsByIdRequestMethodPost {
	return &ByProjectKeyApplicationsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/applications/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyApplicationsByIdRequestBuilder) Delete() *ByProjectKeyApplicationsByIdRequestMethodDelete {
	return &ByProjectKeyApplicationsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/applications/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
