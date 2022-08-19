package platform_test

import (
	"time"

	"github.com/labd/commercetools-go-sdk/ctutils"
	"github.com/labd/commercetools-go-sdk/platform"
)

func createExampleProduct() platform.Product {
	return platform.Product{
		ID:             "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
		Version:        2,
		CreatedAt:      time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		LastModifiedAt: time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		MasterData: platform.ProductCatalogData{
			Current: platform.ProductData{
				Name: platform.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []platform.CategoryReference{
					{
						ID: "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
					},
				},
				Description: &platform.LocalizedString{
					"en": "Sample description",
				},
				Slug: platform.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: platform.ProductVariant{
					ID:         1,
					Sku:        ctutils.StringRef("sku_MB_PREMIUM_TECH_T_variant1_1369226795424"),
					Attributes: []platform.Attribute{},
					Images: []platform.Image{
						{
							Dimensions: platform.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							Url: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []platform.EmbeddedPrice{
						{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: platform.CentPrecisionMoney{
								FractionDigits: 2,
								CentAmount:     10000,
								CurrencyCode:   "EUR",
							},
						},
					},
				},
				Variants: []platform.ProductVariant{},
			},
			// HasStagedChanged: false,
			Published: true,
			Staged: platform.ProductData{
				Name: platform.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []platform.CategoryReference{
					{
						ID: "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
					},
				},
				Description: &platform.LocalizedString{
					"en": "Sample description",
				},
				Slug: platform.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: platform.ProductVariant{
					ID:         1,
					Sku:        ctutils.StringRef("sku_MB_PREMIUM_TECH_T_variant1_1369226795424"),
					Attributes: []platform.Attribute{},
					Images: []platform.Image{
						{
							Dimensions: platform.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							Url: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []platform.EmbeddedPrice{
						{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: platform.CentPrecisionMoney{
								FractionDigits: 2,
								CentAmount:     10000,
								CurrencyCode:   "EUR",
							},
						},
					},
				},
				Variants: []platform.ProductVariant{},
			},
		},
		ProductType: platform.ProductTypeReference{
			ID: "24f510c3-f334-4099-94e2-d6224a8eb919",
		},
		TaxCategory: &platform.TaxCategoryReference{
			ID: "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
		},
	}
}
