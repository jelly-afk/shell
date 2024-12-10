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
		args := strings.Split(input, " ")
		switch args[0] {
		case "exit":
			break repl
		case "echo":
			fmt.Println(input[5:])
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
