// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedZoneGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b",
	  "version": 1,
	  "name": "US",
	  "key": "us-zone",
	  "locations": [
	    {
	      "country": "US"
	    }
	  ],
	  "createdAt": "2015-01-21T09:22:04.275Z",
	  "lastModifiedAt": "2015-01-21T09:22:04.275Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	zone, err := client.ZoneGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, zone)
	assert.NotNil(t, zone.Version)
	assert.NotEmpty(t, zone.Name)
	assert.NotEmpty(t, zone.LastModifiedAt)
	assert.NotEmpty(t, zone.Key)
	assert.NotEmpty(t, zone.ID)
	assert.NotEmpty(t, zone.CreatedAt)

}

func TestGeneratedZoneGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b",
	  "version": 1,
	  "name": "US",
	  "key": "us-zone",
	  "locations": [
	    {
	      "country": "US"
	    }
	  ],
	  "createdAt": "2015-01-21T09:22:04.275Z",
	  "lastModifiedAt": "2015-01-21T09:22:04.275Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	zone, err := client.ZoneGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, zone)
	assert.NotNil(t, zone.Version)
	assert.NotEmpty(t, zone.Name)
	assert.NotEmpty(t, zone.LastModifiedAt)
	assert.NotEmpty(t, zone.Key)
	assert.NotEmpty(t, zone.ID)
	assert.NotEmpty(t, zone.CreatedAt)

}

func TestGeneratedZoneDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b",
	  "version": 1,
	  "name": "US",
	  "key": "us-zone",
	  "locations": [
	    {
	      "country": "US"
	    }
	  ],
	  "createdAt": "2015-01-21T09:22:04.275Z",
	  "lastModifiedAt": "2015-01-21T09:22:04.275Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	zone, err := client.ZoneDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, zone)
	assert.NotNil(t, zone.Version)
	assert.NotEmpty(t, zone.Name)
	assert.NotEmpty(t, zone.LastModifiedAt)
	assert.NotEmpty(t, zone.Key)
	assert.NotEmpty(t, zone.ID)
	assert.NotEmpty(t, zone.CreatedAt)

}

func TestGeneratedZoneDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b",
	  "version": 1,
	  "name": "US",
	  "key": "us-zone",
	  "locations": [
	    {
	      "country": "US"
	    }
	  ],
	  "createdAt": "2015-01-21T09:22:04.275Z",
	  "lastModifiedAt": "2015-01-21T09:22:04.275Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	zone, err := client.ZoneDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, zone)
	assert.NotNil(t, zone.Version)
	assert.NotEmpty(t, zone.Name)
	assert.NotEmpty(t, zone.LastModifiedAt)
	assert.NotEmpty(t, zone.Key)
	assert.NotEmpty(t, zone.ID)
	assert.NotEmpty(t, zone.CreatedAt)

}

func TestGeneratedZoneQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "ebe01381-82be-4e63-9993-d1eb8f8e588b",
	      "version": 1,
	      "name": "US",
	      "locations": [
	        {
	          "country": "US"
	        }
	      ],
	      "createdAt": "2015-01-21T09:22:04.275Z",
	      "lastModifiedAt": "2015-01-21T09:22:04.275Z"
	    },
	    {
	      "id": "5cb532be-c680-43ab-ba14-b664bb03dc35",
	      "version": 3,
	      "name": "Europe",
	      "locations": [
	        {
	          "country": "DE"
	        },
	        {
	          "country": "IT"
	        },
	        {
	          "country": "FR"
	        },
	        {
	          "country": "ES"
	        }
	      ],
	      "createdAt": "2015-01-21T09:22:04.266Z",
	      "lastModifiedAt": "2016-01-26T10:56:32.504Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ZoneQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
