package main

import (
	"bookingApp/helper"
	"fmt"
)

type bookingInfo struct {
	fName    string
	lName    string
	email    string
	ticCount uint
}

var conferenceName = "Go Conference"

const conferenceTicket = 50    // number of availabe seats
var remainingTickets uint = 50 // uint is use to avoid negative values

func main() {

	greetUsers(conferenceName, conferenceTicket, remainingTickets)

	var booking []bookingInfo

	for {

		bi := getUserInput()

		userbooking := bookingInfo{bi.fName, bi.lName, bi.email, bi.ticCount}
		booking = append(booking, userbooking)
		remainingTickets -= bi.ticCount

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", bi.fName, bi.lName, bi.ticCount, bi.email)
		fmt.Printf("These are all our bookings %v\n", booking)
		fmt.Printf("We has %v remaining tickets", remainingTickets)

		if remainingTickets == 0 {
			fmt.Println("All the tickets are received")
			break
		}

	}

}

func getUserInput() bookingInfo {
	var bi bookingInfo
first:
	fmt.Print("Enter your first name : ")
	fmt.Scan(&bi.fName)
	err := helper.NameValidation(bi.fName)
	if err != nil {
		fmt.Println(err)
		goto first
	}

last:
	fmt.Print("Enter your last name : ")
	fmt.Scan(&bi.lName)
	err = helper.NameValidation(bi.lName)
	if err != nil {
		fmt.Println(err)
		goto last
	}

email:
	fmt.Print("Enter your email address : ")
	fmt.Scan(&bi.email)
	err = helper.EmailValidation(bi.email)
	if err != nil {
		fmt.Println(err)
		goto email
	}

	fmt.Print("How many tickets you want : ")
	fmt.Scan(&bi.ticCount)
out:
	if bi.ticCount > remainingTickets {
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
		fmt.Print("Please book tickets accordinly  : ")
		fmt.Scan(&bi.ticCount)
		goto out
	}
	return bi
}

func greetUsers(conf string, confTicks int, remainingTics uint) {
	fmt.Println("Welcome to our", conf, "booking application.")
	fmt.Println("We have total of", confTicks, "tickets and", remainingTics, "are still available.")
	fmt.Println("Get your tickets here to attend\n", conf)
}

// go run  main.go helper.go
// or
// go run .
