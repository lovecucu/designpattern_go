package factory

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	store := NewPizzaStore(SimplePizzaFactory{})
	store.OrderPizza(CHEESE)
	fmt.Println()
	store.OrderPizza(CLAM)
	fmt.Println()
}

func TestFuncFactory(t *testing.T) {
	store1 := GetPizzaStoreByStyle(NY)
	store1.OrderPizza(CHEESE)
	fmt.Println()
	store2 := GetPizzaStoreByStyle(CHICAGO)
	store2.OrderPizza(CLAM)
	fmt.Println()
}
