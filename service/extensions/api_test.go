package extensions_test

import (
	"fmt"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/extensions"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestExtensionCreate(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := extensions.New(client)

	draft := &extensions.ExtensionDraft{
		Key: "test",
		Destination: extensions.DestinationHTTP{
			URL: "http://example.com",
		},
		Triggers: []extensions.Trigger{
			extensions.Trigger{
				ResourceTypeID: "product",
				Actions:        []string{"Create"},
			},
		},
	}

	_, err := svc.Create(draft)
	assert.Equal(t, nil, err)
}

func TestExtensionUpdate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), &output, nil)
	defer server.Close()
	svc := extensions.New(client)

	input := &extensions.UpdateInput{
		Version: 2,
		Actions: commercetools.UpdateActions{
			extensions.ChangeDestination{
				Destination: extensions.DestinationAWSLambda{
					ARN:          "arn:aws:lambda:<region>:<accountid>:function:<functionName>",
					AccessKey:    "qwer",
					AccessSecret: "secret",
				},
			},
		},
	}

	fmt.Println(output)

	_, err := svc.Update(input)
	assert.Equal(t, nil, err)

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
	svc := extensions.New(client)

	_, err := svc.DeleteByID("1234", 2)
	assert.Equal(t, nil, err)
}

func TestExtensionDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := extensions.New(client)

	_, err := svc.DeleteByKey("1234", 2)
	assert.Equal(t, nil, err)
}
