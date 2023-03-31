package main

import "fmt"

func main() {

	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// var bookings [50]string ### This is how we declare arrays but this is not the convinient way to do so.
	bookings := []string{} // We use Slice instead

	// fmt.Println("Welcome to our", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remaininTickets, "are still available")

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var (
		userFirstName string
		userLastName  string
		userTickets   uint
		userEmail     string
		userMobile    uint
	)

	fmt.Println("Enter your first name:")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your email-id:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter your Mobile number:")
	fmt.Scan(&userMobile)

	fmt.Println("Number of tickets you want to book:")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	// bookings[0] = userFirstName + " " + userLastName ### This is not the right way to do.
	bookings = append(bookings, userFirstName+" "+userLastName) // "Append" adds the elements at the end of the slice

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve an confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets are remaining for the %v, Share to your friends\n", remainingTickets, conferenceName)
	fmt.Printf("This is all our bookings: %v\n", bookings)
}
