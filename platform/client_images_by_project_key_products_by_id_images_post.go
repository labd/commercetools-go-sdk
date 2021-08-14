// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyProductsByIdImagesRequestMethodPost struct {
	body   io.Reader
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyProductsByIdImagesRequestMethodPostInput
}

func (r *ByProjectKeyProductsByIdImagesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyProductsByIdImagesRequestMethodPostInput struct {
	Filename *string
	Variant  *int
	Sku      *string
	Staged   *bool
}

func (input *ByProjectKeyProductsByIdImagesRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Filename != nil {
		values.Add("filename", *input.Filename)
	}
	if input.Variant != nil {
		values.Add("variant", strconv.Itoa(*input.Variant))
	}
	if input.Sku != nil {
		values.Add("sku", *input.Sku)
	}
	if input.Staged != nil {
		if *input.Staged == true {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyProductsByIdImagesRequestMethodPost) WithQueryParams(input *ByProjectKeyProductsByIdImagesRequestMethodPostInput) *ByProjectKeyProductsByIdImagesRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Uploads a binary image file to a given product variant. The supported image formats are JPEG, PNG and GIF.
*
 */
func (rb *ByProjectKeyProductsByIdImagesRequestMethodPost) Execute() (result *Product, err error) {
	data := rb.body
	resp, err := rb.client.post(
		context.Background(),
		rb.url,
		rb.query,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
