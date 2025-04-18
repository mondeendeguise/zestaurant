package menus

type strings []string
type ints []int

type TimeRanges []TimeRange
type TimeRange struct {
	Start string
	End string
}

type Schedules []Schedule
type Schedule struct {
	Days strings
	TimeRanges TimeRanges
}

type Availability struct {
	AlwaysAvailable bool
	Schedule Schedules
}

type ItemTags []ItemTag
type ItemTag struct {
	Name string
	GUID string
}

type SequencePrices []SequencePrice
type SequencePrice struct {
	Sequence int32
	Price float64
}

type SizeSequencePricingRules []SizeSequencePricingRule
type SizeSequencePricingRule struct {
	SizeName string
	SizeGUID string
	SequencePrices SequencePrices
}

type TimeSpecificPrices []TimeSpecificPrice
type TimeSpecificPrice struct {
	TimeSpecificPrice float64
	BasePrice float64
	Schedule Schedules
}

type PricingRules struct {
	TimeSpecificPricingRules TimeSpecificPrices
	SizeSpecificPricingGUID string
	SizeSequencePricingRules SizeSequencePricingRules
}

type SalesCategory struct {
	Name string
	GUID string
}

type Alcohol struct {
	ContainsAlcohol string
}

type ContentAdvisories struct {
	Alcohol Alcohol
}

type Portions []Portion
type Portion struct {
	Name string
	GUID string
	ModifierGroupReferences ints
}

type PreModifiers []PreModifier
type PreModifier struct {
	Name string
	GUID string
	MultiLocationID string
	FixedPrice float64
	MultiplicationFactor float64
	DisplayMode string
	PosName string
	PosButtonColorLight string
	PosButtonColorDark string
}

type PreModifierGroup struct {
	Name string
	GUID string
	MultiLocationID string
	PreModifiers PreModifiers
}

type ModifierGroup struct {
	Name string
	GUID string
	ReferenceID int
	MultiLocationID string
	MasterID int64 // DEPRECATED
	PosName string
	PosButtonColorLight string
	PosButtonColorDark string
	PricingStrategy string
	PricingRules PricingRules
	DefaultOptionsChargePrice string
	DefaultOptionsSubstitutionPricing string
	MinSelections int
	MaxSelections int
	RequiredMode string
	IsMultiSelect bool
	PreModifierGroupReference int
	ModifierOptionReferences ints
}

type Metadata struct {
	RestaurantGUID string
	LastUpdated string
}
