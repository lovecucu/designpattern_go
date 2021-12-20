package iteratorcomposite

import "testing"

func TestComposite(t *testing.T) {
	pancake := NewSubMenu("PANCAKE HOUSE MENU", "Breakfast")
	diner := NewSubMenu("DINER HOUSE MENU", "Lunch")
	dessert := NewSubMenu("DESSERT MENU", "Dessert of course!")
	all := NewSubMenu("ALL MENUS", "All menus combined")

	all.add(pancake)
	all.add(diner)

	pancake.add(NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99))
	diner.add(NewMenuItem("Pasta", "Spaghetti with Marinara Sauce, and a slice of sourdough bread", true, 3.89))

	dessert.add(NewMenuItem("Apple Pie", "Apple pie with a flakey crust, topped with vanilla ice cream", true, 1.59))
	diner.add(dessert)

	waitress := NewWaitressNew(all)
	waitress.printMenu()
}
