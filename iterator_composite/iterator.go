package iteratorcomposite

import "fmt"

type MenuItem struct {
	name, description string
	vegetarian        bool
	price             float64
}

func (m MenuItem) getName() string {
	return m.name
}
func (m MenuItem) getDescription() string {
	return m.name
}
func (m MenuItem) getPrice() float64 {
	return m.price
}
func (m MenuItem) isVegetarian() bool {
	return m.vegetarian
}

type Iterator interface {
	hasNext() bool
	Next() *MenuItem
}

type PancakeHouseIterator struct {
	position int
	items    []*MenuItem
}

func (p PancakeHouseIterator) hasNext() bool {
	if p.position >= len(p.items) {
		return false
	} else {
		return true
	}
}

func (p *PancakeHouseIterator) Next() *MenuItem {
	menuItem := p.items[p.position]
	p.position++
	return menuItem
}

type DinerMenuIterator struct {
	position int
	items    [4]*MenuItem
}

func (p DinerMenuIterator) hasNext() bool {
	if p.position >= len(p.items) || p.items[p.position] == nil {
		return false
	} else {
		return true
	}
}

func (p *DinerMenuIterator) Next() *MenuItem {
	menuItem := p.items[p.position]
	p.position++
	return menuItem
}

type Menu interface {
	createInterator() Iterator
}

type PancakeHouseMenu struct {
	menuItems []*MenuItem
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menuItems := []*MenuItem{
		&MenuItem{"K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99},
		&MenuItem{"Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99},
		&MenuItem{"Blueberry Pancake", "Pancakes made with fresh blueberres", true, 3.49},
		&MenuItem{"Waffles", "Waffles, with your choice of blueberries or strawberries", true, 3.59},
	}
	return &PancakeHouseMenu{menuItems: menuItems}
}

func (m PancakeHouseMenu) createInterator() Iterator {
	return &PancakeHouseIterator{items: m.menuItems}
}

var _ Menu = (*PancakeHouseMenu)(nil)

type DinerMenu struct {
	menuItems [4]*MenuItem
}

func NewDinerMenu() *DinerMenu {
	menuItems := [4]*MenuItem{
		&MenuItem{"Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99},
		&MenuItem{"BLT", "Bacon with lettuce & tomato on whole whea", false, 2.99},
		&MenuItem{"Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29},
		&MenuItem{"Hotdog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05},
	}
	return &DinerMenu{menuItems: menuItems}
}

func (m DinerMenu) createInterator() Iterator {
	return &DinerMenuIterator{items: m.menuItems}
}

var _ Menu = (*DinerMenu)(nil)

type Waitress struct {
	panmenu, dinermenu Menu
}

func NewWaitress(panmenu, dinermenu Menu) *Waitress {
	return &Waitress{panmenu: panmenu, dinermenu: dinermenu}
}

func (w Waitress) printMenu() {
	panIterator := w.panmenu.createInterator()
	dinerIterator := w.dinermenu.createInterator()
	fmt.Print("MENU\n----\nBREAKFAST\n")
	w.print(panIterator, false)
	fmt.Print("\nLUNCH\n")
	w.print(dinerIterator, false)
}

func (w Waitress) printVegetarion() {
	panIterator := w.panmenu.createInterator()
	dinerIterator := w.dinermenu.createInterator()
	fmt.Print("MENU\n----\nBREAKFAST\n")
	w.print(panIterator, true)
	fmt.Print("\nLUNCH\n")
	w.print(dinerIterator, true)
}

func (w Waitress) print(iterator Iterator, vegetarian bool) {
	for iterator.hasNext() {
		menuItem := iterator.Next()
		if !vegetarian || menuItem.isVegetarian() {
			fmt.Printf("%s, %.2f -- %s\n", menuItem.getName(), menuItem.getPrice(), menuItem.getDescription())
		}
	}
}
