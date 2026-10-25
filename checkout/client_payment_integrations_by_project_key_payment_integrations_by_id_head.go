package checkout

// Generated file, please do not change!!!

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyPaymentIntegrationsByIdRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyPaymentIntegrationsByIdRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyPaymentIntegrationsByIdRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if a PaymentIntegration with a given `id` exists. Returns `200 OK` if the PaymentIntegration exists, or `404 Not Found` if it doesn't. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyPaymentIntegrationsByIdRequestMethodHead) Execute(ctx context.Context) error {
	queryParams := url.Values{}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
