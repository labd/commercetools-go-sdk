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

type ByProjectKeyProductsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyProductsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsKeyByKeyRequestMethodGetInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Expand             []string
}

func (input *ByProjectKeyProductsKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
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

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) LocaleProjection(v string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) StoreProjection(v string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyProductsKeyByKeyRequestMethodGetInput) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Gets the full representation of a product by Key.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *Product, err error) {
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
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
