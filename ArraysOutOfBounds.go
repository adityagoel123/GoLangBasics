// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var companies = [4]string{"Amazon", "Microsoft", "Yahoo", "Google"}
	fmt.Println("Pre addition, Array of Strings is: ", companies)
	companies[4] = "Adobe"
	fmt.Println("Post addition, Array of Strings is: ", companies)
}

// Above code will NOT WORK.
