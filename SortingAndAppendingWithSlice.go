// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {

	var intArrays = []int{98, 76, 59, 63, 85}
	fmt.Println("Appending an Integer in the end, Current Array: ", intArrays)
	intArrays = append(intArrays, 99)
	fmt.Println("Before sorting, Array of Integers is: ", intArrays)
	sort.Ints(intArrays)
	fmt.Println("After sorting, Array of Integers is: ", intArrays)
}
