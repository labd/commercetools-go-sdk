package frontend

// Generated file, please do not change!!!

import (
	"fmt"
)

type FrontasticPageRequestBuilder struct {
	client *Client
}

/**
*	Returns the structure and data of the published page displayed from commercetools Frontend. \
*
*	Headers&#58;
*
*	- `Frontastic-Path` - `^/.*$` - Required
*	- `Accept` - `application/json` - Required
*	- `Frontastic-Locale` - [Locale](ctp:frontend-api:type:Locale) - Required
*
 */
func (rb *FrontasticPageRequestBuilder) Get() *FrontasticPageRequestMethodGet {
	return &FrontasticPageRequestMethodGet{
		url:    fmt.Sprintf("/frontastic/page"),
		client: rb.client,
	}
}
