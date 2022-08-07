package main

import (
	"deepThought"
	"fmt"
)

// main helps us to find the answer to the ultimate question about life, the universe and everything
func main() {
	answer, onlyConstant := deepThought.Ask("life", "the universe", "everything")
	fmt.Printf("Press [Enter] to see the answer!")
	fmt.Scanln()
	fmt.Printf("The answer is %v.\n", answer)
	fmt.Println("BTW: The only constant in life is", onlyConstant)
}
