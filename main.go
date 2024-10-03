package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/qukuqhd/Interpreter/repl"
	"github.com/qukuqhd/Interpreter/run"
)

func main() {
	if len(os.Args) == 2 {
		scriptFile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		} else {
			run.Run(scriptFile, os.Stdout)
		}
	} else {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

}
