package main

import "fmt"

func main() {

	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remaininTickets = 50

	// fmt.Println("Welcome to our", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remaininTickets, "are still available")

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets, remaininTickets)
	fmt.Println("Get your tickets here to attend")

}
