package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName string = "Go Conference"

const confenrenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			fmt.Printf("These are all the bookings for the %v: %v\n", conferenceName, getFirstNames())

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				// end Program
				fmt.Println("Our conference is book out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or las name is too short. Try again!")
			}
			if !isValidEmail {
				fmt.Println("Email adress you entered is invalid. Try again!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid. Try again!")
			}
		}
	}
}

func validateUserInput(firstName, lastName, email string, userTickets uint) {
	panic("unimplemented")
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confenrenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
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

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a MAp for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
