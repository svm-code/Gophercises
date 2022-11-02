package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	TaskDBName = "taskdb-unit.db"
	main()
}
