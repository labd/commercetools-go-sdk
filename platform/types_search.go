package platform

// Generated file, please do not change!!!

import (
	"time"
)

/**
*	Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
 */
type SearchFieldType string

const (
	SearchFieldTypeBoolean      SearchFieldType = "boolean"
	SearchFieldTypeText         SearchFieldType = "text"
	SearchFieldTypeLtext        SearchFieldType = "ltext"
	SearchFieldTypeEnum         SearchFieldType = "enum"
	SearchFieldTypeLenum        SearchFieldType = "lenum"
	SearchFieldTypeNumber       SearchFieldType = "number"
	SearchFieldTypeMoney        SearchFieldType = "money"
	SearchFieldTypeDate         SearchFieldType = "date"
	SearchFieldTypeDatetime     SearchFieldType = "datetime"
	SearchFieldTypeTime         SearchFieldType = "time"
	SearchFieldTypeReference    SearchFieldType = "reference"
	SearchFieldTypeSetBoolean   SearchFieldType = "set_boolean"
	SearchFieldTypeSetText      SearchFieldType = "set_text"
	SearchFieldTypeSetLtext     SearchFieldType = "set_ltext"
	SearchFieldTypeSetEnum      SearchFieldType = "set_enum"
	SearchFieldTypeSetLenum     SearchFieldType = "set_lenum"
	SearchFieldTypeSetNumber    SearchFieldType = "set_number"
	SearchFieldTypeSetMoney     SearchFieldType = "set_money"
	SearchFieldTypeSetDate      SearchFieldType = "set_date"
	SearchFieldTypeSetDatetime  SearchFieldType = "set_datetime"
	SearchFieldTypeSetTime      SearchFieldType = "set_time"
	SearchFieldTypeSetReference SearchFieldType = "set_reference"
)

type SearchMatchType string

const (
	SearchMatchTypeAny SearchMatchType = "any"
	SearchMatchTypeAll SearchMatchType = "all"
)

type SearchMatchingVariant struct {
	// Unique identifier of the variant.
	ID int `json:"id"`
	// SKU of the matching variant.
	Sku *string `json:"sku,omitempty"`
}

type SearchQuery map[string]interface{}
type SearchCompoundExpression map[string]interface{}
type SearchAndExpression struct {
	And []SearchQuery `json:"and"`
}

type SearchFilterExpression struct {
	Filter []SearchQueryExpression `json:"filter"`
}

type SearchNotExpression struct {
	Not []SearchQuery `json:"not"`
}

type SearchOrExpression struct {
	Or []SearchQuery `json:"or"`
}

type SearchQueryExpression map[string]interface{}
type SearchDateRangeExpression struct {
	Range SearchDateRangeValue `json:"range"`
}

type SearchDateTimeRangeExpression struct {
	Range SearchDateTimeRangeValue `json:"range"`
}

type SearchExactExpression struct {
	Exact SearchAnyValue `json:"exact"`
}

type SearchExistsExpression struct {
	Exists SearchExistsValue `json:"exists"`
}

type SearchFullTextExpression struct {
	FullText SearchFullTextValue `json:"fullText"`
}

type SearchFullTextPrefixExpression struct {
	FullTextPrefix SearchFullTextPrefixValue `json:"fullTextPrefix"`
}

type SearchLongRangeExpression struct {
	Range SearchLongRangeValue `json:"range"`
}

type SearchNumberRangeExpression struct {
	Range SearchNumberRangeValue `json:"range"`
}

type SearchPrefixExpression struct {
	Prefix SearchAnyValue `json:"prefix"`
}

type SearchQueryExpressionValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
}

type SearchAnyValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Value     interface{}      `json:"value"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language        *string `json:"language,omitempty"`
	CaseInsensitive *bool   `json:"caseInsensitive,omitempty"`
}

type SearchDateRangeValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Gte       *Date            `json:"gte,omitempty"`
	Gt        *Date            `json:"gt,omitempty"`
	Lte       *Date            `json:"lte,omitempty"`
	Lt        *Date            `json:"lt,omitempty"`
}

type SearchDateTimeRangeValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Gte       *time.Time       `json:"gte,omitempty"`
	Gt        *time.Time       `json:"gt,omitempty"`
	Lte       *time.Time       `json:"lte,omitempty"`
	Lt        *time.Time       `json:"lt,omitempty"`
}

type SearchExistsValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language *string `json:"language,omitempty"`
}

type SearchFullTextPrefixValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Value     interface{}      `json:"value"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language  *string          `json:"language,omitempty"`
	MustMatch *SearchMatchType `json:"mustMatch,omitempty"`
}

type SearchFullTextValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Value     interface{}      `json:"value"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language  *string          `json:"language,omitempty"`
	MustMatch *SearchMatchType `json:"mustMatch,omitempty"`
}

type SearchLongRangeValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Gte       *int             `json:"gte,omitempty"`
	Gt        *int             `json:"gt,omitempty"`
	Lte       *int             `json:"lte,omitempty"`
	Lt        *int             `json:"lt,omitempty"`
}

type SearchNumberRangeValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Gte       *float64         `json:"gte,omitempty"`
	Gt        *float64         `json:"gt,omitempty"`
	Lte       *float64         `json:"lte,omitempty"`
	Lt        *float64         `json:"lt,omitempty"`
}

/**
*	For set-type fields, only a single value of the set is taken into account for sorting.
*	The sort mode determines whether the minimum or maximum value, or a calculated statistical value should be used as sorting value.
*
 */
type SearchSortMode string

const (
	SearchSortModeMin SearchSortMode = "min"
	SearchSortModeMax SearchSortMode = "max"
	SearchSortModeAvg SearchSortMode = "avg"
	SearchSortModeSum SearchSortMode = "sum"
)

type SearchSortOrder string

const (
	SearchSortOrderAsc  SearchSortOrder = "asc"
	SearchSortOrderDesc SearchSortOrder = "desc"
)

/**
*	Sorting parameters provided with a Search request.
*
 */
type SearchSorting struct {
	// Use any searchable field of the resource as sort criterion, or `"score"` to sort by relevance score calculated by the API.
	Field string `json:"field"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language *string `json:"language,omitempty"`
	// Specify the order in which the search results should be sorted.
	// Can be `asc` for ascending, or `desc` for descending order.
	Order SearchSortOrder `json:"order"`
	// Specify the sort mode to be applied for a set-type `field`.
	Mode *SearchSortMode `json:"mode,omitempty"`
	// Provide the data type of the given `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	// Allows you to apply a [sort filter](/../api/search-query-language#sort-filter).
	Filter *SearchQueryExpression `json:"filter,omitempty"`
}

type SearchTimeRangeExpression struct {
	Range SearchTimeRangeValue `json:"range"`
}

type SearchTimeRangeValue struct {
	Field string   `json:"field"`
	Boost *float64 `json:"boost,omitempty"`
	// Possible values for the `fieldType` property on [query expressions](/../api/search-query-language#query-expressions) indicating the data type of the `field`.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	Gte       *time.Time       `json:"gte,omitempty"`
	Gt        *time.Time       `json:"gt,omitempty"`
	Lte       *time.Time       `json:"lte,omitempty"`
	Lt        *time.Time       `json:"lt,omitempty"`
}

type SearchWildCardExpression struct {
	Wildcard SearchAnyValue `json:"wildcard"`
}
