# commercetools-go-sdk

[![Travis Build Status](https://travis-ci.org/labd/commercetools-go-sdk.svg?branch=master)](https://travis-ci.org/labd/commercetools-go-sdk)
[![codecov](https://codecov.io/gh/LabD/commercetools-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/LabD/commercetools-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/labd/commercetools-go-sdk)](https://goreportcard.com/report/github.com/labd/commercetools-go-sdk)

Note: This is currently **NOT** ready for production usage

## Using the SDK


```go
package main

import (
  "log"

  "github.com/labd/commercetools-go-sdk/commercetools"
  "github.com/labd/commercetools-go-sdk/commercetools/credentials"
  "github.com/labd/commercetools-go-sdk/service/products"
)

client, err := commercetools.NewClient(&commercetools.Config{
  ProjectKey:   "<your-project>",
  Region:       "https://api.sphere.io",
  AuthProvider: credentials.NewClientCredentialsProvider(
    "<clientId>", "<clientSecret>", "manage_project:<your-project>"),
})

if err != nil {
  log.Fatal(err)
}

svc := products.New(client)
product, err := svc.ProductCreate(&products.ProductDraft{
  Key: "test-product",
  Name: commercetools.LocalizedString{
    "nl": "Een test product",
    "en": "A test product",
  },
  ProductType: commercetools.Reference{
    TypeID: "product-type",
    ID:     "8750e1fd-f431-481f-9296-967b1e56bf49",
  },
  Slug: commercetools.LocalizedString{
    "nl": "een-test-product",
    "en": "a-test-product",
  },
}

log.Print(product)

```

## Service implementation

### Product Catalog

 - [ ] Categories
 - [x] Products
 - [ ] Product Projections
 - [ ] Product Projections Search
 - [ ] Product Suggestions
 - [ ] Inventory
 - [ ] Reviews

### Pricing & Discounts

 - [x] Tax Categories
 - [ ] Product Discounts
 - [ ] Cart Discounts
 - [ ] Discount Codes

### Carts, Orders & Shopping Lists

 - [ ] Shopping Lists beta
 - [ ] My Shopping Lists beta
 - [ ] Carts
 - [ ] My Carts beta
 - [ ] Shipping Methods
 - [ ] Shipping Zones
 - [ ] Payments
 - [ ] My Payments beta
 - [ ] Orders
 - [ ] My Orders beta
 - [ ] Order Import

### Customers

 - [ ] Customers
 - [ ] My Customer Profile beta
 - [ ] Customer Groups

### Configuration

 - [x] Project Settings
 - [x] Channels
 - [x] State Machines

### Customize Data

 - [x] Product Types
 - [x] Types
 - [x] Custom Fields
 - [ ] Custom Objects

### Customize Behavior

 - [x] API Extensions beta
 - [x] Subscriptions beta
 - [ ] Messages Query
