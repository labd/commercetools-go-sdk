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

type ByProjectKeyRecurringPaymentJobsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurringPaymentJobsRequestMethodGetInput
}

func (r *ByProjectKeyRecurringPaymentJobsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurringPaymentJobsRequestMethodGetInput struct {
	Sort        []string
	Limit       *int
	Offset      *int
	WithTotal   *bool
	StatusState *string
}

func (input *ByProjectKeyRecurringPaymentJobsRequestMethodGetInput) Values() url.Values {
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
	if input.StatusState != nil {
		values.Add("status.state", fmt.Sprintf("%v", *input.StatusState))
	}
	return values
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) Sort(v []string) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentJobsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) Limit(v int) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentJobsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) Offset(v int) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentJobsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) WithTotal(v bool) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentJobsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) StatusState(v string) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentJobsRequestMethodGetInput{}
	}
	rb.params.StatusState = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) WithQueryParams(input ByProjectKeyRecurringPaymentJobsRequestMethodGetInput) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves Recurring Payment Jobs in a [Project](ctp:api:type:Project).
*
*	The results are [paginated](/api/general-concepts#pagination).
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsRequestMethodGet) Execute(ctx context.Context) (result *PaginatedRecurringPaymentJob, err error) {
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
