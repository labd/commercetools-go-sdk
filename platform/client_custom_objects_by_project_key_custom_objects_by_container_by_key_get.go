// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput
}

func (r *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) WithQueryParams(input *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Get CustomObject by container and key
 */
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) Execute(ctx context.Context) (result *CustomObject, err error) {
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
