package main

import (
	"fmt"
)

func main() {
	bulldog := Dog{"BullDog", 17, "Bhow !!"}
	fmt.Printf("%+v\n", bulldog)
	bulldog.Speakplease()
	bulldog.Speakplease()
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (thisDog Dog) Speakplease() {
	thisDog.Sound = fmt.Sprintf("%v %v %v", thisDog.Sound, thisDog.Sound, thisDog.Sound)
	fmt.Println("This Dog speaks as: ", thisDog.Sound)
}

func (thisDog Dog) Speak() {
	fmt.Println("This Dog speaks as: ", thisDog.Sound)
}
