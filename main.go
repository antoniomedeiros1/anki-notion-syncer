package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
)

func main() {

	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(opts); return nil })

	menu.Option("Add a new deck", 0, true, nil)
	menu.Option("Update deck (sync with Notion)", 1, true, nil)
	menu.Option("List decks", 2, true, nil)
	menu.Option("Delete deck", 3, true, nil)

	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}

}

func handleFunc(opts []wmenu.Opt) {

	switch opts[0].Value {

	case 0:
		fmt.Println("Adding deck")
	case 1:
		fmt.Println("Updating deck")
	case 2:
		fmt.Println("Listing decks")
	case 3:
		fmt.Println("Deleting deck")

	}

}
