# commercetools-go-sdk

[![Travis Build Status](https://travis-ci.org/labd/commercetools-go-sdk.svg?branch=master)](https://travis-ci.org/labd/commercetools-go-sdk)
[![codecov](https://codecov.io/gh/LabD/commercetools-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/LabD/commercetools-go-sdk)

Note: This is currently **NOT** ready for production usage

## Using the SDK


```go
package main

import (
  "log"

  "github.com/labd/commercetools-go-sdk/catalog"
  "github.com/labd/commercetools-go-sdk/common"
  "github.com/labd/commercetools-go-sdk/credentials"
)

client, err := common.NewClient(&common.Config{
  ProjectKey:   "<your-project>",
  Region:       "https://api.sphere.io",
  AuthProvider: credentials.NewClientCredentialsProvider(
    "<clientId>", "<clientSecret>", "manage_project:<your-project>"),
})

if err != nil {
  log.Fatal(err)
}

svc := catalog.New(client)
product, err := svc.ProductCreate(&catalog.ProductDraft{
  Key: "test-product",
  Name: common.LocalizedString{
    "nl": "Een test product",
    "en": "A test product",
  },
  ProductType: common.Reference{
    TypeID: "product-type",
    ID:     "8750e1fd-f431-481f-9296-967b1e56bf49",
  },
  Slug: common.LocalizedString{
    "nl": "een-test-product",
    "en": "a-test-product",
  },
}

log.Print(product)

```
