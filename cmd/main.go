package main

import (
	"bufio"
	"fmt"
	"os"
	l "todo/internal/logger"
	s "todo/internal/storage"
)

func main() {
	path := "data.txt"
	l.Log("the todo app was launched")
	fmt.Print("\033[1m> \033[0m")
	reader := bufio.NewReader(os.Stdin)
	task, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	s.SaveToFile(task, path)
}