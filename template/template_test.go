package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	cb1 := NewTea()
	fmt.Println("------ prepare Tea ------")
	cb1.prepareRecipe()
	fmt.Printf("------ Tea is ready ------\n\n")

	cb2 := NewTeaWithNoCondiments()
	fmt.Println("------ prepare TeaWithNoCondiments ------")
	cb2.prepareRecipe()
	fmt.Printf("------ TeaWithNoCondiments is ready ------\n\n")

	cb3 := NewCoffee()
	fmt.Println("------ prepare Coffee ------")
	cb3.prepareRecipe()
	fmt.Printf("------ Coffee is ready ------\n\n")
}
