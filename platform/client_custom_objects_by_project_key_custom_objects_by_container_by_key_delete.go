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

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput struct {
	Version     *int
	DataErasure *bool
	Expand      []string
}

func (input *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	if input.Version != nil {
		values.Add("version", strconv.Itoa(*input.Version))
	}
	if input.DataErasure != nil {
		if *input.DataErasure == true {
			values.Add("dataErasure", "true")
		} else {
			values.Add("dataErasure", "false")
		}
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) WithQueryParams(input *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	rb.query = input.Values()
	return rb
}

/**
*	Delete CustomObject by container and key
 */
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) Execute(ctx context.Context) (result *CustomObject, err error) {
	resp, err := rb.client.delete(
		ctx,
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
