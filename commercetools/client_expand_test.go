package commercetools_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
)

func TestClientCreateWithReferenceExpansion(t *testing.T) {
	responseData := "{}"
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, responseData, &output, nil)
	defer server.Close()

	productDraft := &commercetools.ProductDraft{
		Key: "test-product",
		Name: &commercetools.LocalizedString{
			"nl": "Een test product",
			"en": "A test product",
		},
		ProductType: &commercetools.ProductTypeResourceIdentifier{
			ID: "my-product-type-id",
		},
		Slug: &commercetools.LocalizedString{
			"nl": "een-test-product",
			"en": "a-test-product",
		},
	}

	product, err := client.ProductCreate(
		context.TODO(), productDraft,
		commercetools.WithReferenceExpansion("productType"))
	assert.NotNil(t, product)
	assert.NoError(t, err)
	assert.Equal(t, url.Values(url.Values{"expand": []string{"productType"}}), output.URL.Query())
}


func TestClientGetWithReferenceExpansion(t *testing.T) {
	responseData := "{}"
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, responseData, &output, nil)
	defer server.Close()

	product, err := client.ProductGetWithID(
		context.TODO(), "my-id",
		commercetools.WithReferenceExpansion("productType"))
	assert.NotNil(t, product)
	assert.NoError(t, err)
	assert.Equal(t, url.Values(url.Values{"expand": []string{"productType"}}), output.URL.Query())
}


func TestClientUpdateWithReferenceExpansion(t *testing.T) {
	responseData := "{}"
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, responseData, &output, nil)
	defer server.Close()

	product, err := client.ProductUpdateWithID(
		context.TODO(), &commercetools.ProductUpdateWithIDInput{
			ID: "foobar",
			Version: 10,
		},
		commercetools.WithReferenceExpansion("productType", "taxCategory"))
	assert.NotNil(t, product)
	assert.NoError(t, err)
	assert.Equal(t, url.Values(url.Values{"expand": []string{"productType", "taxCategory"}}), output.URL.Query())
}


func TestClientDeleteWithReferenceExpansion(t *testing.T) {
	responseData := "{}"
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, responseData, &output, nil)
	defer server.Close()

	product, err := client.ProductDeleteWithID(
		context.TODO(), "my-id", 10,
		commercetools.WithReferenceExpansion("productType"))
	assert.NotNil(t, product)
	assert.NoError(t, err)
	assert.Equal(t, url.Values(url.Values{"expand": []string{"productType"}, "version":[]string{"10"}}), output.URL.Query())
}
