package booking

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type BookingList struct {
	Customer       string
	Customer_Email string
	Trip           string
	TicketAmount   int32
	BookedAt       time.Time
}

type Bookings []BookingList

func (b *Bookings) Create(customer, customerEmail, trip string, ticketAmount int32) {
	bookingData := BookingList{
		Customer:       customer,
		Customer_Email: customerEmail,
		TicketAmount:   ticketAmount,
		BookedAt:       time.Now(),
	}

	*b = append(*b, bookingData)
}

func (b *Bookings) Delete(index int) error {
	list := *b
	if index <= 0 || index > len(list) {
		return errors.New("Invalid Index")
	}

	*b = append(list[:index-1], list[index:]...)

	return nil
}

func (b *Bookings) Save(filename string) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, data, 0644)
	}
}

func (b *Bookings) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	} else if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, b)
	if err != nil {
		return nil
	}

	return nil
}

func (b *Bookings) Print() {
	for i, item := range *b {
		i++
		fmt.Printf("%d - Customer: %s || Email: %s || Trip: %s || TicketAmount: %v\n", i, item.Customer, item.Customer_Email, item.Trip, item.TicketAmount)
	}
}
