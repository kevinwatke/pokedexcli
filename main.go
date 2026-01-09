package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			cleaned := CleanInput(text)
			fmt.Printf("Your command was: %v\n", cleaned[0])
		}
	}
}
