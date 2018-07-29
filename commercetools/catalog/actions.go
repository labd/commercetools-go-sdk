package catalog

import (
	"github.com/labd/commercetools-go-sdk/commercetools"
)


type ProductChangeName struct {
	commercetools.UpdateAction
	Name *commercetools.LocalizedString
}
type ProductAddPrice struct {
	commercetools.UpdateAction
	VariantID int
	Price     int
	Staged    bool
}

