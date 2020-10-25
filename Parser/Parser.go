package parser

import (
	"strconv"

	"../Token"
)

const (
	NodeBody = iota
	FLOAT
	INT

	ADD
	SUB
	MUL
	DIV
	MOD
	POW
)

type IToken = Token.Token

var __top_level_node *Node
var stream *TokenStream

type TokenStream struct {
	index  int
	tokens []*IToken
	next   *IToken
}

func CreateStream(tokens []*IToken) *TokenStream {
	stream := &TokenStream{index: 0, tokens: tokens}
	return stream
}

func (this *TokenStream) Move() *IToken {
	if this.index < len(this.tokens) {
		next := this.tokens[this.index]
		this.index = this.index + 1
		if this.index < len(this.tokens)-1 {
			this.next = this.tokens[this.index]
		} else {
			this.next = nil
		}
		return next
	} else {
		return nil
	}
}

// Node - struct containing all da shit
type Node struct {
	name     int
	_type    bool
	body     interface{}
	children []*Node
	parent   *Node
}

func (this *Node) AddChild(node *Node) {
	node.parent = this
	this.children = append(this.children, node)
}

func (this *Node) RemoveChild(index int) *Node {
	node := this.children[index]
	for i := index; i < len(this.children)-1; i++ {
		this.children[i] = this.children[i+1]
	}
	this.children = this.children[:len(this.children)-1]
	return node
}

func (this *Node) GetChildern() []*Node {
	return this.children
}

func (this *Node) GetChild(index int) *Node {
	if len(this.children)-1 < index {
		return nil
	}
	return this.children[index]
}

func (this *Node) GetParent() *Node {
	return this.parent
}

func Parse(tokens []*IToken) *Node {
	stream = CreateStream(tokens)
	__top_level_node = &Node{name: NodeBody, _type: true, parent: nil}

	for {
		currentToken := stream.Move()
		if currentToken == nil {
			break
		}

		if currentToken.Type == Token.INT || currentToken.Type == Token.FLOAT {
			var node *Node
			switch stream.next.Type {
			case Token.ADD:
				node = &Node{name: ADD, _type: true}
			case Token.SUB:
				node = &Node{name: SUB, _type: true}
			case Token.MUL:
				node = &Node{name: MUL, _type: true}
			case Token.DIV:
				node = &Node{name: DIV, _type: true}
			case Token.MOD:
				node = &Node{name: MOD, _type: true}
			case Token.POW:
				node = &Node{name: POW, _type: true}
			}
			operand := stream.tokens[stream.index+1]
			if operand.Type != currentToken.Type {
				panic("Invalid type")
			}
			firstOperandNode := &Node{_type: false}
			secondOperandNode := &Node{_type: false}
			if currentToken.Type == Token.FLOAT {
				firstOperandNode.name = FLOAT
				secondOperandNode.name = FLOAT
				firstOperandNode.body, _ = strconv.ParseFloat(currentToken.Value, 64)
				secondOperandNode.body, _ = strconv.ParseFloat(operand.Value, 64)
			} else {
				firstOperandNode.name = INT
				secondOperandNode.name = INT
				firstOperandNode.body, _ = strconv.ParseInt(currentToken.Value, 10, 64)
				secondOperandNode.body, _ = strconv.ParseInt(operand.Value, 10, 64)
			}
			node.AddChild(firstOperandNode)
			node.AddChild(secondOperandNode)

			// will change in future
			__top_level_node.AddChild(node)
			stream.index += 2
		}
	}

	return __top_level_node
}
