package statemachines

import (
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// StateType enumeration contains values for each of the objects
// that can have states assigned.
type StateType string

const (
	// OrderStateType is used by Orders
	OrderStateType StateType = "OrderState"
	// LineItemStateType is used by LineItems and CustomLineItems in Carts and Orders
	LineItemStateType StateType = "LineItemState"
	// ProductStateType is used by Products
	ProductStateType StateType = "ProductState"
	// ReviewStateType is used by Reviews
	ReviewStateType StateType = "ReviewState"
	// PaymentStateType is used by paymentStatus in Payments
	PaymentStateType StateType = "PaymentState"
)

// StateRole defines the role of a State
type StateRole string

const (
	// ReviewIncludedInStatisticsRole is used by Reviews
	// This role only applies for the ReviewState StateType.
	// When the state of a Review has this role, the review will be taken into
	// account for rating statistics.
	ReviewIncludedInStatisticsRole StateRole = "ReviewIncludedInStatistics"
	// ReturnRole is used by Orders in the update action transitionLineItemState.
	// This role only applies for the LineItemState StateType.
	ReturnRole StateRole = "Return"
)

// State allows you to model states of certain objects, such as orders,
// line items, products, reviews, and payments in order to define finite
// state machines reflecting the business logic you'd like to implement.
type State struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// A unique identifier for the state.
	Key  string    `json:"key"`
	Type StateType `json:"type"`
	// A human-readable name of the state.
	Name commercetools.LocalizedString `json:"name"`
	// A human-readable description of the state.
	Description commercetools.LocalizedString `json:"description"`
	// A state can be declared as an initial state for any state machine.
	// When a workflow starts, this first state must be an initial state.
	Initial bool `json:"initial"`
	// Builtin states are integral parts of the project that cannot be deleted
	// nor the key can be changed.
	BuiltIn bool        `json:"builtIn"`
	Roles   []StateRole `json:"roles,omitempty"`
	// Transitions are a way to describe possible transformations of the current
	// state to other states of the same type (e.g.: Initial -> Shipped).
	// When performing a transitionState update action and transitions is set,
	// the currently referenced state must have a transition to the new state.
	//  If transitions is an empty list, it means the current state is a final
	// state and no further transitions are allowed.
	// If transitions is not set, the validation is turned off.
	// When performing a transitionState update action, any other state of
	// the same type can be transitioned to.
	Transitions []commercetools.Reference `json:"transitions,omitempty"`
}

// StateDraft is the representation to be given with a Create State request.
type StateDraft struct {
	Key         string                        `json:"key"`
	Type        StateType                     `json:"type"`
	Name        commercetools.LocalizedString `json:"name,omitempty"`
	Description commercetools.LocalizedString `json:"description,omitempty"`
	Initial     bool                          `json:"initial,omitempty"`
	Transitions []commercetools.Reference     `json:"transitions,omitempty"`
}
