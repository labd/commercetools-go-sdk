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

type ByProjectKeyProductsKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput struct {
	PriceCurrency                 *string
	PriceCountry                  *string
	PriceCustomerGroup            *string
	PriceCustomerGroupAssignments []string
	PriceChannel                  *string
	PriceRecurrencePolicy         *string
	Version                       int
	Expand                        []string
}

func (input *ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput) Values() url.Values {
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
	for _, v := range input.PriceCustomerGroupAssignments {
		values.Add("priceCustomerGroupAssignments", fmt.Sprintf("%v", v))
	}
	if input.PriceChannel != nil {
		values.Add("priceChannel", fmt.Sprintf("%v", *input.PriceChannel))
	}
	if input.PriceRecurrencePolicy != nil {
		values.Add("priceRecurrencePolicy", fmt.Sprintf("%v", *input.PriceRecurrencePolicy))
	}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCurrency(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCurrency = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCountry(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCountry = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCustomerGroup(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCustomerGroup = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceCustomerGroupAssignments(v []string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceCustomerGroupAssignments = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceChannel(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceChannel = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) PriceRecurrencePolicy(v string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.PriceRecurrencePolicy = &v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Expand(v []string) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	If [Product price selection query parameters](/../api/pricing-and-discounts-overview#product-price-selection) are provided, the selected Prices are added to the response.
*	Produces the [ProductDeleted](/projects/messages/product-catalog-messages#product-deleted) Message.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *Product, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
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
