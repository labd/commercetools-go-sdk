// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyDiscountCodesByIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyDiscountCodesByIdRequestMethodGetInput
}

func (r *ByProjectKeyDiscountCodesByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyDiscountCodesByIdRequestMethodGetInput struct {
	Expand *[]string
}

func (input *ByProjectKeyDiscountCodesByIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyDiscountCodesByIdRequestMethodGet) WithQueryParams(input *ByProjectKeyDiscountCodesByIdRequestMethodGetInput) *ByProjectKeyDiscountCodesByIdRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Get DiscountCode by ID
 */
func (rb *ByProjectKeyDiscountCodesByIdRequestMethodGet) Execute() (result *DiscountCode, err error) {
	resp, err := rb.client.get(
		context.Background(),
		rb.url,
		rb.query,
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
	case 400, 401, 403, 500, 503:
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
