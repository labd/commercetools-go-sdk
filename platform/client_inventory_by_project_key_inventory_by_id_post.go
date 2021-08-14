// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyInventoryByIdRequestMethodPost struct {
	body   InventoryEntryUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyInventoryByIdRequestMethodPostInput
}

func (r *ByProjectKeyInventoryByIdRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyInventoryByIdRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInventoryByIdRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInventoryByIdRequestMethodPost) WithQueryParams(input *ByProjectKeyInventoryByIdRequestMethodPostInput) *ByProjectKeyInventoryByIdRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Update InventoryEntry by ID
 */
func (rb *ByProjectKeyInventoryByIdRequestMethodPost) Execute(ctx context.Context) (result *InventoryEntry, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		rb.query,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, fmt.Errorf("StatusCode %d returend", resp.StatusCode)
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
