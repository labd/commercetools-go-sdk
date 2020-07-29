// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedOrderGetWithID(t *testing.T) {
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
	order, err := client.OrderGetWithID(context.TODO(), "dummy-id")
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

func TestGeneratedOrderGetWithOrderNumber(t *testing.T) {
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
	order, err := client.OrderGetWithOrderNumber(context.TODO(), "dummy-id")
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

func TestGeneratedOrderDeleteWithID(t *testing.T) {
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
	order, err := client.OrderDeleteWithID(context.TODO(), "dummy-id", 1, true)
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

func TestGeneratedOrderDeleteWithOrderNumber(t *testing.T) {
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
	order, err := client.OrderDeleteWithOrderNumber(context.TODO(), "dummy-id", 1, true)
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

func TestGeneratedOrderQuery(t *testing.T) {
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
	queryResult, err := client.OrderQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}

func TestGeneratedOrderEditGetWithID(t *testing.T) {
	responseData := ` {
	  "id": "df20c5ea-b114-4aab-b330-740b0e9f3099",
	  "version": 1,
	  "resource": {
	    "typeId": "order",
	    "id": "ed454f4e-c43a-485f-a86f-046c691b1363"
	  },
	  "key": "order-edit-key",
	  "createdAt": "2018-10-04T15:22:31.639Z",
	  "lastModifiedAt": "2018-10-04T15:22:31.639Z",
	  "stagedActions": [
	    {
	      "action": "setCustomerEmail",
	      "email": "user@localhost"
	    }
	  ],
	  "result": {
	    "preview": {
	      "type": "Order",
	      "id": "ed454f4e-c43a-485f-a86f-046c691b1363",
	      "version": 3,
	      "customerId": "bf5d96ce-4704-45b2-8842-d409dd34cdfc",
	      "customerEmail": "user@localhost",
	      "createdAt": "2018-05-15T12:40:17.301Z",
	      "lastModifiedAt": "2018-05-15T12:40:17.301Z",
	      "totalPrice": {
	        "type": "centPrecision",
	        "currencyCode": "EUR",
	        "centAmount": 3970,
	        "fractionDigits": 2
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3336,
	          "fractionDigits": 2
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3970,
	          "fractionDigits": 2
	        },
	        "taxPortions": [
	          {
	            "rate": 0.19,
	            "amount": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 634,
	              "fractionDigits": 2
	            },
	            "name": "19% MwSt"
	          }
	        ]
	      },
	      "country": "DE",
	      "orderState": "Open",
	      "syncInfo": [],
	      "returnInfo": [],
	      "refusedGifts": [],
	      "shippingInfo": {
	        "shippingMethodName": "DHL",
	        "price": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 570,
	          "fractionDigits": 2
	        },
	        "shippingRate": {
	          "price": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          },
	          "tiers": []
	        },
	        "taxRate": {
	          "name": "19% MwSt",
	          "amount": 0.19,
	          "includedInPrice": true,
	          "country": "DE",
	          "id": "rrsT1Jbw",
	          "subRates": []
	        },
	        "taxCategory": {
	          "typeId": "tax-category",
	          "id": "fdeb9625-10f8-476c-a549-5d5c6d1bd412"
	        },
	        "deliveries": [],
	        "shippingMethod": {
	          "typeId": "shipping-method",
	          "id": "d18b3f77-92de-4893-b6e3-b5c9c8c1eb96"
	        },
	        "taxedPrice": {
	          "totalNet": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 479,
	            "fractionDigits": 2
	          },
	          "totalGross": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          }
	        },
	        "shippingMethodState": "MatchesCart"
	      },
	      "taxMode": "Platform",
	      "inventoryMode": "None",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "origin": "Customer",
	      "lineItems": [
	        {
	          "id": "31099128-dba8-40a7-bb6c-d12857149ff8",
	          "productId": "d6d0c517-572e-4d26-b80e-ffce825334a4",
	          "name": {
	            "en": "GIRLS HARTBREAK CREW"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "6dc75271-b0e5-4ae9-9158-faa1fff65f7b",
	            "version": 2
	          },
	          "productSlug": {
	            "en": "girls-hartbreak-crew1522841378290"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1522841378290",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "currencyCode": "EUR",
	                  "centAmount": 3400,
	                  "fractionDigits": 2
	                },
	                "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253234387_1.jpg",
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
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            },
	            "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "taxRate": {
	            "name": "19% MwSt",
	            "amount": 0.19,
	            "includedInPrice": true,
	            "country": "DE",
	            "id": "rrsT1Jbw",
	            "subRates": []
	          },
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "60a64e06-b4e8-4205-a0f3-94bc203e2d6d"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "totalPrice": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 3400,
	            "fractionDigits": 2
	          },
	          "taxedPrice": {
	            "totalNet": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 2857,
	              "fractionDigits": 2
	            },
	            "totalGross": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            }
	          },
	          "lineItemMode": "Standard"
	        }
	      ],
	      "customLineItems": [],
	      "transactionFee": true,
	      "discountCodes": [],
	      "lastMessageSequenceNumber": 3,
	      "cart": {
	        "typeId": "cart",
	        "id": "6c97e772-9769-4419-8adc-501c7c5b6088"
	      },
	      "shippingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "billingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "itemShippingAddresses": []
	    },
	    "messagePayloads": [
	      {
	        "email": "user@localhost",
	        "type": "OrderCustomerEmailSet"
	      },
	      {
	        "edit": {
	          "typeId": "order-edit",
	          "id": "df20c5ea-b114-4aab-b330-740b0e9f3099"
	        },
	        "result": {
	          "type": "Applied",
	          "appliedAt": "2018-10-04T15:29:14.091Z",
	          "excerptBeforeEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 1
	          },
	          "excerptAfterEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 3
	          }
	        },
	        "type": "OrderEditApplied"
	      }
	    ],
	    "type": "PreviewSuccess"
	  },
	  "comment": "sample-comment"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order_edit, err := client.OrderEditGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order_edit)
	assert.NotNil(t, order_edit.Version)
	assert.NotEmpty(t, order_edit.LastModifiedAt)
	assert.NotEmpty(t, order_edit.Key)
	assert.NotEmpty(t, order_edit.ID)
	assert.NotEmpty(t, order_edit.CreatedAt)
	assert.NotEmpty(t, order_edit.Comment)

}

func TestGeneratedOrderEditGetWithKey(t *testing.T) {
	responseData := ` {
	  "id": "df20c5ea-b114-4aab-b330-740b0e9f3099",
	  "version": 1,
	  "resource": {
	    "typeId": "order",
	    "id": "ed454f4e-c43a-485f-a86f-046c691b1363"
	  },
	  "key": "order-edit-key",
	  "createdAt": "2018-10-04T15:22:31.639Z",
	  "lastModifiedAt": "2018-10-04T15:22:31.639Z",
	  "stagedActions": [
	    {
	      "action": "setCustomerEmail",
	      "email": "user@localhost"
	    }
	  ],
	  "result": {
	    "preview": {
	      "type": "Order",
	      "id": "ed454f4e-c43a-485f-a86f-046c691b1363",
	      "version": 3,
	      "customerId": "bf5d96ce-4704-45b2-8842-d409dd34cdfc",
	      "customerEmail": "user@localhost",
	      "createdAt": "2018-05-15T12:40:17.301Z",
	      "lastModifiedAt": "2018-05-15T12:40:17.301Z",
	      "totalPrice": {
	        "type": "centPrecision",
	        "currencyCode": "EUR",
	        "centAmount": 3970,
	        "fractionDigits": 2
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3336,
	          "fractionDigits": 2
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3970,
	          "fractionDigits": 2
	        },
	        "taxPortions": [
	          {
	            "rate": 0.19,
	            "amount": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 634,
	              "fractionDigits": 2
	            },
	            "name": "19% MwSt"
	          }
	        ]
	      },
	      "country": "DE",
	      "orderState": "Open",
	      "syncInfo": [],
	      "returnInfo": [],
	      "refusedGifts": [],
	      "shippingInfo": {
	        "shippingMethodName": "DHL",
	        "price": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 570,
	          "fractionDigits": 2
	        },
	        "shippingRate": {
	          "price": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          },
	          "tiers": []
	        },
	        "taxRate": {
	          "name": "19% MwSt",
	          "amount": 0.19,
	          "includedInPrice": true,
	          "country": "DE",
	          "id": "rrsT1Jbw",
	          "subRates": []
	        },
	        "taxCategory": {
	          "typeId": "tax-category",
	          "id": "fdeb9625-10f8-476c-a549-5d5c6d1bd412"
	        },
	        "deliveries": [],
	        "shippingMethod": {
	          "typeId": "shipping-method",
	          "id": "d18b3f77-92de-4893-b6e3-b5c9c8c1eb96"
	        },
	        "taxedPrice": {
	          "totalNet": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 479,
	            "fractionDigits": 2
	          },
	          "totalGross": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          }
	        },
	        "shippingMethodState": "MatchesCart"
	      },
	      "taxMode": "Platform",
	      "inventoryMode": "None",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "origin": "Customer",
	      "lineItems": [
	        {
	          "id": "31099128-dba8-40a7-bb6c-d12857149ff8",
	          "productId": "d6d0c517-572e-4d26-b80e-ffce825334a4",
	          "name": {
	            "en": "GIRLS HARTBREAK CREW"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "6dc75271-b0e5-4ae9-9158-faa1fff65f7b",
	            "version": 2
	          },
	          "productSlug": {
	            "en": "girls-hartbreak-crew1522841378290"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1522841378290",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "currencyCode": "EUR",
	                  "centAmount": 3400,
	                  "fractionDigits": 2
	                },
	                "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253234387_1.jpg",
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
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            },
	            "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "taxRate": {
	            "name": "19% MwSt",
	            "amount": 0.19,
	            "includedInPrice": true,
	            "country": "DE",
	            "id": "rrsT1Jbw",
	            "subRates": []
	          },
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "60a64e06-b4e8-4205-a0f3-94bc203e2d6d"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "totalPrice": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 3400,
	            "fractionDigits": 2
	          },
	          "taxedPrice": {
	            "totalNet": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 2857,
	              "fractionDigits": 2
	            },
	            "totalGross": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            }
	          },
	          "lineItemMode": "Standard"
	        }
	      ],
	      "customLineItems": [],
	      "transactionFee": true,
	      "discountCodes": [],
	      "lastMessageSequenceNumber": 3,
	      "cart": {
	        "typeId": "cart",
	        "id": "6c97e772-9769-4419-8adc-501c7c5b6088"
	      },
	      "shippingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "billingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "itemShippingAddresses": []
	    },
	    "messagePayloads": [
	      {
	        "email": "user@localhost",
	        "type": "OrderCustomerEmailSet"
	      },
	      {
	        "edit": {
	          "typeId": "order-edit",
	          "id": "df20c5ea-b114-4aab-b330-740b0e9f3099"
	        },
	        "result": {
	          "type": "Applied",
	          "appliedAt": "2018-10-04T15:29:14.091Z",
	          "excerptBeforeEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 1
	          },
	          "excerptAfterEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 3
	          }
	        },
	        "type": "OrderEditApplied"
	      }
	    ],
	    "type": "PreviewSuccess"
	  },
	  "comment": "sample-comment"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order_edit, err := client.OrderEditGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order_edit)
	assert.NotNil(t, order_edit.Version)
	assert.NotEmpty(t, order_edit.LastModifiedAt)
	assert.NotEmpty(t, order_edit.Key)
	assert.NotEmpty(t, order_edit.ID)
	assert.NotEmpty(t, order_edit.CreatedAt)
	assert.NotEmpty(t, order_edit.Comment)

}

func TestGeneratedOrderEditDeleteWithID(t *testing.T) {
	responseData := ` {
	  "id": "df20c5ea-b114-4aab-b330-740b0e9f3099",
	  "version": 1,
	  "resource": {
	    "typeId": "order",
	    "id": "ed454f4e-c43a-485f-a86f-046c691b1363"
	  },
	  "key": "order-edit-key",
	  "createdAt": "2018-10-04T15:22:31.639Z",
	  "lastModifiedAt": "2018-10-04T15:22:31.639Z",
	  "stagedActions": [
	    {
	      "action": "setCustomerEmail",
	      "email": "user@localhost"
	    }
	  ],
	  "result": {
	    "preview": {
	      "type": "Order",
	      "id": "ed454f4e-c43a-485f-a86f-046c691b1363",
	      "version": 3,
	      "customerId": "bf5d96ce-4704-45b2-8842-d409dd34cdfc",
	      "customerEmail": "user@localhost",
	      "createdAt": "2018-05-15T12:40:17.301Z",
	      "lastModifiedAt": "2018-05-15T12:40:17.301Z",
	      "totalPrice": {
	        "type": "centPrecision",
	        "currencyCode": "EUR",
	        "centAmount": 3970,
	        "fractionDigits": 2
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3336,
	          "fractionDigits": 2
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3970,
	          "fractionDigits": 2
	        },
	        "taxPortions": [
	          {
	            "rate": 0.19,
	            "amount": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 634,
	              "fractionDigits": 2
	            },
	            "name": "19% MwSt"
	          }
	        ]
	      },
	      "country": "DE",
	      "orderState": "Open",
	      "syncInfo": [],
	      "returnInfo": [],
	      "refusedGifts": [],
	      "shippingInfo": {
	        "shippingMethodName": "DHL",
	        "price": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 570,
	          "fractionDigits": 2
	        },
	        "shippingRate": {
	          "price": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          },
	          "tiers": []
	        },
	        "taxRate": {
	          "name": "19% MwSt",
	          "amount": 0.19,
	          "includedInPrice": true,
	          "country": "DE",
	          "id": "rrsT1Jbw",
	          "subRates": []
	        },
	        "taxCategory": {
	          "typeId": "tax-category",
	          "id": "fdeb9625-10f8-476c-a549-5d5c6d1bd412"
	        },
	        "deliveries": [],
	        "shippingMethod": {
	          "typeId": "shipping-method",
	          "id": "d18b3f77-92de-4893-b6e3-b5c9c8c1eb96"
	        },
	        "taxedPrice": {
	          "totalNet": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 479,
	            "fractionDigits": 2
	          },
	          "totalGross": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          }
	        },
	        "shippingMethodState": "MatchesCart"
	      },
	      "taxMode": "Platform",
	      "inventoryMode": "None",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "origin": "Customer",
	      "lineItems": [
	        {
	          "id": "31099128-dba8-40a7-bb6c-d12857149ff8",
	          "productId": "d6d0c517-572e-4d26-b80e-ffce825334a4",
	          "name": {
	            "en": "GIRLS HARTBREAK CREW"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "6dc75271-b0e5-4ae9-9158-faa1fff65f7b",
	            "version": 2
	          },
	          "productSlug": {
	            "en": "girls-hartbreak-crew1522841378290"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1522841378290",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "currencyCode": "EUR",
	                  "centAmount": 3400,
	                  "fractionDigits": 2
	                },
	                "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253234387_1.jpg",
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
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            },
	            "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "taxRate": {
	            "name": "19% MwSt",
	            "amount": 0.19,
	            "includedInPrice": true,
	            "country": "DE",
	            "id": "rrsT1Jbw",
	            "subRates": []
	          },
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "60a64e06-b4e8-4205-a0f3-94bc203e2d6d"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "totalPrice": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 3400,
	            "fractionDigits": 2
	          },
	          "taxedPrice": {
	            "totalNet": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 2857,
	              "fractionDigits": 2
	            },
	            "totalGross": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            }
	          },
	          "lineItemMode": "Standard"
	        }
	      ],
	      "customLineItems": [],
	      "transactionFee": true,
	      "discountCodes": [],
	      "lastMessageSequenceNumber": 3,
	      "cart": {
	        "typeId": "cart",
	        "id": "6c97e772-9769-4419-8adc-501c7c5b6088"
	      },
	      "shippingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "billingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "itemShippingAddresses": []
	    },
	    "messagePayloads": [
	      {
	        "email": "user@localhost",
	        "type": "OrderCustomerEmailSet"
	      },
	      {
	        "edit": {
	          "typeId": "order-edit",
	          "id": "df20c5ea-b114-4aab-b330-740b0e9f3099"
	        },
	        "result": {
	          "type": "Applied",
	          "appliedAt": "2018-10-04T15:29:14.091Z",
	          "excerptBeforeEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 1
	          },
	          "excerptAfterEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 3
	          }
	        },
	        "type": "OrderEditApplied"
	      }
	    ],
	    "type": "PreviewSuccess"
	  },
	  "comment": "sample-comment"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order_edit, err := client.OrderEditDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order_edit)
	assert.NotNil(t, order_edit.Version)
	assert.NotEmpty(t, order_edit.LastModifiedAt)
	assert.NotEmpty(t, order_edit.Key)
	assert.NotEmpty(t, order_edit.ID)
	assert.NotEmpty(t, order_edit.CreatedAt)
	assert.NotEmpty(t, order_edit.Comment)

}

func TestGeneratedOrderEditDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "id": "df20c5ea-b114-4aab-b330-740b0e9f3099",
	  "version": 1,
	  "resource": {
	    "typeId": "order",
	    "id": "ed454f4e-c43a-485f-a86f-046c691b1363"
	  },
	  "key": "order-edit-key",
	  "createdAt": "2018-10-04T15:22:31.639Z",
	  "lastModifiedAt": "2018-10-04T15:22:31.639Z",
	  "stagedActions": [
	    {
	      "action": "setCustomerEmail",
	      "email": "user@localhost"
	    }
	  ],
	  "result": {
	    "preview": {
	      "type": "Order",
	      "id": "ed454f4e-c43a-485f-a86f-046c691b1363",
	      "version": 3,
	      "customerId": "bf5d96ce-4704-45b2-8842-d409dd34cdfc",
	      "customerEmail": "user@localhost",
	      "createdAt": "2018-05-15T12:40:17.301Z",
	      "lastModifiedAt": "2018-05-15T12:40:17.301Z",
	      "totalPrice": {
	        "type": "centPrecision",
	        "currencyCode": "EUR",
	        "centAmount": 3970,
	        "fractionDigits": 2
	      },
	      "taxedPrice": {
	        "totalNet": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3336,
	          "fractionDigits": 2
	        },
	        "totalGross": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 3970,
	          "fractionDigits": 2
	        },
	        "taxPortions": [
	          {
	            "rate": 0.19,
	            "amount": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 634,
	              "fractionDigits": 2
	            },
	            "name": "19% MwSt"
	          }
	        ]
	      },
	      "country": "DE",
	      "orderState": "Open",
	      "syncInfo": [],
	      "returnInfo": [],
	      "refusedGifts": [],
	      "shippingInfo": {
	        "shippingMethodName": "DHL",
	        "price": {
	          "type": "centPrecision",
	          "currencyCode": "EUR",
	          "centAmount": 570,
	          "fractionDigits": 2
	        },
	        "shippingRate": {
	          "price": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          },
	          "tiers": []
	        },
	        "taxRate": {
	          "name": "19% MwSt",
	          "amount": 0.19,
	          "includedInPrice": true,
	          "country": "DE",
	          "id": "rrsT1Jbw",
	          "subRates": []
	        },
	        "taxCategory": {
	          "typeId": "tax-category",
	          "id": "fdeb9625-10f8-476c-a549-5d5c6d1bd412"
	        },
	        "deliveries": [],
	        "shippingMethod": {
	          "typeId": "shipping-method",
	          "id": "d18b3f77-92de-4893-b6e3-b5c9c8c1eb96"
	        },
	        "taxedPrice": {
	          "totalNet": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 479,
	            "fractionDigits": 2
	          },
	          "totalGross": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 570,
	            "fractionDigits": 2
	          }
	        },
	        "shippingMethodState": "MatchesCart"
	      },
	      "taxMode": "Platform",
	      "inventoryMode": "None",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "origin": "Customer",
	      "lineItems": [
	        {
	          "id": "31099128-dba8-40a7-bb6c-d12857149ff8",
	          "productId": "d6d0c517-572e-4d26-b80e-ffce825334a4",
	          "name": {
	            "en": "GIRLS HARTBREAK CREW"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "6dc75271-b0e5-4ae9-9158-faa1fff65f7b",
	            "version": 2
	          },
	          "productSlug": {
	            "en": "girls-hartbreak-crew1522841378290"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_GIRLS_HARTBREAK_CREW_variant1_1522841378290",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "currencyCode": "EUR",
	                  "centAmount": 3400,
	                  "fractionDigits": 2
	                },
	                "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253234387_1.jpg",
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
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            },
	            "id": "b5595b13-bbb8-44ab-a2e2-89c18edf1b22"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "taxRate": {
	            "name": "19% MwSt",
	            "amount": 0.19,
	            "includedInPrice": true,
	            "country": "DE",
	            "id": "rrsT1Jbw",
	            "subRates": []
	          },
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "60a64e06-b4e8-4205-a0f3-94bc203e2d6d"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "totalPrice": {
	            "type": "centPrecision",
	            "currencyCode": "EUR",
	            "centAmount": 3400,
	            "fractionDigits": 2
	          },
	          "taxedPrice": {
	            "totalNet": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 2857,
	              "fractionDigits": 2
	            },
	            "totalGross": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3400,
	              "fractionDigits": 2
	            }
	          },
	          "lineItemMode": "Standard"
	        }
	      ],
	      "customLineItems": [],
	      "transactionFee": true,
	      "discountCodes": [],
	      "lastMessageSequenceNumber": 3,
	      "cart": {
	        "typeId": "cart",
	        "id": "6c97e772-9769-4419-8adc-501c7c5b6088"
	      },
	      "shippingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "billingAddress": {
	        "id": "51RkSh-E",
	        "salutation": "Mr",
	        "firstName": "user",
	        "lastName": "lastname",
	        "streetName": "streetname",
	        "streetNumber": "2",
	        "postalCode": "101256",
	        "city": "Berlin",
	        "country": "DE",
	        "building": "34",
	        "pOBox": "12344",
	        "email": "user@example.org"
	      },
	      "itemShippingAddresses": []
	    },
	    "messagePayloads": [
	      {
	        "email": "user@localhost",
	        "type": "OrderCustomerEmailSet"
	      },
	      {
	        "edit": {
	          "typeId": "order-edit",
	          "id": "df20c5ea-b114-4aab-b330-740b0e9f3099"
	        },
	        "result": {
	          "type": "Applied",
	          "appliedAt": "2018-10-04T15:29:14.091Z",
	          "excerptBeforeEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 1
	          },
	          "excerptAfterEdit": {
	            "totalPrice": {
	              "type": "centPrecision",
	              "currencyCode": "EUR",
	              "centAmount": 3970,
	              "fractionDigits": 2
	            },
	            "taxedPrice": {
	              "totalNet": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3336,
	                "fractionDigits": 2
	              },
	              "totalGross": {
	                "type": "centPrecision",
	                "currencyCode": "EUR",
	                "centAmount": 3970,
	                "fractionDigits": 2
	              },
	              "taxPortions": [
	                {
	                  "rate": 0.19,
	                  "amount": {
	                    "type": "centPrecision",
	                    "currencyCode": "EUR",
	                    "centAmount": 634,
	                    "fractionDigits": 2
	                  },
	                  "name": "19% MwSt"
	                }
	              ]
	            },
	            "version": 3
	          }
	        },
	        "type": "OrderEditApplied"
	      }
	    ],
	    "type": "PreviewSuccess"
	  },
	  "comment": "sample-comment"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	order_edit, err := client.OrderEditDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, order_edit)
	assert.NotNil(t, order_edit.Version)
	assert.NotEmpty(t, order_edit.LastModifiedAt)
	assert.NotEmpty(t, order_edit.Key)
	assert.NotEmpty(t, order_edit.ID)
	assert.NotEmpty(t, order_edit.CreatedAt)
	assert.NotEmpty(t, order_edit.Comment)

}

func TestGeneratedOrderEditQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	      "id": "df20c5ea-b114-4aab-b330-740b0e9f3099",
	      "version": 1,
	      "resource": {
	        "typeId": "order",
	        "id": "ed454f4e-c43a-485f-a86f-046c691b1363"
	      },
	      "key": "order-edit-key",
	      "createdAt": "2018-10-04T15:22:31.639Z",
	      "lastModifiedAt": "2018-10-04T15:22:31.639Z",
	      "stagedActions": [
	        {
	          "action": "setCustomerEmail",
	          "email": "user@localhost"
	        }
	      ],
	      "result": {
	        "type": "NotProcessed"
	      },
	      "comment": "sample-comment"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.OrderEditQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
