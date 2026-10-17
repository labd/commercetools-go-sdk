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

type ByProjectKeyRecurringPaymentsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurringPaymentsRequestMethodGetInput
}

func (r *ByProjectKeyRecurringPaymentsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurringPaymentsRequestMethodGetInput struct {
	Sort             []string
	Limit            *int
	Offset           *int
	WithTotal        *bool
	RecurringOrderId *string
}

func (input *ByProjectKeyRecurringPaymentsRequestMethodGetInput) Values() url.Values {
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
	if input.RecurringOrderId != nil {
		values.Add("recurringOrderId", fmt.Sprintf("%v", *input.RecurringOrderId))
	}
	return values
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) Sort(v []string) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) Limit(v int) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) Offset(v int) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) WithTotal(v bool) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) RecurringOrderId(v string) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringPaymentsRequestMethodGetInput{}
	}
	rb.params.RecurringOrderId = &v
	return rb
}

func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) WithQueryParams(input ByProjectKeyRecurringPaymentsRequestMethodGetInput) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRecurringPaymentsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves Recurring Payments in a [Project](ctp:api:type:Project).
*
*	The results are [paginated](/api/general-concepts#pagination).
*
 */
func (rb *ByProjectKeyRecurringPaymentsRequestMethodGet) Execute(ctx context.Context) (result *PaginatedRecurringPayment, err error) {
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
