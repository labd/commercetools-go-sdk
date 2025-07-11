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

type ByProjectKeyProductProjectionsSearchRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductProjectionsSearchRequestMethodGetInput
}

func (r *ByProjectKeyProductProjectionsSearchRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductProjectionsSearchRequestMethodGetInput struct {
	MarkMatchingVariants          *bool
	Text                          map[string][]string
	Fuzzy                         *bool
	FuzzyLevel                    *int
	FilterQuery                   []string
	Filter                        []string
	Facet                         []string
	FilterFacets                  []string
	Expand                        []string
	Sort                          []string
	Limit                         *int
	Offset                        *int
	Staged                        *bool
	PriceCurrency                 *string
	PriceCountry                  *string
	PriceCustomerGroup            *string
	PriceCustomerGroupAssignments []string
	PriceChannel                  *string
	PriceRecurrencePolicy         *string
	LocaleProjection              []string
	StoreProjection               *string
}

func (input *ByProjectKeyProductProjectionsSearchRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.MarkMatchingVariants != nil {
		if *input.MarkMatchingVariants {
			values.Add("markMatchingVariants", "true")
		} else {
			values.Add("markMatchingVariants", "false")
		}
	}
	for k, v := range input.Text {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	if input.Fuzzy != nil {
		if *input.Fuzzy {
			values.Add("fuzzy", "true")
		} else {
			values.Add("fuzzy", "false")
		}
	}
	if input.FuzzyLevel != nil {
		values.Add("fuzzyLevel", strconv.Itoa(*input.FuzzyLevel))
	}
	for _, v := range input.FilterQuery {
		values.Add("filter.query", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Filter {
		values.Add("filter", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Facet {
		values.Add("facet", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterFacets {
		values.Add("filter.facets", fmt.Sprintf("%v", v))
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
	if input.StoreProjection != nil {
		values.Add("storeProjection", fmt.Sprintf("%v", *input.StoreProjection))
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) MarkMatchingVariants(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.MarkMatchingVariants = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Text(v map[string][]string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Text = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Fuzzy(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Fuzzy = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) FuzzyLevel(v int) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.FuzzyLevel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) FilterQuery(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.FilterQuery = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Filter(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Filter = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Facet(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Facet = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) FilterFacets(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.FilterFacets = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Expand(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Sort(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Limit(v int) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Offset(v int) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceCurrency(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceCountry(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceCustomerGroup(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceCustomerGroupAssignments(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceRecurrencePolicy(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) LocaleProjection(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) StoreProjection(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) WithQueryParams(input ByProjectKeyProductProjectionsSearchRequestMethodGetInput) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	This method appends query parameters to the URL.
*	The maximum allowed URL length is 8192 characters.
*	Exceeding this limit will result in URL truncation, potentially leading to unexpected results.
*	For funnel searches on Product Listing Pages, where users select multiple filters, we recommend the [POST](ctp:api:endpoint:/{projectKey}/product-projections/search:POST) method which passes the query parameters within the request body, avoiding URL length restrictions.
*
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Execute(ctx context.Context) (result *ProductProjectionPagedSearchResponse, err error) {
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
