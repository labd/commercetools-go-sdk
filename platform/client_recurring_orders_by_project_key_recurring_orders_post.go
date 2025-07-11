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

type ByProjectKeyRecurringOrdersRequestMethodPost struct {
	body    RecurringOrderDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecurringOrdersRequestMethodPostInput
}

func (r *ByProjectKeyRecurringOrdersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecurringOrdersRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyRecurringOrdersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyRecurringOrdersRequestMethodPost) Expand(v []string) *ByProjectKeyRecurringOrdersRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecurringOrdersRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyRecurringOrdersRequestMethodPost) WithQueryParams(input ByProjectKeyRecurringOrdersRequestMethodPostInput) *ByProjectKeyRecurringOrdersRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecurringOrdersRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyRecurringOrdersRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a Recurring Order in the Project.
*	The Cart is validated to ensure that it is convertible to an [Order](ctp:api:type:Order). If validation fails, an error is returned.
*
*	Produces the [RecurringOrderCreated](ctp:api:type:RecurringOrderCreatedMessage) message.
*
*	If a server-side problem occurs, indicated by a 500 Internal Server Error HTTP response, the Recurring Order creation may still successfully complete after the error is returned.
*	If you receive this error, you should verify the status of the Recurring Order by querying a unique identifier supplied during the creation request, such as the key.
*
 */
func (rb *ByProjectKeyRecurringOrdersRequestMethodPost) Execute(ctx context.Context) (result *RecurringOrder, err error) {
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
