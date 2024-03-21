package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput struct {
	Version     *int
	Expand      []string
	DataErasure *bool
}

func (input *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	if input.Version != nil {
		values.Add("version", strconv.Itoa(*input.Version))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	if input.DataErasure != nil {
		if *input.DataErasure {
			values.Add("dataErasure", "true")
		} else {
			values.Add("dataErasure", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) Version(v int) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) Expand(v []string) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) DataErasure(v bool) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDeleteInput) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete) Execute(ctx context.Context) (result *CustomObject, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
