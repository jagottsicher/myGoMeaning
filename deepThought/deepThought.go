// deepThought serves the answer to the the question about life, the universe and everything
package deepThought

import "fmt"

// question is a type with an underlying struct, of three elements of type string
type question struct {
	thing1 string
	thing2 string
	thing3 string
}

// InLife is the only constant in life
const InLife string = "change"

// Uqlue is a variable of type question holding strings asked for
var Uqlue question

// Ask checks, if three elements asked for are identical the one stored
//
// It returns 42 in case if the the strings are identical, else 0
func Ask(life, universe, everything string) (int8, string) {

	Uqlue = question{
		thing1: "life",
		thing2: "universe",
		thing3: "everything",
	}

	fmt.Println("What is the answer to the ultimate question of %v, %v and %v?", life, universe, everything)

	if life == Uqlue.thing1 && universe == Uqlue.thing2 && everything == Uqlue.thing3 {
		return 42, InLife
	}

	return 0, "nothing"
}
