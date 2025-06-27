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

type ByProjectKeyDiscountGroupsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDiscountGroupsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyDiscountGroupsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDiscountGroupsByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyDiscountGroupsByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodDelete) Version(v int) *ByProjectKeyDiscountGroupsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountGroupsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyDiscountGroupsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountGroupsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyDiscountGroupsByIDRequestMethodDeleteInput) *ByProjectKeyDiscountGroupsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyDiscountGroupsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes a DiscountGroup in the Project.
*	This request generates the [DiscountGroupDeleted](ctp:api:type:DiscountGroupDeletedMessage) Message.
*
*	If the DiscountGroup is referenced by a CartDiscount, a [ReferenceExists](ctp:api:type:ReferenceExistsError) error is returned.
*
 */
func (rb *ByProjectKeyDiscountGroupsByIDRequestMethodDelete) Execute(ctx context.Context) (result *DiscountGroup, err error) {
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
