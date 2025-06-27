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

type ByProjectKeyDiscountGroupsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDiscountGroupsByIDRequestMethodGetInput
}

func (r *ByProjectKeyDiscountGroupsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDiscountGroupsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyDiscountGroupsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyDiscountGroupsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountGroupsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyDiscountGroupsByIDRequestMethodGetInput) *ByProjectKeyDiscountGroupsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDiscountGroupsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves a DiscountGroup with the provided `id`.
*
 */
func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodGet) Execute(ctx context.Context) (result *DiscountGroup, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
