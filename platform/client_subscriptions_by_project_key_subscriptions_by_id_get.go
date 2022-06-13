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

type ByProjectKeySubscriptionsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeySubscriptionsByIDRequestMethodGetInput
}

func (r *ByProjectKeySubscriptionsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeySubscriptionsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeySubscriptionsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeySubscriptionsByIDRequestMethodGet) Expand(v []string) *ByProjectKeySubscriptionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeySubscriptionsByIDRequestMethodGet) WithQueryParams(input ByProjectKeySubscriptionsByIDRequestMethodGetInput) *ByProjectKeySubscriptionsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeySubscriptionsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeySubscriptionsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the representation of a subscription by its id.
 */
func (rb *ByProjectKeySubscriptionsByIDRequestMethodGet) Execute(ctx context.Context) (result *Subscription, err error) {
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
