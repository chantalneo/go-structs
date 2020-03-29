package main

/* Commented out to avoid seeing error in my project
import (
	"fmt"
)

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
} */

// This reveals a big gotcha with Go. Even though we did not make use of pointer's nothing like that no memory addresses whatsoever, it appears that with a slice when we
// modified it inside this function, it actually modified the original value which is completely opposite of how our struct works. In other words, this is not quite working
// the same way with a struct as with a slice

// To understand the strange behaviors, we need to understand something that may seems a little bit of a tangent:
// Arrays vs Slices...
// Arrays:
// - Primitive data structure
// - Can't be resized
// - Rarely used directly
// Slices:
// - Can grow and shrink
// - Used 99% of the time for lists of elements
//
// So at this point in the course we've only been making direct use of slices. We rarely use arrays directly because it is seen as a very primitive data structure mostly because arrays
// cannot be resized. And so it's not incredibly often that we ever want to make a list of items that it is a fixed length. Usually we want to make a list of items that can grow or
// shrink over time as required by our code. So 99 percent of time we are always making use of slices and we rarely ever make arrays directly. Now you can think of a slice as being
// like a fancy array. And in truth behind the scenes a slice actually kind of is an array.
//
// In this scenario, when we declared a slice with mySlice := []string{"Hi", "There", "How", "Are", "You"}, Go internally created two separate data structures for us.
// It created a slice, which is a data structure that has three elements inside of it.
// - A pointer. A pointer over to the underlying array that represents the actual list of items
// - A capacity number. The capacity is how many elements it can contain at present
// - A length number the length is represents how many elements currently exist inside the slice
// It also created the aforementioned array that represents the actual list of items
//
// Thus, a light illustration of the RAM can be seen:
// 			RAM
//    Address | Value
//    0000    |
//    0001    | [length][cap][ptr to head]     					<--- mySlice
//    0002    | []string{"Hi", "There", "How", "Are", "You"}   	<--- Where the ptr to head points to
//    0003    |
//    0004    |
// Therefore, whenever we refer to mySlice, we're actually returning the slice data structure not the array
//
// So, when we call a function and pass mySlice into it. For example, when we call an update function and pass mySlice into it, Go is still behaving as a pass by value language
// Go still makes a copy of that underlain of that value. This means that when we call the update size function and pass in our slice we take the slice data structure and
// copy it off to another address in memory:
// 			RAM
//    Address | Value
//    0000    |
//    0001    | [length][cap][ptr to head]     					<--- mySlice
//    0002    | []string{"Hi", "There", "How", "Are", "You"}   	<--- Where the ptr to head points to
//    0003    |
//    0004    | [length][cap][ptr to head]                      <-- Its ptr also points to 0002's value
// But here's the crux, the really important thing.
// Even though the slice data structure is copied it is still pointing at the original array in memory
// because the slice data structure and the array data structure are two separate elements in memory.
// So yeah we are copying the slice but it is still pointing at the same array.
// Therefore, whenever we modify this array or when we're inside the function when we attempt to modify the slice, we're modifying same array that both copies of the slice are now pointing to

// Important!
// Now here's the real gotcha, in Go, slices are not the only data structure that behave in this fashion.
// Values Types | Reference Types
// int          | slices
// float        | maps
// string       | channels
// bool         | pointers
// structs      | functions
// We NEED to use pointers to change value types in a function
// On the other hand, we don't need to worry about pointers for the Reference Types
