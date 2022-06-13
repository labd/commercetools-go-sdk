package importapi

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Execute(ctx context.Context) (result *ImportOperation, err error) {
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
		return result, nil
	case 404, 503:
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
