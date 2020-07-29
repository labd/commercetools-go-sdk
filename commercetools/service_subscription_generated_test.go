// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedSubscriptionGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "destination": {
	    "type": "IronMQ",
	    "uri": "https://queue-uri"
	  },
	  "messages": [
	    {
	      "resourceTypeId": "product",
	      "types": []
	    }
	  ],
	  "changes": [],
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "key": "test-queue",
	  "format": {
	    "type": "Platform"
	  },
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "lastMessageSequenceNumber": 0,
	  "status": "Healthy"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	subscription, err := client.SubscriptionGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, subscription)
	assert.NotNil(t, subscription.Version)
	assert.NotEmpty(t, subscription.Status)
	assert.NotEmpty(t, subscription.LastModifiedAt)
	assert.NotEmpty(t, subscription.Key)
	assert.NotEmpty(t, subscription.ID)
	assert.NotEmpty(t, subscription.CreatedAt)

}

func TestGeneratedSubscriptionGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "destination": {
	    "type": "IronMQ",
	    "uri": "https://queue-uri"
	  },
	  "messages": [
	    {
	      "resourceTypeId": "product",
	      "types": []
	    }
	  ],
	  "changes": [],
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "key": "test-queue",
	  "format": {
	    "type": "Platform"
	  },
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "lastMessageSequenceNumber": 0,
	  "status": "Healthy"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	subscription, err := client.SubscriptionGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, subscription)
	assert.NotNil(t, subscription.Version)
	assert.NotEmpty(t, subscription.Status)
	assert.NotEmpty(t, subscription.LastModifiedAt)
	assert.NotEmpty(t, subscription.Key)
	assert.NotEmpty(t, subscription.ID)
	assert.NotEmpty(t, subscription.CreatedAt)

}

func TestGeneratedSubscriptionDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "destination": {
	    "type": "IronMQ",
	    "uri": "https://queue-uri"
	  },
	  "messages": [
	    {
	      "resourceTypeId": "product",
	      "types": []
	    }
	  ],
	  "changes": [],
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "key": "test-queue",
	  "format": {
	    "type": "Platform"
	  },
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "lastMessageSequenceNumber": 0,
	  "status": "Healthy"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	subscription, err := client.SubscriptionDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, subscription)
	assert.NotNil(t, subscription.Version)
	assert.NotEmpty(t, subscription.Status)
	assert.NotEmpty(t, subscription.LastModifiedAt)
	assert.NotEmpty(t, subscription.Key)
	assert.NotEmpty(t, subscription.ID)
	assert.NotEmpty(t, subscription.CreatedAt)

}

func TestGeneratedSubscriptionDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	  "version": 1,
	  "destination": {
	    "type": "IronMQ",
	    "uri": "https://queue-uri"
	  },
	  "messages": [
	    {
	      "resourceTypeId": "product",
	      "types": []
	    }
	  ],
	  "changes": [],
	  "createdAt": "2017-01-25T14:14:22.417Z",
	  "key": "test-queue",
	  "format": {
	    "type": "Platform"
	  },
	  "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	  "lastMessageSequenceNumber": 0,
	  "status": "Healthy"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	subscription, err := client.SubscriptionDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, subscription)
	assert.NotNil(t, subscription.Version)
	assert.NotEmpty(t, subscription.Status)
	assert.NotEmpty(t, subscription.LastModifiedAt)
	assert.NotEmpty(t, subscription.Key)
	assert.NotEmpty(t, subscription.ID)
	assert.NotEmpty(t, subscription.CreatedAt)

}

func TestGeneratedSubscriptionQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "8062243c-46fc-40b5-88a4-75e2216aef75",
	      "version": 1,
	      "destination": {
	        "type": "IronMQ",
	        "uri": "https://queue-uri"
	      },
	      "messages": [
	        {
	          "resourceTypeId": "product",
	          "types": []
	        }
	      ],
	      "changes": [],
	      "createdAt": "2017-01-25T14:14:22.417Z",
	      "key": "test-queue",
	      "format": {
	        "type": "Platform"
	      },
	      "lastModifiedAt": "2017-01-25T14:14:22.417Z",
	      "lastMessageSequenceNumber": 0,
	      "status": "Healthy"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.SubscriptionQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
