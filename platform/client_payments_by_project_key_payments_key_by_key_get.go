// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyPaymentsKeyByKeyRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyPaymentsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyPaymentsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyPaymentsKeyByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyPaymentsKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyPaymentsKeyByKeyRequestMethodGet) WithQueryParams(input *ByProjectKeyPaymentsKeyByKeyRequestMethodGetInput) *ByProjectKeyPaymentsKeyByKeyRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Get Payment by key
 */
func (rb *ByProjectKeyPaymentsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *Payment, err error) {
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
