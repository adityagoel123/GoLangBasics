package main

import (
	"fmt"
)

func main() {

	bulldog := Dog{"BullDog", 17, "Buzo"}
	fmt.Printf("%+v\n", bulldog)
	fmt.Println("Dog we have is: ", bulldog)
	fmt.Printf("Breed is %v AND Weight is %v\n", bulldog.Breed, bulldog.Weight)

}

type Dog struct {
	Breed  string
	Weight int
	Name   string
}
