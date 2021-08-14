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

type ByProjectKeyProductsRequestMethodPost struct {
	body    ProductDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsRequestMethodPostInput
}

func (r *ByProjectKeyProductsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsRequestMethodPostInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Expand             []string
}

func (input *ByProjectKeyProductsRequestMethodPostInput) Values() url.Values {
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

func (rb *ByProjectKeyProductsRequestMethodPost) PriceCurrency(v string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) PriceCountry(v string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) PriceCustomerGroup(v string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) PriceChannel(v string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) LocaleProjection(v string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) StoreProjection(v string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) Expand(v []string) *ByProjectKeyProductsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsRequestMethodPost) WithQueryParams(input ByProjectKeyProductsRequestMethodPostInput) *ByProjectKeyProductsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	To create a new product, send a representation that is going to become the initial staged representation
*	of the new product in the master catalog. If price selection query parameters are provided,
*	the selected prices will be added to the response.
*
 */
func (rb *ByProjectKeyProductsRequestMethodPost) Execute(ctx context.Context) (result *Product, err error) {
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 201:
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
