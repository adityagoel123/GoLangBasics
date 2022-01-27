// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	thisArrayUsingMake := make([]int, 5)

	thisArrayUsingMake[0] = 91
	thisArrayUsingMake[1] = 89
	thisArrayUsingMake[2] = 75
	thisArrayUsingMake[3] = 67
	thisArrayUsingMake[4] = 55
	fmt.Println("Elements of the Array looks like: ", thisArrayUsingMake)
	thisArrayUsingMake = append(thisArrayUsingMake, 312)
	fmt.Println("Elements of the Array now are: ", thisArrayUsingMake)
	print("Address allocated to the Array is: ", thisArrayUsingMake)
}
