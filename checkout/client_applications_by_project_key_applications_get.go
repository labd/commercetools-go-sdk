package checkout

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

type ByProjectKeyApplicationsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyApplicationsRequestMethodGetInput
}

func (r *ByProjectKeyApplicationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyApplicationsRequestMethodGetInput struct {
	Sort      []string
	Limit     *int
	Offset    *int
	WithTotal *bool
}

func (input *ByProjectKeyApplicationsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
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
	return values
}

func (rb *ByProjectKeyApplicationsRequestMethodGet) Sort(v []string) *ByProjectKeyApplicationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApplicationsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyApplicationsRequestMethodGet) Limit(v int) *ByProjectKeyApplicationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApplicationsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyApplicationsRequestMethodGet) Offset(v int) *ByProjectKeyApplicationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApplicationsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyApplicationsRequestMethodGet) WithTotal(v bool) *ByProjectKeyApplicationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApplicationsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyApplicationsRequestMethodGet) WithQueryParams(input ByProjectKeyApplicationsRequestMethodGetInput) *ByProjectKeyApplicationsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyApplicationsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyApplicationsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves Applications
*
 */
func (rb *ByProjectKeyApplicationsRequestMethodGet) Execute(ctx context.Context) (result *PaginatedApplication, err error) {
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
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
