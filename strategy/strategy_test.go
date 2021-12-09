package strategy

import "testing"

func TestDuck(t *testing.T) {
	// mallardDuck实例化
	duck := NewMallardDuck()
	duck.Swim()
	duck.Display()
	duck.PerformFly()
	duck.PerformQuack()

	// 改变mallardDuck的行为
	duck.setFlyBehavior(new(FlyNoWay))
	duck.setQuackBehavior(new(Squeak))
	duck.Swim()
	duck.Display()
	duck.PerformFly()
	duck.PerformQuack()
}
