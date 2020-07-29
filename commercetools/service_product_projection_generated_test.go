// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedProductProjectionGetWithID(t *testing.T) {
	responseData := ` {
	      "id": "080feded-4f74-4d31-9309-f7ef6b7f1279",
	      "version": 3,
	      "productType": {
	        "typeId": "product-type",
	        "id": "1c095f1b-e638-4c7e-86c4-c58df873fca6"
	      },
	      "name": {
	        "en": "Some Products"
	      },
	      "categories": [],
	      "slug": {
	        "en": "product_slug_jxeutnxwkswk"
	      },
	      "masterVariant": {
	        "id": 1,
	        "prices": [],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text1",
	            "value": {
	              "it": "italian1",
	              "de": "german1",
	              "en": "englisch1"
	            }
	          }
	        ]
	      },
	      "variants": [],
	      "searchKeywords": {},
	      "hasStagedChanges": false,
	      "published": true,
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "2014-01-07T15:25:50.034Z"
	    } `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_projection, err := client.ProductProjectionGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_projection)
	assert.NotEmpty(t, product_projection.LastModifiedAt)
	assert.NotEmpty(t, product_projection.CreatedAt)
	assert.NotNil(t, product_projection.Version)
	assert.True(t, product_projection.Published)
	assert.NotEmpty(t, product_projection.ID)
	assert.False(t, product_projection.HasStagedChanges)

}

func TestGeneratedProductProjectionGetWithKey(t *testing.T) {
	responseData := ` {
	      "id": "080feded-4f74-4d31-9309-f7ef6b7f1279",
	      "version": 3,
	      "productType": {
	        "typeId": "product-type",
	        "id": "1c095f1b-e638-4c7e-86c4-c58df873fca6"
	      },
	      "name": {
	        "en": "Some Products"
	      },
	      "categories": [],
	      "slug": {
	        "en": "product_slug_jxeutnxwkswk"
	      },
	      "masterVariant": {
	        "id": 1,
	        "prices": [],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text1",
	            "value": {
	              "it": "italian1",
	              "de": "german1",
	              "en": "englisch1"
	            }
	          }
	        ]
	      },
	      "variants": [],
	      "searchKeywords": {},
	      "hasStagedChanges": false,
	      "published": true,
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "2014-01-07T15:25:50.034Z"
	    } `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product_projection, err := client.ProductProjectionGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product_projection)
	assert.NotEmpty(t, product_projection.LastModifiedAt)
	assert.NotEmpty(t, product_projection.CreatedAt)
	assert.NotNil(t, product_projection.Version)
	assert.True(t, product_projection.Published)
	assert.NotEmpty(t, product_projection.ID)
	assert.False(t, product_projection.HasStagedChanges)

}

func TestGeneratedProductProjectionQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "id": "080feded-4f74-4d31-9309-f7ef6b7f1279",
	      "version": 3,
	      "productType": {
	        "typeId": "product-type",
	        "id": "1c095f1b-e638-4c7e-86c4-c58df873fca6"
	      },
	      "name": {
	        "en": "Some Products"
	      },
	      "categories": [],
	      "slug": {
	        "en": "product_slug_jxeutnxwkswk"
	      },
	      "masterVariant": {
	        "id": 1,
	        "prices": [],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text1",
	            "value": {
	              "it": "italian1",
	              "de": "german1",
	              "en": "englisch1"
	            }
	          }
	        ]
	      },
	      "variants": [],
	      "searchKeywords": {},
	      "hasStagedChanges": false,
	      "published": true,
	      "createdAt": "1970-01-01T00:00:00.001Z",
	      "lastModifiedAt": "2014-01-07T15:25:50.034Z"
	    },
	    {
	      "id": "e779ec1a-0a98-4135-8344-d51bdafd4fe6",
	      "version": 6,
	      "productType": {
	        "typeId": "product-type",
	        "id": "bac5ef74-e0f9-4792-8439-ab84b270599e"
	      },
	      "name": {
	        "en": "product with dates"
	      },
	      "description": {
	        "en": "<p>Used to test the various date attributes<br></p>"
	      },
	      "categories": [],
	      "slug": {
	        "aa": "product-with-dates-1432797256269",
	        "de": "product-with-dates-1432797256269",
	        "en": "product-with-dates-1432797256269"
	      },
	      "masterVariant": {
	        "id": 1,
	        "prices": [],
	        "images": [],
	        "attributes": [
	          {
	            "name": "aboolean",
	            "value": true
	          }
	        ]
	      },
	      "variants": [],
	      "searchKeywords": {},
	      "hasStagedChanges": true,
	      "published": true,
	      "createdAt": "2015-05-28T07:15:48.738Z",
	      "lastModifiedAt": "2015-05-28T07:19:38.624Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ProductProjectionQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
