package main

import (
	"fmt"
	"os"
	"strings"
)

// countryFormats is a map of country codes to a boolean value indicating whether the country's naming order is
var countryFormats = map[string]bool{
	"VN": false,
	"US": true,
	"GB": true,
	"DE": true,
	"FR": true,
	"JP": false,
	"CN": false,
}

func reorderName(firstName, lastName, middleName, countryCode string) string {
	namingOrder := countryFormats[countryCode]

	if namingOrder {
		if middleName == "" {
			return fmt.Sprintf("%s %s", firstName, lastName)
		}

		return fmt.Sprintf("%s %s %s", firstName, middleName, lastName)
	}

	if middleName == "" {
		return fmt.Sprintf("%s %s", lastName, firstName)
	}

	return fmt.Sprintf("%s %s %s", lastName, firstName, middleName)
}

func main() {
	args := os.Args[1:] // Get the command-line arguments excluding the program name
	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Usage: go run main.go <first name> <last name> [middle name] <country code>")
		return
	}

	countryCode := strings.ToUpper(args[len(args)-1])

	if _, ok := countryFormats[countryCode]; !ok {
		fmt.Println("Invalid country code")
		acceptableCountryCodes := make([]string, 0, len(countryFormats))
		for countryCode := range countryFormats {
			acceptableCountryCodes = append(acceptableCountryCodes, countryCode)
		}
		fmt.Println("Supported country codes: ", acceptableCountryCodes)
		return
	}

	firstName := args[0]
	lastName := args[1]
	middleName := ""

	if len(args) == 4 {
		middleName = args[2]
	}

	reorderedName := reorderName(firstName, lastName, middleName, countryCode)
	fmt.Println(reorderedName)
}
