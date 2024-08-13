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

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput) Values() url.Values {
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

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) DataErasure(v bool) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Version(v int) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	Deleting an Order produces the [OrderDeleted](ctp:api:type:OrderDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Execute(ctx context.Context) (result *Order, err error) {
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
