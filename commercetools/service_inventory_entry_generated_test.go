// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedInventoryEntryGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "aea4caed-accf-4667-adfe-be08ba6fdf91",
	  "version": 1,
	  "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1421832124541",
	  "quantityOnStock": 4,
	  "availableQuantity": 4,
	  "createdAt": "2015-03-11T13:36:20.720Z",
	  "lastModifiedAt": "2015-03-11T13:36:20.720Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	inventory_entry, err := client.InventoryEntryGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, inventory_entry)
	assert.NotNil(t, inventory_entry.Version)
	assert.NotEmpty(t, inventory_entry.SKU)
	assert.NotNil(t, inventory_entry.QuantityOnStock)
	assert.NotEmpty(t, inventory_entry.LastModifiedAt)
	assert.NotEmpty(t, inventory_entry.ID)
	assert.NotEmpty(t, inventory_entry.CreatedAt)
	assert.NotNil(t, inventory_entry.AvailableQuantity)

}

func TestGeneratedInventoryEntryDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "aea4caed-accf-4667-adfe-be08ba6fdf91",
	  "version": 1,
	  "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1421832124541",
	  "quantityOnStock": 4,
	  "availableQuantity": 4,
	  "createdAt": "2015-03-11T13:36:20.720Z",
	  "lastModifiedAt": "2015-03-11T13:36:20.720Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	inventory_entry, err := client.InventoryEntryDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, inventory_entry)
	assert.NotNil(t, inventory_entry.Version)
	assert.NotEmpty(t, inventory_entry.SKU)
	assert.NotNil(t, inventory_entry.QuantityOnStock)
	assert.NotEmpty(t, inventory_entry.LastModifiedAt)
	assert.NotEmpty(t, inventory_entry.ID)
	assert.NotEmpty(t, inventory_entry.CreatedAt)
	assert.NotNil(t, inventory_entry.AvailableQuantity)

}

func TestGeneratedInventoryEntryQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "aea4caed-accf-4667-adfe-be08ba6fdf91",
	      "version": 1,
	      "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1421832124541",
	      "quantityOnStock": 4,
	      "availableQuantity": 4,
	      "createdAt": "2015-03-11T13:36:20.720Z",
	      "lastModifiedAt": "2015-03-11T13:36:20.720Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.InventoryEntryQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
