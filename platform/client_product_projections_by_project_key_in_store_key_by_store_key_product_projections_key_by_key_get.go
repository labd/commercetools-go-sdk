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

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput struct {
	Staged                        *bool
	PriceCurrency                 *string
	PriceCountry                  *string
	PriceCustomerGroup            *string
	PriceCustomerGroupAssignments []string
	PriceChannel                  *string
	PriceRecurrencePolicy         *string
	LocaleProjection              []string
	Expand                        []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput) Values() url.Values {
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
	for _, v := range input.PriceCustomerGroupAssignments {
		values.Add("priceCustomerGroupAssignments", fmt.Sprintf("%v", v))
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", fmt.Sprintf("%v", *input.PriceChannel))
	}
	if input.PriceRecurrencePolicy != nil {
		values.Add("priceRecurrencePolicy", fmt.Sprintf("%v", *input.PriceRecurrencePolicy))
	}
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCurrency(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCountry(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroupAssignments(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) PriceChannel(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) PriceRecurrencePolicy(v string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) by its key in the specified [Store](ctp:api:type:Store).
*	If the Store has defined some languages, countries, distribution, supply Channels, and/or Product Selection,
*	they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices),
*	and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
*
*	If [ProductSelection](ctp:api:type:ProductSelection) is used, it affects the [availability of the Product](/projects/stores#products-available-in-store) in the specified Store.
*
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for the Product with the given `key` and the given Store, this endpoint returns the ProductProjection with tailored data.
*
*	When used with an API Client that has the `view_published_products:{projectKey}` scope, this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *ProductProjection, err error) {
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
