package helper

import (
	"strings"
)

func ValidateUserInputs(customer string, customerEmail string, ticketAmount int32, remainingTickets int32) (bool, bool, bool) {
	isValidName := len(customer) >= 2
	isValidEmail := strings.Contains(customerEmail, "@") //check if email input have @ character
	isValidTicketAmount := ticketAmount > 0 && ticketAmount <= remainingTickets
	return isValidName, isValidEmail, isValidTicketAmount
}
