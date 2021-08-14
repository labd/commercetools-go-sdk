// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyCustomObjectsByContainerRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput
}

func (r *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyCustomObjectsByContainerRequestMethodGetInput struct {
	Where        []string
	PredicateVar map[string][]string
	Expand       []string
}

func (input *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}

	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) WithQueryParams(input *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Get CustomObjectPagedQueryResponse by container
 */
func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Execute(ctx context.Context) (result *CustomObjectPagedQueryResponse, err error) {
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
