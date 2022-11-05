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

type ByProjectKeyAttributeGroupsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyAttributeGroupsRequestMethodGetInput
}

func (r *ByProjectKeyAttributeGroupsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyAttributeGroupsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyAttributeGroupsRequestMethodGetInput) Values() url.Values {
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
			values.Set(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) Expand(v []string) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) Sort(v []string) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) Limit(v int) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) Offset(v int) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) WithTotal(v bool) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) Where(v []string) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyAttributeGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyAttributeGroupsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) WithQueryParams(input ByProjectKeyAttributeGroupsRequestMethodGetInput) *ByProjectKeyAttributeGroupsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyAttributeGroupsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyAttributeGroupsRequestMethodGet) Execute(ctx context.Context) (result *AttributeGroupPagedQueryResponse, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
