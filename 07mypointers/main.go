package main

import "fmt"

func main() {
	fmt.Println("Pointers")
// A pointer which holds integer
	// var ptr  *int 
	// fmt.Println("Value of pinter is ", ptr)

	myNumber :=23;
// Previously just created a pointer
// Here creating to pointer which is referencing to some memory
//  Whenver referencing => use ampersand
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	// *ptr means the value inside that memory location
	 
	*ptr  = *ptr * 2

	fmt.Println("New Value of pointer is ", *ptr, myNumber)
	//  Pointer makes it a guarantee that the operations are actually performed on the value
	//  and not on the copy
}