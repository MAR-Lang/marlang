package main

import (
	"os"

	"./Repl"
)

func main() {
	if len(os.Args) > 1 {
		Repl.ReadFile(os.Args[1])
		return
	}

	Repl.Start(os.Stdin, os.Stdout)

}
