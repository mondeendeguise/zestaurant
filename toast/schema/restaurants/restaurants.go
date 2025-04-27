package restaurants

type Restaurant struct {
	GUID string
}

type General struct {
	Name string
	LocationName string
	LocationCode string
	Description string
	TimeZone string
	CloseoutHour int
	ManagementGroupGUID string
	CurrencyCode string
	FirstBusinessDate int
	Archived bool
}

type URLs struct {
	Website string
	Facebook string
	Twitter string
	OrderOnline string
	PurchaseGiftCard string
	CheckGiftCard string
}

type Location struct {
	Address1 string
	Address2 string
	City string
	StateCode string
	AdministrativeArea string
	ZipCode string
	Country string
	Phone string
	PhoneCountryCode string
	Latitude float64
	Longitude float64
}

type Hours struct {
	StartTime string
	EndTime string
}

type Services []Service
type Service struct {
	Name string
	Hours Hours
	Overnight bool
}

type DaySchedule struct {
	ScheduleName string
	Services Services
	OpenTime string
	CloseTime string
}

type WeekSchedule struct {
	Monday string
	Tuesday string
	Wednesday string
	Thursday string
	Friday string
	Saturday string
	Sunday string
}

type Schedules struct {
	DaySchedules map[string]DaySchedule
	WeekSchedule WeekSchedule
}

type Delivery struct {
	Enabled bool
	Minimum float64
	Area string
}

type DeliveryPaymentOptions struct {
	Cash bool
	CCSameDay bool
	CCFuture bool
}

type TakeoutPaymentOptions struct {
	Cash bool
	CCSameDay bool
	CCFuture bool
	CCInStore bool
}

type PaymentOptions struct {
	Delivery DeliveryPaymentOptions
	Takeout TakeoutPaymentOptions
	CCTip bool
}

type OnlineOrdering struct {
	Enabled bool
	Scheduling bool
	SpecialRequests bool
	SpecialRequestsMessage string
	PaymentOptions PaymentOptions
}

type PrepTimes struct {
	DeliveryPrepTime int32
	DeliveryTimeAfterOpen int32
	DeliveryTimeBeforeClose int32
	TakeoutPrepTime int32
	TakeoutTimeAfterOpen int32
	TakeoutTimeBeforeClose int32
	TakeoutThrottlingTime int
	DeliveryThrottlingTime int
}

type RestaurantInfo struct {
	GUID string
	General General
	URLs URLs
	Location Location
	Schedules Schedules
	Delivery Delivery
	OnlineOrdering OnlineOrdering
	PrepTimes PrepTimes
}
