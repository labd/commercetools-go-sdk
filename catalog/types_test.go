package catalog_test

import (
	"time"

	"github.com/labd/commercetools-go-sdk/catalog"
	"github.com/labd/commercetools-go-sdk/common"
)

func createExampleProduct() *catalog.Product {
	return &catalog.Product{
		ID:             "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
		Version:        2,
		CreatedAt:      time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		LastModifiedAt: time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		MasterData: catalog.ProductCatalogData{
			Current: catalog.ProductData{
				Name: common.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []common.Reference{
					{
						ID:     "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
						TypeID: "category",
					},
				},
				Description: common.LocalizedString{
					"en": "Sample description",
				},
				Slug: common.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: catalog.ProductVariant{
					ID:         1,
					Sku:        "sku_MB_PREMIUM_TECH_T_variant1_1369226795424",
					Attributes: []catalog.Attribute{},
					Images: []catalog.Image{
						catalog.Image{
							Dimensions: catalog.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							URL: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []catalog.Price{
						catalog.Price{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: common.Money{
								Type:           "centPrecision",
								FractionDigits: 2,
								CentAmount:     10000,
								CurrencyCode:   "EUR",
							},
						},
					},
				},
				Variants: []catalog.ProductVariant{},
			},
			HasStagedChanged: false,
			Published:        true,
			Staged: catalog.ProductData{
				Name: common.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []common.Reference{
					{
						ID:     "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
						TypeID: "category",
					},
				},
				Description: common.LocalizedString{
					"en": "Sample description",
				},
				Slug: common.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: catalog.ProductVariant{
					ID:         1,
					Sku:        "sku_MB_PREMIUM_TECH_T_variant1_1369226795424",
					Attributes: []catalog.Attribute{},
					Images: []catalog.Image{
						catalog.Image{
							Dimensions: catalog.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							URL: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []catalog.Price{
						catalog.Price{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: common.Money{
								Type:           "centPrecision",
								FractionDigits: 2,
								CentAmount:     10000,
								CurrencyCode:   "EUR",
							},
						},
					},
				},
				Variants: []catalog.ProductVariant{},
			},
		},
		ProductType: common.Reference{
			ID:     "24f510c3-f334-4099-94e2-d6224a8eb919",
			TypeID: "product-type",
		},
		TaxCategory: common.Reference{
			ID:     "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
			TypeID: "tax-category",
		},
	}
}
