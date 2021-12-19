package adapterfacade

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	duck := &MallardDuck{}
	turkey := &WildTurkey{}
	turkeyAdapter := NewTurkeyAdapter(turkey)
	fmt.Print("\nThe Turkey says...\n")
	turkey.gobble()
	turkey.fly()

	fmt.Print("\nThe Duck says...\n")
	testDuck(duck)

	fmt.Print("\nThe TurkeyAdapter says...\n")
	testDuck(turkeyAdapter)
}

func testDuck(duck Duck) {
	duck.quack()
	duck.fly()
}
