package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	A single ProductTailoring representation contains the _current_ and the _staged_ representation of its product data tailored per Store.
*
 */
type ProductTailoring struct {
	// Unique identifier of the ProductTailoring.
	ID string `json:"id"`
	// Current version of the ProductTailoring.
	Version int `json:"version"`
	// Date and time (UTC) the ProductTailoring was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ProductTailoring was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the ProductTailoring.
	Key *string `json:"key,omitempty"`
	// The Store to which the ProductTailoring belongs.
	Store StoreKeyReference `json:"store"`
	// Reference to the Product the ProductTailoring belongs to.
	Product ProductReference `json:"product"`
	// `true` if the ProductTailoring is published.
	Published bool `json:"published"`
	// Current (published) data of the ProductTailoring.
	Current ProductTailoringData `json:"current"`
	// Staged (unpublished) data of the ProductTailoring.
	Staged ProductTailoringData `json:"staged"`
	// `true` if the `staged` data is different from the `current` data.
	HasStagedChanges bool `json:"hasStagedChanges"`
}

/**
*	Contains all the tailored data of a Product.
*
 */
type ProductTailoringData struct {
	// Tailored name of the Product.
	Name *LocalizedString `json:"name,omitempty"`
	// Tailored description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Tailored title of the Product used by external search engines for improved search engine performance.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Tailored description of the Product used by external search engines for improved search engine performance.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Tailored keywords related to the Product used by external search engines for improved search engine performance.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// User-defined identifier used in a deep-link URL for the ProductTailoring.
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
}

/**
*	Contains all the tailored data of a Product.
*
 */
type ProductTailoringDraft struct {
	// User-defined unique identifier of the ProductTailoring.
	Key *string `json:"key,omitempty"`
	// The Store to which the ProductTailoring belongs.
	Store StoreResourceIdentifier `json:"store"`
	// ResourceIdentifier of the Product the ProductTailoring belongs to.
	Product ProductResourceIdentifier `json:"product"`
	// Tailored name of the Product.
	Name *LocalizedString `json:"name,omitempty"`
	// Tailored description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Tailored title of the Product used by external search engines for improved search engine performance.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Tailored description of the Product used by external search engines for improved search engine performance.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Tailored keywords related to the Product used by external search engines for improved search engine performance.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// User-defined identifier used in a deep-link URL for the ProductTailoring.
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// If `true`, the ProductTailoring is published immediately.
	Publish *bool `json:"publish,omitempty"`
}

/**
*	Contains all the tailored data of a Product for a specific Store.
*
 */
type ProductTailoringInStoreDraft struct {
	// User-defined unique identifier of the ProductTailoring.
	Key *string `json:"key,omitempty"`
	// ResourceIdentifier of the Product the ProductTailoring belongs to.
	Product ProductResourceIdentifier `json:"product"`
	// Tailored name of the Product.
	Name *LocalizedString `json:"name,omitempty"`
	// Tailored description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Tailored title of the Product used by external search engines for improved search engine performance.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Tailored description of the Product used by external search engines for improved search engine performance.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Tailored keywords related to the Product used by external search engines for improved search engine performance.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// User-defined identifier used in a deep-link URL for the ProductTailoring.
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// If `true`, the ProductTailoring is published immediately.
	Publish *bool `json:"publish,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ProductTailoring](ctp:api:type:ProductTailoring).
*
 */
type ProductTailoringPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [ProductTailoring](ctp:api:type:ProductTailoring) list matching the query.
	Results []ProductTailoring `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductTailoring](ctp:api:type:ProductTailoring).
*
 */
type ProductTailoringReference struct {
	// Unique identifier of the referenced [ProductTailoring](ctp:api:type:ProductTailoring).
	ID string `json:"id"`
	// Contains the representation of the expanded ProductTailoring. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ProductTailoring.
	Obj *ProductTailoring `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-tailoring", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ProductTailoring](ctp:api:type:ProductTailoring).
*
 */
type ProductTailoringResourceIdentifier struct {
	// Unique identifier of the referenced [ProductTailoring](ctp:api:type:ProductTailoring). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ProductTailoring](ctp:api:type:ProductTailoring). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-tailoring", Alias: (*Alias)(&obj)})
}

type ProductTailoringUpdateAction interface{}

func mapDiscriminatorProductTailoringUpdateAction(input interface{}) (ProductTailoringUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "publish":
		obj := ProductTailoringPublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ProductTailoringSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaAttributes":
		obj := ProductTailoringSetMetaAttributesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaDescription":
		obj := ProductTailoringSetMetaDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaKeywords":
		obj := ProductTailoringSetMetaKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaTitle":
		obj := ProductTailoringSetMetaTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := ProductTailoringSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSlug":
		obj := ProductTailoringSetSlugAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unpublish":
		obj := ProductTailoringUnpublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Publishes the `staged` data of the ProductTailoring to `current`. Sets `hasStagedChanges` to `false`.
*	Generates the [ProductTailoringPublished](ctp:api:type:ProductTailoringPublishedMessage) Message.
*
 */
type ProductTailoringPublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

/**
*	Generates the [ProductTailoringDescriptionSet](ctp:api:type:ProductTailoringDescriptionSetMessage) Message.
*
 */
type ProductTailoringSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
	// If `true`, only the staged `description` is updated. If `false`, both the current and staged `description` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

/**
*	Updates all meta attributes at the same time.
 */
type ProductTailoringSetMetaAttributesAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// If `true`, only the staged attributes are updated. If `false`, both the current and staged attributes are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaAttributesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaAttributesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaAttributes", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetMetaDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// If `true`, only the staged `metaDescription` is updated. If `false`, both the current and staged `metaDescription` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetMetaKeywordsAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// If `true`, only the staged `metaKeywords` is updated. If `false`, both the current and staged `metaKeywords` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetMetaTitleAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// If `true`, only the staged `metaTitle` is updated. If `false`, both the current and staged `metaTitle` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

/**
*	Generates the [ProductTailoringNameSet](ctp:api:type:ProductTailoringNameSetMessage) Message.
*
 */
type ProductTailoringSetNameAction struct {
	// Value to set. If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
	// If `true`, only the staged `name` is updated. If `false`, both the current and staged `name` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Generates the [ProductTailoringSlugSet](ctp:api:type:ProductTailoringSlugSetMessage) Message.
*
 */
type ProductTailoringSetSlugAction struct {
	// Value to set. If empty, any existing value will be removed.
	Slug *LocalizedString `json:"slug,omitempty"`
	// If `true`, only the staged `slug` is updated. If `false`, both the current and staged `slug` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSlug", Alias: (*Alias)(&obj)})
}

/**
*	Unpublishes the `current` data of the ProductTailoring. Sets the `published` field to `false`.
*	Generates the [ProductTailoringUnpublished](ctp:api:type:ProductTailoringUnpublishedMessage) Message.
*
 */
type ProductTailoringUnpublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}
