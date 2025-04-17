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
