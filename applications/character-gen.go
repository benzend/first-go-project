package applications

import (
	"fmt"
	"time"
)

func CharacterGenApp() {
	fmt.Println("Welcome to the number gen application!")

	for {
		fmt.Println("Would you like to create a character? [y/n]")
		var start string
		fmt.Scan(&start)
	
		if start == "y" {
			Start()
		} else {
			continue
		}

	}
}

func Start() {
	name, height, build, race, primary := GetCharacterDetails()

	fmt.Print("Generating character")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Print(".\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Character generated!")
	time.Sleep(2 * time.Second)
	
	fmt.Printf("Welcome %v! \n", name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Don't forget your %v!\n", primary)
	time.Sleep(1 * time.Second)
	GenerateItem(primary)
	time.Sleep(2 * time.Second)
	fmt.Println("Here is a summary of all of your stats!")
	time.Sleep(1 * time.Second)
	fmt.Printf("Height: %v\nBuild: %v\nRace: %v\n", height, build, race)
}

func GetCharacterDetails() (name string, height int, build string, race string, primary string) {
	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Height: ")
	fmt.Scan(&height)
	fmt.Print("Build [strong/weak]: ")
	fmt.Scan(&build)
	fmt.Print("Race [dwarvish/human]: ")
	fmt.Scan(&race)
	fmt.Print("Primary weapon [sword/bow]: ")
	fmt.Scan(&primary)

	return name, height, build, race, primary
}

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
	default:
		fmt.Println("Looks like I don't have that item in stock :(")
	}
}