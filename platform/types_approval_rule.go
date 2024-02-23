package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ApprovalRule struct {
	// Unique identifier of the Approval Rule.
	ID string `json:"id"`
	// Current version of the Approval Rule.
	Version int `json:"version"`
	// Date and time (UTC) the Approval Rule was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Approval Rule was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// User-defined unique identifier of the Approval Rule. Must be unique within a [Business Unit](ctp:api:type:BusinessUnit).
	Key *string `json:"key,omitempty"`
	// Name of the Approval Rule.
	Name string `json:"name"`
	// Description of the Approval Rule.
	Description *string `json:"description,omitempty"`
	// Indicates whether the Approval Rule should be matched against [Orders](ctp:api:type:Order) or not.
	Status ApprovalRuleStatus `json:"status"`
	// The [Order Predicate](/../api/projects/predicates#order-predicates) describing the [Orders](ctp:api:type:Order) the Approval Rule should match against.
	Predicate string `json:"predicate"`
	// The hierarchy of approvers within the Approval Rule.
	Approvers ApproverHierarchy `json:"approvers"`
	// The [Associate Roles](ctp:api:type:AssociateRole) customers must hold for their Order to require approval.
	Requesters []RuleRequester `json:"requesters"`
	// The [Business Unit](ctp:api:type:BusinessUnit) the Approval Rule belongs to.
	BusinessUnit BusinessUnitKeyReference `json:"businessUnit"`
}

type ApprovalRuleDraft struct {
	// User-defined unique identifier of the Approval Rule. Uniqueness is enforced within the Business Unit.
	Key *string `json:"key,omitempty"`
	// Name of the Approval Rule.
	Name string `json:"name"`
	// Description of the Approval Rule.
	Description *string `json:"description,omitempty"`
	// Indicates whether the Approval Rule should be matched against [Orders](ctp:api:type:Order) or not.
	Status ApprovalRuleStatus `json:"status"`
	// The [predicate](/../api/predicates/predicate-operators) describing the [Orders](ctp:api:type:Order) the Approval Rule should match against.
	Predicate string `json:"predicate"`
	// The hierarchy of approvers within the Approval Rule.
	Approvers ApproverHierarchyDraft `json:"approvers"`
	// The [Associate Roles](ctp:api:type:AssociateRole) customers must hold for their Order to require approval.
	Requesters []RuleRequesterDraft `json:"requesters"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ApprovalRule](ctp:api:type:ApprovalRule).
*
 */
type ApprovalRulePagedQueryResponse struct {
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
	// Approval Rules matching the query.
	Results []ApprovalRule `json:"results"`
}

/**
*	Indicates whether the Approval Rule should be matched against [Orders](ctp:api:type:Order) or not.
*
 */
type ApprovalRuleStatus string

const (
	ApprovalRuleStatusActive   ApprovalRuleStatus = "Active"
	ApprovalRuleStatusInactive ApprovalRuleStatus = "Inactive"
)

type ApprovalRuleUpdate struct {
	// Expected version of the [ApprovalRule](ctp:api:type:ApprovalRule) to which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the [ApprovalRule](ctp:api:type:ApprovalRule).
	Actions []ApprovalRuleUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ApprovalRuleUpdate) UnmarshalJSON(data []byte) error {
	type Alias ApprovalRuleUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorApprovalRuleUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ApprovalRuleUpdateAction interface{}

func mapDiscriminatorApprovalRuleUpdateAction(input interface{}) (ApprovalRuleUpdateAction, error) {
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
	case "setApprovers":
		obj := ApprovalRuleSetApproversAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ApprovalRuleSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ApprovalRuleSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := ApprovalRuleSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPredicate":
		obj := ApprovalRuleSetPredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRequesters":
		obj := ApprovalRuleSetRequestersAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatus":
		obj := ApprovalRuleSetStatusAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Setting the approvers for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRuleApproversSet](ctp:api:type:ApprovalRuleApproversSetMessage) Message.
*
 */
type ApprovalRuleSetApproversAction struct {
	// New approvers to set for the Approval Rule.
	Approvers ApproverHierarchyDraft `json:"approvers"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetApproversAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetApproversAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setApprovers", Alias: (*Alias)(&obj)})
}

/**
*	Setting the description for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRuleDescriptionSet](ctp:api:type:ApprovalRuleDescriptionSetMessage) Message.
*
 */
type ApprovalRuleSetDescriptionAction struct {
	// New description to set for the Approval Rule.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

/**
*	Setting the key for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRuleKeySet](ctp:api:type:ApprovalRuleKeySetMessage) Message.
*
 */
type ApprovalRuleSetKeyAction struct {
	// Value to set. Must be unique within a Business Unit. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	Setting the name for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRuleNameSet](ctp:api:type:ApprovalRuleNameSetMessage) Message.
*
 */
type ApprovalRuleSetNameAction struct {
	// New name to set for the Approval Rule.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Setting the [Order Predicate](/../api/projects/predicates#order-predicates) for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRulePredicateSet](ctp:api:type:ApprovalRulePredicateSetMessage) Message.
*
 */
type ApprovalRuleSetPredicateAction struct {
	// A valid [Order Predicate](/../api/projects/predicates#order-predicates) to set for the Approval Rule.
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}

/**
*	Sets the requesters for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRuleRequestersSet](ctp:api:type:ApprovalRuleRequestersSetMessage) Message.
*
 */
type ApprovalRuleSetRequestersAction struct {
	// New requesters to set for the Approval Rule.
	Requesters []RuleRequesterDraft `json:"requesters"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetRequestersAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetRequestersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRequesters", Alias: (*Alias)(&obj)})
}

/**
*	Setting the status for an [Approval Rule](ctp:api:type:ApprovalRule) generates an [ApprovalRuleStatusSet](ctp:api:type:ApprovalRuleStatusSetMessage) Message.
*
 */
type ApprovalRuleSetStatusAction struct {
	// New status to set for the Approval Rule.
	Status ApprovalRuleStatus `json:"status"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalRuleSetStatusAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalRuleSetStatusAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatus", Alias: (*Alias)(&obj)})
}

type ApproverConjunction struct {
	// All of the nested disjunctions must be approved in order for the conjunction to be considered approved.
	And []ApproverDisjunction `json:"and"`
}

type ApproverConjunctionDraft struct {
	// All of the nested disjunctions must be approved in order for the conjunction to be considered approved.
	// The total count of approvers across the nested disjunctions must not exceed 10.
	And []ApproverDisjunctionDraft `json:"and"`
}

type ApproverDisjunction struct {
	// Any of the nested approvers must approve in order for the disjunction to be considered approved.
	Or []RuleApprover `json:"or"`
}

type ApproverDisjunctionDraft struct {
	// Any of the nested approvers must approve in order for the disjunction to be considered approved.
	Or []RuleApproverDraft `json:"or"`
}

/**
*	Describes the order in which [Associates](ctp:api:type:Associate) can approve the matched [Order](ctp:api:type:Order).
*
 */
type ApproverHierarchy struct {
	// All of the nested conjunctions must be approved in order for the hierarchy to be considered approved.
	Tiers []ApproverConjunction `json:"tiers"`
}

/**
*	Describes the sequence in which [Associates](ctp:api:type:Associate) can approve an [Order](ctp:api:type:Order).
*
 */
type ApproverHierarchyDraft struct {
	// Nested conjunctions representing tiers of approvers in a hierarchy.
	Tiers []ApproverConjunctionDraft `json:"tiers"`
}

type RuleApprover struct {
	// The Associate Role that is allowed to approve at a given stage in the approval process.
	AssociateRole AssociateRoleKeyReference `json:"associateRole"`
}

type RuleApproverDraft struct {
	// Any Associate with this Role can approve.
	AssociateRole AssociateRoleResourceIdentifier `json:"associateRole"`
}

type RuleRequester struct {
	// The [Associate Role](ctp:api:type:AssociateRole) that an [Associate](ctp:api:type) must hold for the Approval Rule to apply to the Orders they create.
	AssociateRole AssociateRoleKeyReference `json:"associateRole"`
}

type RuleRequesterDraft struct {
	// The [Associate Role](ctp:api:type:AssociateRole) that an [Associate](ctp:api:type) must hold for the Approval Rule to apply to the Orders they create.
	AssociateRole AssociateRoleResourceIdentifier `json:"associateRole"`
}
