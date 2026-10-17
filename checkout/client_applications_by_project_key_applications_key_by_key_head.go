package checkout

// Generated file, please do not change!!!

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyApplicationsKeyByKeyRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyApplicationsKeyByKeyRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyApplicationsKeyByKeyRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyApplicationsKeyByKeyRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if an [Application](ctp:checkout:type:Application) with a given `key` exists. Returns `200 OK` if the Application exists, or `404 Not Found` if it doesn't. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyApplicationsKeyByKeyRequestMethodHead) Execute(ctx context.Context) error {
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
