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

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput struct {
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
	Sort                          []string
	Limit                         *int
	Offset                        *int
	WithTotal                     *bool
	Where                         []string
	PredicateVar                  map[string][]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput) Values() url.Values {
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
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Staged(v bool) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PriceCurrency(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PriceCountry(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PriceCustomerGroupAssignments(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PriceChannel(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PriceRecurrencePolicy(v string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) FilterAttributes(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Sort(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Limit(v int) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Offset(v int) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) WithTotal(v bool) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Where(v []string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Queries Variant Projections in the specified [Store](ctp:api:type:Store).
*
*	Only Variants belonging to Products distributed through the Store's configured channels are returned.
*	By default, this endpoint returns the `current` representation of Variants where the `published` flag is `true`.
*
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for a Product and the given Store, the returned Variant Projections include the tailored Product name, slug, and description, along with the tailored Variant images, assets, and attributes. The `staged` query parameter selects either the current or staged tailored data. When no Product Tailoring exists for a Product and Store, the Variant Projection is returned without Product Tailoring data.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Variants, the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation, the API Client must have the `view_products:{projectKey}` scope.
*
*	- To access Variant Projections in the context of a Store, the `view_products:{projectKey}:{storeKey}` scope is required.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet) Execute(ctx context.Context) (result *VariantProjectionPagedQueryResponse, err error) {
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
