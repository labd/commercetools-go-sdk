package commercetools_test

import (
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
		ProductType: &commercetools.ProductTypeReference{
			ID: "8750e1fd-f431-481f-9296-967b1e56bf49",
		},
		Slug: &commercetools.LocalizedString{
			"nl": "een-product",
			"en": "some-product",
		},
	}
	product, err := client.Products.Create(draft)
	assert.Nil(t, err)
	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", (*product.MasterData.Current.Description)["en"])
}

func TestGetProductByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	product, err := client.Products.GetByID("foo-bar")
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

	product, err := client.Products.Update(&commercetools.ProductUpdateInput{
		ID:      "1",
		Version: 2,
		PriceSelection: commercetools.ProductPriceSelection{
			Currency: "EUR",
		},
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

	product, err := client.Products.DeleteByID(&commercetools.ProductDeleteInput{
		ID:      "foobar",
		Version: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
}

func TestProductDeleteByKey(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	product, err := client.Products.DeleteByKey(&commercetools.ProductDeleteInput{
		Key:     "foobar",
		Version: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
}
