package commercetools_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestCreateProductNew(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	draft := &commercetools.ProductDraft{
		Key: "product-test",
		Name: &commercetools.LocalizedString{
			"nl": "Een product",
			"en": "Some product",
		},
		ProductType: &commercetools.ProductTypeResourceIdentifier{
			ID: "8750e1fd-f431-481f-9296-967b1e56bf49",
		},
		Slug: &commercetools.LocalizedString{
			"nl": "een-product",
			"en": "some-product",
		},
	}
	product, err := client.ProductCreate(draft)
	assert.Nil(t, err)
	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", (*product.MasterData.Current.Description)["en"])
}

func TestGetProductByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	product, err := client.ProductGetWithId("foo-bar")
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
	assert.Equal(t, product, expected)
}

func TestProductUpdate(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	product, err := client.ProductUpdateWithId(&commercetools.ProductUpdateWithIdInput{
		ID:      "1",
		Version: 2,
		/*
		TODO: trait priceSelecting should add this...
		PriceSelection: commercetools.ProductPriceSelection{
			Currency: "EUR",
		},
		*/
		Actions: []commercetools.ProductUpdateAction{
			commercetools.ProductAddPriceAction{
				VariantID: 1,
				Price:     &commercetools.PriceDraft{Value: &commercetools.Money{CentAmount: 1000}},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", (*product.MasterData.Current.Description)["en"])
}

func TestProductDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	product, err := client.ProductDeleteWithId("foobar", 2)
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
}

func TestProductDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	product, err := client.ProductDeleteWithKey("foobar", 2)
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
}

func TestProductQuery(t *testing.T) {
	output := testutil.RequestData{}
	client, server := testutil.MockClient(t, "{}", &output, nil)
	defer server.Close()

	queryInput := commercetools.QueryInput{
		Limit: 500,
	}
	_, err := client.ProductQuery(&queryInput)
	assert.Nil(t, err)

	assert.Equal(t, url.Values{"limit": []string{"500"}}, output.URL.Query())
}
