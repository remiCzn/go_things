package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init ToDo project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.Mkdir(".todo", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		data := Data{
			NextId: 0,
			Tasks:  []Task{},
		}

		writeData(data)
	},
}
