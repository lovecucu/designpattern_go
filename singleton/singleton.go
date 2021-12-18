package singleton

import (
	"sync"
)

type ChocolateBoilerInterface interface {
	fill()
	boil()
	drain()
	isEmpty() bool
	isBoiled() bool
}

var (
	ChocolateBoilerFactory ChocolateBoilerInterface
	once                   sync.Once
)

type ChocolateBoiler struct {
	empty, boiled bool
}

func GetChocolateBoilerFactory() ChocolateBoilerInterface {
	once.Do(func() {
		ChocolateBoilerFactory = &ChocolateBoiler{}
	})
	return ChocolateBoilerFactory
}

func (c *ChocolateBoiler) fill() {
	if c.isEmpty() {
		c.empty = false
		c.boiled = false
	}
}

func (c *ChocolateBoiler) boil() {
	if !c.isEmpty() && !c.isBoiled() {
		c.boiled = true
	}
}

func (c *ChocolateBoiler) drain() {
	if !c.isEmpty() && c.isBoiled() {
		c.empty = true
	}
}

func (c ChocolateBoiler) isEmpty() bool {
	return c.empty
}

func (c ChocolateBoiler) isBoiled() bool {
	return c.boiled
}

var _ ChocolateBoilerInterface = (*ChocolateBoiler)(nil)
