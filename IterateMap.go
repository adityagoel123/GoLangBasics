// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// UnOrdered Collection (MAP) of elements.
	states := make(map[string]string)
	fmt.Println("Elements of the Map looks like: ", states)

	states["HR"] = "Haryana"
	states["DL"] = "Delhi"
	states["MP"] = "Madhya Pradesh"
	states["WB"] = "West Bengal"
	//	fmt.Println("Elements of the Map looks like: ", states)

	for k, v := range states {
		fmt.Printf("This Key: %v and val: %v\n", k, v)
	}
}
