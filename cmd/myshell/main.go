package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	valid := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
		"pwd":  true,
	}

repl:
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSpace(input)
		args := joinString(input)
		switch args[0] {
		case "exit":
			break repl
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "type":
			if valid[args[1]] {
				fmt.Printf("%s is a shell builtin\n", args[1])
			} else {
				paths := strings.Split(os.Getenv("PATH"), ":")
				filePath := findFile(paths, args[1])
				if filePath == "" {
					fmt.Printf("%s: not found\n", args[1])
				} else {
					fmt.Printf("%s is %s\n", args[1], filePath)
				}
			}
		case "pwd":
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(cwd)
		case "cd":
			if args[1] == "~" {
				os.Chdir(os.Getenv("HOME"))
				break
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", args[1])
			}

		default:
			if len(args) < 2 {
				fmt.Printf("%s: command not found\n", args[0])
				break
			}
			path := os.Getenv("PATH")
			paths := strings.Split(path, ":")
			filePath := findFile(paths, args[0])
			if filePath == "" {
				fmt.Printf("%s: not found\n", args[0])
			} else {
				cmd := exec.Command(filePath, args[1:]...)
				cmd.Stderr = os.Stderr
				cmd.Stdout = os.Stdout
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}

			}
		}
	}
}

func findFile(paths []string, fileName string) string {
	for _, path := range paths {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == fileName {
				return fmt.Sprintf("%s/%s", path, entry.Name())
			}
		}
	}
	return ""
}

func joinString(args string) []string {
	res := make([]string, 0)
	for i := 0; i < len(args); i++ {
		c := args[i]
		if c == ' ' {
			continue
		}
		if c == '\'' {
			idx := strings.Index(args[i+1:], "'")
			res = append(res, args[i+1:i+idx+1])
			i += idx + 1
		} else if c == '"' {
			idx := strings.Index(args[i+1:], "\"")
			res = append(res, args[i+1:i+idx+1])
			i += idx + 1

		} else {
			j := i
			for ; j < len(args); j++ {
				if args[j] == ' ' {
					break
				}
				if args[j] == '\\' {
					args = args[:j] + args[j+1:]
				}
			}
			res = append(res, args[i:j])
			i = j
		}
	}
	return res

}
