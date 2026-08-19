package storage

import (
	"fmt"
	"os"
	"strings"
	l "todo/internal/logger"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SaveToFile(task string, path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()

	_, err = file.WriteString(task + "\n")
	check(err)

	l.Log("a new task has been added")
}

func GetTasks(path string) {
	data, err := os.ReadFile(path)
	check(err)
	l.Log("receiving your list...")
	l.Log("your tasks:")

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("\033[1m[%d] %s\033[0m\n", i+1, line)
		}
	}
}

func DeleteTask(path string, index int) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")
	
	if index < 1 || index > len(lines) {
		return fmt.Errorf("task [%d] not found", index)
	}
	
	lines = append(lines[:index-1], lines[index:]...)
	
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}