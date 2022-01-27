// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var companies = [4]string{"Amazon", "Microsoft", "Yahoo", "Google"}
	fmt.Println("Array of Strings is: ", companies)

	var intArrays [5]int
	intArrays[0] = 78
	intArrays[1] = 85
	intArrays[2] = 91
	intArrays[3] = 93
	intArrays[4] = 97

	fmt.Println("Array of Integers is: ", intArrays)

	fmt.Println("Size of IntegersArray is: ", len(intArrays))
	fmt.Println("Size of StringArray is: ", len(companies))
}
