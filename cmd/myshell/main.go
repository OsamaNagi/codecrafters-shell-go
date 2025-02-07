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

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		command = strings.TrimSpace(command)

		if command == "exit 0" {
			os.Exit(0)
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command", err)
			os.Exit(1)
		}

		fmt.Println(command + ": command not found")
	}
}
