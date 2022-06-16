package main

import (
	"fmt"
	"go-booking-app/helper"
	"strings"
)

// package level scope variable

var tripName = "Bandung - Jakarta"

const tripTickets = 50

var remainingTickets = 50
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstName, lastName, email, userTicket := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicked(userTicket, firstName, lastName, email)

			// call print first name function
			firstNames := getFirstName()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Printf("%v tickets has been sold out", tripName)

				break // end the loop
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you have entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address should contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of ticket(s) you entered is invalid")
			}
			// the continue statement should not have to be used the next line of code (going back to for loops) will be executed
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v Train Booking Application.\n", tripName)
	fmt.Printf("We have total of %v and we have %v remaining tickets available.\n", tripTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getFirstName() []string {
	listUserFirtNames := []string{}

	// _ (Blank Identifier) is used to ignore unused variable
	for _, booking := range bookings {
		var names = strings.Fields(booking)

		listUserFirtNames = append(listUserFirtNames, names[0])
	}
	return listUserFirtNames
}

func getUserInputs() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTicket int

	// & or pointer is a variable that points to the memory address to other variable
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Please enter the number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicked(userTicket int, firstName string, lastName string, email string) {
	remainingTickets -= userTicket
	// adds the element(s) at the end of the slice
	// slice will grows its capacity if needed and returns the updated slice value
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, tripName)
}
