package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyOrdersSearchRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyOrdersSearchRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyOrdersSearchRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyOrdersSearchRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks whether a search index for the Project's Orders exists.
 */
func (rb *ByProjectKeyOrdersSearchRequestMethodHead) Execute(ctx context.Context) error {
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
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return result
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		return fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
