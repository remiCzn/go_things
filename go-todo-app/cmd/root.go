package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Todo-Cli",
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(taskRoot)
	rootCmd.AddCommand(destroyCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
