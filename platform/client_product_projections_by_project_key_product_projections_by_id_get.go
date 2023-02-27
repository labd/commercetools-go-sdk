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

type ByProjectKeyProductProjectionsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductProjectionsByIDRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductProjectionsByIDRequestMethodGetInput struct {
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	StoreProjection    *string
	Expand             []string
}

func (input *ByProjectKeyProductProjectionsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	if input.PriceCurrency != nil {
		values.Add("priceCurrency", fmt.Sprintf("%v", *input.PriceCurrency))
	}
	if input.PriceCountry != nil {
		values.Add("priceCountry", fmt.Sprintf("%v", *input.PriceCountry))
	}
	if input.PriceCustomerGroup != nil {
		values.Add("priceCustomerGroup", fmt.Sprintf("%v", *input.PriceCustomerGroup))
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", fmt.Sprintf("%v", *input.PriceChannel))
	}
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	if input.StoreProjection != nil {
		values.Add("storeProjection", fmt.Sprintf("%v", *input.StoreProjection))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) StoreProjection(v string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyProductProjectionsByIDRequestMethodGetInput) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) by its ID. When used with an API Client that has the `view_published_products:{projectKey}` scope, this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyProductProjectionsByIDRequestMethodGet) Execute(ctx context.Context) (result *ProductProjection, err error) {
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
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
