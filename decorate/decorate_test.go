package observer

import (
	"fmt"
	"testing"
)

func TestDecorate(t *testing.T) {
	es := NewEspresso()
	fmt.Printf("desc: %s cost: %.2f\n", es.GetDesc(), es.Cost())
	hb := NewHouseBlend()
	fmt.Printf("desc: %s cost: %.2f\n", hb.GetDesc(), hb.Cost())

	esMochaWhip := NewWhip(NewMocha(es))
	fmt.Printf("desc: %s cost: %.2f\n", esMochaWhip.GetDesc(), esMochaWhip.Cost())

	hbWhip := NewWhip(hb)
	fmt.Printf("desc: %s cost: %.2f\n", hbWhip.GetDesc(), hbWhip.Cost())
}
