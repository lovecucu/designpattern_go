package iteratorcomposite

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	gumball := NewGumballMachine(5)

	gumball.insertQuarter()
	gumball.turnCrank()
	fmt.Println()
	gumball.insertQuarter()
	gumball.turnCrank()
	fmt.Println()
	gumball.insertQuarter()
	gumball.turnCrank()
}
