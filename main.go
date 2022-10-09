package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	// Synthatic Sugar
	// conferenceName := "Go Conference"

	const confenrenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confenrenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var bookings = [50]string{}

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

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
