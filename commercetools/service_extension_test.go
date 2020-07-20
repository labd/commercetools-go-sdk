package commercetools_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestExtensionCreate(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), nil, nil)
	defer server.Close()

	draft := &commercetools.ExtensionDraft{
		Key: "test",
		Destination: commercetools.ExtensionHTTPDestination{
			URL: "http://example.com",
		},
		Triggers: []commercetools.ExtensionTrigger{
			{
				ResourceTypeID: "product",
				Actions:        []commercetools.ExtensionAction{"Create"},
			},
		},
	}

	_, err := client.ExtensionCreate(context.TODO(), draft)
	assert.Nil(t, err)
}

func TestExtensionUpdate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), &output, nil)
	defer server.Close()

	input := &commercetools.ExtensionUpdateWithIDInput{
		ID:      "foobar",
		Version: 2,
		Actions: []commercetools.ExtensionUpdateAction{
			commercetools.ExtensionChangeDestinationAction{
				Destination: commercetools.ExtensionAWSLambdaDestination{
					Arn:          "arn:aws:lambda:<region>:<accountid>:function:<functionName>",
					AccessKey:    "qwer",
					AccessSecret: "secret",
				},
			},
		},
	}

	fmt.Println(output)

	_, err := client.ExtensionUpdateWithID(context.TODO(), input)
	assert.Nil(t, err)

	expectedBody := `{
		"version": 2,
		"actions": [
			{
				"action": "changeDestination",
				"destination": {
					"type": "AWSLambda",
					"arn": "arn:aws:lambda:<region>:<accountid>:function:<functionName>",
					"accessKey": "qwer",
					"accessSecret": "secret"
				}
			}
		]
	 }`
	assert.JSONEq(t, expectedBody, output.JSON)
}

func TestExtensionDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), nil, nil)
	defer server.Close()

	_, err := client.ExtensionDeleteWithID(context.TODO(), "1234", 2)
	assert.Nil(t, err)
}

func TestExtensionDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), nil, nil)
	defer server.Close()

	_, err := client.ExtensionDeleteWithKey(context.TODO(), "1234", 2)
	assert.Nil(t, err)
}

func TestExtensionQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.ExtensionQuery(context.TODO(), &queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
