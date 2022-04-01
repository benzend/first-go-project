package tickets

import (
	"fmt"
	"sync"
	"time"
)

func SendTicket(tickets uint, name string, email string, wg sync.WaitGroup) {
	var ticket = fmt.Sprintf("%v tickets for %v", tickets, name)
	
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	
	fmt.Println("Sending ticket...")

	time.Sleep(1 * time.Second)
	
	fmt.Println("Sending ticket...")
	
	time.Sleep(1 * time.Second)

	fmt.Println("Ticket sent!")
	wg.Done()
}