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

type ByProjectKeyProductTypesRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductTypesRequestMethodHeadInput
}

func (r *ByProjectKeyProductTypesRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductTypesRequestMethodHeadInput struct {
	Where *string
}

func (input *ByProjectKeyProductTypesRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	if input.Where != nil {
		values.Add("where", fmt.Sprintf("%v", *input.Where))
	}
	return values
}

func (rb *ByProjectKeyProductTypesRequestMethodHead) Where(v string) *ByProjectKeyProductTypesRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductTypesRequestMethodHeadInput{}
	}
	rb.params.Where = &v
	return rb
}

func (rb *ByProjectKeyProductTypesRequestMethodHead) WithQueryParams(input ByProjectKeyProductTypesRequestMethodHeadInput) *ByProjectKeyProductTypesRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductTypesRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyProductTypesRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Check if Product Types exist. Responds with a `200 OK` status if any Product Types match the Query Predicate, or `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductTypesRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
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

	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}