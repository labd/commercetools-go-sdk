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

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDeleteInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete) Version(v int) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDeleteInput) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Generates the [ProductTailoringDeleted](ctp:api:type:ProductTailoringDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete) Execute(ctx context.Context) (result *ProductTailoring, err error) {
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
