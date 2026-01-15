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

type ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete) Expand(v []string) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDeleteInput) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes a Recurrence Policy in the Project.
*
*	A Recurrence Policy can be deleted only if it is not referenced by any Embedded Price, Standalone Price, or (Custom) Line Item, otherwise a [ReferenceExists](ctp:api:type:ReferenceExistsError) error is returned.
*
*	The `manage_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `manage_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *RecurrencePolicy, err error) {
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
