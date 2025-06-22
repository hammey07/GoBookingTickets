package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Language Conference" // var conferenceBame stgring = "Go Conference 2023".
// above := wont work with connstants
const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{} // var bookings []string // var bookings = []string{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValididateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidEmail && isValidName {
			bookTicket(userTickets, firstName, lastName, conferenceName, email)

			firstNames := getFirstNames()
			fmt.Printf("Debug Message: Booking by Firstnames only:  %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("Your email is invalid. Please try again")
			}
			if !isValidName {
				fmt.Println("Your name is invalid. Please try again")
			}
			if !isValidTicketNumber {
				fmt.Printf("You can't book %v tickets as there are only %v tickets available.\n", userTickets, remainingTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v application!\n", conferenceName)
	fmt.Printf("Hurry up! Out of %v tickets, only %v are available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Please fill in the details to book your slot.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)
	// The & is a pointer operator, it allows us to store the value in the variable. Printing & will print the address of the variable.
	// fmt.Println(bookings)
	// fmt.Println(bookings[0])
	// fmt.Println("length = ", len(bookings))
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, conferenceName string, email string) {
	remainingTickets = remainingTickets - userTickets
	fullName := firstName + " " + lastName
	bookings = append(bookings, fullName)

	fmt.Printf("Thank you %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, userTickets, email)
	fmt.Printf("Debug Message: Booking by Full name:  %v\n", bookings)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
