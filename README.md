# commercetools-go-sdk

[![Build Status](https://github.com/labd/commercetools-go-sdk/workflows/Go%20Tests/badge.svg)](https://github.com/labd/commercetools-go-sdk/workflows/)
[![codecov](https://codecov.io/gh/LabD/commercetools-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/LabD/commercetools-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/labd/commercetools-go-sdk)](https://goreportcard.com/report/github.com/labd/commercetools-go-sdk)
[![GoDoc](https://godoc.org/github.com/labd/commercetools-go-sdk?status.svg)](https://godoc.org/github.com/labd/commercetools-go-sdk)

The Commercetools Go SDK is automatically generated based on the official [API specifications](https://github.com/commercetools/commercetools-api-reference) 
of Commercetools. It should therefore be nearly feature complete.

The SDK was initially created for enabling the creation of the
[Terraform Provider for Commercetools](https://github.com/labd/terraform-provider-commercetools) 
That provider enables you to use infrastructure-as-code principles with Commercetools.


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

    // The last argument is optional for reference expansion
    product, err := client.ProductCreate(
        ctx, productDrafti, commercetools.WithReferenceExpansion("taxCategory"))
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
- CustomObject service implementation
- Implement all traits (f.e. priceSelecting)
