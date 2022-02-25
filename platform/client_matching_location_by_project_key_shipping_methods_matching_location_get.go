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

type ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput struct {
	Country  string
	State    *string
	Currency *string
	Expand   []string
	Sort     []string
}

func (input *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("country", fmt.Sprintf("%v", input.Country))
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
	}
	if input.Currency != nil {
		values.Add("currency", fmt.Sprintf("%v", *input.Currency))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Country(v string) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput{}
	}
	rb.params.Country = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) State(v string) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Currency(v string) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput{}
	}
	rb.params.Currency = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Expand(v []string) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Sort(v []string) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) WithQueryParams(input ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
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
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
