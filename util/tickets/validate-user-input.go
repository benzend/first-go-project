package tickets

import "strings"

func ValidateUserInput(first string, 
	last string, 
	email string, 
	tickets uint, 
	remainingTickets uint) (bool, bool, bool) {
	return len(first) >= 2 && len(last) >= 2,
		strings.Contains(email, "@"), 
		tickets > 0 && 
		tickets <= remainingTickets
}