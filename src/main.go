package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	numArgs := len(os.Args)

	// Check for arguments
	if numArgs > 2 {
		fmt.Println("Incorrect number of arguments")
		os.Exit(64)
	} else if numArgs == 2 {
		// Run the code from a file
		runFile(os.Args[1])
	} else {
		// Run the code from the command line
		runPrompt()
	}
}

func runFile(path string) {

	// Read from file
	data, err := os.ReadFile(path)

	checkError(err)

	// TODO: Create the run function
	run(string(data))

}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(">")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}
		run(line)
		fmt.Println(">")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin!")
	}
}

func run(source string) {
	fmt.Println(source)
}

// func reportError(line int, message string) {
// 	report(line, "", message)
// }

// func report(line int, where string, message string) {
// 	fmt.Fprintf(os.Stderr, "[line: %d] Error %s: %s", line, where, message)
// }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
