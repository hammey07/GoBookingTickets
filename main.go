package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Language Conference" // var conferenceBame stgring = "Go Conference 2023".
var remainingTickets uint = 50

// above := wont work with connstants
const conferenceTickets int = 50

// var bookings = []string{} // var bookings []string // var bookings = []string{} //slice
// var bookings = make ([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData = struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValididateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidEmail && isValidName {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// firstNames := getFirstNames()
			// fmt.Printf("Debug Message: Booking by Firstnames only:  %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year")
				wg.Wait()
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
				fmt.Printf("Input field error / You can't book %v tickets as there are only %v tickets available.\n", userTickets, remainingTickets)
			}
		}

		wg.Wait()
	}

}

func greetUsers() {
	fmt.Println("------------------------------------------------")
	fmt.Printf("Welcome to %v application!\n", conferenceName)
	fmt.Println("------------------------------------------------")
	fmt.Printf("Total Tickets = %v tickets\n", conferenceTickets)
	fmt.Printf("Available Tickets = %v tickets\n", remainingTickets)
	fmt.Println("Please fill in the details to book your slot.")
	fmt.Println("------------------------------------------------")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Print(bookings)

	fmt.Printf("Thank you %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, userTickets, email)
	fmt.Println("--------------- Debugging ----------------------")
	fmt.Printf("Booking Details: \n%v\n------\n", bookings)
	fmt.Println("------------------------------------------------")
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("--------------- Debugging ----------------------")
	fmt.Printf("Sending Ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("------------------------------------------------")
	wg.Done()
}
