package Repl

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"

	"../Lexer"
	parser "../Parser"
)

type INode = parser.Node

const PROMPT = "> "

func Start(stdin io.Reader, stdout io.Writer) {
	scanner := bufio.NewScanner(stdin)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "" {
			return
		}

		Parse(line)
	}
}

func ReadFile(filename string) {
	code, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	Parse(string(code))
	return
}

func Parse(code string) {
	lexer := Lexer.New(code)
	tokens := lexer.Lex()
	tree := parser.Parse(tokens)

	fmt.Println("Tokens:")
	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("Abstract tree:")
	showNode(tree, "", 0)
}

func showNode(node *INode, prefix string, index int) {
	if prefix == "" {
		fmt.Printf("%d: ", index)
		fmt.Println(node)
	} else {
		fmt.Printf("%v%d: ", prefix, index)
		fmt.Println(node)
	}
	for i, n := range node.GetChildren() {
		showNode(n, prefix+"   ", i)
	}
}
