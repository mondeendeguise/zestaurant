package schema

type MenuItem struct {
	Name string
	Price float64
	Description string
	InStock bool
}
type MenuItems []MenuItem

type MenuGroup struct {
	Name string
	Items MenuItems
	SubGroups MenuGroups
}
type MenuGroups []MenuGroup

type Location struct {
	Name string
	Menu MenuGroups
}
type Locations []Location

type Directory struct {
	Locations Locations
}
