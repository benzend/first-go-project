package main

import (
	"first-project/applications"
	"fmt"
)

type App struct {
	name string 
	fn func()
}

var apps = []App{{name: "tickets", fn: applications.TicketApp}, {name: "character-gen", fn: applications.CharacterGenApp}}

func main() {
	app := getChosenApp()
	app.fn()
}

func getChosenApp() App {
	path := ""
	appsStr := ""

	for _, app := range apps {
		appsStr += "/" + app.name
	}

	fmt.Printf("Which application would you like to use? [%v]\n", appsStr)
	fmt.Scan(&path)

	// * Check if the selected path is a valid selection
	for _, app := range apps {
		if path == app.name {
			return app
		}
	}
	
	fmt.Printf("\"%v\" is not a valid application, please try again.\n", path)
	return getChosenApp()
}