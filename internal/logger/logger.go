package logger

import (
	"fmt"
)

func Log(data string) {
	fmt.Printf("\033[1;34m[info]\033[0m \033[1m%s\033[0m\n", data)
}