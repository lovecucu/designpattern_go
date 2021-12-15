package observer

type BeverageInterface interface {
	GetDesc() string
	Cost() float64
}

type Beverage struct {
	desc string
}

func (b Beverage) GetDesc() string {
	return b.desc
}

func (b Beverage) Cost() float64 {
	return 0.00
}

var _ BeverageInterface = (*Beverage)(nil)

type Espresso struct {
	Beverage
}

func NewEspresso() *Espresso {
	return &Espresso{Beverage: Beverage{desc: "Espresso"}}
}

func (e Espresso) Cost() float64 {
	return 1.99
}

var _ BeverageInterface = (*Espresso)(nil)

type HouseBlend struct {
	Beverage
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{Beverage: Beverage{desc: "HouseBlend"}}
}

func (e HouseBlend) Cost() float64 {
	return 0.89
}

var _ BeverageInterface = (*HouseBlend)(nil)

type CondimentDecorator struct {
	Beverage
}

type Mocha struct {
	CondimentDecorator
	beverage BeverageInterface
}

func NewMocha(beverage BeverageInterface) *Mocha {
	return &Mocha{beverage: beverage}
}

func (m Mocha) GetDesc() string {
	return m.beverage.GetDesc() + ", Mocha"
}

func (m Mocha) Cost() float64 {
	return 0.20 + m.beverage.Cost()
}

var _ BeverageInterface = (*Mocha)(nil)

type Whip struct {
	CondimentDecorator
	beverage BeverageInterface
}

func NewWhip(beverage BeverageInterface) *Whip {
	return &Whip{beverage: beverage}
}

func (m Whip) GetDesc() string {
	return m.beverage.GetDesc() + ", Whip"
}

func (m Whip) Cost() float64 {
	return 0.10 + m.beverage.Cost()
}

var _ BeverageInterface = (*Whip)(nil)
