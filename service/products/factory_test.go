package products_test

import (
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/products"
)

func createExampleProduct() *products.Product {
	return &products.Product{
		ID:             "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
		Version:        2,
		CreatedAt:      time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		LastModifiedAt: time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		MasterData: products.ProductCatalogData{
			Current: products.ProductData{
				Name: commercetools.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []commercetools.Reference{
					{
						ID:     "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
						TypeID: "category",
					},
				},
				Description: commercetools.LocalizedString{
					"en": "Sample description",
				},
				Slug: commercetools.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: products.ProductVariant{
					ID:         1,
					Sku:        "sku_MB_PREMIUM_TECH_T_variant1_1369226795424",
					Attributes: []products.Attribute{},
					Images: []products.Image{
						products.Image{
							Dimensions: products.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							URL: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []products.Price{
						products.Price{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: commercetools.Money{
								Type:           "centPrecision",
								FractionDigits: 2,
								CentAmount:     10000,
								CurrencyCode:   "EUR",
							},
						},
					},
				},
				Variants: []products.ProductVariant{},
			},
			HasStagedChanged: false,
			Published:        true,
			Staged: products.ProductData{
				Name: commercetools.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []commercetools.Reference{
					{
						ID:     "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
						TypeID: "category",
					},
				},
				Description: commercetools.LocalizedString{
					"en": "Sample description",
				},
				Slug: commercetools.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: products.ProductVariant{
					ID:         1,
					Sku:        "sku_MB_PREMIUM_TECH_T_variant1_1369226795424",
					Attributes: []products.Attribute{},
					Images: []products.Image{
						products.Image{
							Dimensions: products.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							URL: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []products.Price{
						products.Price{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: commercetools.Money{
								Type:           "centPrecision",
								FractionDigits: 2,
								CentAmount:     10000,
								CurrencyCode:   "EUR",
							},
						},
					},
				},
				Variants: []products.ProductVariant{},
			},
		},
		ProductType: commercetools.Reference{
			ID:     "24f510c3-f334-4099-94e2-d6224a8eb919",
			TypeID: "product-type",
		},
		TaxCategory: commercetools.Reference{
			ID:     "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
			TypeID: "tax-category",
		},
	}
}
