package platform_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestSubscriptionCreate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       platform.SubscriptionDraft
		requestBody string
	}{
		{
			desc: "SQS",
			input: platform.SubscriptionDraft{
				Key: ctutils.StringRef("test"),
				Destination: platform.SqsDestination{
					QueueUrl:     "http://example.com/",
					AccessKey:    "A1234567890",
					AccessSecret: "S1234567800",
					Region:       "eu-central-1",
				},
				Messages: []platform.MessageSubscription{
					{
						ResourceTypeId: "product",
					},
				},
			},
			requestBody: `{
				"key": "test",
				"destination": {
					"type": "SQS",
					"queueUrl": "http://example.com/",
					"accessKey": "A1234567890",
					"accessSecret": "S1234567800",
					"region": "eu-central-1"
				},
				"messages": [
					{
						"resourceTypeId": "product"
					}
				]
			}`,
		},
		{
			desc: "SNS",
			input: platform.SubscriptionDraft{
				Key: ctutils.StringRef("test"),
				Destination: platform.SnsDestination{
					TopicArn:     "x",
					AccessKey:    "A1234567890",
					AccessSecret: "S1234567800",
				},
				Changes: []platform.ChangeSubscription{
					{
						ResourceTypeId: "product",
					},
				},
				Messages: []platform.MessageSubscription{
					{
						ResourceTypeId: "product",
					},
				},
			},
			requestBody: `{
				"key": "test",
				"destination": {
					"type": "SNS",
					"topicArn": "x",
					"accessKey": "A1234567890",
					"accessSecret": "S1234567800"
				},
				"messages": [
					{
						"resourceTypeId": "product"
					}
				],
				"changes": [
					{
						"resourceTypeId": "product"
					}
				]
			}`,
		},
		{
			desc: "GooglePubSub",
			input: platform.SubscriptionDraft{
				Key: ctutils.StringRef("test"),
				Destination: platform.GoogleCloudPubSubDestination{
					ProjectId: "project-id",
					Topic:     "topic",
				},
				Changes: []platform.ChangeSubscription{
					{
						ResourceTypeId: "product",
					},
				},
				Messages: []platform.MessageSubscription{
					{
						ResourceTypeId: "product",
					},
				},
			},
			requestBody: `{
				"key": "test",
				"destination": {
					"type": "GoogleCloudPubSub",
					"topic": "topic",
					"projectId": "project-id"
				},
				"messages": [
					{
						"resourceTypeId": "product"
					}
				],
				"changes": [
					{
						"resourceTypeId": "product"
					}
				]
			}`,
		},
		{
			desc: "AzureServiceBus",
			input: platform.SubscriptionDraft{
				Key: ctutils.StringRef("test"),
				Destination: platform.AzureServiceBusDestination{
					ConnectionString: "Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;EntityPath=<queue-name>",
				},
				Changes: []platform.ChangeSubscription{
					{
						ResourceTypeId: "product",
					},
				},
				Messages: []platform.MessageSubscription{
					{
						ResourceTypeId: "product",
					},
				},
			},
			requestBody: `{
				"key": "test",
				"destination": {
					"type": "AzureServiceBus",
					"connectionString": "Endpoint=sb:\/\/<namespace>.servicebus.windows.net\/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;EntityPath=<queue-name>"
				},
				"messages": [
					{
						"resourceTypeId": "product"
					}
				],
				"changes": [
					{
						"resourceTypeId": "product"
					}
				]
			}`,
		},
		{
			desc: "AzureEventGrid",
			input: platform.SubscriptionDraft{
				Key: ctutils.StringRef("test"),
				Destination: platform.AzureEventGridDestination{
					Uri:       "https://example.com/",
					AccessKey: "exampleAccessKey",
				},
				Changes: []platform.ChangeSubscription{
					{
						ResourceTypeId: "product",
					},
				},
				Format: platform.DeliveryCloudEventsFormat{
					CloudEventsVersion: "0.1",
				},
				Messages: []platform.MessageSubscription{
					{
						ResourceTypeId: "product",
					},
				},
			},
			requestBody: `{
				"key": "test",
				"destination": {
					"type": "EventGrid",
					"uri": "https://example.com/",
					"accessKey": "exampleAccessKey"
				},
				"messages": [
					{
						"resourceTypeId": "product"
					}
				],
				"changes": [
					{
						"resourceTypeId": "product"
					}
				],
				"format": {
					"type": "CloudEvents",
					"cloudEventsVersion": "0.1"
				}
			}`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}
			response := testutil.Fixture("subscription.sns.json", 201)

			client, server := testutil.MockClient(t, response, &output, nil)
			defer server.Close()

			_, err := client.WithProjectKey("unittest").Subscriptions().Post(tC.input).Execute(context.TODO())
			assert.Nil(t, err)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}

}

func TestSubscriptionUpdate(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("subscription.sns.json", 200)

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	input := platform.SubscriptionUpdate{
		Version: 2,
		Actions: []platform.SubscriptionUpdateAction{
			&platform.SubscriptionSetKeyAction{
				Key: ctutils.StringRef("123456"),
			},
			&platform.SubscriptionSetMessagesAction{
				Messages: []platform.MessageSubscription{
					{
						ResourceTypeId: "product",
					},
				},
			},
			&platform.SubscriptionChangeDestinationAction{
				Destination: platform.SnsDestination{
					TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
					AccessKey:    "AKIAIOSFODNN7EXAMPLE",
					AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				},
			},
		},
	}

	_, err := client.WithProjectKey("unittest").Subscriptions().WithId("1234").Post(input).Execute(context.TODO())
	assert.Nil(t, err)

	expectedBody := `{
	"actions": [
		{
			"action": "setKey",
			"key": "123456"
		},
		{
			"action": "setMessages",
			"messages": [
				{
					"resourceTypeId": "product"
				}
			]
		},
		{
			"action": "changeDestination",
			"destination": {
				"type": "SNS",
				"topicArn": "arn:aws:sns:eu-central-1:123456789012345678:example:1",
				"accessKey": "AKIAIOSFODNN7EXAMPLE",
				"accessSecret": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
			}
		}
	],
	"version": 2
}`
	assert.JSONEq(t, expectedBody, output.JSON)
	assert.Equal(t, "/unittest/subscriptions/1234", output.URL.Path)
}

func TestSubscriptionDeleteById(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("subscription.sns.json", 200)
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Subscriptions().
		WithId("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeySubscriptionsByIDRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/subscriptions/1234", output.URL.Path)
}

func TestSubscriptionDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.Fixture("subscription.sns.json", 200)
	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.WithProjectKey("unittest").
		Subscriptions().
		WithKey("1234").
		Delete().
		WithQueryParams(platform.ByProjectKeySubscriptionsKeyByKeyRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/subscriptions/key=1234", output.URL.Path)
}

func TestSubscriptionGetNonExisting(t *testing.T) {
	response := testutil.ResponseData{
		StatusCode: 404,
		Body: `
			{
				"statusCode": 404,
				"message": "The Resource with ID 'a5e3540e-6742-4bef-b239-37dda425a5d9' was not found.",
				"errors": [
					{
						"code": "ResourceNotFound",
						"message": "The Resource with ID 'a5e3540e-6742-4bef-b239-37dda425a5d9' was not found."
					}
				]
			}`,
	}

	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	subscription, err := client.
		WithProjectKey("unittest").
		Subscriptions().
		WithId("a5e3540e-6742-4bef-b239-37dda425a5d9").
		Get().
		Execute(context.TODO())
	assert.Nil(t, subscription)
	assert.NotNil(t, err)
}

func TestSubscriptionGetDestinationInvalid(t *testing.T) {
	response := testutil.Fixture("subscription.invalid.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	subscription, err := client.WithProjectKey("unittest").Subscriptions().WithId("100").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.Nil(t, subscription.Destination)
}

func TestSubscriptionGetDestinationUnknown(t *testing.T) {
	response := testutil.Fixture("subscription.unknown.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	subscription, err := client.WithProjectKey("unittest").Subscriptions().WithId("100").Get().Execute(context.TODO())
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
	assert.Nil(t, subscription.Destination)
}

func TestSubscriptionGetDestinationIronMQ(t *testing.T) {
	response := testutil.Fixture("subscription.ironmq.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	subscription, err := client.WithProjectKey("unittest").Subscriptions().WithId("100").Get().Execute(context.TODO())
	assert.Nil(t, err)

	destination := subscription.Destination.(platform.IronMqDestination)
	expected := platform.IronMqDestination{
		Uri: "https://queue-uri",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSNS(t *testing.T) {
	response := testutil.Fixture("subscription.sns.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	subscription, err := client.WithProjectKey("unittest").Subscriptions().WithId("100").Get().Execute(context.TODO())
	assert.Nil(t, err)

	destination := subscription.Destination.(platform.SnsDestination)
	expected := platform.SnsDestination{
		TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSQS(t *testing.T) {
	response := testutil.Fixture("subscription.sqs.json", 200)
	client, server := testutil.MockClient(t, response, nil, nil)
	defer server.Close()

	subscription, err := client.WithProjectKey("unittest").Subscriptions().WithId("100").Get().Execute(context.TODO())
	assert.Nil(t, err)

	destination := subscription.Destination.(platform.SqsDestination)
	expected := platform.SqsDestination{
		QueueUrl:     "https://queue.amazonaws.com/80398EXAMPLE/MyQueue",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionsQuery(t *testing.T) {
	output := testutil.RequestData{}
	response := testutil.ResponseData{StatusCode: 200, Body: "{}"}

	client, server := testutil.MockClient(t, response, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Subscriptions().
		Get().
		WithQueryParams(platform.ByProjectKeySubscriptionsRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
