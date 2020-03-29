package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // Different syntax from before, but this declares a field name of contactInfo and it also says that it's supposed to have a type contactInfo. This is equivalent to writing: contactInfo contactInfo
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
		contactInfo: contactInfo{
			email:   "chantal.neo@golang.com",
			zipCode: 5412177,
		},
	} // This approach is better, because it means that we can always change the order of fields in the definition up there
	var jerry person // Third approach of declaring a new struct in Go. At this point, Go would prepopulate the different fields with zero value. Refer to point 2 for more information
	jerry.firstName = "Jerry"
	jerry.lastName = "Gan"
	jerry.contactInfo = contactInfo{
		email:   "jerry.gan@golang.com",
		zipCode: 5201314,
	}

	alex.print()
	chantal.print()
	jerry.print()

	alex.updateName("Alec")
	alex.print() // Although we attempted to update Alex's first name to Alec, the update didn't seem to take effect. Explained in notes' point 3
}

func (pointerToPerson *person) updateName(newFirstName string) { // Check out notes' point 5
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p) // Percent plus v right here would help us print out all the different field names and their values from variable p
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
//
// 3. Sweeping broad strokes overview on how RAM works
//    Address | Value
//    0000    |
//    0001    |
//    0002    |
//    0003    |
//    0004    |
// 	  Memory on your local machine can be thought of as like a bunch of little coveys or a bunch of little slots or a bunch of little boxes
//    Each box in your computer's memory can store some data. And each one of these little boxes there are these little value containers has some discrete address.
//    So whenever your program says oh I want to retrieve some information from the computer's memory. It looks at the look goes in find some address and then it pulls the value out of there.
//    And so each of these little boxes right here can contain some amount of information.
//
//    Using this to understand why updating Alex's first name to Alec didn't work...
//    #1 When we did something like alex := person{}
// 	  	 So when we do this when we create this new struct of type person, Go will create that struct.
//       It will then go to the local memory on our own laptop or our own local machine and it will attempt to
//       find some container or some spot that is free and has the ability to accept some data so we can imagine
//       that Go takes this structure right here it goes and finds some space or some location to place that
//       struct and then it shoves that data into this little container right here.
//       And so we can imagine that this alex structure or this person is sitting at the address of 0001,
//       so whenever we look at the variable alex is pointing directly at that little container right there.
//       And so if we print out the values alex, we're always going to see exactly this value right there
//
//    Address | Value
//    0000    |
//    0001    | person{firstName: "Alex"...}   <--- alex - #1
//    0002    |
//    0003    |
//    0004    |
//
//    We tried to update the name with updateName function: func (p person) updateName() {
//    #2 However, Go is what we refer to as a pass by value. Language passed by value means that whenever we pass some value into a function, Go will take that value or take that struct
//    and copy all of that data inside that struct and then place it inside of new some new container inside of our computers memory
//
//    Address | Value
//    0000    |
//    0001    | person{firstName: "Alex"...}   <--- alex - #1
//    0002    |
//    0003    | person{firstName: "Alec"...}   <--- p - #2
//    0004    |
//
//    Therefore, when we pass alex into this updateName function, alex still exists by itself as a struct with the first name of Alex at address 0001.
//    But Go copy's that value it finds some other container that's empty and it stuffs that copy into that container and then it runs the code inside of updateName with this receiver
//    pointing at that copy. And so when we modify that field of firstName inside of that function when we run this code right there where it says p.firstName is going to be newFirstName,
//    we are not updating the original struct of alex. We were simplu updating the copy that was just made for our particular function call
//
//    4. Pointer Operators:
//    &variable means to give me the memory address of the value this variable is pointing at. E.g. from Point 3, alexPointer of &alex would give me 0001
//    *pointer  means to give me the value this memory is pointing at. E.g. from above's case, *alexPointer give me person{firstName: "Alex"...}
//
//    TL;DR:
//    0001, the address, has a value of a person{firstName: "Alex"...}
//    We turn address into value with *address
//    We turn value into address with &value
//
//    5. func (pointerToPerson *person) updateName() {
//		    *pointerToPerson
//       }
//
//       *person is a type description - it means we're working with a pointer to a person. A.k.a. it means that this update main function can only be called with the receiver of a pointer to a person.
//       *pointerToPerson is an operator - it means we want to manipulate the value the pointer is referencing
//
//    6.  Go allows us to take the shortcut and says hey if you have a variable that's just type of person but then your receiver is pointer to a person that's totally fine we'll just gloss
//        over that fact for you and go will automatically turn your variable of type person into pointer person for you. That's why even though we removed alexPointer := &alex, the program
//        still worked as intended

// Quiz 7: Test Your Knowledge: Pointers
// 1. Whenever you pass an integer, float, string, or struct into a function, what does Go do with that argument? It creates a copy of each argument, and these copies are used inside of the function
// 2. What will the following program print out? The memory address that "Bill" is stored at
/* 	package main
	import "fmt"

	func main() {
   		name := "Bill"

   	fmt.Println(&name)
	} */
// 3. What is the & operator used for? Turning a value into a pointer
// 4. When you see a * operator in front of a pointer, what will it turn the pointer into? A value
// 5. When the following program runs, the fmt.Println call reports that the latitude field of newYork is still equal to 40.73.
// What changes should we make to get the latitude of newYork to update to 41.0? Change the receiver type of changeLatitude to *location, then replace lo with (*lo)
// in the function body. This will turn the pointer lo into a value type and then update it
/* 	package main
	import "fmt"

	type location struct {
 		longitude float64
 		latitude float64
	}

	func main() {
 		newYork := location{
   			latitude: 40.73,
   			longitude: -73.93,
 		}

 		newYork.changeLatitude()

 		fmt.Println(newYork)
	}

	func (lo location) changeLatitude() {
 		lo.latitude = 41.0
	} */
// 6. Take a look at the following snippet of code.
/* 	package main

	import "fmt"

	type location struct {
 		longitude float64
 		latitude float64
	}

	func main() {
 		newYork := location{
 			latitude: 40.73,
 			longitude: -73.93,
 		}

 		newYork.changeLatitude()

 		fmt.Println(newYork)
	}

	func (lo *location) changeLatitude() {
 		(*lo).latitude = 41.0
	} */
// In the 'changeLatitude' function, what is *location in the receiver list (after the word 'func') communicating to us? It specifies the type of the receiver that the function expects
// 7. Take a look at the following program. What will the Println function in the main function print out? "Bill"
/* 	package main

	import "fmt"

	func main() {
    	name := "Bill"
    	updateValue(name)
    	fmt.Println(name)
	}

	func updateValue(n string) {
    	n = "Alex"
	} */
// 8. Take a look at the following program. The changeLatitude function expects a receiver of type pointer to a location struct, but in the main function the receiver is a value type
//    of a struct. What will happen when this code is executed? This program uses a shortcut, where Go will automatically assume that even though newYork.changeLatitude() is using a value
//    type, we probably meant to pass in a pointer to the newYork struct
/* 	package main

	import "fmt"

	type location struct {
 		longitude float64
 		latitude float64
	}

	func main() {
 		newYork := location{
   			latitude: 40.73,
   			longitude: -73.93,
 		}

 		newYork.changeLatitude()

 		fmt.Println(newYork)
	}

	func (lo *location) changeLatitude() {
 		(*lo).latitude = 41.0
	} */
// 9. Here's a tricky one! What will the following program print out? The string "Bill"
/* 	package main

	import "fmt"

	func main() {
    	name := "Bill"

    	fmt.Println(*&name)
	} */
