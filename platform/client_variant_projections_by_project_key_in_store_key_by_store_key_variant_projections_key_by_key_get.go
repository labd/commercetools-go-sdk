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

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput struct {
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

func (input *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCurrency(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCountry(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceCustomerGroupAssignments(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceChannel(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) PriceRecurrencePolicy(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) FilterAttributes(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves a Variant Projection by its key in the specified [Store](ctp:api:type:Store).
*
*	Only returns Variants belonging to Products distributed through the Store's configured channels.
*
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for the Product and the given Store, the returned Variant Projection includes the tailored Product name, slug, and description, along with the tailored Variant images, assets, and attributes. The `staged` query parameter selects either the current or staged tailored data. When no Product Tailoring exists for the Product and Store, the Variant Projection is returned without Product Tailoring data.
*
*	Required access scopes:
*
*	- To retrieve the current representation, the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation, the API Client must have the `view_products:{projectKey}` scope.
*
*	- To access Variant Projections in the context of a Store, the `view_products:{projectKey}:{storeKey}` scope is required.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *VariantProjection, err error) {
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
