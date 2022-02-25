package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyProductsKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Version            int
	Expand             []string
}

func (input *ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput) Values() url.Values {
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
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCurrency(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCountry(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCustomerGroup(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceChannel(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) LocaleProjection(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) StoreProjection(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Expand(v []string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *Product, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
