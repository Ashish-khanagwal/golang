package main

import "strings"

func ValidUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@") && strings.Contains(userEmail, ".")
	// strings.Contains() <- this is used to check if a string contains particular word or char
	isValidTicketNumber := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTicketNumber
	// In Go we can return multiple values
}
