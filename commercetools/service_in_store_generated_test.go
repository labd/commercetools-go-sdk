// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedStoreCartGetWithCustomerID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreCartGetWithCustomerID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreCartGetWithID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreCartGetWithID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreCartGetWithKey(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreCartGetWithKey(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreCartDeleteWithID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreCartDeleteWithID(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreCartDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreCartDeleteWithKey(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreCartQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "type": "Cart",
	      "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	      "version": 5,
	      "createdAt": "2015-09-22T15:36:17.510Z",
	      "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	      "lineItems": [
	        {
	          "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	          "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	          "name": {
	            "en": "SAPPHIRE"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	            "version": 8
	          },
	          "productSlug": {
	            "en": "sapphire1421832124423"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_SAPPHIRE_variant1_1421832124423",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 2800
	                },
	                "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	                "dimensions": {
	                  "w": 1400,
	                  "h": 1400
	                }
	              }
	            ],
	            "attributes": [],
	            "assets": []
	          },
	          "price": {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          },
	          "quantity": 2,
	          "discountedPricePerQuantity": [],
	          "state": [
	            {
	              "quantity": 2,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 5600
	          }
	        }
	      ],
	      "cartState": "Active",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      },
	      "customLineItems": [],
	      "discountCodes": [],
	      "inventoryMode": "None",
	      "taxMode": "Platform",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "refusedGifts": [],
	      "origin": "Customer"
	    },
	    {
	      "type": "Cart",
	      "id": "668e5783-73c8-4f2d-91f4-3c90b872c700",
	      "version": 3,
	      "createdAt": "2015-10-07T07:33:05.894Z",
	      "lastModifiedAt": "2015-10-07T07:33:06.070Z",
	      "lineItems": [
	        {
	          "id": "90dff06c-272e-47fa-b8de-923dce092474",
	          "productId": "7b1203f4-66c0-438c-9a30-f4fb6be79bdf",
	          "name": {
	            "de": "WB ATHLETIC PANZER",
	            "en": "WB ATHLETIC TANK"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	            "version": 8
	          },
	          "productSlug": {
	            "en": "wb-athletic-tank1421832124574"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_WB_ATHLETIC_TANK_variant1_1421832124574",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 8400
	                },
	                "id": "37696f7c-8260-4941-a921-68e6aa76b4a3"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253265444_1.jpg",
	                "dimensions": {
	                  "w": 1400,
	                  "h": 1400
	                }
	              }
	            ],
	            "attributes": [],
	            "assets": []
	          },
	          "price": {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 8400
	            },
	            "id": "37696f7c-8260-4941-a921-68e6aa76b4a3"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 8400
	          },
	          "custom": {
	            "type": {
	              "typeId": "type",
	              "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa"
	            },
	            "fields": {
	              "offer_name": "SuperMax"
	            }
	          }
	        }
	      ],
	      "cartState": "Active",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 8400
	      },
	      "country": "DE",
	      "customLineItems": [],
	      "discountCodes": [],
	      "inventoryMode": "None",
	      "taxMode": "Platform",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "refusedGifts": [],
	      "origin": "Customer"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreCartQuery(context.TODO(), "dummy-id", &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}

func TestGeneratedStoreCustomerGetWithEmailToken(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.StoreCustomerGetWithEmailToken(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedStoreCustomerGetWithID(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.StoreCustomerGetWithID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedStoreCustomerGetWithKey(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.StoreCustomerGetWithKey(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedStoreCustomerGetWithPasswordToken(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.StoreCustomerGetWithPasswordToken(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedStoreCustomerDeleteWithID(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.StoreCustomerDeleteWithID(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedStoreCustomerDeleteWithKey(t *testing.T) {
	responseData := ` {
	      "addresses": [],
	      "email": "<random>@example.com",
	      "firstName": "John",
	      "id": "some_123_id",
	      "isEmailVerified": false,
	      "lastName": "Doe",
	      "password": "secret123",
	      "version": 1,
	      "createdAt": "2015-07-06T13:22:33.339Z",
	      "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	customer, err := client.StoreCustomerDeleteWithKey(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, customer)
	assert.NotNil(t, customer.Version)
	assert.NotEmpty(t, customer.Password)
	assert.NotEmpty(t, customer.LastName)
	assert.NotEmpty(t, customer.LastModifiedAt)
	assert.False(t, customer.IsEmailVerified)
	assert.NotEmpty(t, customer.ID)
	assert.NotEmpty(t, customer.FirstName)
	assert.NotEmpty(t, customer.Email)
	assert.NotEmpty(t, customer.CreatedAt)

}

func TestGeneratedStoreCustomerQuery(t *testing.T) {
	responseData := ` {
	      "limit": 20,
	      "offset": 0,
	      "count": 1,
	      "total": 1,
	      "results": [{
	            "addresses": [],
	            "email": "<random>@example.com",
	            "firstName": "John",
	            "id": "some_123_id",
	            "isEmailVerified": false,
	            "lastName": "Doe",
	            "password": "secret123",
	            "version": 1,
	            "createdAt": "2015-07-06T13:22:33.339Z",
	            "lastModifiedAt": "2015-07-06T13:22:33.339Z"
	      }]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreCustomerQuery(context.TODO(), "dummy-id", &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}

func TestGeneratedStoreMyCartGetWithID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreMyCartGetWithID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreMyCartDeleteWithID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.StoreMyCartDeleteWithID(context.TODO(), "dummy-id", "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedStoreMyCartQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "type": "Cart",
	      "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	      "version": 5,
	      "createdAt": "2015-09-22T15:36:17.510Z",
	      "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	      "lineItems": [
	        {
	          "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	          "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	          "name": {
	            "en": "SAPPHIRE"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	            "version": 8
	          },
	          "productSlug": {
	            "en": "sapphire1421832124423"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_SAPPHIRE_variant1_1421832124423",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 2800
	                },
	                "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	                "dimensions": {
	                  "w": 1400,
	                  "h": 1400
	                }
	              }
	            ],
	            "attributes": [],
	            "assets": []
	          },
	          "price": {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          },
	          "quantity": 2,
	          "discountedPricePerQuantity": [],
	          "state": [
	            {
	              "quantity": 2,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 5600
	          }
	        }
	      ],
	      "cartState": "Active",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      },
	      "customLineItems": [],
	      "discountCodes": [],
	      "inventoryMode": "None",
	      "taxMode": "Platform",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "refusedGifts": [],
	      "origin": "Customer"
	    },
	    {
	      "type": "Cart",
	      "id": "668e5783-73c8-4f2d-91f4-3c90b872c700",
	      "version": 3,
	      "createdAt": "2015-10-07T07:33:05.894Z",
	      "lastModifiedAt": "2015-10-07T07:33:06.070Z",
	      "lineItems": [
	        {
	          "id": "90dff06c-272e-47fa-b8de-923dce092474",
	          "productId": "7b1203f4-66c0-438c-9a30-f4fb6be79bdf",
	          "name": {
	            "de": "WB ATHLETIC PANZER",
	            "en": "WB ATHLETIC TANK"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	            "version": 8
	          },
	          "productSlug": {
	            "en": "wb-athletic-tank1421832124574"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_WB_ATHLETIC_TANK_variant1_1421832124574",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 8400
	                },
	                "id": "37696f7c-8260-4941-a921-68e6aa76b4a3"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253265444_1.jpg",
	                "dimensions": {
	                  "w": 1400,
	                  "h": 1400
	                }
	              }
	            ],
	            "attributes": [],
	            "assets": []
	          },
	          "price": {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 8400
	            },
	            "id": "37696f7c-8260-4941-a921-68e6aa76b4a3"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 8400
	          },
	          "custom": {
	            "type": {
	              "typeId": "type",
	              "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa"
	            },
	            "fields": {
	              "offer_name": "SuperMax"
	            }
	          }
	        }
	      ],
	      "cartState": "Active",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 8400
	      },
	      "country": "DE",
	      "customLineItems": [],
	      "discountCodes": [],
	      "inventoryMode": "None",
	      "taxMode": "Platform",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "refusedGifts": [],
	      "origin": "Customer"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreMyCartQuery(context.TODO(), "dummy-id", &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}

func TestGeneratedStoreMyOrderGetWithID(t *testing.T) {
	responseData := ` {
	  "type": "Order",
	  "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	  "version": 1,
	  "createdAt": "2017-01-04T19:54:49.797Z",
	  "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 1000
	  },
	  "orderState": "Open",
	  "syncInfo": [],
	  "returnInfo": [],
	  "refusedGifts": [],
	  "taxMode": "External",
	  "inventoryMode": "None",
	  "taxRoundingMode": "HalfEven",
	  "lineItems": [
	    {
	      "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	      "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	      "name": {
	        "en": "MyProduct"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	        "version": 21
	      },
	      "productSlug": {
	        "de": "neues-produkt-slug-1234678",
	        "en": "new-product-slug-12345678"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "some-sku-123",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	          }
	        ],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text",
	            "value": "attribute-value"
	          },
	          {
	            "name": "enum",
	            "value": {
	              "key": "test",
	              "label": "test"
	            }
	          }
	        ],
	        "assets": []
	      },
	      "price": {
	        "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        }
	      },
	      "quantity": 1,
	      "discountedPricePerQuantity": [],
	      "taxRate": {
	        "name": "Bla",
	        "amount": 0.1,
	        "includedInPrice": false,
	        "country": "DE",
	        "subRates": []
	      },
	      "state": [
	        {
	          "quantity": 1,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1100
	        }
	      }
	    }
	  ],
	  "customLineItems": [],
	  "discountCodes": [],
	  "lastMessageSequenceNumber": 1,
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order, err := client.StoreMyOrderGetWithID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order)
	assert.NotNil(t, order.Version)
	assert.NotEmpty(t, order.TaxRoundingMode)
	assert.NotEmpty(t, order.TaxMode)
	assert.NotEmpty(t, order.Origin)
	assert.NotEmpty(t, order.OrderState)
	assert.NotEmpty(t, order.LastModifiedAt)
	assert.NotNil(t, order.LastMessageSequenceNumber)
	assert.NotEmpty(t, order.InventoryMode)
	assert.NotEmpty(t, order.ID)
	assert.NotEmpty(t, order.CreatedAt)

}

func TestGeneratedStoreMyOrderQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "type": "Order",
	      "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	      "version": 1,
	      "createdAt": "2017-01-04T19:54:49.797Z",
	      "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "orderState": "Open",
	      "syncInfo": [],
	      "returnInfo": [],
	      "refusedGifts": [],
	      "taxMode": "External",
	      "inventoryMode": "None",
	      "taxRoundingMode": "HalfEven",
	      "lineItems": [
	        {
	          "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	          "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	          "name": {
	            "en": "MyProduct"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	            "version": 21
	          },
	          "productSlug": {
	            "de": "neues-produkt-slug-1234678",
	            "en": "new-product-slug-12345678"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "some-sku-123",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 1000
	                },
	                "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	              }
	            ],
	            "images": [],
	            "attributes": [
	              {
	                "name": "text",
	                "value": "attribute-value"
	              },
	              {
	                "name": "enum",
	                "value": {
	                  "key": "test",
	                  "label": "test"
	                }
	              }
	            ],
	            "assets": []
	          },
	          "price": {
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            }
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "taxRate": {
	            "name": "Bla",
	            "amount": 0.1,
	            "includedInPrice": false,
	            "country": "DE",
	            "subRates": []
	          },
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 1000
	          },
	          "taxedPrice": {
	            "totalNet": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "totalGross": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1100
	            }
	          }
	        }
	      ],
	      "customLineItems": [],
	      "discountCodes": [],
	      "lastMessageSequenceNumber": 1,
	      "origin": "Customer"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreMyOrderQuery(context.TODO(), "dummy-id", &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}

func TestGeneratedStoreOrderGetWithID(t *testing.T) {
	responseData := ` {
	  "type": "Order",
	  "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	  "version": 1,
	  "createdAt": "2017-01-04T19:54:49.797Z",
	  "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 1000
	  },
	  "orderState": "Open",
	  "syncInfo": [],
	  "returnInfo": [],
	  "refusedGifts": [],
	  "taxMode": "External",
	  "inventoryMode": "None",
	  "taxRoundingMode": "HalfEven",
	  "lineItems": [
	    {
	      "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	      "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	      "name": {
	        "en": "MyProduct"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	        "version": 21
	      },
	      "productSlug": {
	        "de": "neues-produkt-slug-1234678",
	        "en": "new-product-slug-12345678"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "some-sku-123",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	          }
	        ],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text",
	            "value": "attribute-value"
	          },
	          {
	            "name": "enum",
	            "value": {
	              "key": "test",
	              "label": "test"
	            }
	          }
	        ],
	        "assets": []
	      },
	      "price": {
	        "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        }
	      },
	      "quantity": 1,
	      "discountedPricePerQuantity": [],
	      "taxRate": {
	        "name": "Bla",
	        "amount": 0.1,
	        "includedInPrice": false,
	        "country": "DE",
	        "subRates": []
	      },
	      "state": [
	        {
	          "quantity": 1,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1100
	        }
	      }
	    }
	  ],
	  "customLineItems": [],
	  "discountCodes": [],
	  "lastMessageSequenceNumber": 1,
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order, err := client.StoreOrderGetWithID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order)
	assert.NotNil(t, order.Version)
	assert.NotEmpty(t, order.TaxRoundingMode)
	assert.NotEmpty(t, order.TaxMode)
	assert.NotEmpty(t, order.Origin)
	assert.NotEmpty(t, order.OrderState)
	assert.NotEmpty(t, order.LastModifiedAt)
	assert.NotNil(t, order.LastMessageSequenceNumber)
	assert.NotEmpty(t, order.InventoryMode)
	assert.NotEmpty(t, order.ID)
	assert.NotEmpty(t, order.CreatedAt)

}

func TestGeneratedStoreOrderGetWithOrderNumber(t *testing.T) {
	responseData := ` {
	  "type": "Order",
	  "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	  "version": 1,
	  "createdAt": "2017-01-04T19:54:49.797Z",
	  "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 1000
	  },
	  "orderState": "Open",
	  "syncInfo": [],
	  "returnInfo": [],
	  "refusedGifts": [],
	  "taxMode": "External",
	  "inventoryMode": "None",
	  "taxRoundingMode": "HalfEven",
	  "lineItems": [
	    {
	      "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	      "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	      "name": {
	        "en": "MyProduct"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	        "version": 21
	      },
	      "productSlug": {
	        "de": "neues-produkt-slug-1234678",
	        "en": "new-product-slug-12345678"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "some-sku-123",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	          }
	        ],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text",
	            "value": "attribute-value"
	          },
	          {
	            "name": "enum",
	            "value": {
	              "key": "test",
	              "label": "test"
	            }
	          }
	        ],
	        "assets": []
	      },
	      "price": {
	        "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        }
	      },
	      "quantity": 1,
	      "discountedPricePerQuantity": [],
	      "taxRate": {
	        "name": "Bla",
	        "amount": 0.1,
	        "includedInPrice": false,
	        "country": "DE",
	        "subRates": []
	      },
	      "state": [
	        {
	          "quantity": 1,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1100
	        }
	      }
	    }
	  ],
	  "customLineItems": [],
	  "discountCodes": [],
	  "lastMessageSequenceNumber": 1,
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order, err := client.StoreOrderGetWithOrderNumber(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order)
	assert.NotNil(t, order.Version)
	assert.NotEmpty(t, order.TaxRoundingMode)
	assert.NotEmpty(t, order.TaxMode)
	assert.NotEmpty(t, order.Origin)
	assert.NotEmpty(t, order.OrderState)
	assert.NotEmpty(t, order.LastModifiedAt)
	assert.NotNil(t, order.LastMessageSequenceNumber)
	assert.NotEmpty(t, order.InventoryMode)
	assert.NotEmpty(t, order.ID)
	assert.NotEmpty(t, order.CreatedAt)

}

func TestGeneratedStoreOrderDeleteWithID(t *testing.T) {
	responseData := ` {
	  "type": "Order",
	  "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	  "version": 1,
	  "createdAt": "2017-01-04T19:54:49.797Z",
	  "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 1000
	  },
	  "orderState": "Open",
	  "syncInfo": [],
	  "returnInfo": [],
	  "refusedGifts": [],
	  "taxMode": "External",
	  "inventoryMode": "None",
	  "taxRoundingMode": "HalfEven",
	  "lineItems": [
	    {
	      "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	      "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	      "name": {
	        "en": "MyProduct"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	        "version": 21
	      },
	      "productSlug": {
	        "de": "neues-produkt-slug-1234678",
	        "en": "new-product-slug-12345678"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "some-sku-123",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	          }
	        ],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text",
	            "value": "attribute-value"
	          },
	          {
	            "name": "enum",
	            "value": {
	              "key": "test",
	              "label": "test"
	            }
	          }
	        ],
	        "assets": []
	      },
	      "price": {
	        "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        }
	      },
	      "quantity": 1,
	      "discountedPricePerQuantity": [],
	      "taxRate": {
	        "name": "Bla",
	        "amount": 0.1,
	        "includedInPrice": false,
	        "country": "DE",
	        "subRates": []
	      },
	      "state": [
	        {
	          "quantity": 1,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1100
	        }
	      }
	    }
	  ],
	  "customLineItems": [],
	  "discountCodes": [],
	  "lastMessageSequenceNumber": 1,
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order, err := client.StoreOrderDeleteWithID(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order)
	assert.NotNil(t, order.Version)
	assert.NotEmpty(t, order.TaxRoundingMode)
	assert.NotEmpty(t, order.TaxMode)
	assert.NotEmpty(t, order.Origin)
	assert.NotEmpty(t, order.OrderState)
	assert.NotEmpty(t, order.LastModifiedAt)
	assert.NotNil(t, order.LastMessageSequenceNumber)
	assert.NotEmpty(t, order.InventoryMode)
	assert.NotEmpty(t, order.ID)
	assert.NotEmpty(t, order.CreatedAt)

}

func TestGeneratedStoreOrderDeleteWithOrderNumber(t *testing.T) {
	responseData := ` {
	  "type": "Order",
	  "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	  "version": 1,
	  "createdAt": "2017-01-04T19:54:49.797Z",
	  "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 1000
	  },
	  "orderState": "Open",
	  "syncInfo": [],
	  "returnInfo": [],
	  "refusedGifts": [],
	  "taxMode": "External",
	  "inventoryMode": "None",
	  "taxRoundingMode": "HalfEven",
	  "lineItems": [
	    {
	      "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	      "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	      "name": {
	        "en": "MyProduct"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	        "version": 21
	      },
	      "productSlug": {
	        "de": "neues-produkt-slug-1234678",
	        "en": "new-product-slug-12345678"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "some-sku-123",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	          }
	        ],
	        "images": [],
	        "attributes": [
	          {
	            "name": "text",
	            "value": "attribute-value"
	          },
	          {
	            "name": "enum",
	            "value": {
	              "key": "test",
	              "label": "test"
	            }
	          }
	        ],
	        "assets": []
	      },
	      "price": {
	        "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        }
	      },
	      "quantity": 1,
	      "discountedPricePerQuantity": [],
	      "taxRate": {
	        "name": "Bla",
	        "amount": 0.1,
	        "includedInPrice": false,
	        "country": "DE",
	        "subRates": []
	      },
	      "state": [
	        {
	          "quantity": 1,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1000
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 1100
	        }
	      }
	    }
	  ],
	  "customLineItems": [],
	  "discountCodes": [],
	  "lastMessageSequenceNumber": 1,
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order, err := client.StoreOrderDeleteWithOrderNumber(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order)
	assert.NotNil(t, order.Version)
	assert.NotEmpty(t, order.TaxRoundingMode)
	assert.NotEmpty(t, order.TaxMode)
	assert.NotEmpty(t, order.Origin)
	assert.NotEmpty(t, order.OrderState)
	assert.NotEmpty(t, order.LastModifiedAt)
	assert.NotNil(t, order.LastMessageSequenceNumber)
	assert.NotEmpty(t, order.InventoryMode)
	assert.NotEmpty(t, order.ID)
	assert.NotEmpty(t, order.CreatedAt)

}

func TestGeneratedStoreOrderQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "type": "Order",
	      "id": "92f5a867-bf19-47ab-982c-6720a03a3921",
	      "version": 1,
	      "createdAt": "2017-01-04T19:54:49.797Z",
	      "lastModifiedAt": "2017-01-04T19:54:49.797Z",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 1000
	      },
	      "orderState": "Open",
	      "syncInfo": [],
	      "returnInfo": [],
	      "refusedGifts": [],
	      "taxMode": "External",
	      "inventoryMode": "None",
	      "taxRoundingMode": "HalfEven",
	      "lineItems": [
	        {
	          "id": "7297c742-d8b0-484b-aadc-1d0ba1869dc9",
	          "productId": "c5c75935-595a-4cc3-a80e-24e9a9c55094",
	          "name": {
	            "en": "MyProduct"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "a5c7e173-9754-4f02-9b18-a600896ae0e1",
	            "version": 21
	          },
	          "productSlug": {
	            "de": "neues-produkt-slug-1234678",
	            "en": "new-product-slug-12345678"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "some-sku-123",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 1000
	                },
	                "id": "6d36dc85-6131-495d-9f20-d00f411d4124"
	              }
	            ],
	            "images": [],
	            "attributes": [
	              {
	                "name": "text",
	                "value": "attribute-value"
	              },
	              {
	                "name": "enum",
	                "value": {
	                  "key": "test",
	                  "label": "test"
	                }
	              }
	            ],
	            "assets": []
	          },
	          "price": {
	            "id": "6d36dc85-6131-495d-9f20-d00f411d4124",
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            }
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "taxRate": {
	            "name": "Bla",
	            "amount": 0.1,
	            "includedInPrice": false,
	            "country": "DE",
	            "subRates": []
	          },
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 1000
	          },
	          "taxedPrice": {
	            "totalNet": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1000
	            },
	            "totalGross": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 1100
	            }
	          }
	        }
	      ],
	      "customLineItems": [],
	      "discountCodes": [],
	      "lastMessageSequenceNumber": 1,
	      "origin": "Customer"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreOrderQuery(context.TODO(), "dummy-id", &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}

func TestGeneratedStoreShoppingListGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.StoreShoppingListGetWithID(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedStoreShoppingListGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.StoreShoppingListGetWithKey(context.TODO(), "dummy-id", "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedStoreShoppingListDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.StoreShoppingListDeleteWithID(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedStoreShoppingListDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	  "version": 1,
	  "name": {
	    "en": "test"
	  },
	  "key": "test",
	  "lineItems": [],
	  "textLineItems": [],
	  "createdAt": "2017-03-30T11:49:40.904Z",
	  "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	shopping_list, err := client.StoreShoppingListDeleteWithKey(context.TODO(), "dummy-id", "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, shopping_list)
	assert.NotNil(t, shopping_list.Version)
	assert.NotEmpty(t, shopping_list.LastModifiedAt)
	assert.NotEmpty(t, shopping_list.Key)
	assert.NotEmpty(t, shopping_list.ID)
	assert.NotEmpty(t, shopping_list.CreatedAt)

}

func TestGeneratedStoreShoppingListQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "9693f04b-5aec-467f-baa1-fc74da7d0c3d",
	      "version": 1,
	      "name": {
	        "en": "test"
	      },
	      "key": "test",
	      "lineItems": [],
	      "textLineItems": [],
	      "createdAt": "2017-03-30T11:49:40.904Z",
	      "lastModifiedAt": "2017-03-30T11:49:40.904Z"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.StoreShoppingListQuery(context.TODO(), "dummy-id", &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
