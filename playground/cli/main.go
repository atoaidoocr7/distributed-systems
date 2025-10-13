package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "playground",
	Short: "playground is a simple raft server",
	Long:  `playground is a simple raft server`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err = fmt.Fprintf(os.Stderr, "An error occure while executing zero %s\n", err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
