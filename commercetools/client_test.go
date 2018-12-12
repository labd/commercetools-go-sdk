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
	assert.Equal(t, "invalid character ',' looking for beginning of value", err.Error())
}

func TestClientNotFound(t *testing.T) {
	body := ``
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusNotFound, body, "application/json"))
	defer server.Close()

	output := OutputData{}

	err := client.Get("/", nil, &output)
	assert.Equal(t, "Not Found (404): ResourceNotFound", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)
	assert.Equal(t, 404, ctErr.StatusCode)
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

	assert.Equal(t, "Insufficient scope", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)

	ctChildErr, ok := ctErr.Errors[0].(commercetools.InsufficientScopeError)
	assert.Equal(t, true, ok)
	assert.Equal(t, "Insufficient scope", ctChildErr.Message)
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

	assert.Equal(t, "Request body does not contain valid JSON.", err.Error())

	ctErr, ok := err.(commercetools.ErrorResponse)
	assert.Equal(t, true, ok)

	ctChildErr, ok := ctErr.Errors[0].(commercetools.InvalidJSONInputError)
	assert.Equal(t, "Request body does not contain valid JSON.", ctChildErr.Error())
	// assert.Equal(t, commercetools.ErrInvalidJSONInput, ctErr.Errors[0].Code())
}
