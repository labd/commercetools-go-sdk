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
	"strconv"
)

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost struct {
	body    io.Reader
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput struct {
	Filename *string
	Variant  *int
	Sku      *string
	Staged   *bool
}

func (input *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Filename != nil {
		values.Add("filename", fmt.Sprintf("%v", *input.Filename))
	}
	if input.Variant != nil {
		values.Add("variant", strconv.Itoa(*input.Variant))
	}
	if input.Sku != nil {
		values.Add("sku", fmt.Sprintf("%v", *input.Sku))
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

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) Filename(v string) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput{}
	}
	rb.params.Filename = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) Variant(v int) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput{}
	}
	rb.params.Variant = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) Sku(v string) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput{}
	}
	rb.params.Sku = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Uploads a JPEG, PNG and GIF file to a [ProductVariantTailoring](ctp:api:type:ProductVariantTailoring).
*	The maximum file size of the image is **10MB**.
*	Either `variant` or `sku` is required to update a specific ProductVariant.
*	If neither is provided, the image is uploaded to the Master Variant of the Product.
*
*	The response status code depends on the size of the original image.
*	If the image is small, the API responds with `200 OK`, and if the image is larger, it responds with `202 Accepted`.
*	The Product returned with a `202 Accepted` status code contains a `warnings` field with an [ImageProcessingOngoing](ctp:api:type:ImageProcessingOngoingWarning) Warning.
*
*	Produces the [ProductTailoringImageAdded](/projects/messages/product-catalog-messages#product-tailoring-image-added) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost) Execute(ctx context.Context) (result *ProductTailoring, err error) {
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