package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var taskRoot = &cobra.Command{
	Use:   "task",
	Short: "Task services",
}

var createTask = &cobra.Command{
	Use:   "create",
	Short: "Create task inside this project",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		name = strings.Trim(name, " ")
		if name == "" {
			name = getInput("Task name: ")
			if strings.Trim(name, " ") == "" {
				log.Fatal("Empty task name")
			}
		}

		data := readData()
		data.Tasks = append(data.Tasks, Task{Name: name, Created: time.Now(), Id: data.NextId})
		data.NextId += 1
		writeData(data)

	},
}

var listTask = &cobra.Command{
	Use:   "ls",
	Short: "List all created tasks",
	Run: func(cmd *cobra.Command, args []string) {
		data := readData()
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleLight)
		t.AppendHeader(table.Row{"#", "Name", "Created at", "Done at"})
		for i := 0; i < len(data.Tasks); i++ {
			task := data.Tasks[i]
			t.AppendRow(table.Row{task.Id, task.Name, task.Created.Format("02-Jan-2006 03:04 PM"), task.Done.Format("02-Jan-2006 03:04 PM")})
		}
		t.Render()
	},
}

var doneTask = &cobra.Command{
	Use:   "done",
	Short: "Set a task to done",
	Run: func(cmd *cobra.Command, args []string) {
		var idString string
		if len(args) == 0 {
			idString = getInput("Task Id:")
			idString = strings.Trim(idString, " ")
			if idString == "" {
				log.Fatal("Empty task id")
			}
		} else {
			idString = args[0]
			idString = strings.Trim(idString, " ")
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatal("Incorrect task id:", idString)
		}
		fmt.Println(id)
		data := readData()
		changes := false
		for _, a := range data.Tasks {
			if a.Id == id {
				changes = true

			}
		}

		if changes == false {
			log.Fatal("No task with id: ", id)
		}
	},
}

func init() {
	createTask.PersistentFlags().StringP("name", "n", "", "Name of the created task")
	taskRoot.AddCommand(createTask)
	taskRoot.AddCommand(listTask)
	taskRoot.AddCommand(doneTask)
}
