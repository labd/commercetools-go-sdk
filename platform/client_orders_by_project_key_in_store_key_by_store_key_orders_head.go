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

type ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHeadInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead) Where(v []string) *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHeadInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if an Order exists for a given Query Predicate. Returns a `200 OK` status if any Orders match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead) Execute(ctx context.Context) error {
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
