# commercetools-go-sdk

[![Travis Build Status](https://travis-ci.org/labd/commercetools-go-sdk.svg?branch=master)](https://travis-ci.org/labd/commercetools-go-sdk)
[![codecov](https://codecov.io/gh/LabD/commercetools-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/LabD/commercetools-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/labd/commercetools-go-sdk)](https://goreportcard.com/report/github.com/labd/commercetools-go-sdk)
[![GoDoc](https://godoc.org/github.com/labd/commercetools-go-sdk?status.svg)](https://godoc.org/github.com/labd/commercetools-go-sdk)

The Commercetools Go SDK was created for enabling the creation of the [Terraform Provider for Commercetools](https://github.com/labd/terraform-provider-commercetools). That provider enables you to use infrastructure-as-code principles with Commercetools.

This means that the SDK is not feature complete at the moment. The SDK is currently not meant for building an e-commerce front-ends with it, but aimed at maintaining the configuration of such a site. A front-end can be built using [one of the existing SDK's, provided and maintained by commercetools](https://docs.commercetools.com/software-development-kits).

## Using the SDK


```go
package main

import (
    "log"

    "golang.org/x/oauth2/clientcredentials"
    "github.com/labd/commercetools-go-sdk/commercetools"
    "github.com/labd/commercetools-go-sdk/service/products"
)

func main() {

    oauth2Config := &clientcredentials.Config{
        ClientID:     "<client-id>",
        ClientSecret: "<client-secret>",
        Scopes:       []string{"manage_project:<scopes>"},
        TokenURL:     "https://auth.sphere.io/oauth/token",
    }
    httpClient := oauth2Config.Client(context.TODO())

    // Create the new client. When an empty value is passed it will use the CTP_*
    // environment variables to get the value. The HTTPClient arg is optional,
    // and when empty will automatically be created using the env values.
    client := commercetools.New(&commercetools.Config{
        ProjectKey: "<project-key>",
        URL:        "https://api.sphere.io",
        HTTPClient: httpClient,
    })

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
    if err != nil {
        log.Fatal(err)
    }

    log.Print(product)
}
```

## Service implementation

At the moment the SDK has service coverage for primarily Terraform configuration use-cases, for example declaring API Extensions, Subscriptions and Project Settings. Broader coverage will be implemented if our use-cases require it or via contributions of course!

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
