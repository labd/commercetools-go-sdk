package platform

// Generated file, please do not change!!!

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if a Project exists for a given `projectKey`. Returns a `200 OK` status if the Project exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyRequestMethodHead) Execute(ctx context.Context) error {
	queryParams := url.Values{}
	resp, err := rb.client.head(
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
