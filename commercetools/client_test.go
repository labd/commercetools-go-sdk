package commercetools_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

type OutputData struct{}

func errorHandler(statusCode int, returnValue string, encoding string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(returnValue))
	}
}

func TestClientGetBadRequestJson(t *testing.T) {
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusBadRequest, "{}", "application/json"))
	defer server.Close()

	output := OutputData{}

	err := client.Get("/", nil, &output)
	assert.Equal(t, "HTTP Bad Request: {}", err.Error())
}

func TestClientGetBadRequestStructuredJson(t *testing.T) {
	body, err := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: "Failure reason",
	})
	client, server := testutil.MockClient(
		t, "", nil, errorHandler(http.StatusBadRequest, string(body), "application/json"))
	defer server.Close()

	output := OutputData{}

	err = client.Get("/", nil, &output)
	assert.Equal(t, "HTTP Bad Request: Failure reason", err.Error())
}
