package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentIntegrationsByIdRequestMethodDeleteInput
}

func (r *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentIntegrationsByIdRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyPaymentIntegrationsByIdRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete) Version(v int) *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentIntegrationsByIdRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete) WithQueryParams(input ByProjectKeyPaymentIntegrationsByIdRequestMethodDeleteInput) *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete) Execute(ctx context.Context) (result *PaymentIntegration, err error) {
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
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
