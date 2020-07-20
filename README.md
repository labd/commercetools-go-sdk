# commercetools-go-sdk

[![Build Status](https://github.com/labd/commercetools-go-sdk/workflows/Go%20Tests/badge.svg)](https://github.com/labd/commercetools-go-sdk/workflows/)
[![codecov](https://codecov.io/gh/LabD/commercetools-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/LabD/commercetools-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/labd/commercetools-go-sdk)](https://goreportcard.com/report/github.com/labd/commercetools-go-sdk)
[![GoDoc](https://godoc.org/github.com/labd/commercetools-go-sdk?status.svg)](https://godoc.org/github.com/labd/commercetools-go-sdk)

The Commercetools Go SDK was created for enabling the creation of the
[Terraform Provider for Commercetools](https://github.com/labd/terraform-provider-commercetools).
That provider enables you to use infrastructure-as-code principles with Commercetools.

This means that the SDK is not feature complete at the moment. The SDK is
currently not meant for building e-commerce front-ends with it, but aimed at
maintaining the configuration of such a site. A front-end can be built using
[one of the existing SDK's, provided and maintained by commercetools](https://docs.commercetools.com/software-development-kits). Or
use our [unofficial Python SDK for Commercetools](https://github.com/labd/commercetools-python-sdk)!

## Using the SDK


```go
package main

import (
    "context"
    "fmt"
    "log"
    "math/rand"
    "time"
    
    "github.com/labd/commercetools-go-sdk/commercetools"
)

func main() {

    // Create the new client. When an empty value is passed it will use the CTP_*
    // environment variables to get the value. The HTTPClient arg is optional,
    // and when empty will automatically be created using the env values.
    client, err := commercetools.NewClient(&commercetools.ClientConfig{
        ProjectKey: "<project-key>",
        Endpoints:  commercetools.NewClientEndpoints("europe-west1", "gcp"),
        Credentials: &commercetools.ClientCredentials{
            ClientID:     "<client-id>",
            ClientSecret: "<client-secret>",
            Scopes:       []string{"<scope>"},
        },
    })
    
    ctx := context.Background()

    // Get or Createa product type
    productTypeDraft := commercetools.ProductTypeDraft{
        Name: "a-product-type",
        Key:  "a-product-type",
    }

    productType, err := client.ProductTypeGetWithKey(ctx, productTypeDraft.Key)
    if productType == nil {
        productType, err = client.ProductTypeCreate(ctx, &productTypeDraft)
        if err != nil {
            log.Println(err)
        }
    }

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomID := r.Int()
    productDraft := &commercetools.ProductDraft{
        Key: fmt.Sprintf("test-product-%d", randomID),
        Name: &commercetools.LocalizedString{
            "nl": "Een test product",
            "en": "A test product",
        },
        ProductType: &commercetools.ProductTypeResourceIdentifier{
            ID: productType.ID,
        },
        Slug: &commercetools.LocalizedString{
            "nl": fmt.Sprintf("een-test-product-%d", randomID),
            "en": fmt.Sprintf("a-test-product-%d", randomID),
        },
    }

    product, err := client.ProductCreate(ctx, productDraft)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Print(product)
}
```

## Generating code

To generate code do the following steps:
- `pip install pyyaml`
- install yq (f.e. `brew install yq` on macOS)
- `git clone https://github.com/commercetools/commercetools-api-reference.git` in the folder on the same level as the commercetools-go-sdk folder
- `make generate`


### TODO for code generating the services:

- Currently only /{ID} and /key={key} is supported, support other actions (f.e. Replicate cart)
- Parse and use trait definitions (traits folder)
- CustomObject service implementation
- Implement other HTTPMethods/actions apart from get/post/delete (f.e. /images on product)
- Implement all traits (f.e. priceSelecting)
- Implement reference expansion
