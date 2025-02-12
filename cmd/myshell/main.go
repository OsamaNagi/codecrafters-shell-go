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
	"cd":   true,
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command := readCommand()
		args := parseArguments(command)

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
		case "cd":
			if len(args) > 1 {
				changeDirectory(args[1])
			} else {
				fmt.Println("cd: missing argument")
			}
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

// $ echo 'hello' 'world'
// 'hello' 'world'
// $ echo hello world

func parseArguments(input string) []string {
	var tokens []string
	var token strings.Builder
	inQuote := false

	i := 0
	for i < len(input) {
		c := input[i]

		if c == '\'' {
			inQuote = !inQuote
			i++
			continue
		}

		if !inQuote && (c == ' ' || c == '\t') {
			if token.Len() > 0 {
				tokens = append(tokens, token.String())
				token.Reset()
			}

			for i < len(input) && (input[i] == ' ' || input[i] == '\t') {
				i++
			}
			continue
		}

		token.WriteByte(c)
		i++
	}

	if token.Len() > 0 {
		tokens = append(tokens, token.String())
	}
	return tokens
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

func changeDirectory(path string) {
	if path == "~" {
		path = os.Getenv("HOME")
	}

	if !strings.HasPrefix(path, "/") {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting current directory", err)
			os.Exit(1)
		}
		path = dir + "/" + path
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
	}
}
