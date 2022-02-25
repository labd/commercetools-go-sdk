package importapi

// Generated file, please do not change!!!

import (
	"time"
)

/**
*	Serves as the entry point of resources.
*
 */
type ImportSink struct {
	// User-defined unique identifier for the ImportSink.
	// Keys can only contain alphanumeric characters (a-Z, 0-9), underscores and hyphens (_, -).
	Key string `json:"key"`
	// The [resource type](#importresourcetype) the ImportSink is able to handle.
	// If not present, the ImportSink is able to import all of the supported [ImportResourceTypes](#importresourcetype).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
	// The version of the ImportSink.
	Version int `json:"version"`
	// The time when the ImportSink was created.
	CreatedAt time.Time `json:"createdAt"`
	// The last time when the ImportSink was modified.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

/**
*	The representation sent to the server when creating an [ImportSink](#importsink).
*
 */
type ImportSinkDraft struct {
	// User-defined unique identifier of the ImportSink.
	// Keys can only contain alphanumeric characters (a-Z, 0-9), underscores and hyphens (_, -).
	Key string `json:"key"`
	// The [resource type](#importresourcetype) to be imported.
	// If not given, the ImportSink is able to import all of the supported [ImportResourceTypes](#importresourcetype).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
}

/**
*	The representation sent to the server when updating an [ImportSink](#importsink).
*
 */
type ImportSinkUpdateDraft struct {
	// Current version of the ImportSink.
	Version int `json:"version"`
	// The [resource type](#importresourcetype) to be imported.
	// If not given, the ImportSink is able to import all of the supported [ImportResourceTypes](#importresourcetype).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) for [ImportSinks](#importsink).
*	Used as a response to a query request for [ImportSinks](#importsink).
*
 */
type ImportSinkPagedResponse struct {
	// The number of results requested in the query request.
	Limit int `json:"limit"`
	// The number of elements skipped, not a page number.
	// Supplied by the client or the server default.
	Offset int `json:"offset"`
	// The actual number of results returned.
	Count int `json:"count"`
	// The total number of results matching the query.
	Total int `json:"total"`
	// The array of Import Sinks matching the query.
	Results []ImportSink `json:"results"`
}
