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

type ByProjectKeyProductsByIDRequestMethodPost struct {
	body    ProductUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsByIDRequestMethodPostInput
}

func (r *ByProjectKeyProductsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsByIDRequestMethodPostInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Expand             []string
}

func (input *ByProjectKeyProductsByIDRequestMethodPostInput) Values() url.Values {
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

func (rb *ByProjectKeyProductsByIDRequestMethodPost) PriceCurrency(v string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) PriceCountry(v string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) PriceCustomerGroup(v string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) PriceChannel(v string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) LocaleProjection(v string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) StoreProjection(v string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyProductsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyProductsByIDRequestMethodPostInput) *ByProjectKeyProductsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductsByIDRequestMethodPost) Execute(ctx context.Context) (result *Product, err error) {
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
		return result, nil
	case 409, 400, 401, 403, 500, 502, 503:
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