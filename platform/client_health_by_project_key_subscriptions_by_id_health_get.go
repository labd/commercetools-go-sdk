package platform

// Generated file, please do not change!!!

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeySubscriptionsByIDHealthRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeySubscriptionsByIDHealthRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeySubscriptionsByIDHealthRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeySubscriptionsByIDHealthRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	This endpoint can be polled by a monitoring or alerting system that checks the health of your Subscriptions. To ease integration with such systems this endpoint does not require [Authorization](/../api/authorization).
*
 */
func (rb *ByProjectKeySubscriptionsByIDHealthRequestMethodGet) Execute(ctx context.Context) error {
	queryParams := url.Values{}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
