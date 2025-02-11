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
	"pwd":  true,
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
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error getting current directory", err)
				os.Exit(1)
			}
			fmt.Println(dir)
		case "exit":
			os.Exit(0)
		default:
			runCommand(args)
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
	if err == nil {
		fmt.Printf("%s is %s\n", cmd, path)
	} else {
		fmt.Printf("%s: not found\n", cmd)
	}
}

func runCommand(args []string) {
	if _, err := exec.LookPath(args[0]); err != nil {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", args[0])
		return
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: command failed: %v\n", args[0], err)
	}
}
