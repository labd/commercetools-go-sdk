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

type ByProjectKeyPaymentIntegrationsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentIntegrationsRequestMethodGetInput
}

func (r *ByProjectKeyPaymentIntegrationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentIntegrationsRequestMethodGetInput struct {
	Sort      []string
	Limit     *int
	Offset    *int
	WithTotal *bool
}

func (input *ByProjectKeyPaymentIntegrationsRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) Sort(v []string) *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentIntegrationsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) Limit(v int) *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentIntegrationsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) Offset(v int) *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentIntegrationsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) WithTotal(v bool) *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentIntegrationsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) WithQueryParams(input ByProjectKeyPaymentIntegrationsRequestMethodGetInput) *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves Payment Integrations in a [Project](ctp:api:type:Project).
*
*	The results are [paginated](/api/general-concepts#pagination).
*
 */
func (rb *ByProjectKeyPaymentIntegrationsRequestMethodGet) Execute(ctx context.Context) (result *PaginatedPaymentIntegration, err error) {
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
