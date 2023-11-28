package main

import (
	"bufio"
	"fmt"
	"horse-lang/src/lexer"
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
	scanner := lexer.New(source)
	scanner.ScanTokens()
	for _, token := range scanner.Tokens {
		fmt.Printf(token.Lexeme)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
