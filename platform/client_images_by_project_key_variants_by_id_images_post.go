package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyVariantsByIDImagesRequestMethodPost struct {
	body    io.Reader
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyVariantsByIDImagesRequestMethodPostInput
}

func (r *ByProjectKeyVariantsByIDImagesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyVariantsByIDImagesRequestMethodPostInput struct {
	Filename *string
	Staged   *bool
}

func (input *ByProjectKeyVariantsByIDImagesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Filename != nil {
		values.Add("filename", fmt.Sprintf("%v", *input.Filename))
	}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyVariantsByIDImagesRequestMethodPost) Filename(v string) *ByProjectKeyVariantsByIDImagesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantsByIDImagesRequestMethodPostInput{}
	}
	rb.params.Filename = &v
	return rb
}

func (rb *ByProjectKeyVariantsByIDImagesRequestMethodPost) Staged(v bool) *ByProjectKeyVariantsByIDImagesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantsByIDImagesRequestMethodPostInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyVariantsByIDImagesRequestMethodPost) WithQueryParams(input ByProjectKeyVariantsByIDImagesRequestMethodPostInput) *ByProjectKeyVariantsByIDImagesRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyVariantsByIDImagesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyVariantsByIDImagesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Uploads a JPEG, PNG, or a GIF image file to a [Variant](ctp:api:type:Variant) identified by its ID.
*	The maximum file size of the image is **10MB**.
*
*	The response status code depends on the size of the original image.
*	If the image is small, the API responds with `200 OK`, and if the image is larger, it responds with `202 Accepted`.
*	The Variant returned with a `202 Accepted` status code contains a `warnings` field with an [ImageProcessingOngoing](ctp:api:type:ImageProcessingOngoingWarning) Warning.
*
*	Produces the [VariantImageAdded](/projects/messages/product-catalog-messages#variant-image-added) Message.
*
 */
func (rb *ByProjectKeyVariantsByIDImagesRequestMethodPost) Execute(ctx context.Context) (result *Variant, err error) {
	data := rb.body
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
	case 200:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
