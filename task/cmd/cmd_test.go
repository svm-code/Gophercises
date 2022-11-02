package cmd

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/svm-code/Gophercises/task/taskdb"
)

var TestDBFileName = "taskdb-unit-test.db"

func TestAdd(t *testing.T) {
	taskdb.Init(TestDBFileName)
	arg := []string{
		"add",
		"hi",
	}
	op, err := executeCommandC(RootCmd, arg...)
	if err != nil {
		deleteTempDB()
		log.Fatal(err)
	}
	op, err = executeCommandC(RootCmd, []string{
		"add",
		"somthings",
	}...)
	op, err = executeCommandC(RootCmd, []string{
		"add",
		"by",
	}...)

	log.Println(op)
}

func TestDo(t *testing.T) {
	arg := []string{
		"do",
		"1",
	}
	op, err := executeCommandC(RootCmd, arg...)
	if err != nil {
		deleteTempDB()
		log.Fatal(err)
	}
	log.Println(op)
}

func TestList(t *testing.T) {
	arg := []string{
		"list",
	}
	op, err := executeCommandC(RootCmd, arg...)
	if err != nil {
		deleteTempDB()
		log.Fatal(err)
	}
	log.Println(op)
}

func TestRoot(t *testing.T) {

	op, err := executeCommandC(RootCmd)
	if err != nil {
		deleteTempDB()
		log.Fatal(err)
	}
	log.Println(op)
}

func Test_Errors(t *testing.T) {
	executeCommandC(RootCmd, []string{
		"do",
		"1",
		"wrong",
		"2",
		"3",
		"4",
		"5",
	}...)

	executeCommandC(RootCmd, []string{
		"list",
	}...)
	taskdb.Close()
	executeCommandC(RootCmd, []string{
		"add",
		"tets",
	}...)

	executeCommandC(RootCmd, []string{
		"do",
		"1",
	}...)

	executeCommandC(RootCmd, []string{
		"list",
	}...)

	deleteTempDB()
}

func executeCommandC(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err = root.ExecuteC()

	return buf.String(), err
}

func deleteTempDB() {
	taskdb.Close()
	if _, err := os.Stat(TestDBFileName); err == nil {
		log.Println("File Removed:", TestDBFileName, ", any err:", os.Remove(TestDBFileName))
	}
}
