package schema

type MenuItem struct {
	Name string
	Price float64
	Description string
	InStock bool
}
type MenuItems []MenuItem

type SubMenu struct {
	Name string
	Items MenuItems
}
type SubMenus []SubMenu

type Location struct {
	Name string
	Menu SubMenus
}
type Locations []Location

type Directory struct {
	Locations Locations
}
