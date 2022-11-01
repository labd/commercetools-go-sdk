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

type ByProjectKeySubscriptionsRequestMethodPost struct {
	body    SubscriptionDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeySubscriptionsRequestMethodPostInput
}

func (r *ByProjectKeySubscriptionsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeySubscriptionsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeySubscriptionsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeySubscriptionsRequestMethodPost) Expand(v []string) *ByProjectKeySubscriptionsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodPost) WithQueryParams(input ByProjectKeySubscriptionsRequestMethodPostInput) *ByProjectKeySubscriptionsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeySubscriptionsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeySubscriptionsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	A test message is sent to ensure the correct configuration of the Destination. If the message cannot be delivered, the Subscription will not be created. The payload of the test message is a notification of type [ResourceCreated](/../api/projects/subscriptions#resourcecreateddeliverypayload) for the `resourceTypeId` `subscription`.
*
 */
func (rb *ByProjectKeySubscriptionsRequestMethodPost) Execute(ctx context.Context) (result *Subscription, err error) {
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
