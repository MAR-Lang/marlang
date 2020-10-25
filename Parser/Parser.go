package parser

import (
	"../Token"
)

const (
	NodeBody = iota
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
		this.next = this.tokens[this.index]
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
	if len(this.children) <= index {
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

	return __top_level_node
}
