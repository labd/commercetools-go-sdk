package channels

import (
	"encoding/json"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/service/customfields"
	"github.com/labd/commercetools-go-sdk/service/products"
)

// ChannelRole describes the purpose and the type of this channel. Each channel
// can have one or more roles.
type ChannelRole string

const (
	// InventorySupply tells that this channel can be used to track inventory
	// entries. (e.g. channels with this role can be treated as warehouses)
	InventorySupply ChannelRole = "InventorySupply"

	// ProductDistribution tells that this channel can be used to expose
	// products to a specific distribution channel. It can be used by the cart
	// to select a product price.
	ProductDistribution ChannelRole = "ProductDistribution"

	// OrderExport tells that this channel can be used to track order export
	// activities.
	OrderExport ChannelRole = "OrderExport"

	// OrderImport tells that this channel can be used to track order import
	// activities.
	OrderImport ChannelRole = "OrderImport"

	// Primary role can be combined with some other roles (e.g. with
	// InventorySupply) to represent the fact that this particular channel is
	// the primary/master channel among the channels of the same type.
	Primary ChannelRole = "Primary"
)

// Channel represent a source or destination of different entities. They can be
// used to model warehouses or stores.
type Channel struct {
	// The unique ID of the channel.
	ID string `json:"id"`

	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`

	// Any arbitrary string key that uniquely identifies this channel within the
	// project.
	Key string `json:"key"`

	// The roles of this channel. Each channel must have at least one role.
	Roles []ChannelRole `json:"roles"`

	// A human-readable name of the channel.
	Name commercetools.LocalizedString `json:"name,omitempty"`

	// A human-readable description of the channel.
	Description commercetools.LocalizedString `json:"description,omitempty"`

	// The address where this channel is located (e.g. if the channel is a
	// physical store).
	Address commercetools.Address `json:"address,omitempty"`

	// Statistics about the review ratings taken into account for this channel.
	ReviewRatingStatistics products.ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`

	Custom customfields.CustomFields `json:"custom,omitempty"`

	// A GeoJSON geometry object encoding the geo location of the channel.
	GeoLocation commercetools.GeoJSONGeometry `json:"geoLocation,omitempty"`
}

// UnmarshalJSON override to map the shipping rate input type to the
// corresponding struct.
func (c *Channel) UnmarshalJSON(data []byte) error {
	type Alias Channel
	if err := json.Unmarshal(data, (*Alias)(c)); err != nil {
		return err
	}
	if c.GeoLocation != nil {
		c.GeoLocation = commercetools.GeoJSONGeometryMapping(c.GeoLocation)
	}
	return nil
}

// ChannelDraft is given as payload for Create Channel requests.
type ChannelDraft struct {
	Key string `json:"key"`

	// If not specified, then channel will get InventorySupply role by default.
	Roles []ChannelRole `json:"roles,omitempty"`

	Name        commercetools.LocalizedString `json:"name,omitempty"`
	Description commercetools.LocalizedString `json:"description,omitempty"`
	Address     commercetools.Address         `json:"address,omitempty"`

	// The custom fields.
	Custom customfields.CustomFieldsDraft `json:"custom,omitempty"`

	GeoLocation commercetools.GeoJSONGeometry `json:"geoLocation,omitempty"`
}
