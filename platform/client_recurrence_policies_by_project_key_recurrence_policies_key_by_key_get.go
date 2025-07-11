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

type ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGetInput) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves a Recurrence Policy with the provided `key`.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *RecurrencePolicy, err error) {
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
