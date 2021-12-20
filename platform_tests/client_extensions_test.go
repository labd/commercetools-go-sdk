package platform_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestExtensionCreate(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json", 201), nil, nil)
	defer server.Close()

	draft := platform.ExtensionDraft{
		Key: ctutils.StringRef("test"),
		Destination: platform.ExtensionHttpDestination{
			Url: "http://example.com",
		},
		Triggers: []platform.ExtensionTrigger{
			{
				ResourceTypeId: "product",
				Actions:        []platform.ExtensionAction{"Create"},
			},
		},
	}

	_, err := client.WithProjectKey("unittest").Extensions().Post(draft).Execute(context.TODO())
	assert.Nil(t, err)
}

func TestExtensionUpdate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json", 200), &output, nil)
	defer server.Close()

	input := platform.ExtensionUpdate{
		Version: 2,
		Actions: []platform.ExtensionUpdateAction{
			platform.ExtensionChangeDestinationAction{
				Destination: platform.ExtensionAWSLambdaDestination{
					Arn:          "arn:aws:lambda:<region>:<accountid>:function:<functionName>",
					AccessKey:    "qwer",
					AccessSecret: "secret",
				},
			},
		},
	}

	_, err := client.WithProjectKey("unittest").
		Extensions().
		WithId("1234").
		Post(input).
		Execute(context.TODO())

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
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json", 200), nil, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Extensions().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyExtensionsByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())

	assert.Nil(t, err)
}

func TestExtensionDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json", 200), nil, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Extensions().
		WithKey("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeyExtensionsKeyByKeyRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)
}

func TestExtensionQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Extensions().
		Get().
		WithQueryParams(platform.ByProjectKeyExtensionsRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
