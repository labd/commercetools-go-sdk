// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyProductsByIdRequestMethodPost struct {
	body   ProductUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyProductsByIdRequestMethodPostInput
}

func (r *ByProjectKeyProductsByIdRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyProductsByIdRequestMethodPostInput struct {
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   *string
	StoreProjection    *string
	Expand             *[]string
}

func (input *ByProjectKeyProductsByIdRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.PriceCurrency != nil {
		values.Add("priceCurrency", *input.PriceCurrency)
	}
	if input.PriceCountry != nil {
		values.Add("priceCountry", *input.PriceCountry)
	}
	if input.PriceCustomerGroup != nil {
		values.Add("priceCustomerGroup", *input.PriceCustomerGroup)
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", *input.PriceChannel)
	}
	if input.LocaleProjection != nil {
		values.Add("localeProjection", *input.LocaleProjection)
	}
	if input.StoreProjection != nil {
		values.Add("storeProjection", *input.StoreProjection)
	}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyProductsByIdRequestMethodPost) WithQueryParams(input *ByProjectKeyProductsByIdRequestMethodPostInput) *ByProjectKeyProductsByIdRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Update Product by ID
 */
func (rb *ByProjectKeyProductsByIdRequestMethodPost) Execute() (result *Product, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		context.Background(),
		rb.url,
		rb.query,
		data,
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
	case 409, 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, fmt.Errorf("StatusCode %d returend", resp.StatusCode)
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
