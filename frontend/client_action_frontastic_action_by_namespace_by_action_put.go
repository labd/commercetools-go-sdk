package frontend

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FrontasticActionByNamespaceByActionRequestMethodPut struct {
	body    interface{}
	url     string
	client  *Client
	headers http.Header
}

func (r *FrontasticActionByNamespaceByActionRequestMethodPut) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *FrontasticActionByNamespaceByActionRequestMethodPut) WithHeaders(headers http.Header) *FrontasticActionByNamespaceByActionRequestMethodPut {
	rb.headers = headers
	return rb
}

/**
*	Use the PUT method to write data to a backend system. Any JSON serializable payload is accepted. The following request example adds a product to a cart. For the response, we recommend to use standard HTTP codes and `application/json` encoded content. The response will be structured [as defined by the `body` property of the action](/../frontend-development/developing-an-action-extension#1-implement-the-action). The following response example contains the updated cart information, which includes the added product.
 */
func (rb *FrontasticActionByNamespaceByActionRequestMethodPut) Execute(ctx context.Context) (result *interface{}, err error) {
	queryParams := url.Values{}
	resp, err := rb.client.put(
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
	case 400:
		errorObj := Error{}
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
