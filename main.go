// Type "go run main.go " at the terminal to run this application
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Variable declaration is like var conferenceName = "Go Conference"

	// Alternative syntax for defining a variable and assigning it a value
	// Just a syntactic sugar
	// Only applies to var and not const. And it also doesnt work, if you  want to include the type 
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	
	var remainingTickets uint  = 50
	// fmt function is used to accept or print 
	// println is used for printing a line
	fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Println("We have a total of",conferenceTickets,"tickets and", remainingTickets, "are available.")
	fmt.Println("Get your tickets here to attend.")
	
	// slices
	// Abstraction of an array. i.e. variable length array
	// It can get a sub-array of its own
	//  Slices are also index-based and have a size but its resized when needed
	// var bookingList []string
	bookingList := []string {}
	// For loop to be executed only when a particular set of conditions are satisfied
	for remainingTickets > 0 && len(bookingList) < 50 {
		// Variable declaration without initialising and we need to specify the datatype if its not initialised
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var city string

		println("Enter your first name")
		// scan is used to accept the user input and & indicates pointer
		fmt.Scan(&firstName)

		println("Enter your last name")
		fmt.Scan(&lastName)

		println("Enter your Email")
		fmt.Scan(&email)

		println("Enter the number of tickets")
		fmt.Scan (&userTickets)
		
		println("Enter the city")
		fmt.Scan (&city)
		
		// User input validation examples

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		// isInvalidCity := city != "Singapore" && city != "London" 
		// This is the same as above
		// isValidCity := city == "Singapore" or city == "London"
		// And then checking for !isValidCity
		if isValidName && isValidEmail && isValidTicketNumber {

			// Adding value to the slice
			//  We dont need to know the current index.use the append
			bookingList = append(bookingList, firstName + " " + lastName)
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets.You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining", remainingTickets)

			firstNames := []string{}

			// for index, booking := range bookingList
			// _ is a Blank Identifier in Go
			// To ignor a variable that you dont want to use
			// i.e. with Go, we need to make unused variables explicit
			for _, booking := range bookingList {
				// splitting a string by spaces
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
		} else {
			// fmt.Println("Your input data is invalid, Try again!")
			if !isValidName{
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("The number of tickets you entered is invalid")
			}
		}
	}
}