package adapterfacade

import "testing"

func TestFacade(t *testing.T) {
	homeTheater := NewHomeTheater(&amplifier{}, &lights{}, &dvd{}, &projector{}, &screen{}, &popper{})
	homeTheater.watchMovie("Raiders of the Lost Ark")
	homeTheater.endMovie()
}
