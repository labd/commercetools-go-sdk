package commercetools_test

import (
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

func createExampleProduct() *commercetools.Product {
	return &commercetools.Product{
		ID:             "e7ba4c75-b1bb-483d-94d8-2c4a10f78472",
		Version:        2,
		CreatedAt:      time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		LastModifiedAt: time.Date(1970, 1, 1, 0, 0, 0, 1000000, time.UTC),
		MasterData: &commercetools.ProductCatalogData{
			Current: &commercetools.ProductData{
				Name: &commercetools.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []commercetools.CategoryReference{
					{
						ID: "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
					},
				},
				Description: &commercetools.LocalizedString{
					"en": "Sample description",
				},
				Slug: &commercetools.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: &commercetools.ProductVariant{
					ID:         1,
					SKU:        "sku_MB_PREMIUM_TECH_T_variant1_1369226795424",
					Attributes: []commercetools.Attribute{},
					Images: []commercetools.Image{
						{
							Dimensions: &commercetools.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							URL: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []commercetools.Price{
						{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: &commercetools.Money{
								CentAmount:   10000,
								CurrencyCode: "EUR",
							},
						},
					},
				},
				Variants: []commercetools.ProductVariant{},
			},
			// HasStagedChanged: false,
			Published: true,
			Staged: &commercetools.ProductData{
				Name: &commercetools.LocalizedString{
					"en": "MB PREMIUM TECH T",
				},
				Categories: []commercetools.CategoryReference{
					{
						ID: "cf6d790a-f027-4f46-9a2b-4bc9a31066fb",
					},
				},
				Description: &commercetools.LocalizedString{
					"en": "Sample description",
				},
				Slug: &commercetools.LocalizedString{
					"en": "mb-premium-tech-t1369226795424",
				},
				MasterVariant: &commercetools.ProductVariant{
					ID:         1,
					SKU:        "sku_MB_PREMIUM_TECH_T_variant1_1369226795424",
					Attributes: []commercetools.Attribute{},
					Images: []commercetools.Image{
						{
							Dimensions: &commercetools.ImageDimensions{
								H: 1400,
								W: 1400,
							},
							URL: "https://sphere.io/cli/data/253245821_1.jpg",
						},
					},
					Prices: []commercetools.Price{
						{
							ID: "753472a3-ddff-4e0f-a93b-2eb29c90ba54",
							Value: &commercetools.Money{
								CentAmount:   10000,
								CurrencyCode: "EUR",
							},
						},
					},
				},
				Variants: []commercetools.ProductVariant{},
			},
		},
		ProductType: &commercetools.ProductTypeReference{
			ID: "24f510c3-f334-4099-94e2-d6224a8eb919",
		},
		TaxCategory: &commercetools.TaxCategoryReference{
			ID: "f1e10e3a-45eb-49d8-ad0b-fdf984202f59",
		},
	}
}
