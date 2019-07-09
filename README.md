# commercetools-go-sdk

[![Build Status](https://dev.azure.com/lab-digital/commercetools-go-sdk/_apis/build/status/labd.commercetools-go-sdk?branchName=master)](https://dev.azure.com/lab-digital/commercetools-go-sdk/_build/latest?definitionId=4&branchName=master)
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
    "log"

    "golang.org/x/oauth2/clientcredentials"
    "github.com/labd/commercetools-go-sdk/commercetools"
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
        LibraryName: "some-library",
        LibraryVersion: "v1.0.0",
        ContactURL: "https://example.org",
        ContactEmail: "opensource@example.org",
    })

    product, err := client.Products.Create(&commercetools.ProductDraft{
        Key: "test-product",
        Name: commercetools.LocalizedString{
            "nl": "Een test product",
            "en": "A test product",
        },
        ProductType: commercetools.ProductTypeReference{
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

## Generating code

To generate code do the following steps:
- `pip install pyyaml`
- install yq (f.e. `brew install yq` on macOS)
- `git clone https://github.com/commercetools/commercetools-api-reference.git` in the folder on the same level as the commercetools-go-sdk folder
- `make generate`

## Service implementation

At the moment the SDK has service coverage for primarily Terraform configuration use-cases, for example declaring API Extensions, Subscriptions and Project Settings. Broader coverage will be implemented if our use-cases require it or via contributions of course!

### TODO for code generating the services:

- Currently only /{ID} and /key={key} is supported, support other actions (f.e. Replicate cart)
- CustomObject service implementation
- Implement other HTTPMethods/actions apart from get/post/delete (f.e. /images on product)
- Implement all traits (f.e. priceSelecting)
- Implement reference expansion
