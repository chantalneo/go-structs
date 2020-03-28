package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{"Alex", "Anderson"}                       // This approach right here of kind of relying upon the order of definition of fields. This can be catastrophic if we go along and then accidentally swap the order of those two fields
	chantal := person{firstName: "Chantal", lastName: "Neo"} // This approach is better, because it means that we can always change the order of fields in the definition up there

	fmt.Println(alex)
	fmt.Println(chantal)
}

// Notes:
// 1. Whenever we make a struct we have to first define all of the different properties that a struct might have
//    We're going to provide this in some type of rule set to Go and then we can create a value that matches that type of structure definition.
