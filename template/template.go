package template

import "fmt"

// 示例场景：咖啡因饮料制作，茶和咖啡

type CaffeineMake interface {
	brew()
	addCondiments()
	customerWantsCondiments() bool
}

type CaffeineBeverage struct {
	cm CaffeineMake
}

func (cb *CaffeineBeverage) prepareRecipe() { // TODO: 理论上，该模板方法是不允许子类重写的，但go在这块没有好方法，不知咋解决
	cb.boilWater()
	cb.cm.brew()
	cb.pourInCup()
	if cb.cm.customerWantsCondiments() { // 该方法为hook，可由子类自由选择是否添加调料
		cb.cm.addCondiments()
	}
}

func (cb *CaffeineBeverage) boilWater() {
	fmt.Println(`Boiling water`)
}
func (cb *CaffeineBeverage) pourInCup() {
	fmt.Println(`Pouring into cup`)
}
func (cb *CaffeineBeverage) customerWantsCondiments() bool {
	return true
}

type Tea struct {
	CaffeineBeverage
}

func NewTea() *Tea {
	s := &Tea{}
	s.CaffeineBeverage.cm = s
	return s
}

func (c *Tea) brew() {
	fmt.Println(`Steeping the tea`)
}

func (c *Tea) addCondiments() {
	fmt.Println(`Adding lemon`)
}

type TeaWithNoCondiments struct {
	CaffeineBeverage
}

func NewTeaWithNoCondiments() *TeaWithNoCondiments {
	s := &TeaWithNoCondiments{}
	s.CaffeineBeverage.cm = s
	return s
}

func (c *TeaWithNoCondiments) brew() {
	fmt.Println(`Steeping the tea`)
}

func (c *TeaWithNoCondiments) addCondiments() {
	fmt.Println(`Adding nothing`)
}

func (c *TeaWithNoCondiments) customerWantsCondiments() bool {
	return false
}

type Coffee struct {
	CaffeineBeverage
}

func NewCoffee() *Coffee {
	s := &Coffee{}
	s.CaffeineBeverage.cm = s
	return s
}
func (c *Coffee) brew() {
	fmt.Println(`Dripping Coffee through filter`)
}

func (c *Coffee) addCondiments() {
	fmt.Println(`Adding Sugar and Milk`)
}

var _ CaffeineMake = (*Coffee)(nil)
var _ CaffeineMake = (*Tea)(nil)
var _ CaffeineMake = (*TeaWithNoCondiments)(nil)
