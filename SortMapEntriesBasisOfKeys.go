package main

import (
	"fmt"
	"sort"
)

func main() {
	// UnOrdered Collection (MAP) of elements.
	states := make(map[string]string)
	fmt.Println("Elements of the Map looks like: ", states)
	states["HR"] = "Haryana"
	states["DL"] = "Delhi"
	states["MP"] = "Madhya Pradesh"
	states["WB"] = "West Bengal"
	//	fmt.Println("Elements of the Map looks like: ", states)
	onlyKeys := make([]string, len(states))
	index := 0
	for k := range states {
		onlyKeys[index] = k
		index++
	}
	fmt.Println("Before sorting, Keys we have are: ", onlyKeys)
	sort.Strings(onlyKeys)
	fmt.Println("Post Sorting, Keys now looks like: ", onlyKeys)

	for k := range onlyKeys {
		fmt.Println("Key here is: ", onlyKeys[k], " And Value is: ", states[onlyKeys[k]])
	}
}
