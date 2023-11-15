package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ApprovalFlow struct {
	// Unique identifier of the Approval Flow.
	ID string `json:"id"`
	// Current version of the Approval Flow.
	Version int `json:"version"`
	// Date and time (UTC) the Approval Flow was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Approval Flow was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// [Order](ctp:api:type:Order) that needs to be approved.
	Order OrderReference `json:"order"`
	// [Business Unit](ctp:api:type:BusinessUnit) the Approval Flow belongs to.
	BusinessUnit BusinessUnitKeyReference `json:"businessUnit"`
	// Approval Rules that matched the [Order](ctp:api:type:Order).
	Rules []ApprovalRule `json:"rules"`
	// Indicates whether the Approval Flow is under review, approved, or rejected.
	Status ApprovalFlowStatus `json:"status"`
	// Present when the [status](ctp:api:type:ApprovalFlowStatus) of the Approval Flow is `Rejected`.
	Rejection *ApprovalFlowRejection `json:"rejection,omitempty"`
	// Existing approvals in the Approval Flow.
	Approvals []ApprovalFlowApproval `json:"approvals"`
	// Associate Roles that can approve according to the approver hierarchy tiers defined in `rules`.
	// Associates are allowed to reject even after they have given approval, as long as the current approver hierarchy tier still contains their role.
	EligibleApprovers []RuleApprover `json:"eligibleApprovers"`
	// Associate Roles required for approval based on the approver hierarchy tiers defined in `rules` across all remaining tiers.
	PendingApprovers []RuleApprover `json:"pendingApprovers"`
	// Associate Roles required for approval based on the approver hierarchy tiers defined in `rules` only for the currently active tier(s).
	CurrentTierPendingApprovers []RuleApprover `json:"currentTierPendingApprovers"`
}

type ApprovalFlowApproval struct {
	// Associate who approved the [Approval Flow](ctp:api:type:ApprovalFlow).
	Approver Associate `json:"approver"`
	// Date and time (UTC) when the [Approval Flow](ctp:api:type:ApprovalFlow) was approved at.
	ApprovedAt time.Time `json:"approvedAt"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ApprovalFlow](ctp:api:type:ApprovalFlow).
*
 */
type ApprovalFlowPagedQueryResponse struct {
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
	// Approval Flows matching the query.
	Results []ApprovalFlow `json:"results"`
}

type ApprovalFlowRejection struct {
	// Associate who rejected the [Approval Flow](ctp:api:type:ApprovalFlow).
	Rejecter Associate `json:"rejecter"`
	// Date and time (UTC) when the [Approval Flow](ctp:api:type:ApprovalFlow) was rejected at.
	RejectedAt time.Time `json:"rejectedAt"`
	// The reason for the rejection of the [Approval Flow](ctp:api:type:ApprovalFlow).
	Reason *string `json:"reason,omitempty"`
}

/**
*	Indicates whether the [Approval Flow](ctp:api:type:ApprovalFlow) is under review, approved, or rejected.
*
 */
type ApprovalFlowStatus string

const (
	ApprovalFlowStatusPending  ApprovalFlowStatus = "Pending"
	ApprovalFlowStatusApproved ApprovalFlowStatus = "Approved"
	ApprovalFlowStatusRejected ApprovalFlowStatus = "Rejected"
)

type ApprovalFlowUpdate struct {
	// Expected version of the [Approval Flow](ctp:api:type:ApprovalFlow) to which the changes should be applied.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the [Approval Flow](ctp:api:type:ApprovalFlow).
	Actions []ApprovalFlowUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ApprovalFlowUpdate) UnmarshalJSON(data []byte) error {
	type Alias ApprovalFlowUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorApprovalFlowUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ApprovalFlowUpdateAction interface{}

func mapDiscriminatorApprovalFlowUpdateAction(input interface{}) (ApprovalFlowUpdateAction, error) {
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
	case "approve":
		obj := ApprovalFlowApproveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "reject":
		obj := ApprovalFlowRejectAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	This update action allows an [Associate](ctp:api:type:Associate) to approve an Approval Flow. The process takes into account all [Associate Roles](ctp:api:type:AssociateRole) held by the Associate, aligning with the matched [Approval Rules](ctp:api:type:ApprovalRule) and their respective approver hierarchies.
*
*	When every required Associate has given their approval, the Approval Flow achieves a fully approved state, automatically updating its status to `Approved`. An Associate is eligible to approve only if their roles are within tiers of the Approval Rule hierarchy that are yet to be fully approved or rejected. As such, an Associate may be able to give their approval more than once.
*
 */
type ApprovalFlowApproveAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalFlowApproveAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalFlowApproveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "approve", Alias: (*Alias)(&obj)})
}

/**
*	This update action allows an [Associate](ctp:api:type:Associate) to reject an Approval Flow, setting its status to `Rejected`.
*	The process takes into account all [Associate Roles](ctp:api:type:AssociateRole) held by the Associate, aligning with the matched [Approval Rules](ctp:api:type:ApprovalRule) and their respective approver hierarchies.
*	Even a single rejection in the process will result in the rejection of the entire Approval Flow.
*
*	An Associate is eligible to reject only if their roles are within tiers of the Approval Rule hierarchy that are yet to be rejected. An Associate may alter a prior approval into a rejection.
*
 */
type ApprovalFlowRejectAction struct {
	// The reason for the rejection of the [Approval Flow](ctp:api:type:ApprovalFlow).
	Reason *string `json:"reason,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ApprovalFlowRejectAction) MarshalJSON() ([]byte, error) {
	type Alias ApprovalFlowRejectAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "reject", Alias: (*Alias)(&obj)})
}
