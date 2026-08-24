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

type ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput struct {
	Staged                        *bool
	PriceCurrency                 *string
	PriceCountry                  *string
	PriceCustomerGroup            *string
	PriceCustomerGroupAssignments []string
	PriceChannel                  *string
	PriceRecurrencePolicy         *string
	LocaleProjection              []string
	FilterAttributes              []string
	Expand                        []string
}

func (input *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput) Values() url.Values {
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
	for _, v := range input.PriceCustomerGroupAssignments {
		values.Add("priceCustomerGroupAssignments", fmt.Sprintf("%v", v))
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", fmt.Sprintf("%v", *input.PriceChannel))
	}
	if input.PriceRecurrencePolicy != nil {
		values.Add("priceRecurrencePolicy", fmt.Sprintf("%v", *input.PriceRecurrencePolicy))
	}
	for _, v := range input.LocaleProjection {
		values.Add("localeProjection", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterAttributes {
		values.Add("filter[attributes]", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) Staged(v bool) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCurrency(v string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCountry(v string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroupAssignments(v []string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceChannel(v string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceRecurrencePolicy(v string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) FilterAttributes(v []string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGetInput) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves a Variant Projection by its key.
*
 */
func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *VariantProjection, err error) {
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
