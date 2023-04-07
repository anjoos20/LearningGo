// Define which package this file belongs to
package helper

import "strings"

//  Multiple returns included inside a paranthesis
//Capitalise the name means we are just exporting the function
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint,RemainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

