package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command := readCommand()
		args := strings.Fields(command)

		switch args[0] {
		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(command + ": command not found")
		}
	}
}

func readCommand() string {
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading command", err)
		os.Exit(1)
	}
	return strings.TrimSpace(command)
}
