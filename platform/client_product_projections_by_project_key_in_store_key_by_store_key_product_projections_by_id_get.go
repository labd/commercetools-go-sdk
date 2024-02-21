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

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput struct {
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	Expand             []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput) Values() url.Values {
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
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) PriceCurrency(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) PriceCountry(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) PriceChannel(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) by its ID from the specified [Store](ctp:api:type:Store).
*	If the Store has defined some languages, countries, distribution, supply Channels, and/or Product Selection,
*	they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices)
*	and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
*
*	If [ProductSelection](ctp:api:type:ProductSelection) is used, it affects the [availability of the Product](/projects/stores#products-available-in-store) in the specified Store.
*
*	When used with an API Client that has the `view_published_products:{projectKey}` scope, this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet) Execute(ctx context.Context) (result *ProductProjection, err error) {
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
