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
		case "int":
			arr, err := parser.ParseInt(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			res = sorter.SortIntArray(arr)
		case "float":
			arr, err := parser.ParseFloat(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			res = sorter.SortFloatArray(arr)
		case "string":
			res = sorter.SortStringArray(args)
			break
		default:
			fmt.Println("Invalid data type. Supported data types are 'int', 'float' or 'string'")
		}

		fmt.Println(res)
	},
}

var dataType string

func init() {
	rootCmd.PersistentFlags().StringVarP(&dataType, "type", "t", "", "Data type of the input. Supported data types are 'int', 'float' or 'string'")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
