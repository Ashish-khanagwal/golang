package main

import (
	"fmt"
	"strconv"
	// "strings"
)

// These are known as package variables, it means we just have to declare them once and they'll be availabel to all fiunctions

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = []string{}
var bookings = make([]map[string]string, 0) // This will create a list of maps.

func main() {

	// var conferenceName = "Go Conference"
	// conferenceName := "Go Conference"
	// const conferenceTickets int = 50
	// var remainingTickets uint = 50

	// var bookings [50]string ### This is how we declare arrays but this is not the convinient way to do so.
	// bookings := []string{} // We use Slice instead

	// Grerting user using a fucntion declared outside the main func
	greetUsers()

	// fmt.Println("Welcome to our", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remaininTickets, "are still available")

	for {

		// Get user Input
		userFirstName, userLastName, userTickets, userEmail := getUserInput()

		// Validating User's Input
		isValidName, isValidEmail, isValidTicketNumber := ValidUserInput(userFirstName, userLastName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Booking tickets
			bookTicket(userFirstName, userLastName, userEmail, userTickets)

			// Getting first names using funciton declared outside the main func
			firstNames := getFirstNames()
			fmt.Printf("The First name in all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Tickets for Go Conference are booked, come back next year\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid Firstname or Lastname")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid Email")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter valid number of tickets")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Below, parameter string is for input purpose and string declare outside parameter is for output purpose
func getFirstNames() []string {
	firstNames := []string{}
	// we have used "_" <- blank identifire and it is used to ignore a variable that we don't use.
	// for index, chr := range str {}
	for _, booking := range bookings {
		// var firstName = strings.Fields(booking) // strings.Fields() <- This basically splits the sring.
		firstNames = append(firstNames, booking["firstName"])
		// booking["firstNames"] way to extract firstname from map's key-value pair
	}
	return firstNames
}

func getUserInput() (string, string, uint, string) {
	var (
		userFirstName string
		userLastName  string
		userTickets   uint
		userEmail     string
	)

	fmt.Println("Enter your first name:")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email-id:")
	fmt.Scan(&userEmail)

	fmt.Println("Number of tickets you want to book:")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userTickets, userEmail
}

func bookTicket(userFirstName string, userLastName string, userEmail string, userTickets uint) {
	remainingTickets -= userTickets
	// bookings[0] = userFirstName + " " + userLastName ### This is not the right way to do.

	// Creates a map for a user
	var userData = make(map[string]string) // This will create a empty map.
	userData["firstName"] = userFirstName
	userData["lastName"] = userLastName
	userData["email"] = userEmail
	userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10) // this is used to convert uint to string

	// bookings = append(bookings, userFirstName+" "+userLastName) // "Append" adds the elements at the end of the slice
	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve an confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for the %v, Share to your friends\n", remainingTickets, conferenceName)
}
