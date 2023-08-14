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

Note that since this SDK is automatically generated we cannot guarantee backwards
compatibility between releases. Please pin the dependency correctly and be aware
of potential changes when updating

## Using the SDK


```go
package main

import (
   "context"
   "errors"
   "fmt"
   "github.com/labd/commercetools-go-sdk/ctutils"
   "log"
   "math/rand"
   "time"

   "github.com/davecgh/go-spew/spew"
   "github.com/labd/commercetools-go-sdk/platform"
   "golang.org/x/oauth2/clientcredentials"
)

func main() {

   // Create the new client. When an empty value is passed it will use the CTP_*
   // environment variables to get the value. The HTTPClient arg is optional,
   // and when empty will automatically be created using the env values.
   client, err := platform.NewClient(&platform.ClientConfig{
      URL: "https://api.europe-west1.gcp.commercetools.com",
      Credentials: &clientcredentials.Config{
         TokenURL:     "https://auth.europe-west1.gcp.commercetools.com/oauth/token",
         ClientID:     "<client-id>",
         ClientSecret: "<client-secret>",
         Scopes:       []string{"manage_project:<project-key>"},
      },
   })

   projectClient := client.WithProjectKey("<project-key>")

   ctx := context.Background()

   // Get or Createa product type
   productTypeDraft := platform.ProductTypeDraft{
      Name: "a-product-type",
      Key:  ctutils.StringRef("a-product-type"),
   }

   productType, err := projectClient.
      ProductTypes().
      WithKey(*productTypeDraft.Key).
      Get().
      Execute(ctx)

   if err != nil {
      if errors.As(err, &platform.ErrNotFound) {
         productType, err = projectClient.
            ProductTypes().
            Post(productTypeDraft).
            Execute(ctx)
      }

      if err != nil {
         log.Fatal(err)
      }
   }

   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   randomID := r.Int()
   productDraft := platform.ProductDraft{
      Key: ctutils.StringRef(fmt.Sprintf("test-product-%d", randomID)),
      Name: platform.LocalizedString{
         "nl": "Een test product",
         "en": "A test product",
      },
      ProductType: platform.ProductTypeResourceIdentifier{
         ID: ctutils.StringRef(productType.ID),
      },
      Slug: platform.LocalizedString{
         "nl": fmt.Sprintf("een-test-product-%d", randomID),
         "en": fmt.Sprintf("a-test-product-%d", randomID),
      },
   }

   // The last argument is optional for reference expansion
   product, err := projectClient.
      Products().
      Post(productDraft).
           WithQueryParams(
              platform.ByProjectKeyProductsRequestMethodPostInput{
                 Expand: []string{"foobar"},
              },
           ).
      Execute(ctx)

   // Alternatively you can pass query params via methods
   //product, err := projectClient.
   //	Products().
   //	Post(productDraft).
   //	Expand([]string{"foobar"}).
   //	Execute(ctx)

   if err != nil {
      log.Fatal(err)
   }

   spew.Dump(product)
}

```

## Generating code

To re-generate the API based on new RAML files take the following steps:
 - Install rmf-codegen (see https://github.com/commercetools/rmf-codegen)
 - Clone the API specifications https://github.com/commercetools/commercetools-api-reference
   in the parent directory
 - Run `make generate`
