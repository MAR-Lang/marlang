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

func (this *Node) GetChildren() []*Node {
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

		if IsNumder(currentToken) {
			if IsArithmeticOperator(stream.next) {
				node := &Node{_type: true}
				switch stream.next.Type {
				case Token.ADD:
					node.name = ADD
				case Token.SUB:
					node.name = SUB
				case Token.MUL:
					node.name = MUL
				case Token.DIV:
					node.name = DIV
				case Token.MOD:
					node.name = MOD
				case Token.POW:
					node.name = POW
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
	}

	return __top_level_node
}

func IsNumder(token *IToken) bool {
	return token.Type == Token.INT || token.Type == Token.FLOAT
}

func IsArithmeticOperator(token *IToken) bool {
	return token.Type == Token.ADD ||
		token.Type == Token.SUB ||
		token.Type == Token.MUL ||
		token.Type == Token.DIV ||
		token.Type == Token.MOD ||
		token.Type == Token.POW
}
