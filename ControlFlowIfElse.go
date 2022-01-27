package main

import "fmt"

func main() {
	thisVar := 79
	var result string

	if thisVar < 100 {
		result = "Less than HUNDRED"
	} else if thisVar == 100 {
		result = "Equal to HUNDRED"
	} else if thisVar == 100 {
		result = "Equal to HUNDRED"
	}
	fmt.Println("Result thus obtained is: ", result)
}
