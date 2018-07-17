package extension_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/common"
	"github.com/labd/commercetools-go-sdk/extension"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestSubscriptionCreate(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), nil, nil)
	defer server.Close()
	svc := extension.New(client)

	draft := &extension.SubscriptionDraft{
		Key: "test",
		Destination: extension.SubscriptionAWSSQSDestination{
			Type:         "SQS",
			QueueURL:     "http://example.com/",
			AccessKey:    "A1234567890",
			AccessSecret: "S1234567800",
			Region:       "eu-central-1",
		},
		Messages: []extension.MessageSubscription{
			extension.MessageSubscription{
				ResourceTypeID: "product",
			},
		},
	}

	_, err := svc.SubscriptionCreate(draft)
	assert.Equal(t, nil, err)
}

func TestSubscriptionUpdate(t *testing.T) {
	var output map[string]interface{}

	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), &output, nil)
	defer server.Close()
	svc := extension.New(client)

	input := &extension.SubscriptionUpdateInput{
		Version: 2,
		Actions: common.UpdateActions{
			&extension.SubscriptionSetKey{
				Key: "123456",
			},
			&extension.SubscriptionSetMessages{
				Messages: []extension.MessageSubscription{
					extension.MessageSubscription{
						ResourceTypeID: "product",
					},
				},
			},
		},
	}

	fmt.Println(output)

	_, err := svc.SubscriptionUpdate(input)
	assert.Equal(t, nil, err)
}

func TestSubscriptionDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), nil, nil)
	defer server.Close()
	svc := extension.New(client)

	_, err := svc.SubscriptionDeleteByID("1234", 2)
	assert.Equal(t, nil, err)
}

func TestSubscriptionDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), nil, nil)
	defer server.Close()
	svc := extension.New(client)

	_, err := svc.SubscriptionDeleteByKey("1234", 2)
	assert.Equal(t, nil, err)
}

func TestSubscriptionGetDestinationIronMQ(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.ironmq.json"), nil, nil)
	defer server.Close()

	svc := extension.New(client)
	subscription, err := svc.SubscriptionGetByID("100")
	assert.Equal(t, nil, err)

	destination := subscription.Destination.(extension.SubscriptionIronMQDestination)
	expected := extension.SubscriptionIronMQDestination{
		Type: "IronMQ",
		URI:  "https://queue-uri",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSNS(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), nil, nil)
	defer server.Close()

	svc := extension.New(client)
	subscription, err := svc.SubscriptionGetByID("100")
	assert.Equal(t, nil, err)

	destination := subscription.Destination.(extension.SubscriptionAWSSNSDestination)
	expected := extension.SubscriptionAWSSNSDestination{
		Type:         "SNS",
		TopicArn:     "arn:aws:sns:eu-central-1:123456789012345678:example:1",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}

func TestSubscriptionGetDestinationSQS(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.sqs.json"), nil, nil)
	defer server.Close()

	svc := extension.New(client)
	subscription, err := svc.SubscriptionGetByID("100")
	assert.Equal(t, nil, err)

	destination := subscription.Destination.(extension.SubscriptionAWSSQSDestination)
	expected := extension.SubscriptionAWSSQSDestination{
		Type:         "SQS",
		QueueURL:     "https://queue.amazonaws.com/80398EXAMPLE/MyQueue",
		AccessKey:    "AKIAIOSFODNN7EXAMPLE",
		AccessSecret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
	}
	assert.Equal(t, destination, expected)
}
