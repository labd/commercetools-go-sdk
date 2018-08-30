package products_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/products"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func TestCreateProductNew(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()

	svc := products.New(client)

	draft := &products.ProductDraft{
		Key: "product-test",
		Name: commercetools.LocalizedString{
			"nl": "Een product",
			"en": "Some product",
		},
		ProductType: commercetools.Reference{
			TypeID: "product-type",
			ID:     "8750e1fd-f431-481f-9296-967b1e56bf49",
		},
		Slug: commercetools.LocalizedString{
			"nl": "een-product",
			"en": "some-product",
		},
	}
	product, err := svc.Create(draft)
	assert.Nil(t, err)
	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", product.MasterData.Current.Description["en"])
}

func TestGetProductByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()
	svc := products.New(client)

	product, err := svc.GetByID(100)
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
	svc := products.New(client)

	product, err := svc.Update(&products.UpdateInput{
		ID:      "1",
		Version: 2,
		PriceSelection: products.PriceSelection{
			Currency: "EUR",
		},
		Actions: commercetools.UpdateActions{
			products.AddPrice{
				VariantID: 1,
				Price:     1000,
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", product.MasterData.Current.Description["en"])
}

func TestProductDeleteByID(t *testing.T) {
	client, server := testutil.MockClient(t, testutil.Fixture("product.example.json"), nil, nil)
	defer server.Close()
	svc := products.New(client)

	product, err := svc.DeleteByID(&products.DeleteInput{
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
	svc := products.New(client)

	product, err := svc.DeleteByKey(&products.DeleteInput{
		Key:     "foobar",
		Version: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
}
