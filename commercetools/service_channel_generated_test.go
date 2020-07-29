// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedChannelGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "ac390383-370f-43f8-a534-db1604cb96a8",
	  "key": "commercetools",
	  "version": 1,
	  "roles": [
	    "InventorySupply"
	  ],
	  "createdAt": "2015-05-28T09:48:35.023Z",
	  "lastModifiedAt": "2015-05-28T09:48:35.023Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	channel, err := client.ChannelGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, channel)
	assert.NotNil(t, channel.Version)
	assert.NotEmpty(t, channel.LastModifiedAt)
	assert.NotEmpty(t, channel.Key)
	assert.NotEmpty(t, channel.ID)
	assert.NotEmpty(t, channel.CreatedAt)

}

func TestGeneratedChannelDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "ac390383-370f-43f8-a534-db1604cb96a8",
	  "key": "commercetools",
	  "version": 1,
	  "roles": [
	    "InventorySupply"
	  ],
	  "createdAt": "2015-05-28T09:48:35.023Z",
	  "lastModifiedAt": "2015-05-28T09:48:35.023Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	channel, err := client.ChannelDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, channel)
	assert.NotNil(t, channel.Version)
	assert.NotEmpty(t, channel.LastModifiedAt)
	assert.NotEmpty(t, channel.Key)
	assert.NotEmpty(t, channel.ID)
	assert.NotEmpty(t, channel.CreatedAt)

}

func TestGeneratedChannelQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "ac390383-370f-43f8-a534-db1604cb96a8",
	      "key": "channel1",
	      "version": 1,
	      "roles": [
	        "InventorySupply"
	      ],
	      "createdAt": "2015-05-28T09:48:35.023Z",
	      "lastModifiedAt": "2015-05-28T09:48:35.023Z"
	    },
	    {
	      "id": "51323ad2-89f2-4233-b9be-f15c049769c8",
	      "key": "channel2",
	      "version": 2,
	      "roles": [
	        "InventorySupply"
	      ],
	      "createdAt": "2017-01-10T06:51:08.866Z",
	      "lastModifiedAt": "2017-01-10T06:51:08.924Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ChannelQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
