package main

import (
	"fmt"
	"horse-lang/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the REPL for the horse programming language.", user.Username)
	fmt.Println("Type some commands and start hacking!")

	// Start the repl
	repl.Start(os.Stdin, os.Stdout)
}
