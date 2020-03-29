package main

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
}

// This reveals a big gotcha with Go. Even though we did not make use of pointer's nothing like that no memory addresses whatsoever, it appears that with a slice when we
// modified it inside this function, it actually modified the original value which is completely opposite of how our struct works. In other words, this is not quite working
// the same way with a struct as with a slice
