// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedProductGetWithID(t *testing.T) {
	responseData := ` {
	    "id": "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
	    "version": 2,
	    "masterData": {
	        "current": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        },
	        "hasStagedChanges": false,
	        "published": true,
	        "staged": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        }
	    },
	    "productType": {
	        "id": "24f510c3-f334-4099-94e2-d6224a8eb919",
	        "typeId": "product-type"
	    },
	    "taxCategory": {
	        "id": "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
	        "typeId": "tax-category"
	    },
	    "createdAt": "1970-01-01T00:00:00.001Z",
	    "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product, err := client.ProductGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product)
	assert.NotNil(t, product.Version)
	assert.NotEmpty(t, product.LastModifiedAt)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)

}

func TestGeneratedProductGetWithKey(t *testing.T) {
	responseData := ` {
	    "id": "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
	    "version": 2,
	    "masterData": {
	        "current": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        },
	        "hasStagedChanges": false,
	        "published": true,
	        "staged": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        }
	    },
	    "productType": {
	        "id": "24f510c3-f334-4099-94e2-d6224a8eb919",
	        "typeId": "product-type"
	    },
	    "taxCategory": {
	        "id": "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
	        "typeId": "tax-category"
	    },
	    "createdAt": "1970-01-01T00:00:00.001Z",
	    "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product, err := client.ProductGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product)
	assert.NotNil(t, product.Version)
	assert.NotEmpty(t, product.LastModifiedAt)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)

}

func TestGeneratedProductDeleteWithID(t *testing.T) {
	responseData := ` {
	    "id": "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
	    "version": 2,
	    "masterData": {
	        "current": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        },
	        "hasStagedChanges": false,
	        "published": true,
	        "staged": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        }
	    },
	    "productType": {
	        "id": "24f510c3-f334-4099-94e2-d6224a8eb919",
	        "typeId": "product-type"
	    },
	    "taxCategory": {
	        "id": "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
	        "typeId": "tax-category"
	    },
	    "createdAt": "1970-01-01T00:00:00.001Z",
	    "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product, err := client.ProductDeleteWithID(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product)
	assert.NotNil(t, product.Version)
	assert.NotEmpty(t, product.LastModifiedAt)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)

}

func TestGeneratedProductDeleteWithKey(t *testing.T) {
	responseData := ` {
	    "id": "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
	    "version": 2,
	    "masterData": {
	        "current": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        },
	        "hasStagedChanges": false,
	        "published": true,
	        "staged": {
	            "categories": [
	                {
	                    "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                    "typeId": "category"
	                }
	            ],
	            "description": {
	                "en": "Sample description"
	            },
	            "masterVariant": {
	                "attributes": [],
	                "id": 1,
	                "images": [
	                    {
	                        "dimensions": {
	                            "h": 1400,
	                            "w": 1400
	                        },
	                        "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                    }
	                ],
	                "prices": [
	                    {
	                        "value": {
	                            "type": "centPrecision",
	                            "fractionDigits": 2,
	                            "centAmount": 10000,
	                            "currencyCode": "EUR"
	                        },
	                        "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                    }
	                ],
	                "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	            },
	            "name": {
	                "en": "MB PREMIUM TECH T"
	            },
	            "slug": {
	                "en": "mb-premium-tech-t1369226795424"
	            },
	            "variants": [],
	            "searchKeywords": {}
	        }
	    },
	    "productType": {
	        "id": "24f510c3-f334-4099-94e2-d6224a8eb919",
	        "typeId": "product-type"
	    },
	    "taxCategory": {
	        "id": "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
	        "typeId": "tax-category"
	    },
	    "createdAt": "1970-01-01T00:00:00.001Z",
	    "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	product, err := client.ProductDeleteWithKey(context.TODO(), "dummy-id", 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, product)
	assert.NotNil(t, product.Version)
	assert.NotEmpty(t, product.LastModifiedAt)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)

}

func TestGeneratedProductQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 1,
	  "total": 1,
	  "results": [
	    {
	    "id": "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
	    "masterData": {
	      "current": {
	        "categories": [
	            {
	                "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                "typeId": "category"
	            }
	        ],
	        "description": {
	            "en": "Sample description"
	        },
	        "masterVariant": {
	            "attributes": [],
	            "id": 1,
	            "images": [
	                {
	                    "dimensions": {
	                        "h": 1400,
	                        "w": 1400
	                    },
	                    "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                }
	            ],
	            "prices": [
	                {
	                    "value": {
	                        "type": "centPrecision",
	                        "fractionDigits": 2,
	                        "centAmount": 10000,
	                        "currencyCode": "EUR"
	                    },
	                    "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                }
	            ],
	            "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	        },
	        "name": {
	            "en": "MB PREMIUM TECH T"
	        },
	        "slug": {
	            "en": "mb-premium-tech-t1369226795424"
	        },
	        "variants": [],
	        "searchKeywords": {}
	      },
	      "hasStagedChanges": false,
	      "published": true,
	      "staged": {
	        "categories": [
	            {
	                "id": "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
	                "typeId": "category"
	            }
	        ],
	        "description": {
	            "en": "Sample description"
	        },
	        "masterVariant": {
	            "attributes": [],
	            "id": 1,
	            "images": [
	                {
	                    "dimensions": {
	                        "h": 1400,
	                        "w": 1400
	                    },
	                    "url": "https://commercetools.com/cli/data/253245821_1.jpg"
	                }
	            ],
	            "prices": [
	                {
	                    "value": {
	                        "type": "centPrecision",
	                        "fractionDigits": 2,
	                        "centAmount": 10000,
	                        "currencyCode": "EUR"
	                    },
	                    "id": "753472a3-ddff-4e0f-a93b-2eb29c90ba54"
	                }
	            ],
	            "sku": "sku_MB_PREMIUM_TECH_T_variant1_1369226795424"
	        },
	        "name": {
	            "en": "MB PREMIUM TECH T"
	        },
	        "slug": {
	            "en": "mb-premium-tech-t1369226795424"
	        },
	        "variants": [],
	        "searchKeywords": {}
	      }
	    },
	    "productType": {
	      "id": "24f510c3-f334-4099-94e2-d6224a8eb919",
	      "typeId": "product-type"
	    },
	    "taxCategory": {
	      "id": "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
	      "typeId": "tax-category"
	    },
	    "version": 2,
	    "createdAt": "1970-01-01T00:00:00.001Z",
	    "lastModifiedAt": "1970-01-01T00:00:00.001Z"
	  }]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.ProductQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
