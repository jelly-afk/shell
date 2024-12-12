package main

import (
	"bufio"
	"fmt"
	"log"
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
				paths := strings.Split(os.Getenv("PATH"), ":")
				exists := false
				var err error
				for _, path := range paths {
					exists, err = findFile(path, args[1])
					if err != nil {
						log.Fatal(err)
					}
					if exists {
						fmt.Printf("%s is %s/%s\n", args[1], path, args[1])
						break
					}
				}
				if !exists {
					fmt.Printf("%s: not found\n", args[1])
				}

			}
		default:
			fmt.Printf("%s: command not found\n", input)
		}
	}
}

func findFile(dir, fileName string) (bool, error) {
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		if !entry.IsDir() && entry.Name() == fileName {
			return true, nil
		}
	}
	return false, nil
}
