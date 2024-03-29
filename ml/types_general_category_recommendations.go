package ml

// Generated file, please do not change!!!

type GeneralCategoryRecommendation struct {
	// An English category name that is recommended for a product.
	CategoryName string `json:"categoryName"`
	// Probability score for the category recommendation.
	Confidence float64 `json:"confidence"`
}

type GeneralCategoryRecommendationPagedQueryResponse struct {
	Count int `json:"count"`
	Total int `json:"total"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int                             `json:"offset"`
	Results []GeneralCategoryRecommendation `json:"results"`
}
