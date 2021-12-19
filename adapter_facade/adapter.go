package adapterfacade

import "fmt"

type Duck interface {
	quack()
	fly()
}

type MallardDuck struct{}

func (d MallardDuck) quack() {
	fmt.Println(`Quack`)
}

func (d MallardDuck) fly() {
	fmt.Println(`I'm flying`)
}

var _ Duck = (*MallardDuck)(nil)

type Turkey interface {
	gobble()
	fly()
}

type WildTurkey struct{}

func (t WildTurkey) gobble() {
	fmt.Println(`Gobble gobble`)
}

func (d WildTurkey) fly() {
	fmt.Println(`I'm flying a short distance`)
}

var _ Turkey = (*WildTurkey)(nil)

type TurkeyAdapter struct {
	turkey Turkey
}

func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}

func (d TurkeyAdapter) quack() {
	d.turkey.gobble()
}

func (d TurkeyAdapter) fly() {
	for i := 0; i < 5; i++ {
		d.turkey.fly()
	}
}

var _ Duck = (*TurkeyAdapter)(nil)
