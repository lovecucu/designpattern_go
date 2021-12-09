package strategy

import "fmt"

type FlyBehavior interface {
	fly()
}

type FlyWithWings struct{}

func (f FlyWithWings) fly() {
	fmt.Println("fly: with wings")
}

type FlyNoWay struct{}

func (f FlyNoWay) fly() {
	fmt.Println("fly: no way")
}

type QuackBehavior interface {
	quack()
}

type Quack struct{}

func (q Quack) quack() {
	fmt.Println("quack: quack")
}

type Squeak struct{}

func (q Squeak) quack() {
	fmt.Println("quack: squeak")
}

type MuteQuack struct{}

func (q MuteQuack) quack() {
	fmt.Println("quack: mutequack")
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) Swim() {
	fmt.Println("swim: swim")
}

func (d *Duck) Display() {
	fmt.Println("display: duck")
}

func (d *Duck) PerformFly() {
	d.flyBehavior.fly()
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.quack()
}

func (d *Duck) setFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) setQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}

type MallardDuck struct {
	Duck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{Duck{flyBehavior: new(FlyWithWings), quackBehavior: new(Quack)}}
}

func (d MallardDuck) Display() {
	fmt.Println("display: MallardDuck")
}
