package frontend

// Generated file, please do not change!!!

import (
	"fmt"
)

type FrontasticActionByNamespaceByActionRequestBuilder struct {
	namespace string
	action    string
	client    *Client
}

/**
*	Use the GET method to allow the frontend to fetch data from a backend system. For the response, we recommend to use standard HTTP codes and `application/json` encoded content. The response will be structured [as defined by the `body` property of the action](/../frontend-development/developing-an-action-extension#1-implement-the-action). The following response example contains information about a cart.
 */
func (rb *FrontasticActionByNamespaceByActionRequestBuilder) Get() *FrontasticActionByNamespaceByActionRequestMethodGet {
	return &FrontasticActionByNamespaceByActionRequestMethodGet{
		url:    fmt.Sprintf("/frontastic/action/%s/%s", rb.namespace, rb.action),
		client: rb.client,
	}
}

/**
*	Use the POST method to write data to a backend system. Any JSON serializable payload is accepted. The following request example adds a product to a cart. For the response, we recommend to use standard HTTP codes and `application/json` encoded content. The response will be structured [as defined by the `body` property of the action](/../frontend-development/developing-an-action-extension#1-implement-the-action). The following response example contains the updated cart information, which includes the added product.
 */
func (rb *FrontasticActionByNamespaceByActionRequestBuilder) Post(body interface{}) *FrontasticActionByNamespaceByActionRequestMethodPost {
	return &FrontasticActionByNamespaceByActionRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/frontastic/action/%s/%s", rb.namespace, rb.action),
		client: rb.client,
	}
}

/**
*	Use the PUT method to write data to a backend system. Any JSON serializable payload is accepted. The following request example adds a product to a cart. For the response, we recommend to use standard HTTP codes and `application/json` encoded content. The response will be structured [as defined by the `body` property of the action](/../frontend-development/developing-an-action-extension#1-implement-the-action). The following response example contains the updated cart information, which includes the added product.
 */
func (rb *FrontasticActionByNamespaceByActionRequestBuilder) Put(body interface{}) *FrontasticActionByNamespaceByActionRequestMethodPut {
	return &FrontasticActionByNamespaceByActionRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/frontastic/action/%s/%s", rb.namespace, rb.action),
		client: rb.client,
	}
}
