package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Basic user interface to your operating system

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		hostname, err := os.Hostname()
		currentWorkingDirectory, err := os.Getwd()
		directorySplit := strings.Split(currentWorkingDirectory, "/")
		currentPath := directorySplit[len(directorySplit)-1]
		fmt.Print(hostname, " ", currentPath, " % ")
		// read in keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// handle execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// split input to separate command arguments
	args := strings.Split(input, " ")

	// check for built-in commands
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// prepare command to execute
	cmd := exec.Command(args[0], args[1:]...)

	//set correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// execute the command and return the error
	return cmd.Run()

}
