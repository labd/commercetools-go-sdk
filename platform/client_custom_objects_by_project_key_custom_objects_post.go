// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyCustomObjectsRequestMethodPost struct {
	body   CustomObjectDraft
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyCustomObjectsRequestMethodPostInput
}

func (r *ByProjectKeyCustomObjectsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyCustomObjectsRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyCustomObjectsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsRequestMethodPost) WithQueryParams(input *ByProjectKeyCustomObjectsRequestMethodPostInput) *ByProjectKeyCustomObjectsRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Creates a new custom object or updates an existing custom object.
*	If an object with the given container/key exists,
*	the object will be replaced with the new value and the version is incremented.
*	If the request contains a version and an object with the given container/key exists then the version
*	must match the version of the existing object. Concurrent updates for the same custom object still can result
*	in a Conflict (409) even if the version is not provided.
*	Fields with null values will not be saved.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestMethodPost) Execute() (result *CustomObject, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		context.Background(),
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
	case 201:
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
