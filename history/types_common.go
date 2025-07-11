package history

// Generated file, please do not change!!!

type LocalizedString map[string]string
type Address struct {
	// Unique ID of the Address.
	ID                   string `json:"id"`
	Key                  string `json:"key"`
	Title                string `json:"title"`
	Salutation           string `json:"salutation"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	StreetName           string `json:"streetName"`
	StreetNumber         string `json:"streetNumber"`
	AdditionalStreetInfo string `json:"additionalStreetInfo"`
	PostalCode           string `json:"postalCode"`
	City                 string `json:"city"`
	Region               string `json:"region"`
	State                string `json:"state"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country               string `json:"country"`
	Company               string `json:"company"`
	Department            string `json:"department"`
	Building              string `json:"building"`
	Apartment             string `json:"apartment"`
	POBox                 string `json:"pOBox"`
	Phone                 string `json:"phone"`
	Mobile                string `json:"mobile"`
	Email                 string `json:"email"`
	Fax                   string `json:"fax"`
	AdditionalAddressInfo string `json:"additionalAddressInfo"`
	ExternalId            string `json:"externalId"`
}

type Asset struct {
	ID          string          `json:"id"`
	Name        LocalizedString `json:"name"`
	Description LocalizedString `json:"description"`
	Custom      CustomFields    `json:"custom"`
	Key         string          `json:"key"`
}

type AssetDimensions struct {
	W int `json:"w"`
	H int `json:"h"`
}

type AssetSource struct {
	Uri         string          `json:"uri"`
	Key         string          `json:"key"`
	Dimensions  AssetDimensions `json:"dimensions"`
	ContentType string          `json:"contentType"`
}

type Associate struct {
	AssociateRoleAssignments []AssociateRoleAssignment `json:"associateRoleAssignments"`
	Customer                 Reference                 `json:"customer"`
}

type AssociateRoleAssignment struct {
	AssociateRole KeyReference `json:"associateRole"`
	// Determines whether an [AssociateRoleAssignment](ctp:api:type:AssociateRoleAssignment) can be inherited by child Business Units.
	Inheritance AssociateRoleInheritanceMode `json:"inheritance"`
}

/**
*	Determines whether an [AssociateRoleAssignment](ctp:api:type:AssociateRoleAssignment) can be inherited by child Business Units.
 */
type AssociateRoleInheritanceMode string

const (
	AssociateRoleInheritanceModeEnabled  AssociateRoleInheritanceMode = "Enabled"
	AssociateRoleInheritanceModeDisabled AssociateRoleInheritanceMode = "Disabled"
)

type AttributeConstraintEnum string

const (
	AttributeConstraintEnumNone              AttributeConstraintEnum = "None"
	AttributeConstraintEnumUnique            AttributeConstraintEnum = "Unique"
	AttributeConstraintEnumCombinationUnique AttributeConstraintEnum = "CombinationUnique"
	AttributeConstraintEnumSameForAll        AttributeConstraintEnum = "SameForAll"
)

type AttributeDefinition struct {
	Type AttributeType `json:"type"`
	// The unique name of the attribute used in the API. The name must be between two and 256 characters long and can contain the ASCII letters A to Z in lowercase or uppercase, digits, underscores (`_`) and the hyphen-minus (`-`). When using the same `name` for an attribute in two or more product types all fields of the AttributeDefinition of this attribute need to be the same across the product types, otherwise an AttributeDefinitionAlreadyExists error code will be returned. An exception to this are the values of an `enum` or `lenum` type and sets thereof.
	Name  string          `json:"name"`
	Label LocalizedString `json:"label"`
	// Whether the attribute is required to have a value.
	IsRequired          bool                    `json:"isRequired"`
	AttributeConstraint AttributeConstraintEnum `json:"attributeConstraint"`
	InputTip            LocalizedString         `json:"inputTip"`
	InputHint           TextInputHint           `json:"inputHint"`
	// Whether the attribute's values should generally be enabled in product search. This determines whether the value is stored in products for matching terms in the context of full-text search queries  and can be used in facets & filters as part of product search queries. The exact features that are enabled/disabled with this flag depend on the concrete attribute type and are described there. The max size of a searchable field is **restricted to 10922 characters**. This constraint is enforced at both product creation and product update. If the length of the input exceeds the maximum size an InvalidField error is returned.
	IsSearchable bool `json:"isSearchable"`
}

/**
*	A localized enum value must be unique within the enum, else a [DuplicateEnumValues](ctp:api:type:DuplicateEnumValuesError) error is returned.
*
 */
type AttributeLocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier, for example in facets & filters.
	Key string `json:"key"`
	// Descriptive, localized label of the value.
	Label LocalizedString `json:"label"`
}

/**
*	A plain enum value must be unique within the enum, else a [DuplicateEnumValues](ctp:api:type:DuplicateEnumValuesError) error is returned.
*
 */
type AttributePlainEnumValue struct {
	// Key of the value used as a programmatic identifier, for example in facets & filters.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

type AttributeType struct {
	Name string `json:"name"`
}

type AuthenticationMode string

const (
	AuthenticationModePassword     AuthenticationMode = "Password"
	AuthenticationModeExternalAuth AuthenticationMode = "ExternalAuth"
)

/**
*	Determines whether a Business Unit can inherit Associates from a parent.
 */
type BusinessUnitAssociateMode string

const (
	BusinessUnitAssociateModeExplicit              BusinessUnitAssociateMode = "Explicit"
	BusinessUnitAssociateModeExplicitAndFromParent BusinessUnitAssociateMode = "ExplicitAndFromParent"
)

/**
*	Indicates whether the Business Unit can be edited and used in [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [Quote Requests](ctp:api:type:QuoteRequest), or [Quotes](ctp:api:type:Quote).
 */
type BusinessUnitStatus string

const (
	BusinessUnitStatusActive   BusinessUnitStatus = "Active"
	BusinessUnitStatusInactive BusinessUnitStatus = "Inactive"
)

/**
*	Defines whether the Stores of the Business Unit are set directly on the Business Unit or are inherited from its parent unit.
 */
type BusinessUnitStoreMode string

const (
	BusinessUnitStoreModeExplicit   BusinessUnitStoreMode = "Explicit"
	BusinessUnitStoreModeFromParent BusinessUnitStoreMode = "FromParent"
)

type CategoryOrderHints map[string]string

/**
*	Describes the purpose and type of the Channel. A Channel can have one or more roles.
*
 */
type ChannelRoleEnum string

const (
	ChannelRoleEnumInventorySupply     ChannelRoleEnum = "InventorySupply"
	ChannelRoleEnumProductDistribution ChannelRoleEnum = "ProductDistribution"
	ChannelRoleEnumOrderExport         ChannelRoleEnum = "OrderExport"
	ChannelRoleEnumOrderImport         ChannelRoleEnum = "OrderImport"
	ChannelRoleEnumPrimary             ChannelRoleEnum = "Primary"
)

type CustomFields struct {
	Type Reference `json:"type"`
	// A valid JSON object, based on FieldDefinition.
	Fields interface{} `json:"fields"`
}

type CustomLineItem struct {
	// The unique ID of this CustomLineItem.
	ID         string          `json:"id"`
	Name       LocalizedString `json:"name"`
	Money      Money           `json:"money"`
	TaxedPrice TaxedItemPrice  `json:"taxedPrice"`
	TotalPrice Money           `json:"totalPrice"`
	// A unique String in the cart to identify this CustomLineItem.
	Slug string `json:"slug"`
	// The amount of a CustomLineItem in the cart. Must be a positive integer.
	Quantity int `json:"quantity"`
}

type Delivery struct {
	ID        string         `json:"id"`
	CreatedAt string         `json:"createdAt"`
	Items     []DeliveryItem `json:"items"`
	Parcels   []Parcel       `json:"parcels"`
	Address   *Address       `json:"address,omitempty"`
	// Custom Fields for the Transaction.
	Custom *CustomFields `json:"custom,omitempty"`
}

type DeliveryItem struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type DiscountCodeInfo struct {
	DiscountCode Reference         `json:"discountCode"`
	State        DiscountCodeState `json:"state"`
}

type DiscountCodeState string

const (
	DiscountCodeStateNotActive                            DiscountCodeState = "NotActive"
	DiscountCodeStateDoesNotMatchCart                     DiscountCodeState = "DoesNotMatchCart"
	DiscountCodeStateMatchesCart                          DiscountCodeState = "MatchesCart"
	DiscountCodeStateMaxApplicationReached                DiscountCodeState = "MaxApplicationReached"
	DiscountCodeStateApplicationStoppedByPreviousDiscount DiscountCodeState = "ApplicationStoppedByPreviousDiscount"
	DiscountCodeStateApplicationStoppedByGroupBestDeal    DiscountCodeState = "ApplicationStoppedByGroupBestDeal"
	DiscountCodeStateNotValid                             DiscountCodeState = "NotValid"
)

type DiscountedLineItemPortion struct {
	Discount         Reference `json:"discount"`
	DiscountedAmount Money     `json:"discountedAmount"`
}

type DiscountedLineItemPrice struct {
	Value             Money                       `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

type DiscountedLineItemPriceForQuantity struct {
	Quantity        int                     `json:"quantity"`
	DiscountedPrice DiscountedLineItemPrice `json:"discountedPrice"`
}

type FieldDefinition struct {
	Type FieldType `json:"type"`
	// The name of the field. The name must be between two and 36 characters long and can contain the ASCII letters A to Z in lowercase or uppercase, digits, underscores (`_`) and the hyphen-minus (`-`). The name must be unique for a given resource type ID. In case there is a field with the same name in another type it has to have the same FieldType also.
	Name      string          `json:"name"`
	Label     LocalizedString `json:"label"`
	InputHint TextInputHint   `json:"inputHint"`
}

type FieldType struct {
	Name string `json:"name"`
}

type GeoLocation struct {
	Type        string `json:"type"`
	Coordinates []int  `json:"coordinates"`
}

type Image struct {
	Url        string          `json:"url"`
	Dimensions ImageDimensions `json:"dimensions"`
	Label      string          `json:"label"`
}

type ImageDimensions struct {
	W int `json:"w"`
	H int `json:"h"`
}

type InheritedAssociate struct {
	AssociateRoleAssignments []InheritedAssociateRoleAssignment `json:"associateRoleAssignments"`
	Customer                 Reference                          `json:"customer"`
}

type InheritedAssociateRoleAssignment struct {
	AssociateRole KeyReference `json:"associateRole"`
	Source        KeyReference `json:"source"`
}

type ItemShippingDetails struct {
	Targets []ItemShippingTarget `json:"targets"`
	// true if the quantity of the (custom) line item is equal to the sum of the sub-quantities in `targets`, `false` otherwise. A cart cannot be ordered when the value is `false`. The error InvalidItemShippingDetails will be triggered.
	Valid bool `json:"valid"`
}

type ItemShippingTarget struct {
	// The key of the address in the cart's `itemShippingAddresses`
	AddressKey string `json:"addressKey"`
	// The quantity of items that should go to the address with the specified `addressKey`. Only positive values are allowed. Using `0` as quantity is also possible in a draft object, but the element will not be present in the resulting ItemShippingDetails.
	Quantity int `json:"quantity"`
}

type ItemState struct {
	Quantity int       `json:"quantity"`
	State    Reference `json:"state"`
}

type KeyReference struct {
	Key    string          `json:"key"`
	TypeId ReferenceTypeId `json:"typeId"`
}

type LineItem struct {
	AddedAt     string          `json:"addedAt"`
	Custom      CustomFields    `json:"custom"`
	ID          string          `json:"id"`
	Name        LocalizedString `json:"name"`
	ProductId   string          `json:"productId"`
	ProductSlug LocalizedString `json:"productSlug"`
	ProductType Reference       `json:"productType"`
	Quantity    int             `json:"quantity"`
	Variant     Variant         `json:"variant"`
	VariantId   int             `json:"variantId"`
}

/**
*	Shape of the value for `addLocation` and `removeLocation` actions
 */
type Location struct {
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	State   string `json:"state"`
}

type Money struct {
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode   string    `json:"currencyCode"`
	CentAmount     int       `json:"centAmount"`
	FractionDigits int       `json:"fractionDigits"`
	Type           MoneyType `json:"type"`
}

type MoneyType string

const (
	MoneyTypeCentPrecision MoneyType = "centPrecision"
	MoneyTypeHighPrecision MoneyType = "highPrecision"
)

type OrderState string

const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

type Parcel struct {
	ID           string             `json:"id"`
	CreatedAt    string             `json:"createdAt"`
	Measurements ParcelMeasurements `json:"measurements"`
	TrackingData TrackingData       `json:"trackingData"`
	Items        []DeliveryItem     `json:"items"`
}

type ParcelMeasurements struct {
	HeightInMillimeter int `json:"heightInMillimeter"`
	LengthInMillimeter int `json:"lengthInMillimeter"`
	WidthInMillimeter  int `json:"widthInMillimeter"`
	WeightInGram       int `json:"weightInGram"`
}

type PaymentInfo struct {
	Payments []Reference `json:"payments"`
}

type PaymentState string

const (
	PaymentStateBalanceDue PaymentState = "BalanceDue"
	PaymentStateFailed     PaymentState = "Failed"
	PaymentStatePending    PaymentState = "Pending"
	PaymentStateCreditOwed PaymentState = "CreditOwed"
	PaymentStatePaid       PaymentState = "Paid"
)

/**
*	Permissions grant granular access to [Business Units](ctp:api:type:BusinessUnit), [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [Quotes](ctp:api:type:Quote), and [Quote Requests](ctp:api:type:QuoteRequest).
 */
type Permission string

const (
	PermissionAddChildUnits                      Permission = "AddChildUnits"
	PermissionUpdateAssociates                   Permission = "UpdateAssociates"
	PermissionUpdateBusinessUnitDetails          Permission = "UpdateBusinessUnitDetails"
	PermissionUpdateParentUnit                   Permission = "UpdateParentUnit"
	PermissionViewMyCarts                        Permission = "ViewMyCarts"
	PermissionViewOthersCarts                    Permission = "ViewOthersCarts"
	PermissionUpdateMyCarts                      Permission = "UpdateMyCarts"
	PermissionUpdateOthersCarts                  Permission = "UpdateOthersCarts"
	PermissionCreateMyCarts                      Permission = "CreateMyCarts"
	PermissionCreateOthersCarts                  Permission = "CreateOthersCarts"
	PermissionDeleteMyCarts                      Permission = "DeleteMyCarts"
	PermissionDeleteOthersCarts                  Permission = "DeleteOthersCarts"
	PermissionViewMyOrders                       Permission = "ViewMyOrders"
	PermissionViewOthersOrders                   Permission = "ViewOthersOrders"
	PermissionUpdateMyOrders                     Permission = "UpdateMyOrders"
	PermissionUpdateOthersOrders                 Permission = "UpdateOthersOrders"
	PermissionCreateMyOrdersFromMyCarts          Permission = "CreateMyOrdersFromMyCarts"
	PermissionCreateMyOrdersFromMyQuotes         Permission = "CreateMyOrdersFromMyQuotes"
	PermissionCreateOrdersFromOthersCarts        Permission = "CreateOrdersFromOthersCarts"
	PermissionCreateOrdersFromOthersQuotes       Permission = "CreateOrdersFromOthersQuotes"
	PermissionViewMyQuotes                       Permission = "ViewMyQuotes"
	PermissionViewOthersQuotes                   Permission = "ViewOthersQuotes"
	PermissionAcceptMyQuotes                     Permission = "AcceptMyQuotes"
	PermissionAcceptOthersQuotes                 Permission = "AcceptOthersQuotes"
	PermissionDeclineMyQuotes                    Permission = "DeclineMyQuotes"
	PermissionDeclineOthersQuotes                Permission = "DeclineOthersQuotes"
	PermissionRenegotiateMyQuotes                Permission = "RenegotiateMyQuotes"
	PermissionRenegotiateOthersQuotes            Permission = "RenegotiateOthersQuotes"
	PermissionReassignMyQuotes                   Permission = "ReassignMyQuotes"
	PermissionReassignOthersQuotes               Permission = "ReassignOthersQuotes"
	PermissionViewMyQuoteRequests                Permission = "ViewMyQuoteRequests"
	PermissionViewOthersQuoteRequests            Permission = "ViewOthersQuoteRequests"
	PermissionUpdateMyQuoteRequests              Permission = "UpdateMyQuoteRequests"
	PermissionUpdateOthersQuoteRequests          Permission = "UpdateOthersQuoteRequests"
	PermissionCreateMyQuoteRequestsFromMyCarts   Permission = "CreateMyQuoteRequestsFromMyCarts"
	PermissionCreateQuoteRequestsFromOthersCarts Permission = "CreateQuoteRequestsFromOthersCarts"
)

type Price struct {
	ID    string `json:"id"`
	Value Money  `json:"value"`
}

type ProductSelectionSetting struct {
	ProductSelection Reference `json:"productSelection"`
	Active           bool      `json:"active"`
}

type ProductVariantAvailability struct {
	IsOnStock         bool                                 `json:"isOnStock"`
	RestockableInDays int                                  `json:"restockableInDays"`
	AvailableQuantity int                                  `json:"availableQuantity"`
	Channels          ProductVariantChannelAvailabilityMap `json:"channels"`
}

type ProductVariantChannelAvailability struct {
	IsOnStock         bool `json:"isOnStock"`
	RestockableInDays int  `json:"restockableInDays"`
	AvailableQuantity int  `json:"availableQuantity"`
}

type ProductVariantChannelAvailabilityMap map[string]ProductVariantChannelAvailability
type ProductVariantSelection struct {
	Type ProductVariantSelectionTypeEnum `json:"type"`
	Skus []string                        `json:"skus"`
}

type ProductVariantSelectionTypeEnum string

const (
	ProductVariantSelectionTypeEnumInclusion ProductVariantSelectionTypeEnum = "inclusion"
	ProductVariantSelectionTypeEnumExclusion ProductVariantSelectionTypeEnum = "exclusion"
)

type QuoteRequestState string

const (
	QuoteRequestStateSubmitted QuoteRequestState = "Submitted"
	QuoteRequestStateAccepted  QuoteRequestState = "Accepted"
	QuoteRequestStateClosed    QuoteRequestState = "Closed"
	QuoteRequestStateRejected  QuoteRequestState = "Rejected"
	QuoteRequestStateCancelled QuoteRequestState = "Cancelled"
)

type QuoteState string

const (
	QuoteStatePending                  QuoteState = "Pending"
	QuoteStateDeclined                 QuoteState = "Declined"
	QuoteStateDeclinedForRenegotiation QuoteState = "DeclinedForRenegotiation"
	QuoteStateAccepted                 QuoteState = "Accepted"
	QuoteStateFailed                   QuoteState = "Failed"
	QuoteStateWithdrawn                QuoteState = "Withdrawn"
)

type Reference struct {
	ID     string          `json:"id"`
	TypeId ReferenceTypeId `json:"typeId"`
}

type ReferenceTypeId string

const (
	ReferenceTypeIdAssociateRole         ReferenceTypeId = "associate-role"
	ReferenceTypeIdBusinessUnit          ReferenceTypeId = "business-unit"
	ReferenceTypeIdCart                  ReferenceTypeId = "cart"
	ReferenceTypeIdCartDiscount          ReferenceTypeId = "cart-discount"
	ReferenceTypeIdCategory              ReferenceTypeId = "category"
	ReferenceTypeIdChannel               ReferenceTypeId = "channel"
	ReferenceTypeIdCustomer              ReferenceTypeId = "customer"
	ReferenceTypeIdCustomerEmailToken    ReferenceTypeId = "customer-email-token"
	ReferenceTypeIdCustomerGroup         ReferenceTypeId = "customer-group"
	ReferenceTypeIdCustomerPasswordToken ReferenceTypeId = "customer-password-token"
	ReferenceTypeIdDiscountCode          ReferenceTypeId = "discount-code"
	ReferenceTypeIdExtension             ReferenceTypeId = "extension"
	ReferenceTypeIdInventoryEntry        ReferenceTypeId = "inventory-entry"
	ReferenceTypeIdKeyValueDocument      ReferenceTypeId = "key-value-document"
	ReferenceTypeIdOrder                 ReferenceTypeId = "order"
	ReferenceTypeIdOrderEdit             ReferenceTypeId = "order-edit"
	ReferenceTypeIdPaymentMethod         ReferenceTypeId = "payment-method"
	ReferenceTypeIdPayment               ReferenceTypeId = "payment"
	ReferenceTypeIdProduct               ReferenceTypeId = "product"
	ReferenceTypeIdProductDiscount       ReferenceTypeId = "product-discount"
	ReferenceTypeIdProductSelection      ReferenceTypeId = "product-selection"
	ReferenceTypeIdProductType           ReferenceTypeId = "product-type"
	ReferenceTypeIdQuote                 ReferenceTypeId = "quote"
	ReferenceTypeIdQuoteRequest          ReferenceTypeId = "quote-request"
	ReferenceTypeIdReview                ReferenceTypeId = "review"
	ReferenceTypeIdShippingMethod        ReferenceTypeId = "shipping-method"
	ReferenceTypeIdShoppingList          ReferenceTypeId = "shopping-list"
	ReferenceTypeIdStagedQuote           ReferenceTypeId = "staged-quote"
	ReferenceTypeIdState                 ReferenceTypeId = "state"
	ReferenceTypeIdStore                 ReferenceTypeId = "store"
	ReferenceTypeIdSubscription          ReferenceTypeId = "subscription"
	ReferenceTypeIdTaxCategory           ReferenceTypeId = "tax-category"
	ReferenceTypeIdType                  ReferenceTypeId = "type"
	ReferenceTypeIdZone                  ReferenceTypeId = "zone"
)

type Reservation struct {
	Quantity          int       `json:"quantity"`
	Owner             Reference `json:"owner"`
	CreatedAt         string    `json:"createdAt"`
	CheckoutStartedAt string    `json:"checkoutStartedAt"`
}

type ResourceIdentifier struct {
	ID     string          `json:"id"`
	Key    string          `json:"key"`
	TypeId ReferenceTypeId `json:"typeId"`
}

type ReturnInfo struct {
	Items []ReturnItem `json:"items"`
	// Identifies, which return tracking ID is connected to this particular return.
	ReturnTrackingId string `json:"returnTrackingId"`
	ReturnDate       string `json:"returnDate"`
}

type ReturnItem struct {
	ID             string              `json:"id"`
	Quantity       int                 `json:"quantity"`
	Type           string              `json:"type"`
	Comment        string              `json:"comment"`
	ShipmentState  ReturnShipmentState `json:"shipmentState"`
	PaymentState   ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt string              `json:"lastModifiedAt"`
	CreatedAt      string              `json:"createdAt"`
}

type ReturnPaymentState string

const (
	ReturnPaymentStateNonRefundable ReturnPaymentState = "NonRefundable"
	ReturnPaymentStateInitial       ReturnPaymentState = "Initial"
	ReturnPaymentStateRefunded      ReturnPaymentState = "Refunded"
	ReturnPaymentStateNotRefunded   ReturnPaymentState = "NotRefunded"
)

type ReturnShipmentState string

const (
	ReturnShipmentStateAdvised     ReturnShipmentState = "Advised"
	ReturnShipmentStateReturned    ReturnShipmentState = "Returned"
	ReturnShipmentStateBackInStock ReturnShipmentState = "BackInStock"
	ReturnShipmentStateUnusable    ReturnShipmentState = "Unusable"
)

type ReviewRatingStatistics struct {
	// Average rating of one target This number is rounded with 5 decimals.
	AverageRating int `json:"averageRating"`
	// Highest rating of one target
	HighestRating int `json:"highestRating"`
	// Lowest rating of one target
	LowestRating int `json:"lowestRating"`
	// Number of ratings taken into account
	Count int `json:"count"`
	// The full distribution of the ratings. The keys are the different ratings and the values are the count of reviews having this rating. Only the used ratings appear in this object.
	RatingsDistribution interface{} `json:"ratingsDistribution"`
}

type RoundingMode string

const (
	RoundingModeHalfEven RoundingMode = "HalfEven"
	RoundingModeHalfUp   RoundingMode = "HalfUp"
	RoundingModeHalfDown RoundingMode = "HalfDown"
)

type SearchKeyword struct {
	Text             string           `json:"text"`
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer"`
}

type SearchKeywords map[string][]SearchKeyword
type SelectionMode string

const (
	SelectionModeCheapest      SelectionMode = "Cheapest"
	SelectionModeMostExpensive SelectionMode = "MostExpensive"
)

type ShipmentState string

const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
)

type ShippingRate struct {
	Price     Money `json:"price"`
	FreeAbove Money `json:"freeAbove"`
	// Only appears in response to requests for ShippingMethods by Cart or location to mark this shipping rate as one that matches the Cart or location.
	IsMatching bool                    `json:"isMatching"`
	Tiers      []ShippingRatePriceTier `json:"tiers"`
}

type ShippingRatePriceTier struct {
	Type ShippingRateTierType `json:"type"`
}

type ShippingRateTierType string

const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

type StackingMode string

const (
	StackingModeStacking              StackingMode = "Stacking"
	StackingModeStopAfterThisDiscount StackingMode = "StopAfterThisDiscount"
)

type StagedQuoteState string

const (
	StagedQuoteStateInProgress StagedQuoteState = "InProgress"
	StagedQuoteStateSent       StagedQuoteState = "Sent"
	StagedQuoteStateClosed     StagedQuoteState = "Closed"
)

/**
*	For some resource types, a State can fulfill the following predefined roles:
*
 */
type StateRoleEnum string

const (
	StateRoleEnumReviewIncludedInStatistics StateRoleEnum = "ReviewIncludedInStatistics"
	StateRoleEnumReturn                     StateRoleEnum = "Return"
)

/**
*	Resource or object type the State can be assigned to.
*
 */
type StateTypeEnum string

const (
	StateTypeEnumOrderState          StateTypeEnum = "OrderState"
	StateTypeEnumRecurringOrderState StateTypeEnum = "RecurringOrderState"
	StateTypeEnumLineItemState       StateTypeEnum = "LineItemState"
	StateTypeEnumProductState        StateTypeEnum = "ProductState"
	StateTypeEnumReviewState         StateTypeEnum = "ReviewState"
	StateTypeEnumPaymentState        StateTypeEnum = "PaymentState"
	StateTypeEnumQuoteRequestState   StateTypeEnum = "QuoteRequestState"
	StateTypeEnumStagedQuoteState    StateTypeEnum = "StagedQuoteState"
	StateTypeEnumQuoteState          StateTypeEnum = "QuoteState"
)

type StoreCountry struct {
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Code string `json:"code"`
}

type SubRate struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type SuggestTokenizer struct {
	Type string `json:"type"`
}

type SyncInfo struct {
	Channel Reference `json:"channel"`
	// Can be used to reference an external order instance, file etc.
	ExternalId string `json:"externalId"`
	SyncedAt   string `json:"syncedAt"`
}

type TaxCalculationMode string

const (
	TaxCalculationModeLineItemLevel  TaxCalculationMode = "LineItemLevel"
	TaxCalculationModeUnitPriceLevel TaxCalculationMode = "UnitPriceLevel"
)

type TaxMode string

const (
	TaxModePlatform       TaxMode = "Platform"
	TaxModeExternal       TaxMode = "External"
	TaxModeExternalAmount TaxMode = "ExternalAmount"
	TaxModeDisabled       TaxMode = "Disabled"
)

/**
*	Shape of the value for `addTaxRate` and `removeTaxRate` actions
 */
type TaxRate struct {
	// The ID is always set if the tax rate is part of a TaxCategory. The external tax rates in a Cart do not contain an `id`.
	ID   string `json:"id"`
	Name string `json:"name"`
	// Percentage in the range of [0..1]. The sum of the amounts of all `subRates`, if there are any.
	Amount          int  `json:"amount"`
	IncludedInPrice bool `json:"includedInPrice"`
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// The state in the country
	State    string    `json:"state"`
	SubRates []SubRate `json:"subRates"`
}

type TaxedItemPrice struct {
	TotalNet   Money `json:"totalNet"`
	TotalGross Money `json:"totalGross"`
}

type TaxedPrice struct {
	// Total net price of the Order.
	TotalNet Money `json:"totalNet"`
	// Total gross price of the Order.
	TotalGross Money `json:"totalGross"`
}

type TextInputHint string

const (
	TextInputHintSingleLine TextInputHint = "SingleLine"
	TextInputHintMultiLine  TextInputHint = "MultiLine"
)

type TextLineItem struct {
	AddedAt     string          `json:"addedAt"`
	Custom      CustomFields    `json:"custom"`
	Description LocalizedString `json:"description"`
	ID          string          `json:"id"`
	Name        LocalizedString `json:"name"`
	Quantity    int             `json:"quantity"`
}

type TrackingData struct {
	// The ID to track one parcel.
	TrackingId string `json:"trackingId"`
	// The carrier that delivers the parcel.
	Carrier             string `json:"carrier"`
	Provider            string `json:"provider"`
	ProviderTransaction string `json:"providerTransaction"`
	// Flag to distinguish if the parcel is on the way to the customer (false) or on the way back (true).
	IsReturn bool `json:"isReturn"`
}

type Transaction struct {
	// Unique identifier of the Transaction.
	ID string `json:"id"`
	// Time at which the transaction took place.
	Timestamp string          `json:"timestamp"`
	Type      TransactionType `json:"type"`
	Amount    Money           `json:"amount"`
	// Identifier used by the interface that manages the transaction (usually the PSP). If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction should be findable with this ID.
	InteractionId string           `json:"interactionId"`
	State         TransactionState `json:"state"`
}

type TransactionState string

const (
	TransactionStateInitial TransactionState = "Initial"
	TransactionStatePending TransactionState = "Pending"
	TransactionStateSuccess TransactionState = "Success"
	TransactionStateFailure TransactionState = "Failure"
)

type TransactionType string

const (
	TransactionTypeAuthorization       TransactionType = "Authorization"
	TransactionTypeCancelAuthorization TransactionType = "CancelAuthorization"
	TransactionTypeCharge              TransactionType = "Charge"
	TransactionTypeRefund              TransactionType = "Refund"
	TransactionTypeChargeback          TransactionType = "Chargeback"
)

type Variant struct {
	ID  int    `json:"id"`
	Sku string `json:"sku"`
	Key string `json:"key"`
}
