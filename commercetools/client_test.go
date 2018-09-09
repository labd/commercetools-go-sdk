package commercetools_test

import (
	"net/http"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

type OutputData struct{}

func errorHandler(statusCode int, returnValue string, encoding string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(returnValue))
	}
}

func TestClientGetBadRequestJson(t *testing.T) {
	body := `,`
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusBadRequest, body, "application/json"))
	defer server.Close()

	output := OutputData{}

	err := client.Get("/", nil, &output)
	assert.Equal(t, "Unmarshalling error response failed: invalid character ',' looking for beginning of value", err.Error())
}

func TestClientNotFound(t *testing.T) {
	body := ``
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusNotFound, body, "application/json"))
	defer server.Close()

	output := OutputData{}

	err := client.Get("/", nil, &output)
	assert.Equal(t, "Not Found (404): ResourceNotFound", err.Error())

	ctErr, ok := err.(commercetools.Error)
	assert.Equal(t, true, ok)
	assert.Equal(t, commercetools.ErrResourceNotFound, ctErr.Code())
}

func TestAuthError(t *testing.T) {
	body := `{
		"statusCode": 403,
		"message": "Insufficient scope",
		"errors": [
			{
				"code": "insufficient_scope",
				"message": "Insufficient scope"
			}
		],
		"error": "insufficient_scope",
		"error_description": "Insufficient scope"
	}`
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusForbidden, body, "application/json"))

	defer server.Close()

	output := OutputData{}

	err := client.Get("/", nil, &output)

	assert.Equal(t, "Forbidden (403): insufficient_scope - Insufficient scope", err.Error())

	ctErr, ok := err.(commercetools.Error)
	assert.Equal(t, true, ok)

	assert.Equal(t, "insufficient_scope", ctErr.Errors[0].Code())
	assert.Equal(t, "insufficient_scope - Insufficient scope", ctErr.Errors[0].Message())
}

func TestInvalidJsonError(t *testing.T) {
	body := `{
		"statusCode": 400,
		"message": "Request body does not contain valid JSON.",
		"errors": [
			{
				"code": "InvalidJsonInput",
				"message": "Request body does not contain valid JSON.",
				"detailedErrorMessage": "No content to map due to end-of-input"
			}
		]
	}`
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusBadRequest, body, "application/json"))

	defer server.Close()

	output := OutputData{}

	err := client.Get("/", nil, &output)

	assert.Equal(t, "Bad Request (400): InvalidJsonInput - Request body does not contain valid JSON. (No content to map due to end-of-input)", err.Error())

	ctErr, ok := err.(commercetools.Error)
	assert.Equal(t, true, ok)

	assert.Equal(t, commercetools.ErrInvalidJSONInput, ctErr.Code())
	assert.Equal(t, commercetools.ErrInvalidJSONInput, ctErr.Errors[0].Code())
}
