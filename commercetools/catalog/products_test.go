package catalog_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/commercetools/catalog"
	"github.com/labd/commercetools-go-sdk/testutil"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestCreateProductNew(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("product.example.json"), nil, nil)
	defer server.Close()

	svc := catalog.New(client)

	draft := &catalog.ProductDraft{
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
	product, err := svc.ProductCreate(draft)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", product.MasterData.Current.Description["en"])
}

func TestGetProductByID(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("product.example.json"), nil, nil)
	defer server.Close()
	svc := catalog.New(client)

	product, err := svc.ProductGetByID(100)
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
	assert.Equal(t, product, expected)
}

func TestProductUpdate(t *testing.T) {
	client, server := testutil.MockClient(t, fixture("product.example.json"), nil, nil)
	defer server.Close()
	svc := catalog.New(client)

	product, err := svc.ProductUpdate(&catalog.ProductUpdateInput{
		ID:      "1",
		Version: 2,
		PriceSelection: catalog.PriceSelection{
			Currency: "EUR",
		},
		Actions: commercetools.UpdateActions{
			catalog.ProductAddPrice{
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
	client, server := testutil.MockClient(t, fixture("product.example.json"), nil, nil)
	defer server.Close()
	svc := catalog.New(client)

	product, err := svc.ProductDeleteByID(&catalog.ProductDeleteInput{
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
	client, server := testutil.MockClient(t, fixture("product.example.json"), nil, nil)
	defer server.Close()
	svc := catalog.New(client)

	product, err := svc.ProductDeleteByKey(&catalog.ProductDeleteInput{
		Key:     "foobar",
		Version: 2,
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := createExampleProduct()
	assert.Equal(t, expected, product)
}
