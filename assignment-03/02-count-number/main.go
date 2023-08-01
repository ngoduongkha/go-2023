package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func countDifferentIntegers(word string) int {
	// Regular expression to match integers in the string.
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(word, -1)

	// Create a set to store unique integers.
	uniqueIntegers := make(map[string]struct{})

	for _, match := range matches {
		// Remove leading zeros from the integer.
		num, _ := strconv.Atoi(match)
		normalizedNum := strconv.Itoa(num)

		// Store the normalized integer in the set.
		uniqueIntegers[normalizedNum] = struct{}{}
	}

	return len(uniqueIntegers)
}

func main() {
	word := "a123bc34d8ef34"

	count := countDifferentIntegers(word)

	fmt.Println(count)
}
