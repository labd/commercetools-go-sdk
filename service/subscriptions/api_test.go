package subscriptions_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/subscriptions"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestSubscriptionCreate(t *testing.T) {
	testCases := []struct {
		desc        string
		input       *subscriptions.SubscriptionDraft
		requestBody string
	}{
		{
			desc: "SQS",
			input: &subscriptions.SubscriptionDraft{
				Key: "test",
				Destination: subscriptions.DestinationAWSSQS{
					QueueURL:     "http://example.com/",
					AccessKey:    "A1234567890",
					AccessSecret: "S1234567800",
					Region:       "eu-central-1",
				},
				Messages: []subscriptions.MessageSubscription{
					subscriptions.MessageSubscription{
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
			input: &subscriptions.SubscriptionDraft{
				Key: "test",
				Destination: subscriptions.DestinationAWSSNS{
					TopicArn:     "x",
					AccessKey:    "A1234567890",
					AccessSecret: "S1234567800",
				},
				Changes: []subscriptions.ChangeSubscription{
					subscriptions.ChangeSubscription{
						ResourceTypeID: "product",
					},
				},
				Messages: []subscriptions.MessageSubscription{
					subscriptions.MessageSubscription{
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
			input: &subscriptions.SubscriptionDraft{
				Key: "test",
				Destination: subscriptions.DestinationGooglePubSub{
					ProjectID: "project-id",
					Topic:     "topic",
				},
				Changes: []subscriptions.ChangeSubscription{
					subscriptions.ChangeSubscription{
						ResourceTypeID: "product",
					},
				},
				Messages: []subscriptions.MessageSubscription{
					subscriptions.MessageSubscription{
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
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := testutil.RequestData{}

			client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), &output, nil)
			defer server.Close()
			svc := subscriptions.New(client)

			_, err := svc.Create(tC.input)
			assert.Nil(t, err)
			assert.JSONEq(t, tC.requestBody, output.JSON)

		})
	}

}

func TestSubscriptionUpdate(t *testing.T) {
	output := testutil.RequestData{}

	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), &output, nil)
	defer server.Close()
	svc := subscriptions.New(client)

	input := &subscriptions.UpdateInput{
		ID:      "1234",
		Version: 2,
		Actions: commercetools.UpdateActions{
			&subscriptions.SetKey{
				Key: "123456",
			},
			&subscriptions.SetMessages{
				Messages: []subscriptions.MessageSubscription{
					subscriptions.MessageSubscription{
						ResourceTypeID: "product",
					},
				},
			},
			&subscriptions.ChangeDestination{
				Destination: subscriptions.DestinationAWSSNS{
					TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
					AccessKey:    "AKIAIOSFODNN7EXAMPLE",
					AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				},
			},
		},
	}

	_, err := svc.Update(input)
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
	svc := subscriptions.New(client)

	_, err := svc.DeleteByID("1234", 2)
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
	svc := subscriptions.New(client)

	_, err := svc.DeleteByKey("1234", 2)
	assert.Nil(t, err)

	params := url.Values{}
	params.Add("version", "2")
	assert.Equal(t, params, output.URL.Query())
	assert.Equal(t, "/unittest/subscriptions/key=1234", output.URL.Path)
}

func TestSubscriptionGetDestinationIronMQ(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.ironmq.json"), nil, nil)
	defer server.Close()

	svc := subscriptions.New(client)
	subscription, err := svc.GetByID("100")
	assert.Nil(t, err)

	destination := subscription.Destination.(subscriptions.DestinationIronMQ)
	expected := subscriptions.DestinationIronMQ{
		URI: "https://queue-uri",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSNS(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sns.json"), nil, nil)
	defer server.Close()

	svc := subscriptions.New(client)
	subscription, err := svc.GetByID("100")
	assert.Nil(t, err)

	destination := subscription.Destination.(subscriptions.DestinationAWSSNS)
	expected := subscriptions.DestinationAWSSNS{
		TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSQS(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("subscription.sqs.json"), nil, nil)
	defer server.Close()

	svc := subscriptions.New(client)
	subscription, err := svc.GetByID("100")
	assert.Nil(t, err)

	destination := subscription.Destination.(subscriptions.DestinationAWSSQS)
	expected := subscriptions.DestinationAWSSQS{
		QueueURL:     "https://queue.amazonaws.com/80398EXAMPLE/MyQueue",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}
