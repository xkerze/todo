package storage

import (
	"os"
	l "todo/internal/logger"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveToFile(task string, path string) {
	dump := []byte(task)
	err := os.WriteFile(path, dump, 0644)
	check(err)
	l.Log("a new task has been added")
}
