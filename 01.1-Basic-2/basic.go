package main

import "fmt"

func main() {

	// Approach 1:

	a := "welcome"
	b := 24
	fmt.Println(a, b)
	fmt.Printf("a is %T,b is %T\n", a, b)

	// Approach 2:

	var conferenceName = "GoConference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("welcome to", conferenceName, "always")
	fmt.Println("We have total tickets of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("let's meet")

	// Approach 3:

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

	var c int = 12
	fmt.Println(c)

	var (
		id    int    = 23
		name  string = "mahesh"
		sname string = "miriyampalli"
	)
	var (
		count int = 1
	)
	fmt.Println(id, name, sname, count)
}
