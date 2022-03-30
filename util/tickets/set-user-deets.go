package tickets

import "fmt"

func SetUserDeets(name *string, tickets *uint, email *string) {
	fmt.Println("What's your name?")
	fmt.Scan(name)

	fmt.Printf("Alrighty %v, how many tickets would you like to purchase?\n", *name)
	fmt.Scan(tickets)

	fmt.Println("What's your email?")
	fmt.Scan(email)
}