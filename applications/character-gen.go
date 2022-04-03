package applications

import (
	characterUtil "first-project/util/character"
	"fmt"
	"time"
)

type Character struct {
	name string 
	height string 
	build string
	race string 
	primaryWeapon string
}

func (c Character) Details() {
	fmt.Printf("name: %v\nheight: %v\nbuild: %v\nrace: %v\nprimaryWeapon: %v\n", 
		c.name, 
		c.height, 
		c.build, 
		c.race, 
		c.primaryWeapon,
	)
}

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
	character := GetCharacterDetails()

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
	
	fmt.Printf("Welcome %v! \n", character.name)
	time.Sleep(2 * time.Second)
	fmt.Printf("Don't forget your %v!\n", character.primaryWeapon)
	time.Sleep(1 * time.Second)
	characterUtil.GenerateItem(character.primaryWeapon)
	time.Sleep(2 * time.Second)
	fmt.Println("Here is a summary of all of your stats!")
	time.Sleep(1 * time.Second)
	character.Details()
}

func GetCharacterDetails() (c Character) {
	fmt.Print("Name: ")
	fmt.Scan(&c.name)
	fmt.Print("Height: ")
	fmt.Scan(&c.height)
	fmt.Print("Build [strong/weak]: ")
	fmt.Scan(&c.build)
	fmt.Print("Race [dwarvish/human]: ")
	fmt.Scan(&c.race)
	fmt.Print("Primary weapon [sword/bow]: ")
	fmt.Scan(&c.primaryWeapon)

	return c
}

