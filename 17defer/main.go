package main

import "fmt"

func main() {
	// When it finds the keyword defer, it just cuts the first line and puts it towards the end of the function
	// i.e. just before the end of curly braces
	defer fmt.Println("Hello")
	// ie. in the reverse order they are deferred. Printed LIFO
	// McDonalds que is being created
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("World")
}