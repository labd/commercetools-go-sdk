package commercetools_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestAPIClientCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	input := &commercetools.APIClientDraft{
		Scope: "manage_project",
		Name:  "api-client-name",
	}

	fmt.Println(output)

	_, err := client.APIClientCreate(input)
	assert.Nil(t, err)

	expectedBody := `{
		"name": "api-client-name",
		"scope": "manage_project"
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestAPIClientDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	_, err := client.APIClientDeleteWithID("1234")
	assert.Nil(t, err)

	assert.Equal(t, "/unittest/api-clients/1234", output.URL.Path)
}

func TestAPIClientGetByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("api-client.example.json"), nil, nil)
	defer server.Close()

	timestamp, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00.001Z")

	input := &commercetools.APIClient{
		ID:         "api-client-id",
		Scope:      "view_products",
		Name:       "api-client-name",
		CreatedAt:  timestamp,
		LastUsedAt: commercetools.NewDate(2017, 9, 10),
		Secret:     "secret-passphrase",
	}

	result, err := client.APIClientGetWithID("1234")
	assert.Nil(t, err)
	assert.Equal(t, input, result)
}

func TestAPIClientQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.APIClientQuery(&queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
