package main

import (
	"fmt"
	"github.com/interperter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is bullshit programming language!\n", user.Username)
	fmt.Printf("Try to write something!\n")
	repl.Start(os.Stdin, os.Stdout)
}
