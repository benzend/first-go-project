package tickets

import (
	"fmt"
	"sync"
	"time"
)

func SendTicket(tickets uint, name string, email string, wg sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", tickets, name)

	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done()
}