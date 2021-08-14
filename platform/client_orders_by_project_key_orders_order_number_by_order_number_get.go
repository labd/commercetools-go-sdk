// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput
}

func (r *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) WithQueryParams(input *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
*	it should be provided in URL-encoded format.
*
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Execute(ctx context.Context) (result *Order, err error) {
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
