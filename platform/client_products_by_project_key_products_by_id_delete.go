// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyProductsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyProductsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsByIDRequestMethodDeleteInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Version            int
	Expand             []string
}

func (input *ByProjectKeyProductsByIDRequestMethodDeleteInput) Values() url.Values {
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

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) PriceCurrency(v string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) PriceCountry(v string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) PriceCustomerGroup(v string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) PriceChannel(v string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) LocaleProjection(v string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) StoreProjection(v string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) Version(v int) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyProductsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyProductsByIDRequestMethodDeleteInput) *ByProjectKeyProductsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyProductsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductsByIDRequestMethodDelete) Execute(ctx context.Context) (result *Product, err error) {
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
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
