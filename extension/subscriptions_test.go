package extension_test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

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
	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), nil)
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

func TestSubscriptionGetDestinationIronMQ(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.ironmq.json"), nil)
	defer server.Close()

	svc := extension.New(client)
	subscription, err := svc.SubscriptionGetByID("100")
	assert.Equal(t, nil, err)

	destination := subscription.Destination.(extension.SubscriptionIronMQDestination)
	assert.Equal(t, "https://queue-uri", destination.URI)
	log.Println(subscription)
}

func TestSubscriptionGetDestinationSNS(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("subscription.sns.json"), nil)
	defer server.Close()

	svc := extension.New(client)

	subscription, err := svc.SubscriptionGetByID("100")
	assert.Equal(t, nil, err)

	destination := subscription.Destination.(extension.SubscriptionAWSSNSDestination)
	assert.Equal(t, "SNS", destination.Type)
	assert.Equal(t, "arn:aws:sns:eu-central-1:123456789012345678:example:1", destination.TopicArn)
	assert.Equal(t, "AKIAIOSFODNN7EXAMPLE", destination.AccessKey)
	assert.Equal(t, "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY", destination.AccessSecret)

	log.Println(subscription)
}
