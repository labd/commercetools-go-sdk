package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyApplicationsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves an [Application](ctp:checkout:type:Application) with a given `key`.
*	Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyApplicationsKeyByKeyRequestBuilder) Get() *ByProjectKeyApplicationsKeyByKeyRequestMethodGet {
	return &ByProjectKeyApplicationsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/applications/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyApplicationsKeyByKeyRequestBuilder) Post(body ApplicationUpdateActions) *ByProjectKeyApplicationsKeyByKeyRequestMethodPost {
	return &ByProjectKeyApplicationsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/applications/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if an [Application](ctp:checkout:type:Application) with a given `key` exists. Returns `200 OK` if the Application exists, or `404 Not Found` if it doesn't. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyApplicationsKeyByKeyRequestBuilder) Head() *ByProjectKeyApplicationsKeyByKeyRequestMethodHead {
	return &ByProjectKeyApplicationsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/applications/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyApplicationsKeyByKeyRequestBuilder) Delete() *ByProjectKeyApplicationsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyApplicationsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/applications/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
