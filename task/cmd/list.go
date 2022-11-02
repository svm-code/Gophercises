package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/svm-code/Gophercises/task/taskdb"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskdb.FindAllTask()

		if err != nil {
			fmt.Println("Somthing went wrong=", err)
		}

		if len(tasks) == 0 {
			fmt.Println("you have no task to commplete! Why not take a vacation?")
			return
		}
		fmt.Println("You have following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
