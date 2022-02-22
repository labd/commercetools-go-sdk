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
	Fuzzy                *bool
	FuzzyLevel           *float64
	MarkMatchingVariants *bool
	Staged               *bool
	Filter               []string
	FilterFacets         []string
	FilterQuery          []string
	Facet                []string
	Text                 map[string][]string
	Sort                 []string
	Limit                *int
	Offset               *int
	WithTotal            *bool
	PriceCurrency        *string
	PriceCountry         *string
	PriceCustomerGroup   *string
	PriceChannel         *string
	LocaleProjection     *string
	StoreProjection      *string
	Expand               []string
}

func (input *ByProjectKeyProductProjectionsSearchRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Fuzzy != nil {
		if *input.Fuzzy {
			values.Add("fuzzy", "true")
		} else {
			values.Add("fuzzy", "false")
		}
	}
	if input.FuzzyLevel != nil {
		values.Add("fuzzyLevel", fmt.Sprintf("%f", *input.FuzzyLevel))
	}
	if input.MarkMatchingVariants != nil {
		if *input.MarkMatchingVariants {
			values.Add("markMatchingVariants", "true")
		} else {
			values.Add("markMatchingVariants", "false")
		}
	}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	for _, v := range input.Filter {
		values.Add("filter", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterFacets {
		values.Add("filter.facets", fmt.Sprintf("%v", v))
	}
	for _, v := range input.FilterQuery {
		values.Add("filter.query", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Facet {
		values.Add("facet", fmt.Sprintf("%v", v))
	}
	for k, v := range input.Text {
		for _, x := range v {
			values.Set(k, x)
		}
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
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Fuzzy(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Fuzzy = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) FuzzyLevel(v float64) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.FuzzyLevel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) MarkMatchingVariants(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.MarkMatchingVariants = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Staged(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Filter(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Filter = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) FilterFacets(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.FilterFacets = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) FilterQuery(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.FilterQuery = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Facet(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Facet = v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Text(v map[string][]string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Text = v
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

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
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

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) PriceChannel(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) LocaleProjection(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.LocaleProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) StoreProjection(v string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.StoreProjection = &v
	return rb
}

func (rb *ByProjectKeyProductProjectionsSearchRequestMethodGet) Expand(v []string) *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductProjectionsSearchRequestMethodGetInput{}
	}
	rb.params.Expand = v
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
*	Search Product Projection
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
		return result, nil
	case 400, 401, 403, 500, 502, 503:
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
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
