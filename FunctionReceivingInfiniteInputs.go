package main

import (
	"fmt"
)

func main() {
	sumReceived := addInfiniteValues(75, 15, 10, 50)
	fmt.Println("The sum received is ", sumReceived)
}
func addInfiniteValues(values ...int) int {
	sumTotal := 0
	for _, thisVal := range values {
		sumTotal += thisVal
	}
	return (sumTotal)
}
