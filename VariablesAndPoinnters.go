// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	thisParticularInVar := 75
	var thisPointerVar = &thisParticularInVar

	fmt.Println("Before the change, value of variable is: ", thisParticularInVar)
	fmt.Println("Before the change, address of variable is: ", thisPointerVar)

	*thisPointerVar = (*thisPointerVar) * 2

	fmt.Println("ost change, Value of this variable through Pointer is: ", *thisPointerVar)
	fmt.Println("Post change, value of variable is: ", thisParticularInVar)
}
