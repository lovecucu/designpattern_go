package iteratorcomposite

import (
	"testing"
)

func TestIterator(t *testing.T) {
	w := NewWaitress(NewPancakeHouseMenu(), NewDinerMenu())
	w.printMenu()
	w.printVegetarion()
}
