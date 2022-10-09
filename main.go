package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const confenrenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", confenrenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your ticket here to attend")

}
