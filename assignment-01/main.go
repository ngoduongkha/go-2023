package main

import (
	"fmt"
	"github.com/ngoduongkha/go-2023/assignment-01/constants"
	"github.com/ngoduongkha/go-2023/assignment-01/utils"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // Get the command-line arguments excluding the program name
	if len(args) < 3 {
		fmt.Println("Invalid number of arguments")
		fmt.Println("Usage: go run main.go <first name> <last name> [<middle name>] <country code>")
		return
	}

	countryCodeStr := strings.ToUpper(args[len(args)-1])
	countryCode, err := constants.ParseCountryCode(countryCodeStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	reorderFactory := utils.New()
	reorder, err := reorderFactory.GetStrategy(countryCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reorder.Do(args[0], args[1], args[2:len(args)-1]...))
}
