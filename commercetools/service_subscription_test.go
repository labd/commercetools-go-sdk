package commercetools_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestSubscriptionCreate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *commercetools.SubscriptionDraft
		requestBody string
	}{
		{
			desc: "SQS",
			input: &commercetools.SubscriptionDraft{
				Key: "test",
				Destination: commercetools.SqsDestination{
					QueueURL:     "http://example.com/",
					AccessKey:    "A1234567890",
					AccessSecret: "S1234567800",
					Region:       "eu-central-1",
				},
				Messages: []commercetools.MessageSubscription{
					{
						ResourceTypeID: "product",
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
			input: &commercetools.SubscriptionDraft{
				Key: "test",
				Destination: commercetools.SnsDestination{
					TopicArn:     "x",
					AccessKey:    "A1234567890",
					AccessSecret: "S1234567800",
				},
				Changes: []commercetools.ChangeSubscription{
					{
						ResourceTypeID: "product",
					},
				},
				Messages: []commercetools.MessageSubscription{
					{
						ResourceTypeID: "product",
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
			input: &commercetools.SubscriptionDraft{
				Key: "test",
				Destination: commercetools.GoogleCloudPubSubDestination{
					ProjectID: "project-id",
					Topic:     "topic",
				},
				Changes: []commercetools.ChangeSubscription{
					{
						ResourceTypeID: "product",
					},
				},
				Messages: []commercetools.MessageSubscription{
					{
						ResourceTypeID: "product",
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
			input: &commercetools.SubscriptionDraft{
				Key: "test",
				Destination: commercetools.AzureServiceBusDestination{
					ConnectionString: "Endpoint=sb://<namespace>.servicebus.windows.net/;SharedAccessKeyName=<key-name>;SharedAccessKey=<key>;EntityPath=<queue-name>",
				},
				Changes: []commercetools.ChangeSubscription{
					{
						ResourceTypeID: "product",
					},
				},
				Messages: []commercetools.MessageSubscription{
					{
						ResourceTypeID: "product",
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
			input: &commercetools.SubscriptionDraft{
				Key: "test",
				Destination: commercetools.AzureEventGridDestination{
					URI:       "https://example.com/",
					AccessKey: "exampleAccessKey",
				},
				Changes: []commercetools.ChangeSubscription{
					{
						ResourceTypeID: "product",
					},
				},
				Format: commercetools.DeliveryCloudEventsFormat{
					CloudEventsVersion: "0.1",
				},
				Messages: []commercetools.MessageSubscription{
					{
						ResourceTypeID: "product",
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

			client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), &output, nil)
			defer server.Close()

			_, err := client.SubscriptionCreate(context.TODO(), tC.input)
			assert.Nil(t, err)
			assert.JSONEq(t, tC.requestBody, output.JSON)
		})
	}

}

func TestSubscriptionUpdate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), &output, nil)
	defer server.Close()

	input := &commercetools.SubscriptionUpdateWithIDInput{
		ID:      "1234",
		Version: 2,
		Actions: []commercetools.SubscriptionUpdateAction{
			&commercetools.SubscriptionSetKeyAction{
				Key: "123456",
			},
			&commercetools.SubscriptionSetMessagesAction{
				Messages: []commercetools.MessageSubscription{
					{
						ResourceTypeID: "product",
					},
				},
			},
			&commercetools.SubscriptionChangeDestinationAction{
				Destination: commercetools.SnsDestination{
					TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
					AccessKey:    "AKIAIOSFODNN7EXAMPLE",
					AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				},
			},
		},
	}

	_, err := client.SubscriptionUpdateWithID(context.TODO(), input)
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

func TestSubscriptionDeleteByID(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), &output, nil)
	defer server.Close()

	_, err := client.SubscriptionDeleteWithID(context.TODO(), "1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/subscriptions/1234", output.URL.Path)
}

func TestSubscriptionDeleteByKey(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), &output, nil)
	defer server.Close()

	_, err := client.SubscriptionDeleteWithKey(context.TODO(), "1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/subscriptions/key=1234", output.URL.Path)
}
func TestSubscriptionGetDestinationInvalid(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.invalid.json"), nil, nil)
	defer server.Close()

	subscription, err := client.SubscriptionGetWithID(context.TODO(), "100")
	assert.NotNil(t, err)
	assert.Nil(t, subscription)
}

func TestSubscriptionGetDestinationUnknown(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.unknown.json"), nil, nil)
	defer server.Close()

	subscription, err := client.SubscriptionGetWithID(context.TODO(), "100")
	assert.Nil(t, err)
	assert.NotNil(t, subscription)
	assert.Nil(t, subscription.Destination)
}

func TestSubscriptionGetDestinationIronMQ(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.ironmq.json"), nil, nil)
	defer server.Close()

	subscription, err := client.SubscriptionGetWithID(context.TODO(), "100")
	assert.Nil(t, err)

	destination := subscription.Destination.(commercetools.IronMqDestination)
	expected := commercetools.IronMqDestination{
		URI: "https://queue-uri",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSNS(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), nil, nil)
	defer server.Close()

	subscription, err := client.SubscriptionGetWithID(context.TODO(), "100")
	assert.Nil(t, err)

	destination := subscription.Destination.(commercetools.SnsDestination)
	expected := commercetools.SnsDestination{
		TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSQS(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sqs.json"), nil, nil)
	defer server.Close()

	subscription, err := client.SubscriptionGetWithID(context.TODO(), "100")
	assert.Nil(t, err)

	destination := subscription.Destination.(commercetools.SqsDestination)
	expected := commercetools.SqsDestination{
		QueueURL:     "https://queue.amazonaws.com/80398EXAMPLE/MyQueue",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionTypeQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.SubscriptionQuery(context.TODO(), &queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
