package main

import (
	"fmt"
)

func main() {
	sumReceived, countOfValues := addInfiniteValues(75, 15, 10, 50)
	fmt.Println("The sum received is ", sumReceived, " And the count of Values added are :", countOfValues)
}
func addInfiniteValues(values ...int) (int, int) {
	sumTotal := 0
	countOfInputsReceived := 0
	for thisIndex, thisVal := range values {
		countOfInputsReceived = thisIndex
		sumTotal += thisVal
	}
	return sumTotal, countOfInputsReceived
}
