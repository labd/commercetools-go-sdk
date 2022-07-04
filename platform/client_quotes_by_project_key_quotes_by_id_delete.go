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

type ByProjectKeyQuotesByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyQuotesByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyQuotesByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyQuotesByIDRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyQuotesByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	if input.DataErasure != nil {
		if *input.DataErasure {
			values.Add("dataErasure", "true")
		} else {
			values.Add("dataErasure", "false")
		}
	}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyQuotesByIDRequestMethodDelete) DataErasure(v bool) *ByProjectKeyQuotesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesByIDRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyQuotesByIDRequestMethodDelete) Version(v int) *ByProjectKeyQuotesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyQuotesByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyQuotesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyQuotesByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyQuotesByIDRequestMethodDeleteInput) *ByProjectKeyQuotesByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyQuotesByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyQuotesByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyQuotesByIDRequestMethodDelete) Execute(ctx context.Context) (result *Quote, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
	case 409, 400, 401, 403, 500, 502, 503:
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
