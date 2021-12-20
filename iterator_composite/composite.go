package iteratorcomposite

import "fmt"

type MenuComponent interface {
	add(MenuComponent)
	remove(MenuComponent)
	getName() string
	getDescription() string
	getPrice() float64
	isVegetarian() bool
	print()
}

type MenuComponentBase struct{}

func (mc MenuComponentBase) add(m MenuComponent)    {}
func (mc MenuComponentBase) remove(m MenuComponent) {}
func (mc MenuComponentBase) getName() string {
	return ""
}
func (mc MenuComponentBase) getDescription() string {
	return ""
}
func (mc MenuComponentBase) getPrice() float64 {
	return 0
}
func (mc MenuComponentBase) isVegetarian() bool {
	return false
}
func (mc MenuComponentBase) print() {}

var _ MenuComponent = (*MenuComponentBase)(nil)

type MenuItemNew struct {
	MenuComponentBase
	name, description string
	vegetatian        bool
	price             float64
}

func NewMenuItem(name, desc string, vegetatian bool, price float64) *MenuItemNew {
	return &MenuItemNew{name: name, description: desc, vegetatian: vegetatian, price: price}
}

func (mt MenuItemNew) getName() string {
	return mt.name
}

func (mt MenuItemNew) getDescription() string {
	return mt.description
}

func (mt MenuItemNew) getPrice() float64 {
	return mt.price
}

func (mt MenuItemNew) isVegetarian() bool {
	return mt.vegetatian
}

func (mt MenuItemNew) print() {
	fmt.Print(" " + mt.getName())
	if mt.isVegetarian() {
		fmt.Print("(v)")
	}
	fmt.Printf(", %.2f", mt.getPrice())
	fmt.Printf("    -- %s\n", mt.getDescription())
}

type SubMenu struct {
	MenuComponentBase
	components        []MenuComponent
	name, description string
}

func NewSubMenu(name, desc string) *SubMenu {
	return &SubMenu{components: make([]MenuComponent, 0), name: name, description: desc}
}

func (sm *SubMenu) add(mc MenuComponent) {
	sm.components = append(sm.components, mc)
}

func (sm *SubMenu) remove(mc MenuComponent) {
	if len(sm.components) <= 0 {
		return
	}
	index := -1
	for i := 0; i < len(sm.components); i++ {
		if sm.components[i] == mc {
			index = i
			break
		}
	}
	if index >= 0 {
		sm.components[index], sm.components[len(sm.components)] = sm.components[len(sm.components)], sm.components[index]
	}
	sm.components = sm.components[0 : len(sm.components)-1]
}

func (sm SubMenu) getName() string {
	return sm.name
}

func (sm SubMenu) getDescription() string {
	return sm.description
}

func (sm SubMenu) print() {
	fmt.Printf("\n%s, %s\n", sm.getName(), sm.getDescription())
	fmt.Println(`----------------------`)
	for i := 0; i < len(sm.components); i++ {
		sm.components[i].print()
	}
}

type WaitressNew struct {
	allMenus MenuComponent
}

func NewWaitressNew(all MenuComponent) *WaitressNew {
	return &WaitressNew{allMenus: all}
}

func (w WaitressNew) printMenu() {
	w.allMenus.print()
}
