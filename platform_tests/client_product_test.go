package platform_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestCreateProductNew(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json", 201), nil, nil)
	defer server.Close()

	draft := platform.ProductDraft{
		Key: ctutils.StringRef("product-test"),
		Name: platform.LocalizedString{
			"nl": "Een product",
			"en": "Some product",
		},
		ProductType: platform.ProductTypeResourceIdentifier{
			ID: ctutils.StringRef("8750e1fd-f431-481f-9296-967b1e56bf49"),
		},
		Slug: platform.LocalizedString{
			"nl": "een-product",
			"en": "some-product",
		},
	}
	product, err := client.WithProjectKey("unittest").Products().Post(draft).Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", (*product.MasterData.Current.Description)["en"])
}

func TestGetProductByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json", 200), nil, nil)
	defer server.Close()

	product, err := client.WithProjectKey("unittest").Products().WithId("foo-bar").Get().Execute(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, *product)
	assert.Equal(t, *product, expected)
}

func TestProductUpdate(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json", 200), nil, nil)
	defer server.Close()

	input := platform.ProductUpdate{
		Version: 2,
		Actions: []platform.ProductUpdateAction{
			platform.ProductAddPriceAction{
				VariantId: ctutils.IntRef(1),
				Price:     platform.PriceDraft{Value: platform.Money{CentAmount: 1000}},
			},
		},
	}
	product, err := client.WithProjectKey("unittest").Products().WithId("1").Post(input).Execute(context.TODO())

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", (*product.MasterData.Current.Description)["en"])
}

func TestProductDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json", 200), nil, nil)
	defer server.Close()

	product, err := client.WithProjectKey("unittest").Products().WithId("foobar").Delete().WithQueryParams(platform.ByProjectKeyProductsByIDRequestMethodDeleteInput{
		Version: 2,
	}).Execute(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, *product)
}

func TestProductDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json", 200), nil, nil)
	defer server.Close()

	product, err := client.
		WithProjectKey("unittest").
		Products().
		WithKey("foobar").
		Delete().
		WithQueryParams(platform.ByProjectKeyProductsKeyByKeyRequestMethodDeleteInput{
			Version: 2,
		}).
		Execute(context.TODO())

	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, *product)
}

func TestProductQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, testutil.ResponseData{StatusCode: 200, Body: "{}"}, &output, nil)
	defer server.Close()

	_, err := client.
		WithProjectKey("unittest").
		Products().
		Get().
		WithQueryParams(platform.ByProjectKeyProductsRequestMethodGetInput{
			Limit: ctutils.IntRef(500),
		}).
		Execute(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
