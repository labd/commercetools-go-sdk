package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput struct {
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Expand       []string
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) Limit(v int) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) Offset(v int) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) WithTotal(v bool) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) Where(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Queries Product Selection assignments in a specific [Store](ctp:api:type:Store).
*
*	The response will include duplicate Products whenever more than one active Product Selection of the Store
*	includes a Product. To make clear through which Product Selection a Product is available in the Store
*	the response contains assignments including both the Product and the Product Selection.
*	Only Products of Product Selections that are activated in the Store will be returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet) Execute(ctx context.Context) (result *ProductsInStorePagedQueryResponse, err error) {
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
