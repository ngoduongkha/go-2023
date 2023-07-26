package cmd

import (
	"fmt"
	"github.com/ngoduongkha/go-2023/assignment-02/utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "sorter",
	Short: "A command-line tool to sort input provided by the user.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parser := utils.NewParser()
		sorter := utils.NewSorter()

		var res interface{}

		switch dataType {
		case dataTypeChecker{isInt: true}:
			arr, err := parser.ParseInt(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			res = sorter.SortIntArray(arr)
		case dataTypeChecker{isFloat: true}:
			arr, err := parser.ParseFloat(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}
			res = sorter.SortFloatArray(arr)
		case dataTypeChecker{isString: true}:
			res = sorter.SortStringArray(args)
		default:
			fmt.Println("Invalid data type. Supported data types are 'int', 'float' or 'string'")
			os.Exit(1)
		}

		fmt.Println(res)
	},
}

type dataTypeChecker struct {
	isInt    bool
	isFloat  bool
	isString bool
}

var dataType = dataTypeChecker{}

func init() {
	rootCmd.Flags().BoolVarP(&dataType.isInt, "int", "i", false, "Sorts the input as integers")
	rootCmd.Flags().BoolVarP(&dataType.isFloat, "float", "f", false, "Sorts the input as floats")
	rootCmd.Flags().BoolVarP(&dataType.isString, "string", "s", false, "Sorts the input as strings")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
