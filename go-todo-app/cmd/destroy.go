package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy everything",
	Run: func(cmd *cobra.Command, args []string) {
		input := getInput("Are you sure to destroy all ToDo datas? (y): ")
		if input == "y" {
			err := os.RemoveAll(".todo")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Not destroyed")
		}

	},
}
