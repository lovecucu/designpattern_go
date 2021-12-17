package factory

import "fmt"

type PizzaInterface interface {
	prepare()
	bake()
	cut()
	box()
	GetName() string
}

type PizzaType int

const (
	_                = iota
	CHEESE PizzaType = 1
	CLAM             = 2
)

type Pizza struct {
	name, dough, sauce string
	toppings           []string
}

func (p Pizza) prepare() {
	fmt.Println("Preparing " + p.name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings: ")
	for i := 0; i < len(p.toppings); i++ {
		fmt.Println("\t" + p.toppings[i])
	}
}

func (p Pizza) bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p Pizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p Pizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p Pizza) GetName() string {
	return p.name
}

type CheesePizza struct {
	Pizza
}

func NewCheesePizza() *CheesePizza {
	return &CheesePizza{Pizza: Pizza{name: "cheese"}}
}

type ClamPizza struct {
	Pizza
}

func NewClamPizza() *ClamPizza {
	return &ClamPizza{Pizza: Pizza{name: "clam"}}
}

type PizzaCreate interface {
	CreatePizza(t PizzaType) PizzaInterface
}

type PizzaStoreInterface interface {
	OrderPizza(t PizzaType) PizzaInterface
	CreatePizza(t PizzaType) PizzaInterface
}

// 简单工厂
type SimplePizzaFactory struct{}

func (s SimplePizzaFactory) CreatePizza(t PizzaType) PizzaInterface {
	var pizza PizzaInterface
	switch t {
	case CHEESE:
		pizza = NewCheesePizza()
	case CLAM:
		pizza = NewClamPizza()
	}
	return pizza
}

var _ PizzaCreate = (*SimplePizzaFactory)(nil)

type PizzaStore struct {
	factory PizzaCreate
}

func NewPizzaStore(factory PizzaCreate) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (ps PizzaStore) OrderPizza(t PizzaType) PizzaInterface {
	pizza := ps.factory.CreatePizza(t)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return pizza
}

// 工厂方法

type NYCheesePizza struct {
	Pizza
}

func NewNYCheesePizza() *NYCheesePizza {
	topping := []string{}
	topping = append(topping, "Grated Reggiano Cheese")
	return &NYCheesePizza{Pizza: Pizza{name: "NY Style Sauce and Cheesse Pizza", dough: "Thin Crust Dough", sauce: "Marinara Sauce", toppings: topping}}
}

type NYClamPizza struct {
	Pizza
}

func NewNYClamPizza() *NYClamPizza {
	topping := []string{}
	topping = append(topping, "Grated Reggiano Cheese")
	return &NYClamPizza{Pizza: Pizza{name: "NY Style Sauce and Clam Pizza", dough: "Thin Crust Dough", sauce: "Marinara Sauce", toppings: topping}}
}

type ChicagoCheesePizza struct {
	Pizza
}

func NewChicagoCheesePizza() *ChicagoCheesePizza {
	topping := []string{}
	topping = append(topping, "Shredded Mozzarella Cheese")
	return &ChicagoCheesePizza{Pizza: Pizza{name: "Chicago Style Deep Dish Cheesse Pizza", dough: "Extra Thick Crust Dough", sauce: "Plum Tomato Sauce", toppings: topping}}
}

type ChicagoClamPizza struct {
	Pizza
}

func NewChicagoClamPizza() *ChicagoClamPizza {
	topping := []string{}
	topping = append(topping, "Shredded Mozzarella Cheese")
	return &ChicagoClamPizza{Pizza: Pizza{name: "Chicago Style Deep Dish Clam Pizza", dough: "Extra Thick Crust Dough", sauce: "Plum Tomato Sauce", toppings: topping}}
}

type PizzaStoreFunc struct {
	pc PizzaCreate
}

func (ps PizzaStoreFunc) OrderPizza(t PizzaType) PizzaInterface {
	pizza := ps.pc.CreatePizza(t)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return pizza
}

type NYPizzaStoreFunc struct {
	PizzaStoreFunc
}

func (ps NYPizzaStoreFunc) CreatePizza(t PizzaType) PizzaInterface {
	var pizza PizzaInterface
	switch t {
	case CHEESE:
		pizza = NewNYCheesePizza()
	case CLAM:
		pizza = NewNYClamPizza()
	}
	return pizza
}

func NewNYPizzaStoreFunc() PizzaStoreInterface {
	s := new(NYPizzaStoreFunc)
	s.PizzaStoreFunc.pc = s
	return s
}

type ChicagoPizzaStoreFunc struct {
	PizzaStoreFunc
}

func (ps ChicagoPizzaStoreFunc) CreatePizza(t PizzaType) PizzaInterface {
	var pizza PizzaInterface
	switch t {
	case CHEESE:
		pizza = NewChicagoCheesePizza()
	case CLAM:
		pizza = NewChicagoClamPizza()
	}
	return pizza
}

func NewChicagoPizzaStoreFunc() PizzaStoreInterface {
	s := new(ChicagoPizzaStoreFunc)
	s.PizzaStoreFunc.pc = s
	return s
}

var _ PizzaStoreInterface = (*NYPizzaStoreFunc)(nil)
var _ PizzaStoreInterface = (*ChicagoPizzaStoreFunc)(nil)

type PizzaStyle int

const (
	_                  = iota
	NY      PizzaStyle = 1
	CHICAGO            = 2
)

func GetPizzaStoreByStyle(s PizzaStyle) PizzaStoreInterface {
	var pizzaStore PizzaStoreInterface
	switch s {
	case NY:
		pizzaStore = NewNYPizzaStoreFunc()
	case CHICAGO:
		pizzaStore = NewChicagoPizzaStoreFunc()
	}
	return pizzaStore
}
