package frontend

// Generated file, please do not change!!!

import (
	"fmt"
)

type ApiBuildUploadRequestBuilder struct {
	client *Client
}

/**
*	Uploads your backend builds from your custom CI environment to the Studio.
 */
func (rb *ApiBuildUploadRequestBuilder) Post(body BuildUpload) *ApiBuildUploadRequestMethodPost {
	return &ApiBuildUploadRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/api/build/upload"),
		client: rb.client,
	}
}
