// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedProductDiscountGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "d7a204f9-7746-4857-b17e-71f1235a691a",
	  "version": 2,
	  "value": {
	    "type": "absolute",
	    "money": [
	      {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 100
	      }
	    ],
	    "id": "3544d4a3-1279-443c-8699-4220753bf6f5"
	  },
	  "predicate": "1=1",
	  "name": {
	    "en": "test-discount1"
	  },
	  "description": {
	    "en": "test-discount1"
	  },
	  "isActive": false,
	  "sortOrder": "0.9534",
	  "references": [],
	  "createdAt": "2016-02-24T09:44:57.858Z",
	  "lastModifiedAt": "2016-02-24T09:44:57.939Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_discount, err := client.ProductDiscountGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_discount)
	assert.NotNil(t, product_discount.Version)
	assert.NotEmpty(t, product_discount.SortOrder)
	assert.NotEmpty(t, product_discount.Predicate)
	assert.NotEmpty(t, product_discount.LastModifiedAt)
	assert.False(t, product_discount.IsActive)
	assert.NotEmpty(t, product_discount.ID)
	assert.NotEmpty(t, product_discount.CreatedAt)

}

func TestGeneratedProductDiscountGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "d7a204f9-7746-4857-b17e-71f1235a691a",
	  "version": 2,
	  "value": {
	    "type": "absolute",
	    "money": [
	      {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 100
	      }
	    ],
	    "id": "3544d4a3-1279-443c-8699-4220753bf6f5"
	  },
	  "predicate": "1=1",
	  "name": {
	    "en": "test-discount1"
	  },
	  "description": {
	    "en": "test-discount1"
	  },
	  "isActive": false,
	  "sortOrder": "0.9534",
	  "references": [],
	  "createdAt": "2016-02-24T09:44:57.858Z",
	  "lastModifiedAt": "2016-02-24T09:44:57.939Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_discount, err := client.ProductDiscountGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_discount)
	assert.NotNil(t, product_discount.Version)
	assert.NotEmpty(t, product_discount.SortOrder)
	assert.NotEmpty(t, product_discount.Predicate)
	assert.NotEmpty(t, product_discount.LastModifiedAt)
	assert.False(t, product_discount.IsActive)
	assert.NotEmpty(t, product_discount.ID)
	assert.NotEmpty(t, product_discount.CreatedAt)

}

func TestGeneratedProductDiscountDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "d7a204f9-7746-4857-b17e-71f1235a691a",
	  "version": 2,
	  "value": {
	    "type": "absolute",
	    "money": [
	      {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 100
	      }
	    ],
	    "id": "3544d4a3-1279-443c-8699-4220753bf6f5"
	  },
	  "predicate": "1=1",
	  "name": {
	    "en": "test-discount1"
	  },
	  "description": {
	    "en": "test-discount1"
	  },
	  "isActive": false,
	  "sortOrder": "0.9534",
	  "references": [],
	  "createdAt": "2016-02-24T09:44:57.858Z",
	  "lastModifiedAt": "2016-02-24T09:44:57.939Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_discount, err := client.ProductDiscountDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_discount)
	assert.NotNil(t, product_discount.Version)
	assert.NotEmpty(t, product_discount.SortOrder)
	assert.NotEmpty(t, product_discount.Predicate)
	assert.NotEmpty(t, product_discount.LastModifiedAt)
	assert.False(t, product_discount.IsActive)
	assert.NotEmpty(t, product_discount.ID)
	assert.NotEmpty(t, product_discount.CreatedAt)

}

func TestGeneratedProductDiscountDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "d7a204f9-7746-4857-b17e-71f1235a691a",
	  "version": 2,
	  "value": {
	    "type": "absolute",
	    "money": [
	      {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 100
	      }
	    ],
	    "id": "3544d4a3-1279-443c-8699-4220753bf6f5"
	  },
	  "predicate": "1=1",
	  "name": {
	    "en": "test-discount1"
	  },
	  "description": {
	    "en": "test-discount1"
	  },
	  "isActive": false,
	  "sortOrder": "0.9534",
	  "references": [],
	  "createdAt": "2016-02-24T09:44:57.858Z",
	  "lastModifiedAt": "2016-02-24T09:44:57.939Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_discount, err := client.ProductDiscountDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_discount)
	assert.NotNil(t, product_discount.Version)
	assert.NotEmpty(t, product_discount.SortOrder)
	assert.NotEmpty(t, product_discount.Predicate)
	assert.NotEmpty(t, product_discount.LastModifiedAt)
	assert.False(t, product_discount.IsActive)
	assert.NotEmpty(t, product_discount.ID)
	assert.NotEmpty(t, product_discount.CreatedAt)

}

func TestGeneratedProductDiscountQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "d7a204f9-7746-4857-b17e-71f1235a691a",
	      "version": 2,
	      "value": {
	        "type": "absolute",
	        "money": [
	          {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 100
	          }
	        ]
	      },
	      "predicate": "1=1",
	      "name": {
	        "en": "test-discount1"
	      },
	      "description": {
	        "en": "test-discount1"
	      },
	      "isActive": false,
	      "sortOrder": "0.9534",
	      "references": [],
	      "createdAt": "2016-02-24T09:44:57.858Z",
	      "lastModifiedAt": "2016-02-24T09:44:57.939Z"
	    },
	    {
	      "id": "29cc032d-8ee3-4763-b2ab-3fe0dda22a2d",
	      "version": 2,
	      "value": {
	        "type": "absolute",
	        "money": [
	          {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 100
	          }
	        ]
	      },
	      "predicate": "1=1",
	      "name": {
	        "en": "test-discount2"
	      },
	      "description": {
	        "en": "test-discount2"
	      },
	      "isActive": false,
	      "sortOrder": "0.9478",
	      "references": [],
	      "createdAt": "2016-02-24T10:25:04.206Z",
	      "lastModifiedAt": "2016-02-24T10:25:04.295Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ProductDiscountQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
