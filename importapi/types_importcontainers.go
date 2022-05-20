package importapi

// Generated file, please do not change!!!

import (
	"time"
)

/**
*	Serves as the entry point of resources.
*	An Import Container is not resource type-specific.
*
 */
type ImportContainer struct {
	// User-defined unique identifier for the ImportContainer.
	// Keys can only contain alphanumeric characters (a-Z, 0-9), underscores and hyphens (_, -).
	Key string `json:"key"`
	// The [resource type](#importresourcetype) the ImportContainer is able to handle.
	// If not present, the ImportContainer is able to import all of the supported [ImportResourceTypes](#importresourcetype).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
	// The version of the ImportContainer.
	Version int `json:"version"`
	// The time when the ImportContainer was created.
	CreatedAt time.Time `json:"createdAt"`
	// The last time when the ImportContainer was modified.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

/**
*	The representation sent to the server when creating an [ImportContainer](#importcontainer).
*
 */
type ImportContainerDraft struct {
	// User-defined unique identifier of the ImportContainer.
	// Keys can only contain alphanumeric characters (a-Z, 0-9), underscores and hyphens (_, -).
	Key string `json:"key"`
	// The [resource type](#importresourcetype) to be imported.
	// If not given, the ImportContainer is able to import all of the supported [ImportResourceTypes](#importresourcetype).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
}

/**
*	The representation sent to the server when updating an import container.
*
 */
type ImportContainerUpdateDraft struct {
	// Current version of the ImportContainer.
	Version int `json:"version"`
	// The [resource type](#importresourcetype) to be imported.
	// If not given, the ImportContainer is able to import all of the supported [ImportResourceTypes](#importresourcetype).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) for [ImportContainers](#importcontainer).
*	Used as a response to a query request for [ImportContainers](#importcontainer).
*
 */
type ImportContainerPagedResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// The actual number of results returned.
	Count int `json:"count"`
	// The total number of results matching the query.
	Total int `json:"total"`
	// The array of Import Containers matching the query.
	Results []ImportContainer `json:"results"`
}
