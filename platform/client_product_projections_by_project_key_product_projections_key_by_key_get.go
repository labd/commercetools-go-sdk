// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput struct {
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Expand             []string
}

func (input *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Staged != nil {
		if *input.Staged == true {
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
	if input.LocaleProjection != nil {
		values.Add("localeProjection", fmt.Sprintf("%v", *input.LocaleProjection))
	}
	if input.StoreProjection != nil {
		values.Add("storeProjection", fmt.Sprintf("%v", *input.StoreProjection))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) LocaleProjection(v string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) StoreProjection(v string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyProductProjectionsKeyByKeyRequestMethodGetInput) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Gets the current or staged representation of a product found by Key.
*	When used with an API client that has the view_published_products:{projectKey} scope,
*	this endpoint only returns published (current) product projections.
*
 */
func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *ProductProjection, err error) {
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
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
