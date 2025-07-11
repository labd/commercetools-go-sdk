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

type ByProjectKeyMeOrdersRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeOrdersRequestMethodHeadInput
}

func (r *ByProjectKeyMeOrdersRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeOrdersRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyMeOrdersRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeOrdersRequestMethodHead) Where(v []string) *ByProjectKeyMeOrdersRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeOrdersRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyMeOrdersRequestMethodHead) WithQueryParams(input ByProjectKeyMeOrdersRequestMethodHeadInput) *ByProjectKeyMeOrdersRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeOrdersRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyMeOrdersRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more Orders exist for the provided query predicate for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
*
*	- If no Orders exist that match the provided query predicate.
*	- If one or more Orders exist but don't have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeOrdersRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
