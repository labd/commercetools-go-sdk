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

type ByProjectKeyInventoryRequestMethodPost struct {
	body    InventoryEntryDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInventoryRequestMethodPostInput
}

func (r *ByProjectKeyInventoryRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInventoryRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInventoryRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInventoryRequestMethodPost) Expand(v []string) *ByProjectKeyInventoryRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInventoryRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInventoryRequestMethodPost) WithQueryParams(input ByProjectKeyInventoryRequestMethodPostInput) *ByProjectKeyInventoryRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInventoryRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInventoryRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Produces the [InventoryEntryCreatedMessage](ctp:api:type:InventoryEntryCreatedMessage).
 */
func (rb *ByProjectKeyInventoryRequestMethodPost) Execute(ctx context.Context) (result *InventoryEntry, err error) {
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
