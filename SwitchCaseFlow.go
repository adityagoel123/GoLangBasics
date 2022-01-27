package main

import (
	"fmt"
)

func main() {
	var result string
	//rand.Seed(time.Now().Unix())
	thisVar := 1
	switch thisVar {
	case 1:
		result = "Its ONE"
		fallthrough
	case 2:
		result = "Its TWO"
		fallthrough
	case 3:
		result = "Its THREE"
		fallthrough
	case 4:
		result = "Its FOUR"
		fallthrough
	default:
		result = "Its DEFAULT case"
	}
	fmt.Println("Value thus obtained is: ", result)
}
