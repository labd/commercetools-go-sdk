// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedDiscountCodeGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "3b3327a4-06c2-4129-8d87-83d5ce53e038",
	  "version": 2,
	  "code": "SAVE10",
	  "name": {
	    "en": "Save10"
	  },
	  "cartDiscounts": [
	    {
	      "typeId": "cart-discount",
	      "id": "fdbaf4ea-fbc9-4fea-bac4-1d7e6c1995b3"
	    }
	  ],
	  "isActive": true,
	  "cartPredicate": "1=1",
	  "references": [],
	  "groups": [],
	  "createdAt": "2015-09-15T09:02:29.323Z",
	  "lastModifiedAt": "2015-09-15T09:04:06.910Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	discount_code, err := client.DiscountCodeGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, discount_code)
	assert.NotNil(t, discount_code.Version)
	assert.NotEmpty(t, discount_code.LastModifiedAt)
	assert.True(t, discount_code.IsActive)
	assert.NotEmpty(t, discount_code.ID)
	assert.NotEmpty(t, discount_code.CreatedAt)
	assert.NotEmpty(t, discount_code.Code)
	assert.NotEmpty(t, discount_code.CartPredicate)

}

func TestGeneratedDiscountCodeDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "3b3327a4-06c2-4129-8d87-83d5ce53e038",
	  "version": 2,
	  "code": "SAVE10",
	  "name": {
	    "en": "Save10"
	  },
	  "cartDiscounts": [
	    {
	      "typeId": "cart-discount",
	      "id": "fdbaf4ea-fbc9-4fea-bac4-1d7e6c1995b3"
	    }
	  ],
	  "isActive": true,
	  "cartPredicate": "1=1",
	  "references": [],
	  "groups": [],
	  "createdAt": "2015-09-15T09:02:29.323Z",
	  "lastModifiedAt": "2015-09-15T09:04:06.910Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	discount_code, err := client.DiscountCodeDeleteWithID(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, discount_code)
	assert.NotNil(t, discount_code.Version)
	assert.NotEmpty(t, discount_code.LastModifiedAt)
	assert.True(t, discount_code.IsActive)
	assert.NotEmpty(t, discount_code.ID)
	assert.NotEmpty(t, discount_code.CreatedAt)
	assert.NotEmpty(t, discount_code.Code)
	assert.NotEmpty(t, discount_code.CartPredicate)

}

func TestGeneratedDiscountCodeQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "3b3327a4-06c2-4129-8d87-83d5ce53e038",
	      "version": 2,
	      "code": "SAVE10",
	      "name": {
	        "en": "Save10"
	      },
	      "cartDiscounts": [
	        {
	          "typeId": "cart-discount",
	          "id": "fdbaf4ea-fbc9-4fea-bac4-1d7e6c1995b3"
	        }
	      ],
	      "isActive": true,
	      "cartPredicate": "1=1",
	      "references": [],
	      "groups": [],
	      "attributeTypes": {},
	      "cartFieldTypes": {},
	      "lineItemFieldTypes": {},
	      "customLineItemFieldTypes": {},
	      "createdAt": "2015-09-15T09:02:29.323Z",
	      "lastModifiedAt": "2015-09-15T09:04:06.910Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.DiscountCodeQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
