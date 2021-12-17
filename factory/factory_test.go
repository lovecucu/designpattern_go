package factory

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	store := NewPizzaStore(SimplePizzaFactory{})
	store.OrderPizza(CHEESE)
	store.OrderPizza(CLAM)
}

func TestFuncFactory(t *testing.T) {
	store1 := GetPizzaStoreByStyle(NY)
	store1.OrderPizza(CHEESE)
	store2 := GetPizzaStoreByStyle(CHICAGO)
	store2.OrderPizza(CLAM)
}
