// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

// Below code will NOT work.
func main() {

	var intArrays [5]int
	intArrays[0] = 98
	intArrays[1] = 75
	intArrays[2] = 81
	intArrays[3] = 33
	fmt.Println("Before sorting, Array of Integers is: ", intArrays)
	sort.Ints(intArrays)
	fmt.Println("After sorting, Array of Integers is: ", intArrays)
}
