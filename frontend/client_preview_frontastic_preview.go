package frontend

// Generated file, please do not change!!!

import (
	"fmt"
)

type FrontasticPreviewRequestBuilder struct {
	client *Client
}

/**
*	Returns the page structure and data for a specific [preview](/../frontend-studio/using-the-page-builder) by its `previewId`. \
*
*	Headers&#58;
*
*	- `Accept` - `application/json` - Required
*
 */
func (rb *FrontasticPreviewRequestBuilder) Get() *FrontasticPreviewRequestMethodGet {
	return &FrontasticPreviewRequestMethodGet{
		url:    fmt.Sprintf("/frontastic/preview"),
		client: rb.client,
	}
}
