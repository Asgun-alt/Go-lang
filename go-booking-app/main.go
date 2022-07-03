package main

import (
	"flag"
	"fmt"
	"go-booking-app/booking"
	"os"
)

const (
	bookingFile = ".booking_data.json"
)

func main() {

	create := flag.Bool("create", false, "Create a new booker")
	delete := flag.Int("delete", 0, "Delete a booker")
	list := flag.Bool("list", false, "Booking list")

	flag.Parse()

	bookings := &booking.Bookings{}

	if err := bookings.Load(bookingFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *create:
		customer, customerEmail, trip, ticketAmount := getInput()

		bookings.Create(customer, customerEmail, trip, ticketAmount)
		err := bookings.Save(bookingFile)
		fmt.Println("New data has been added")
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *list:
		bookings.Print()

	case *delete > 0:
		err := bookings.Delete(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		} else {
			err = bookings.Save(bookingFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}
}

func getInput() (string, string, string, int32) {
	var (
		customer      string
		customerEmail string
		trip          string
		ticketAmount  int32
	)

	fmt.Print("Enter Customer Name: ")
	fmt.Scan(&customer)

	fmt.Print("Enter Customer Email: ")
	fmt.Scan(&customerEmail)

	fmt.Print("Enter Trip Destination: ")
	fmt.Scan(&trip)

	fmt.Print("Enter Ticket Amount: ")
	fmt.Scan(&ticketAmount)

	return customer, customerEmail, trip, ticketAmount
}
