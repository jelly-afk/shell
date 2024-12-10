package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	 fmt.Fprint(os.Stdout, "$ ")

    input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    input = strings.TrimSpace(input)
    switch input {
        default:
        fmt.Printf("%s: command not found\n", input)
    }
}
