package main

import (
	"fmt"
	"horse-lang/src/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is horse lang\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
