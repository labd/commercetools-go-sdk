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

type ByProjectKeyStandalonePricesRequestMethodPost struct {
	body    StandalonePriceDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStandalonePricesRequestMethodPostInput
}

func (r *ByProjectKeyStandalonePricesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStandalonePricesRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyStandalonePricesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStandalonePricesRequestMethodPost) Expand(v []string) *ByProjectKeyStandalonePricesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodPost) WithQueryParams(input ByProjectKeyStandalonePricesRequestMethodPostInput) *ByProjectKeyStandalonePricesRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStandalonePricesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyStandalonePricesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Produces the [StandalonePriceCreated](ctp:api:type:StandalonePriceCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyStandalonePricesRequestMethodPost) Execute(ctx context.Context) (result *StandalonePrice, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 201:
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
