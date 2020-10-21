package main

import (
	"fmt"

	"./Lexer"
)

func main() {
	l := Lexer.New("3.14*2==t")
	fmt.Println("Result: ")
	a := l.Lex()
	fmt.Println(len(a))
	for _, el := range a {
		fmt.Println(el)
	}
}
