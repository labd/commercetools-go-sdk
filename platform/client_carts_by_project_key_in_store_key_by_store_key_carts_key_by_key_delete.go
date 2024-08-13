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

type ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	if input.DataErasure != nil {
		if *input.DataErasure {
			values.Add("dataErasure", "true")
		} else {
			values.Add("dataErasure", "false")
		}
	}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) DataErasure(v bool) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDeleteInput) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	If the Cart exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *Cart, err error) {
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
