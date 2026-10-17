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

type ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete) Version(v int) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDeleteInput) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes a [Category](ctp:api:type:Category) by its `id` in the specified [Store](ctp:api:type:Store).
*
*	To delete a global Category, use the [Delete Category by ID](ctp:api:endpoint:/{projectKey}/categories/{id}:DELETE) endpoint.
*
*	If you do not have permissions for a Store the Category is assigned to, an [Unauthorized](ctp:api:type:UnauthorizedError) error is returned.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete) Execute(ctx context.Context) (result *Category, err error) {
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
