package ml

// Generated file, please do not change!!!

/**
*	Response format from image search endpoint.
*
 */
type ImageSearchResponse struct {
	// The maximum number of results to return from a query.
	Count int `json:"count"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset float64 `json:"offset"`
	// The total number of product images that were have been analyzed.
	Total int `json:"total"`
	// An array of image URLs of images that are similar to the query image. If no matching images are found, results is empty.
	Results []ResultItem `json:"results"`
}

/**
*	An image URL and the product variants it is contained in. If no matching images are found, ResultItem is not present.
 */
type ResultItem struct {
	// The URL of the image.
	ImageUrl string `json:"imageUrl"`
	// An array of product variants containing the image URL.
	ProductVariants []ProductVariant `json:"productVariants"`
}
