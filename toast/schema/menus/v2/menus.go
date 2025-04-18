package menus

import (
	"github.com/mondeendeguise/zestaurant/toast/schema/menus"
)

type strings []string
type ints []int

type MenuItems []MenuItem
type MenuItem struct {
	Name string
	KitchenName string
	GUID string
	MultiLocationID string
	MasterID int64
	Description string
	PosName string
	PosButtonColorLight string
	PosButtonColorDark string
	Image string
	Visibility strings
	Price float64
	PricingStrategy string
	PricingRules menus.PricingRules
	IsDeferred bool
	IsDiscountable bool
	SalesCategory menus.SalesCategory
	TaxInfo strings
	TaxInclusion string
	ItemTags menus.ItemTags
	Plu string
	Sku string
	Calories int
	ContentAdvisories menus.ContentAdvisories
	UnitOfMeasure string
	Portions menus.Portions
	PrepTime int32
	PrepStations strings
	ModifierGroupReferences ints
	EligiblePaymentAssistancePrograms strings
	Length float32
	Height float32
	Width float32
	DimensionUnitOfMeasure string
	Weight float32
	WeightUnitOfMeasure string
	Images strings
	GuestCount float32
}

type MenuGroups []MenuGroup
type MenuGroup struct {
	Name string
	GUID string
	MultiLocationID string
	MasterID int64
	Description string
	PosName string
	PosButtonColorLight string
	PosButtonColorDark string
	Image string
	Visibility strings
	ItemTags menus.ItemTags
	MenuGroups MenuGroups
	MenuItems MenuItems
}

type Menu struct {
	Name string
	GUID string
	MultiLocationID string
	MasterID int64
	Description string
	PosName string
	PosButtonColorLight string
	PosButtonColorDark string
	HighResImage string
	Image string
	Visibility strings
	Availability menus.Availability
	MenuGroups MenuGroups
}
