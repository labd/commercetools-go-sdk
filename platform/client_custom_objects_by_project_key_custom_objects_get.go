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

type ByProjectKeyCustomObjectsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsRequestMethodGetInput
}

func (r *ByProjectKeyCustomObjectsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyCustomObjectsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
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

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) Expand(v []string) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) Sort(v []string) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) Limit(v int) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) Offset(v int) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) WithTotal(v bool) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) Where(v []string) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyCustomObjectsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodGet) WithQueryParams(input ByProjectKeyCustomObjectsRequestMethodGetInput) *ByProjectKeyCustomObjectsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	For performance reasons, it is highly advisable to query for Custom Objects in a container by using the `container` field in the `where` predicate.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestMethodGet) Execute(ctx context.Context) (result *CustomObjectPagedQueryResponse, err error) {
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
