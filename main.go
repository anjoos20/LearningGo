// Type "go run main.go helper.go " at the terminal to run this application
// or use go run . BOOKING_APP
package main

import (
	"fmt"
	"strconv"
)

//Package level variables
//Variables which are defined outside all thefunctions and are accessible by all the functions
//Cannot use := for this but only var keyword
//They can be accessed by all other files in the same package
//The best practice is define a variable as local as possible and they can be accessed only inside
// that function or block of code
//Create the variable where you need it
const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
// Currently bookings variable is a slice of strings
// var bookings = []string{}
//  We need to make it into a slice of maps
// Now it becoms a list of maps and not a list of strings
// var bookings = []map[string]string{}
// We need to actually create an empty slice of maps using the make function
// This creates an empty slice of maps with an initial length and capacity of 0.
var bookings = make([]map[string]string,0,0)
func main() {

	// Welcome message for the user
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

// within the brackets input arguements and outside the bracket return parameter
func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		// Accessing the first name becomes easy with the map
		firstNames = append(firstNames,booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//  Create map for a user
	// This is just declaring the map type
	// var mymap map[string]string
	// To create an empty map we have a built in function called make
	var userData = make(map[string]string)
	//  Add key value pair to the map
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	//  In Go, maps can only have the same data types as keys and the same data type as values
	// ie. we cannot mix various data types as values
	// so cannot map with the usertickets which is a uint unless its converted to string using the builtin function
	userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)
	
	// Now we have the bookings list which contains all the user information
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings(map) is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
