package main

import "fmt"

func main() {

	// Approach 1:
	var num1 = 2
	var num2 = 3

	// or //

	// Approach 2:

	var num0 int // by default int value is zero
	var num3 int = 2
	var num4 int = 3

	//  or //

	// Approach 3

	var num5, num6 int
	// num5 = 6
	// num6 = 9
	num5, num6 = 6, 9

	var result = num1 + num2
	var result1 = num0 + num3 + num4
	var result2 = num5 + num6

	fmt.Println(result)  // 5
	fmt.Println(result1) //5
	fmt.Println(result2) //15

	// Approach 4:

	var (
		id    int    = 23
		name  string = "mahesh"
		sname string = "miriyampalli"
	)
	var (
		count int = 1
	)
	fmt.Println(id, name, sname, count)

	// without variable declaration

	a := "welcome"
	b := 24
	d := 6
	output := b + d
	fmt.Println(a, b, d)
	fmt.Println(b + d)
	// or //
	fmt.Println(output)

	fmt.Printf("a is %T,b is %T\n", a, b)

	// constants:

	var conferenceName = "GoConference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("welcome to", conferenceName, "always")
	fmt.Println("We have total tickets of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("let's meet")

	// Example:

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

	var c int = 12
	fmt.Println(c)

}
