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

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput struct {
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	FilterAttributes   []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	if input.Staged != nil {
		if *input.Staged {
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
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterAttributes {
		values.Add("filter[attributes]", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) PriceCurrency(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) PriceCountry(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) PriceCustomerGroup(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) PriceChannel(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) LocaleProjection(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) FilterAttributes(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHeadInput) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if a Variant Projection exists with the provided `id` in the specified [Store](ctp:api:type:Store).
*	Returns a `200` status if the Variant Projection exists, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
