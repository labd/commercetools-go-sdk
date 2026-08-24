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

type ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput struct {
	Staged               *bool
	LocaleProjection     []string
	FilterAttributes     []string
	FilterSupplyChannels []string
}

func (input *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterAttributes {
		values.Add("filter[attributes]", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterSupplyChannels {
		values.Add("filter[supplyChannels]", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) FilterAttributes(v []string) *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) FilterSupplyChannels(v []string) *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput{}
	}
	rb.params.FilterSupplyChannels = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) WithQueryParams(input ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGetInput) *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns a lightweight representation of all Variants for a [Product](ctp:api:type:Product), including only the requested variant-level [Attributes](ctp:api:type:Attribute) and minimal availability data.
*
*	Designed for building attribute selectors on product detail pages (PDPs) with large numbers of variants.
*	Only available for Projects with [productCatalogModel](ctp:api:type:ProductCatalogModel) set to `Modular`.
*
*	Product-level Attributes are omitted from the response.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Products (published data), the `view_published_products:{projectKey}` or `view_products:{projectKey}` scope is required.
*	- To retrieve the staged representation of Products (draft data), the API Client must have the `view_products:{projectKey}` scope.
*
 */
func (rb *ByProjectKeyProductProjectionsByIDVariantAttributesRequestMethodGet) Execute(ctx context.Context) (result *VariantAttributes, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
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
