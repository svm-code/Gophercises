package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/svm-code/Gophercises/task/taskdb"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := taskdb.CreateTask(task)
		if err != nil {
			fmt.Println("Somthing went wrong!")
			return
		}
		fmt.Println("Added new Task: ", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
