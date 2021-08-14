// Generated file, please do not change!!!
package ml

/**
*	Response format from image search endpoint.
*
 */
type ImageSearchResponse struct {
	// The maximum number of results to return from a query.
	Count int `json:"count"`
	// The offset into the results matching the query. An offset of 0 is the default value indicating that no results should be skipped.
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
