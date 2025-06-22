package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Language Conference" // var conferenceBame stgring = "Go Conference 2023".
	// above := wont work with connstants
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v application!\n", conferenceName)
	fmt.Printf("Hurry up! Out of %v tickets, only %v are remaining!", conferenceTickets, remainingTickets)
	fmt.Println("Please fill in the details to book your slot")

	var firstName, lastName, userEmail string
	var userTickets uint
	bookings := []string{} // var bookings []string // var bookings = []string{}

	for {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address:")
		fmt.Scan(&userEmail)
		fmt.Println("How many tickets would you like to book?")
		fmt.Scan(&userTickets)
		// The & is a pointer operator, it allows us to store the value in the variable. Printing & will print the address of the variable.
		// fmt.Println(bookings)
		// fmt.Println(bookings[0])
		// fmt.Println("length = ", len(bookings))
		remainingTickets = remainingTickets - userTickets
		fullName := firstName + " " + lastName
		bookings = append(bookings, fullName)

		fmt.Printf("Thank you %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, userTickets, userEmail)
		firstNames := []string{}
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}
		fmt.Printf("Debug Message: Booking by Firstnames only:  %v\n", firstNames)
		fmt.Printf("Debug Message: Booking by Full name:  %v\n", bookings)
		fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end program
			fmt.Println("Our Conference is booked out. Come back next year")
			break
		}

	}
}
