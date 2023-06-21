package frontend

// Generated file, please do not change!!!

import (
	"fmt"
)

type FrontasticContextRequestBuilder struct {
	client *Client
}

/**
*	Returns information about the project locales setup and the environment in which the requested host acts in. \
*
*	Headers&#58;
*
*	- `Accept` - `application/json` - Required
*
 */
func (rb *FrontasticContextRequestBuilder) Get() *FrontasticContextRequestMethodGet {
	return &FrontasticContextRequestMethodGet{
		url:    fmt.Sprintf("/frontastic/context"),
		client: rb.client,
	}
}
