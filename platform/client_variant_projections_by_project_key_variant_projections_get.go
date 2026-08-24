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

type ByProjectKeyVariantProjectionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyVariantProjectionsRequestMethodGetInput
}

func (r *ByProjectKeyVariantProjectionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyVariantProjectionsRequestMethodGetInput struct {
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

func (input *ByProjectKeyVariantProjectionsRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Staged(v bool) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PriceCurrency(v string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PriceCountry(v string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PriceCustomerGroupAssignments(v []string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PriceChannel(v string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PriceRecurrencePolicy(v string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) FilterAttributes(v []string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.FilterAttributes = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Expand(v []string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Sort(v []string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Limit(v int) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Offset(v int) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) WithTotal(v bool) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Where(v []string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyVariantProjectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyVariantProjectionsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) WithQueryParams(input ByProjectKeyVariantProjectionsRequestMethodGetInput) *ByProjectKeyVariantProjectionsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyVariantProjectionsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Queries Variant Projections matching the provided predicates.
*
*	By default, this endpoint returns the `current` representation where variants are published.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Variants, the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation (draft data), the API Client must have the `view_products:{projectKey}` scope.
*
 */
func (rb *ByProjectKeyVariantProjectionsRequestMethodGet) Execute(ctx context.Context) (result *VariantProjectionPagedQueryResponse, err error) {
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
