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

type ByProjectKeyStagedQuotesByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStagedQuotesByIDRequestMethodGetInput
}

func (r *ByProjectKeyStagedQuotesByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStagedQuotesByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyStagedQuotesByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStagedQuotesByIDRequestMethodGet) Expand(v []string) *ByProjectKeyStagedQuotesByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStagedQuotesByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStagedQuotesByIDRequestMethodGet) WithQueryParams(input ByProjectKeyStagedQuotesByIDRequestMethodGetInput) *ByProjectKeyStagedQuotesByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStagedQuotesByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyStagedQuotesByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStagedQuotesByIDRequestMethodGet) Execute(ctx context.Context) (result *StagedQuote, err error) {
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
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
