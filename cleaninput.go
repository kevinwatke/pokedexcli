package main

import "strings"

func CleanInput(input string) (cleanedInput []string) {
	s := strings.ToLower(input)
	cleanedInput = strings.Fields(s)
	return
}
