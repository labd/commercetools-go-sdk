package customize_test

import (
	"fmt"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools/customize"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestExtensionCreate(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := customize.New(client)

	draft := &customize.ExtensionDraft{
		Key: "test",
		Destination: customize.ExtensionDestinationHTTP{
			URL: "http://example.com",
		},
		Triggers: []customize.ExtensionTrigger{
			customize.ExtensionTrigger{
				ResourceTypeID: "product",
				Actions:        []string{"Create"},
			},
		},
	}

	_, err := svc.ExtensionCreate(draft)
	assert.Equal(t, nil, err)
}

func TestExtensionUpdate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, fixture("extension.azure.json"), &output, nil)
	defer server.Close()
	svc := customize.New(client)

	input := &customize.ExtensionUpdateInput{
		Version: 2,
		Actions: commercetools.UpdateActions{
			customize.ExtensionChangeDestination{
				Destination: customize.ExtensionDestinationAWSLambda{
					ARN:          "arn:aws:lambda:<region>:<accountid>:function:<functionName>",
					AccessKey:    "qwer",
					AccessSecret: "secret",
				},
			},
		},
	}

	fmt.Println(output)

	_, err := svc.ExtensionUpdate(input)
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
	assert.JSONEq(t, expectedBody, string(output.Body))
}

func TestExtensionDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.ExtensionDeleteByID("1234", 2)
	assert.Equal(t, nil, err)
}

func TestExtensionDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("extension.azure.json"), nil, nil)
	defer server.Close()
	svc := customize.New(client)

	_, err := svc.ExtensionDeleteByKey("1234", 2)
	assert.Equal(t, nil, err)
}
