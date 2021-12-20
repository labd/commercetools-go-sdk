// Generated file, please do not change!!!
package ml

type ProjectCategoryRecommendation struct {
	// A category that is recommended for a product.
	Category CategoryReference `json:"category"`
	// Probability score for the category recommendation.
	Confidence float64 `json:"confidence"`
	// Breadcrumb path to the recommended category. This only picks up one language, not all available languages for the category. English is prioritized, but if English data is not available, an arbitrary language is selected. Do not use this to identify a category,use the category ID from the category reference instead.
	Path string `json:"path"`
}

type ProjectCategoryRecommendationMeta struct {
	// The product name that was used to generate recommendations.
	ProductName *string `json:"productName,omitempty"`
	// The product image that was used to generate recommendations.
	ProductImageUrl *string `json:"productImageUrl,omitempty"`
	// Top 5 general categories that were used internally to generate the project-specific categories. These category names are not related to the categories defined in the project, but they provide additional information to understand the project-specific categories in the results section.
	GeneralCategoryNames []string `json:"generalCategoryNames"`
}

type ProjectCategoryRecommendationPagedQueryResponse struct {
	Count   int                               `json:"count"`
	Total   int                               `json:"total"`
	Offset  int                               `json:"offset"`
	Results []ProjectCategoryRecommendation   `json:"results"`
	Meta    ProjectCategoryRecommendationMeta `json:"meta"`
}
