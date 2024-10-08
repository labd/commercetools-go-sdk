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
*	Creating a Standalone Price produces the [StandalonePriceCreated](ctp:api:type:StandalonePriceCreatedMessage) Message.
*
*	- If the Standalone Price has the same price scope as an existing Standalone Price, a [DuplicateStandalonePriceScope](ctp:api:type:DuplicateStandalonePriceScopeError) error is returned.
*	- If the Standalone Price has overlapping validity periods within the same price scope, a [OverlappingStandalonePriceValidity](ctp:api:type:OverlappingStandalonePriceValidityError) error is returned. A Price without validity period does not conflict with a Price defined for a time period.
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
