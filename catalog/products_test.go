package catalog_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labd/commercetools-go-sdk/catalog"
	"github.com/labd/commercetools-go-sdk/common"
	"github.com/labd/commercetools-go-sdk/credentials"
)

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestCreateProduct(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("product.example.json"))

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	client, err := common.NewClient(&common.Config{
		ProjectKey:   "unittest",
		Region:       ts.URL,
		AuthProvider: credentials.NewDummyCredentialsProvider("Bearer unittest"),
	})
	if err != nil {
		t.Fatal(err)
	}
	svc := catalog.New(client)

	draft := &catalog.ProductDraft{
		Key: "product-test",
		Name: common.LocalizedString{
			"nl": "Een product",
			"en": "Some product",
		},
		ProductType: common.Reference{
			TypeID: "product-type",
			ID:     "8750e1fd-f431-481f-9296-967b1e56bf49",
		},
		Slug: common.LocalizedString{
			"nl": "een-product",
			"en": "some-product",
		},
	}
	product, err := svc.ProductCreate(draft)

	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", product.MasterData.Current.Description["en"])
}

func TestGetProductByID(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("product.example.json"))
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	client, err := common.NewClient(&common.Config{
		ProjectKey:   "unittest",
		Region:       ts.URL,
		AuthProvider: credentials.NewDummyCredentialsProvider("Bearer unittest"),
	})
	if err != nil {
		t.Fatal(err)
	}
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
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("product.example.json"))

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	client, err := common.NewClient(&common.Config{
		ProjectKey:   "unittest",
		Region:       ts.URL,
		AuthProvider: credentials.NewDummyCredentialsProvider("Bearer unittest"),
	})
	if err != nil {
		t.Fatal(err)
	}
	svc := catalog.New(client)

	product, err := svc.ProductUpdate(&catalog.ProductUpdateInput{
		ID:      "1",
		Version: 2,
		PriceSelection: catalog.PriceSelection{
			Currency: "EUR",
		},
		Actions: common.UpdateActions{
			catalog.ProductAddPrice{
				VariantID: 1,
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, product.Version)
	assert.Equal(t, "Sample description", product.MasterData.Current.Description["en"])
}
