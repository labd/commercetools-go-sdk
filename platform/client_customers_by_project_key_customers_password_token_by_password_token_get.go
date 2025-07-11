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

type ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput
}

func (r *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) Expand(v []string) *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) WithQueryParams(input ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGetInput) *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Use this method to retrieve the details of a global Customer by using the password token during their [password reset process](/../api/customers-overview#customer-password-reset).
*
 */
func (rb *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet) Execute(ctx context.Context) (result *Customer, err error) {
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
	case 404:
		return nil, ErrNotFound
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
