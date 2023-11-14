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

type ByProjectKeyStagedQuotesRequestMethodPost struct {
	body    StagedQuoteDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStagedQuotesRequestMethodPostInput
}

func (r *ByProjectKeyStagedQuotesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStagedQuotesRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyStagedQuotesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStagedQuotesRequestMethodPost) Expand(v []string) *ByProjectKeyStagedQuotesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyStagedQuotesRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStagedQuotesRequestMethodPost) WithQueryParams(input ByProjectKeyStagedQuotesRequestMethodPostInput) *ByProjectKeyStagedQuotesRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStagedQuotesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyStagedQuotesRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStagedQuotesRequestMethodPost) Execute(ctx context.Context) (result *StagedQuote, err error) {
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
