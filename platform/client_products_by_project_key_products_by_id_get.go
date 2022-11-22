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

type ByProjectKeyProductsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsByIDRequestMethodGetInput
}

func (r *ByProjectKeyProductsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsByIDRequestMethodGetInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	Expand             []string
}

func (input *ByProjectKeyProductsByIDRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyProductsByIDRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyProductsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyProductsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyProductsByIDRequestMethodGetInput) *ByProjectKeyProductsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsByIDRequestMethodGet) Execute(ctx context.Context) (result *Product, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
