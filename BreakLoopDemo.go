package main

import (
	"fmt"
)

func main() {
	thisVar := 1

	for thisVar < 10 {
		fmt.Println("Value of thisVar variable now is: ", thisVar)
		if thisVar == 5 {
			break
		}
		thisVar++
	}
}
