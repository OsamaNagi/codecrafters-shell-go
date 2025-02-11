package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtin = map[string]bool{
	"echo": true,
	"type": true,
	"exit": true,
}

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
		case "type":
			if len(args) > 1 {
				for _, cmd := range args[1:] {
					isBuiltin(cmd)
				}
			} else {
				fmt.Println("type: missing argument")
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

func isBuiltin(cmd string) {
	if builtin[cmd] {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("%s: not found\n", cmd)
	} else {
		fmt.Printf("%s is %s\n", cmd, path)
	}
}
