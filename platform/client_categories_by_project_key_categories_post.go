package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyCategoriesRequestMethodPost struct {
	body    CategoryDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCategoriesRequestMethodPostInput
}

func (r *ByProjectKeyCategoriesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCategoriesRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyCategoriesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCategoriesRequestMethodPost) Expand(v []string) *ByProjectKeyCategoriesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyCategoriesRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCategoriesRequestMethodPost) WithQueryParams(input ByProjectKeyCategoriesRequestMethodPostInput) *ByProjectKeyCategoriesRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCategoriesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyCategoriesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Either the [scope](/../api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
*	Creating a Category produces the [CategoryCreated](ctp:api:type:CategoryCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyCategoriesRequestMethodPost) Execute(ctx context.Context) (result *Category, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 502, 503:
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
