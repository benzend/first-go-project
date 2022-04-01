package applications

import (
	"first-project/util/tickets"
	"fmt"
	"sync"
)

type UserData struct {
	name string
	email string 
	tickets uint
}

var wg = sync.WaitGroup{}

func TicketApp() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = make([]UserData, 0)

	tickets.Greet(conferenceName, conferenceTickets, remainingTickets)

	for {
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out, come back next year")
			break
		}
	
		var userName string
		var userTickets uint
		var userEmail string
		var purchaseTickets string
	
		fmt.Println("Would you link to pushase tickets? [y/n]")
		fmt.Scan(&purchaseTickets)

		if purchaseTickets == "y" {
			tickets.SetUserDeets(&userName, &userTickets, &userEmail)
			validName, validEmail, validOrder := tickets.ValidateUserInput(userName, userName, userEmail, userTickets, remainingTickets)

			if validName && validEmail && validOrder {
				fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
			
				remainingTickets = remainingTickets - userTickets

				var userData = UserData {
					name: userName,
					email: userEmail,
					tickets: userTickets,
				}

				bookings = append(bookings, userData)
				
				dontWait := ""

				fmt.Println("Would you like to buy more tickets while you wait? [y/n]")

				fmt.Scan(&dontWait)

				if dontWait == "y" {
					wg.Add(1)
					go tickets.SendTicket(userTickets, userName, userEmail, wg)
				} else if dontWait == "n" {
					tickets.SendTicket(userTickets, userName, userEmail, wg)
				} else {
					fmt.Println("You done fucked up boy")
					continue
				}

				
				
			} else {
				fmt.Println("Invalid details")
				continue
			}

		} else if purchaseTickets == "n" {
			continue
		} else {
			fmt.Println("That isn't a valid response")
			continue
		}
	}
}