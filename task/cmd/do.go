package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/svm-code/Gophercises/task/taskdb"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err != nil {
				fmt.Println("Failed to parse to in id=", a)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := taskdb.FindAllTask()
		if err != nil {
			fmt.Println("Somthing went wrong :: ", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number :", id)
				continue
			}
			task := tasks[id-1]
			taskdb.DeleteTask(task.Key)
			fmt.Printf("Mark %d as Completed.\n", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
