package singleton

import "testing"

func TestSingleton(t *testing.T) {
	cb1 := GetChocolateBoilerFactory()
	cb2 := GetChocolateBoilerFactory()

	if cb1 != cb2 {
		t.Error(`singleton is not working`)
	}
	cb1.fill()
	if cb2.isEmpty() {
		t.Error(`singleton is not working`)
	}
	cb2.boil()
	if cb1.isEmpty() || !cb1.isBoiled() {
		t.Error(`singleton is not working`)
	}
	cb1.drain()
	if !cb2.isEmpty() || !cb1.isBoiled() {
		t.Error(`singleton is not working`)
	}
}
