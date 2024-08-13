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
*	If [Product price selection query parameters](/../api/pricing-and-discounts-overview#product-price-selection) are provided, the selected Prices are added to the response.
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
