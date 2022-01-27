package main

import (
	"fmt"
)

func main() {
	sum := addValues(75, 15)
	fmt.Println("The sum received is ", sum)
}

func addValues(val1 int, val2 int) int {
	return (val1 + val2)

}
