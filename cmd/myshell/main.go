package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

repl:
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {

		case "exit 0":
			break repl
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
