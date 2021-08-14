// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput struct {
	Country  string
	State    *string
	Currency *string
	Expand   *[]string
}

func (input *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("country", input.Country)
	if input.State != nil {
		values.Add("state", *input.State)
	}
	if input.Currency != nil {
		values.Add("currency", *input.Currency)
	}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) WithQueryParams(input *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGetInput) *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	rb.query = input.Values()
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet) Execute() (result *ShippingMethodPagedQueryResponse, err error) {
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
