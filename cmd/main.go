package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	l "todo/internal/logger"
	s "todo/internal/storage"
)

var path string = "data.txt"

func main() {
	fmt.Print("\033[2J\033[H")
	l.Log("the \033[1;35m[todo]\033[0m\033[1m app was launched")
	for {
		fmt.Print("\033[1m> \033[0m")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.TrimSpace(input)
		switch {
		case input == "!new":
			fmt.Println("ok, write your task")
			fmt.Print("\033[1m> \033[0m")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("\033[1mtask can't be empty\033[0m")
				continue
			}
			s.SaveToFile(input, path)
		case input == "!tasks":
			s.GetTasks(path)
		case strings.HasPrefix(input, "!done"):
			parts := strings.Fields(input)
			if len(parts) < 2 {
				fmt.Println("usage: !done [task number]")
				fmt.Println("example: !done 3")
				continue
			}

			index, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("invalid task number: ", parts[1])
				continue
			}

			err = s.DeleteTask(path, index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\033[1mtask [%d] completed\033[0m\n", index)
			}
		case input == "!exit":
			l.Log("goodbye")
		default:
			fmt.Println("please enter the command")
			fmt.Println("available commands: !new [task], !tasks, !done [num], !exit")
		}
	}
}
