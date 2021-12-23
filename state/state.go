package state

import (
	"fmt"
	"math/rand"
)

// 场景：自动售卖机，投币、退币、出货、售空等状态之间的切换

type State interface {
	insertQuarter()
	ejectQuarter()
	turnCrank()
	dispense()
}

type GumballMachine struct {
	soldOutState    State
	noQuarterState  State
	hasQuarterState State
	soldState       State
	winState        State

	state State
	count int
}

func NewGumballMachine(balls int) *GumballMachine {
	machine := GumballMachine{}
	machine.soldOutState = &SoldOutState{&machine}
	machine.noQuarterState = &NoQuarterState{&machine}
	machine.hasQuarterState = &HasQuarterState{&machine}
	machine.soldState = &SoldState{&machine}
	machine.winState = &WinState{&machine}
	machine.count = balls
	if balls > 0 {
		machine.state = machine.noQuarterState
	} else {
		machine.state = machine.soldOutState
	}
	return &machine
}

func (b *GumballMachine) setState(state State) {
	b.state = state
}

func (b *GumballMachine) insertQuarter() {
	b.state.insertQuarter()
}

func (b *GumballMachine) ejectQuarter() {
	b.state.ejectQuarter()
}

func (b *GumballMachine) turnCrank() {
	b.state.turnCrank()
	b.state.dispense()
}

func (b *GumballMachine) releaseBall() {
	fmt.Println(`A gumball comes rolling out the slot...`)
	if b.count > 0 {
		b.count--
	}
}

func (b *GumballMachine) getCount() int {
	return b.count
}

func (b *GumballMachine) getSoldOutState() State {
	return b.soldOutState
}

func (b *GumballMachine) getNoQuarterState() State {
	return b.noQuarterState
}

func (b *GumballMachine) getHasQuarterState() State {
	return b.hasQuarterState
}

func (b *GumballMachine) getSoldState() State {
	return b.soldState
}

func (b *GumballMachine) getWinState() State {
	return b.winState
}

type NoQuarterState struct {
	machine *GumballMachine
}

func (s *NoQuarterState) insertQuarter() {
	fmt.Println(`You inserted a quarter`)
	s.machine.setState(s.machine.getHasQuarterState())
}

func (s *NoQuarterState) ejectQuarter() {
	fmt.Println(`You haven't inserted a quarter`)
}

func (s *NoQuarterState) turnCrank() {
	fmt.Println(`You turned, but there's no quarter`)
}

func (s *NoQuarterState) dispense() {
	fmt.Println(`You need to pay first`)
}

var _ State = (*NoQuarterState)(nil)

type HasQuarterState struct {
	machine *GumballMachine
}

func (s *HasQuarterState) insertQuarter() {
	fmt.Println(`You can't insert another quarter`)
}

func (s *HasQuarterState) ejectQuarter() {
	fmt.Println(`You haven't inserted a quarter`)
	s.machine.setState(s.machine.getNoQuarterState())
}

func (s *HasQuarterState) turnCrank() {
	fmt.Println(`You turned...`)
	random := rand.Intn(3)
	fmt.Printf("random is %d\n", random)
	if random == 0 && s.machine.getCount() > 1 {
		s.machine.setState(s.machine.getWinState())
	} else {
		s.machine.setState(s.machine.getSoldState())
	}
}

func (s *HasQuarterState) dispense() {
	fmt.Println(`No gumball dispensed`)
}

var _ State = (*HasQuarterState)(nil)

type SoldState struct {
	machine *GumballMachine
}

func (s *SoldState) insertQuarter() {
	fmt.Println(`Please wait, we're already giving you a gumall`)
}

func (s *SoldState) ejectQuarter() {
	fmt.Println(`Sorry, you already turned ther crank`)
}

func (s *SoldState) turnCrank() {
	fmt.Println(`Turning twice doesn't get you another gumball!`)
}

func (s *SoldState) dispense() {
	s.machine.releaseBall()
	if s.machine.getCount() > 0 {
		s.machine.setState(s.machine.getNoQuarterState())
	} else {
		fmt.Println(`Oops, out of gumalls!`)
		s.machine.setState(s.machine.getSoldOutState())
	}
}

var _ State = (*SoldState)(nil)

type SoldOutState struct {
	machine *GumballMachine
}

func (s *SoldOutState) insertQuarter() {
	fmt.Println(`You can't insert a quarter, the machine is sold out`)
}

func (s *SoldOutState) ejectQuarter() {
	fmt.Println(`You can't eject, you haven't inserted a quarter yet`)
}

func (s *SoldOutState) turnCrank() {
	fmt.Println(`You turned, but there are no gumballs`)
}

func (s *SoldOutState) dispense() {
	fmt.Println(`No gumball dispensed`)
}

var _ State = (*SoldOutState)(nil)

type WinState struct {
	machine *GumballMachine
}

func (s *WinState) insertQuarter() {
	fmt.Println(`Please wait, we're already giving you a gumall`)
}

func (s *WinState) ejectQuarter() {
	fmt.Println(`Sorry, you already turned ther crank`)
}

func (s *WinState) turnCrank() {
	fmt.Println(`Turning twice doesn't get you another gumball!`)
}

func (s *WinState) dispense() {
	fmt.Println(`YOU'RE A WINNER! You get two gumballs for your quarter`)
	s.machine.releaseBall()
	if s.machine.getCount() == 0 {
		s.machine.setState(s.machine.getSoldOutState())
	} else {
		s.machine.releaseBall()
		if s.machine.getCount() > 0 {
			s.machine.setState(s.machine.getNoQuarterState())
		} else {
			fmt.Println(`Oops, out of gumalls!`)
			s.machine.setState(s.machine.getNoQuarterState())
		}
	}
}

var _ State = (*WinState)(nil)
