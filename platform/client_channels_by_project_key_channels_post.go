// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyChannelsRequestMethodPost struct {
	body    ChannelDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyChannelsRequestMethodPostInput
}

func (r *ByProjectKeyChannelsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyChannelsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyChannelsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyChannelsRequestMethodPost) Expand(v []string) *ByProjectKeyChannelsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyChannelsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyChannelsRequestMethodPost) WithQueryParams(input ByProjectKeyChannelsRequestMethodPostInput) *ByProjectKeyChannelsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyChannelsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyChannelsRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyChannelsRequestMethodPost) Execute(ctx context.Context) (result *Channel, err error) {
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
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
