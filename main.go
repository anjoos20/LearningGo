// Type "go run main.go " at the terminal to run this application
package main

import (
	"fmt"
)

func main() {
	// Variable declaration
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint  = 50
	// fmt function is used to accept or print 
	// println is used for printing a line
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of",conferenceTickets,"tickets and", remainingTickets, "are available.")
	fmt.Println("Get your tickets here to attend.")
	//  Array declaration
	var bookings [50] string

	// slices
	// Abstraction of an array. i.e. variable length array
	// It can get a sub-array of its own
	//  Slices are also index-based and have a size but its resized when needed
	var bookingList []string

	// Variable declaration without initialising and we need to specify the datatype if its not initialised
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	println("Enter your first name")
	// scan is used to accept the user input
	fmt.Scan(&firstName)

	println("Enter your last name")
	fmt.Scan(&lastName)

	println("Enter your Email")
	fmt.Scan(&email)

	println("Enter the number of tickets")
	fmt.Scan (&userTickets)
	// Adding value to array
	bookings[0] = firstName + " " + lastName

	// Adding value to the size
	//  We dont need to know the current index.use the append
	bookingList = append(bookingList, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets


	// %v indicates the value
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	// %T is used to display the Type
	fmt.Printf("Array Type: %T \n", bookings)
	fmt.Printf("Array length: %v \n", len(bookings))

	// %v indicates the value
	fmt.Printf("The whole slice: %v\n", bookingList)
	fmt.Printf("The first value: %v \n", bookingList[0])
	// %T is used to display the Type
	fmt.Printf("Slice Type: %T \n", bookingList)
	fmt.Printf("Slice length: %v \n", len(bookingList))


	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining", remainingTickets)
}