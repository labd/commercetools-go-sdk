// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyZonesKeyByKeyRequestMethodPost struct {
	body   ZoneUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyZonesKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyZonesKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyZonesKeyByKeyRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyZonesKeyByKeyRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyZonesKeyByKeyRequestMethodPost) WithQueryParams(input *ByProjectKeyZonesKeyByKeyRequestMethodPostInput) *ByProjectKeyZonesKeyByKeyRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Update Zone by key
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *Zone, err error) {
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
