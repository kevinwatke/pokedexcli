package main

import (
	"bufio"
	"fmt"
	"os"
)

func Repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		fmt.Printf("You entered: %s\n", input)
	}
}
