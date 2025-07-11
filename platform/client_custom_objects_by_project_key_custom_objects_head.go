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

type ByProjectKeyCustomObjectsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsRequestMethodHeadInput
}

func (r *ByProjectKeyCustomObjectsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyCustomObjectsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsRequestMethodHead) Where(v []string) *ByProjectKeyCustomObjectsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodHead) WithQueryParams(input ByProjectKeyCustomObjectsRequestMethodHeadInput) *ByProjectKeyCustomObjectsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more CustomObjects exist for the provided query predicate. Returns a `200 OK` status if any CustomObjects match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCustomObjectsRequestMethodHead) Execute(ctx context.Context) error {
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
	case 404:
		return ErrNotFound
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 503:
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
