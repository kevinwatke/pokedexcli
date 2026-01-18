package main

import "fmt"

func commandHelp() error {
	//commenting out some of this as the lesson calls for a predefined text to print, we will enable printing out all commands later
	// commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	// for _, value := range commands {
	// 	fmt.Println("Command: ", value.name, " Description: ", value.description)
	// }
	return nil
}
