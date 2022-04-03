package character

import "fmt"

func GenerateItem(item string) {
	switch item {
	case "sword":
		fmt.Print(`
     /\
    /  \
    |  |
    |  |
    |  |
    |  |
    |  | 
    |  |
----    ----
    |  |
    |  |
    |__|
		`)
		fmt.Print("\n")
	case "bow":
		fmt.Print(`
		\\
		 \\
		  \\
			 \\
			  \\
				 \\
				  \\
					 \\
					  \\`)
	default:
		fmt.Println("Looks like I don't have that item in stock :(")
	}
}