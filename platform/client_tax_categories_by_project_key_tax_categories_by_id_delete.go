// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyTaxCategoriesByIdRequestMethodDelete struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyTaxCategoriesByIdRequestMethodDeleteInput
}

func (r *ByProjectKeyTaxCategoriesByIdRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyTaxCategoriesByIdRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyTaxCategoriesByIdRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyTaxCategoriesByIdRequestMethodDelete) WithQueryParams(input *ByProjectKeyTaxCategoriesByIdRequestMethodDeleteInput) *ByProjectKeyTaxCategoriesByIdRequestMethodDelete {
	rb.query = input.Values()
	return rb
}

/**
*	Delete TaxCategory by ID
 */
func (rb *ByProjectKeyTaxCategoriesByIdRequestMethodDelete) Execute(ctx context.Context) (result *TaxCategory, err error) {
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		rb.query,
		nil,
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
