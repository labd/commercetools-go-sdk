// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput struct {
	OrderEditId string
	Country     string
	State       *string
}

func (input *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("orderEditId", fmt.Sprintf("%v", input.OrderEditId))
	values.Add("country", fmt.Sprintf("%v", input.Country))
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) WithQueryParams(input *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	rb.query = input.Values()
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
	resp, err := rb.client.get(
		ctx,
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
