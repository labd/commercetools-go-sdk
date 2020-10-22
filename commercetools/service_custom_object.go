// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// CustomObjectCreate creates a new instance of type CustomObject
func (client *Client) CustomObjectCreate(ctx context.Context, draft *CustomObjectDraft, opts ...RequestOption) (result *CustomObject, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "custom-objects"
	err = client.create(ctx, endpoint, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomObjectQuery allows querying for type CustomObject
/*
The query endpoint allows to retrieve custom objects in a specific container or all custom objects.
For performance reasons, it is highly advisable to query only for custom objects in a container by using
the container field in the where predicate.
*/
func (client *Client) CustomObjectQuery(ctx context.Context, input *QueryInput) (result *CustomObjectPagedQueryResponse, err error) {
	endpoint := "custom-objects"
	err = client.query(ctx, endpoint, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomObjectDeleteWithContainerAndKey Delete CustomObject by container and key
func (client *Client) CustomObjectDeleteWithContainerAndKey(ctx context.Context, container string, key string, version int, dataErasure bool, opts ...RequestOption) (result *CustomObject, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("custom-objects/%s/%s", container, key)
	err = client.delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomObjectGetWithContainer for type CustomObject
func (client *Client) CustomObjectGetWithContainer(ctx context.Context, container string, opts ...RequestOption) (result *CustomObject, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("custom-objects/%s", container)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CustomObjectGetWithContainerAndKey Get CustomObject by container and key
func (client *Client) CustomObjectGetWithContainerAndKey(ctx context.Context, container string, key string, opts ...RequestOption) (result *CustomObject, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("custom-objects/%s/%s", container, key)
	err = client.get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
