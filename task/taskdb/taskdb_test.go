package taskdb

import (
	"log"
	"os"
	"testing"
)

var TestDBFileName = "taskdb-unit-test.db"
var t1 = "Do testing"
var t2 = "Coverage 100%"

func TestInit(t *testing.T) {
	if _, err := os.Stat(TestDBFileName); err == nil {
		log.Println("File Removed err:", os.Remove(TestDBFileName))
	}
	err := Init(TestDBFileName)
	if err != nil {
		log.Fatal(err)
	}
}

func TestFindAllTask(t *testing.T) {
	tasks, err := FindAllTask()
	if err != nil {
		MyFatalln(err)
	}
	lenght := len(tasks)
	if lenght != 0 {
		MyFatalln("Expected length=0, Got lenght=", lenght)
	}
}

func TestCreateTask(t *testing.T) {
	_, err := CreateTask(t1)
	if err != nil {
		MyFatalln("Create task:", t1, ",Error:", err)
	}

	_, err2 := CreateTask(t2)
	if err2 != nil {
		MyFatalln("Create task:", t2, ",Error:", err2)
	}

	tasks, err := FindAllTask()
	if err != nil {
		MyFatalln(err)
	}
	lenght := len(tasks)
	if lenght != 2 {
		MyFatalln("Epected len task: 2, Found:", lenght)
	}

	if tasks[0].Value != t1 {
		MyFatalln("Epected task:", t1, " , Found:", tasks[0].Value)
	}

	if tasks[1].Value != t2 {
		MyFatalln("Epected task:", t2, " , Found:", tasks[1].Value)
	}
}

func TestDeleteTask(t *testing.T) {
	tasks, err := FindAllTask()
	if err != nil {
		MyFatalln(err)
	}
	legnth := len(tasks)
	if legnth != 2 {
		MyFatalln("Expected length greter than 2, find:", legnth)
	}
	err = DeleteTask(tasks[0].Key)
	if err != nil {
		MyFatalln(err)
	}

	err = DeleteTask(tasks[1].Key)
	if err != nil {
		MyFatalln(err)
	}
}

func TestClose(t *testing.T) {
	Close()
	deleteTempDB()
	_, err := CreateTask(t1)
	if err == nil {
		log.Fatalln("Expeted some error, but err:", err)
	}
	_, err = FindAllTask()
	if err == nil {
		log.Fatalln("Expeted some error, but err:", err)
	}
}

func TestInit_error(t *testing.T) {
	err := Init("")
	if err == nil {
		log.Fatalln("Expeted some error, but err:", err)
	}
}

func MyFatalln(v ...any) {
	Close()
	deleteTempDB()
	log.Fatalln(v...)
}

func deleteTempDB() {
	if _, err := os.Stat(TestDBFileName); err == nil {
		log.Println("File Removed:", TestDBFileName, ", any err:", os.Remove(TestDBFileName))
	}
}
