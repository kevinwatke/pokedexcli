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
		scanner.Scan()
		input := scanner.Text()
		fmt.Println()
		cleanedInput := CleanInput(input)
		fmt.Println("Your command was:", cleanedInput[0])
	}
}
