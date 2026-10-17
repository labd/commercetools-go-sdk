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

type ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDeleteInput struct {
	Version int
}

func (input *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete) Version(v int) *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete) WithQueryParams(input ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDeleteInput) *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete) Execute(ctx context.Context) (result *PaymentIntegration, err error) {
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
