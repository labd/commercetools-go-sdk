package platform_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestApiClientCreate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 201, Body: `{}`}, &output, nil)
	defer server.Close()

	input := platform.ApiClientDraft{
		Scope: "manage_project",
		Name:  "api-client-name",
	}

	fmt.Println(input)
	xdata, _ := json.Marshal(input)
	fmt.Println(string(xdata))

	_, err := client.WithProjectKey("test").ApiClients().Post(input).Execute(context.TODO())
	assert.Nil(t, err)

	expectedBody := `{
		"name": "api-client-name",
		"scope": "manage_project"
	}`

	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestAPIClientDelete(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").ApiClients().WithId("1234").Delete().Execute(context.TODO())
	assert.Nil(t, err)

	assert.Equal(t, "/unittest/api-clients/1234", output.URL.Path)
}

func TestAPIClientGetByID(t *testing.T) {
	response := testutil.ResponseData{
		StatusCode: 200,
		Body: `{
			"id": "api-client-id",
			"name": "api-client-name",
			"scope": "view_products",
			"createdAt": "2018-01-01T00:00:00.001Z",
			"lastUsedAt": "2017-09-10",
			"secret": "secret-passphrase"
		}`,
	}
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	timestamp, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00.001Z")

	input := platform.ApiClient{
		ID:         "api-client-id",
		Scope:      "view_products",
		Name:       "api-client-name",
		CreatedAt:  &timestamp,
		LastUsedAt: &platform.Date{Year: 2017, Month: 9, Day: 10},
		Secret:     ctutils.StringRef("secret-passphrase"),
	}

	result, err := client.WithProjectKey("test").ApiClients().WithId("1234").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, input, *result)
}

func TestAPIClientQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("test").Get().Execute(context.TODO())
	assert.Nil(t, err)
}
