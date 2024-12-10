package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	valid := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
	}

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
		case "type":
			if valid[args[1]] {
				fmt.Printf("%s is a shell builtin\n", args[1])
			} else {
				fmt.Printf("%s: not found\n", args[1])
			}
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}
