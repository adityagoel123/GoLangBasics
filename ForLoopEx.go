package main
import (
	"fmt"

func main() {
	states := []string{"Haryana", "Delhi", "Meghalaya", "Manipur", "Tripura"}
	fmt.Println(states)

	fmt.Println("For Loop using range way:")
	for i := range states {
		fmt.Println(states[i])
	}

	fmt.Println("For Loop using traditional way:")
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
}
