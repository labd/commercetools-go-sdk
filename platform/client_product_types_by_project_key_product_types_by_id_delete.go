// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyProductTypesByIdRequestMethodDelete struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyProductTypesByIdRequestMethodDeleteInput
}

func (r *ByProjectKeyProductTypesByIdRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyProductTypesByIdRequestMethodDeleteInput struct {
	Version int
	Expand  *[]string
}

func (input *ByProjectKeyProductTypesByIdRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyProductTypesByIdRequestMethodDelete) WithQueryParams(input *ByProjectKeyProductTypesByIdRequestMethodDeleteInput) *ByProjectKeyProductTypesByIdRequestMethodDelete {
	rb.query = input.Values()
	return rb
}

/**
*	Delete ProductType by ID
 */
func (rb *ByProjectKeyProductTypesByIdRequestMethodDelete) Execute() (result *ProductType, err error) {
	resp, err := rb.client.delete(
		context.Background(),
		rb.url,
		rb.query,
		nil,
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
