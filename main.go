package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		"Alex",
		"Anderson",
		contactInfo{
			"alex.anderson@golang.com",
			123456,
		},
	} // This approach right here of kind of relying upon the order of definition of fields. This can be catastrophic if we go along and then accidentally swap the order of those two fields
	chantal := person{
		firstName: "Chantal",
		lastName:  "Neo",
		contact: contactInfo{
			email:   "chantal.neo@golang.com",
			zipCode: 5412177,
		},
	} // This approach is better, because it means that we can always change the order of fields in the definition up there
	var jerry person // Third approach of declaring a new struct in Go. At this point, Go would prepopulate the different fields with zero value. Refer to point 2 for more information
	jerry.firstName = "Jerry"
	jerry.lastName = "Gan"
	jerry.contact = contactInfo{
		email:   "jerry.gan@golang.com",
		zipCode: 5201314,
	}

	fmt.Println(alex)
	fmt.Println(chantal)
	fmt.Println(jerry)
	fmt.Printf("%+v", jerry) // Percent plus v right here would help us print out all the different field names and their values from Jerry
}

// Notes:
// 1. Whenever we make a struct we have to first define all of the different properties that a struct might have
//    We're going to provide this in some type of rule set to Go and then we can create a value that matches that type of structure definition.
//
// 2. Type   | Zero Value
//    string | ""
//    int    | 0
//    float  | 0
//    bool   | false
