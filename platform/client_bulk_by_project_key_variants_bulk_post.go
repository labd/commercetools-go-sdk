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

type ByProjectKeyVariantsBulkRequestMethodPost struct {
	body    VariantBulkUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyVariantsBulkRequestMethodPostInput
}

func (r *ByProjectKeyVariantsBulkRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyVariantsBulkRequestMethodPostInput struct {
	VersionControl *string
}

func (input *ByProjectKeyVariantsBulkRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.VersionControl != nil {
		values.Add("versionControl", fmt.Sprintf("%v", *input.VersionControl))
	}
	return values
}

func (rb *ByProjectKeyVariantsBulkRequestMethodPost) VersionControl(v string) *ByProjectKeyVariantsBulkRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantsBulkRequestMethodPostInput{}
	}
	rb.params.VersionControl = &v
	return rb
}

func (rb *ByProjectKeyVariantsBulkRequestMethodPost) WithQueryParams(input ByProjectKeyVariantsBulkRequestMethodPostInput) *ByProjectKeyVariantsBulkRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyVariantsBulkRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyVariantsBulkRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Applies update actions to multiple Variants in a single request.
*	Returns a `207 Multi-Status` response with the result for each Variant.
*	Each Variant is updated independently, so some may succeed while others fail.
*
 */
func (rb *ByProjectKeyVariantsBulkRequestMethodPost) Execute(ctx context.Context) (result *VariantBulkUpdateResponse, err error) {
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
	case 207:
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
