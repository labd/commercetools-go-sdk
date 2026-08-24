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

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput struct {
	Where              []string
	Staged             *bool
	PriceCurrency      *string
	PriceCountry       *string
	PriceCustomerGroup *string
	PriceChannel       *string
	LocaleProjection   []string
	FilterAttributes   []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
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

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) Where(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) PriceCurrency(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) PriceCountry(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) PriceCustomerGroup(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) PriceChannel(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) LocaleProjection(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) FilterAttributes(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHeadInput) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more Variant Projections exist for the provided query predicate in the specified [Store](ctp:api:type:Store).
*	Returns a `200` status if any Variant Projections match, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead) Execute(ctx context.Context) error {
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
