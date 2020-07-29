// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedCartDiscountGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z",
	  "name": {
	    "en": "Summer Sale"
	  },
	  "value": {
	    "type": "relative",
	    "permyriad": 1000
	  },
	  "cartPredicate": "1=1",
	  "target": {
	    "type": "lineItems",
	    "predicate": "1=1"
	  },
	  "sortOrder": "0.1",
	  "isActive": true,
	  "requiresDiscountCode": false,
	  "references": [],
	  "stackingMode": "Stacking"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart_discount, err := client.CartDiscountGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart_discount)
	assert.NotNil(t, cart_discount.Version)
	assert.NotEmpty(t, cart_discount.StackingMode)
	assert.NotEmpty(t, cart_discount.SortOrder)
	assert.False(t, cart_discount.RequiresDiscountCode)
	assert.NotEmpty(t, cart_discount.LastModifiedAt)
	assert.True(t, cart_discount.IsActive)
	assert.NotEmpty(t, cart_discount.ID)
	assert.NotEmpty(t, cart_discount.CreatedAt)
	assert.NotEmpty(t, cart_discount.CartPredicate)

}

func TestGeneratedCartDiscountGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z",
	  "name": {
	    "en": "Summer Sale"
	  },
	  "value": {
	    "type": "relative",
	    "permyriad": 1000
	  },
	  "cartPredicate": "1=1",
	  "target": {
	    "type": "lineItems",
	    "predicate": "1=1"
	  },
	  "sortOrder": "0.1",
	  "isActive": true,
	  "requiresDiscountCode": false,
	  "references": [],
	  "stackingMode": "Stacking"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart_discount, err := client.CartDiscountGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart_discount)
	assert.NotNil(t, cart_discount.Version)
	assert.NotEmpty(t, cart_discount.StackingMode)
	assert.NotEmpty(t, cart_discount.SortOrder)
	assert.False(t, cart_discount.RequiresDiscountCode)
	assert.NotEmpty(t, cart_discount.LastModifiedAt)
	assert.True(t, cart_discount.IsActive)
	assert.NotEmpty(t, cart_discount.ID)
	assert.NotEmpty(t, cart_discount.CreatedAt)
	assert.NotEmpty(t, cart_discount.CartPredicate)

}

func TestGeneratedCartDiscountDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z",
	  "name": {
	    "en": "Summer Sale"
	  },
	  "value": {
	    "type": "relative",
	    "permyriad": 1000
	  },
	  "cartPredicate": "1=1",
	  "target": {
	    "type": "lineItems",
	    "predicate": "1=1"
	  },
	  "sortOrder": "0.1",
	  "isActive": true,
	  "requiresDiscountCode": false,
	  "references": [],
	  "stackingMode": "Stacking"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart_discount, err := client.CartDiscountDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart_discount)
	assert.NotNil(t, cart_discount.Version)
	assert.NotEmpty(t, cart_discount.StackingMode)
	assert.NotEmpty(t, cart_discount.SortOrder)
	assert.False(t, cart_discount.RequiresDiscountCode)
	assert.NotEmpty(t, cart_discount.LastModifiedAt)
	assert.True(t, cart_discount.IsActive)
	assert.NotEmpty(t, cart_discount.ID)
	assert.NotEmpty(t, cart_discount.CreatedAt)
	assert.NotEmpty(t, cart_discount.CartPredicate)

}

func TestGeneratedCartDiscountDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	  "version": 1,
	  "createdAt": "1970-01-01T00:00:00.001Z",
	  "lastModifiedAt": "1970-01-01T00:00:00.001Z",
	  "name": {
	    "en": "Summer Sale"
	  },
	  "value": {
	    "type": "relative",
	    "permyriad": 1000
	  },
	  "cartPredicate": "1=1",
	  "target": {
	    "type": "lineItems",
	    "predicate": "1=1"
	  },
	  "sortOrder": "0.1",
	  "isActive": true,
	  "requiresDiscountCode": false,
	  "references": [],
	  "stackingMode": "Stacking"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart_discount, err := client.CartDiscountDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart_discount)
	assert.NotNil(t, cart_discount.Version)
	assert.NotEmpty(t, cart_discount.StackingMode)
	assert.NotEmpty(t, cart_discount.SortOrder)
	assert.False(t, cart_discount.RequiresDiscountCode)
	assert.NotEmpty(t, cart_discount.LastModifiedAt)
	assert.True(t, cart_discount.IsActive)
	assert.NotEmpty(t, cart_discount.ID)
	assert.NotEmpty(t, cart_discount.CreatedAt)
	assert.NotEmpty(t, cart_discount.CartPredicate)

}

func TestGeneratedCartDiscountQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "c2f93298-c967-44af-8c2a-d2220bf39eb2",
	      "version": 1,
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "1970-01-01T00:00:00.001Z",
	      "name": {
	        "en": "Summer Sale"
	      },
	      "value": {
	        "type": "relative",
	        "permyriad": 1000
	      },
	      "cartPredicate": "1=1",
	      "target": {
	        "type": "lineItems",
	        "predicate": "1=1"
	      },
	      "sortOrder": "0.1",
	      "isActive": true,
	      "requiresDiscountCode": false,
	      "references": [],
	      "stackingMode": "Stacking"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.CartDiscountQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
