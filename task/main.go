package main

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/svm-code/Gophercises/task/cmd"
	"github.com/svm-code/Gophercises/task/taskdb"
)

var TaskDBName = "taskdb.db"

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, TaskDBName)
	err := taskdb.Init(dbPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmd.RootCmd.Execute()
}
