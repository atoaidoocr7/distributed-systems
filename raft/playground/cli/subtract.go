package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func Subtract(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}
	return fmt.Sprintf("%f", num1+num2)
}

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"subtraction"},
	Short:   "Subtract 2 numbers",
	Long:    "Carry out subtraction operation on two numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bankudamius")
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
