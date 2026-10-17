package checkout

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyPaymentIntegrationsByIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyPaymentIntegrationsByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyPaymentIntegrationsByIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves Payment Integrations in a [Project](ctp:api:type:Project) with a given `id`.
*
 */
func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodGet) Execute(ctx context.Context) (result *PaymentIntegration, err error) {
	queryParams := url.Values{}
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
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
