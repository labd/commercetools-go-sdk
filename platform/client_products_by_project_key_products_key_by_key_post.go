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

type ByProjectKeyProductsKeyByKeyRequestMethodPost struct {
	body    ProductUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyProductsKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsKeyByKeyRequestMethodPostInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	Expand             []string
}

func (input *ByProjectKeyProductsKeyByKeyRequestMethodPostInput) Values() url.Values {
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
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) PriceCurrency(v string) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) PriceCountry(v string) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) PriceCustomerGroup(v string) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) PriceChannel(v string) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) LocaleProjection(v []string) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) Expand(v []string) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) WithQueryParams(input ByProjectKeyProductsKeyByKeyRequestMethodPostInput) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	A failed response can return a [DuplicatePriceScope](ctp:api:type:DuplicatePriceScopeError), [DuplicateVariantValues](ctp:api:type:DuplicateVariantValuesError), [DuplicateAttributeValue](ctp:api:type:DuplicateAttributeValueError), or [DuplicateAttributeValues](ctp:api:type:DuplicateAttributeValuesError) error.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *Product, err error) {
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
	case 200:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
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
